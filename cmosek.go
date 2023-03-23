// gmsk is an unofficial wrapper for MOSEK, the conic optimizer from [MOSEK ApS].
//
// gmsk is based on mosek's C api, which must be installed and configured before
// using gmsk.
//
// [MOSEK ApS]: https://www.mosek.com
package gmsk

// #cgo LDFLAGS: -lmosek64
//
// #include <stdlib.h> // stdlib.h is required for calloc and free
// #include <mosek.h>
//
// extern void writeStreamToWriter(void*, char*); // for MSK_linkfunctotaskstream
// extern size_t writeFuncToWriter(void*, void*, size_t); // MSKhwritefunc
import "C"

import (
	"fmt"
	"io"
	"runtime/cgo"
	"unsafe"
)

// Env wraps mosek environment
type Env struct {
	env C.MSKenv_t
}

// MakeEnv creates a new mosek environment.
// dbgfile can be zero or one, error when more than two dbgfile are provided.
func MakeEnv(dbgfile ...string) (*Env, error) {
	var env C.MSKenv_t = nil
	var cdbgfile *C.char = nil
	if len(dbgfile) > 1 {
		return nil, fmt.Errorf("too many dbgfiles: %v", dbgfile)
	}
	if len(dbgfile) == 1 {
		cdbgfile = C.CString(dbgfile[0])
		if dbgfile != nil {
			defer C.free(unsafe.Pointer(cdbgfile))
		}
	}

	r := uint32(C.MSK_makeenv(&env, cdbgfile))
	if r != RES_OK {
		return nil, fmtError("failed to create environment: %s %s", r)
	}

	return &Env{env: env}, nil
}

// getContainingEnv gets the containing env. if env itself is nil, return nil
func (env *Env) getEnv() C.MSKenv_t {
	if env == nil {
		return nil
	}
	return env.env
}

// MakeTask creates a new task in this environment
func (env *Env) MakeTask(maxnumcon int32, maxnumvar int32) (*Task, error) {
	return MakeTask(env, maxnumcon, maxnumvar)
}

// DeleteEnv deletes the environment
func DeleteEnv(env *Env) {
	if env.env != nil {
		C.MSK_deleteenv(&env.env)
	}
}

// MSKTask holds a MSKtask_t, MOSEK task
type Task struct {
	task C.MSKtask_t

	// writerHandles are handles to [writerHolder]
	// those are freed when the task is deleted.
	writerHandles []cgo.Handle
}

// MakeTask is the equivalent of MSK_maketask.
// To use the global environment, use nil for env
func MakeTask(env *Env, maxnumcon, maxnumvar int32) (*Task, error) {
	var task C.MSKtask_t = nil
	var e C.MSKenv_t = nil
	if env != nil {
		e = env.env
	}
	r := uint32(C.MSK_maketask(e, C.MSKint32t(maxnumcon), C.MSKint32t(maxnumvar), &task))
	if r != RES_OK {
		return nil, fmtError("failed to create task: %s %s", r)
	}

	return &Task{task: task}, nil
}

// DeleteTask deletes the task
func DeleteTask(task *Task) {
	if task != nil && task.task != nil {
		C.MSK_deletetask(&task.task)
	}

	for _, h := range task.writerHandles {
		h.Delete()
	}
}

// GetCodeDesc set the content of sym and desc just like MSK_getcodedesc.
// However, the size of the input slice must be greater than or equal to MAX_STR_LEN + 1
// The returned bool indicate if the size of sym/desc is greater than [MAX_STR_LEN].
// The function will do nothing if sym and desc are smaller than or equal to [MAX_STR_LEN].
func GetCodeDesc(resCode uint32, sym []byte, desc []byte) (uint32, bool) {
	if len(sym) <= MAX_STR_LEN || len(desc) <= MAX_STR_LEN {
		return 0, false
	}

	symstart := unsafe.Pointer(&sym[0])
	descstart := unsafe.Pointer(&desc[0])
	r := C.MSK_getcodedesc(C.MSKrescodee(resCode), (*C.char)(symstart), (*C.char)(descstart))

	return uint32(r), true
}

// GetCodeDescSimple gets the description of the return code.
// Please check uint32 to see if the operation is successful.
// The first returned string is the symbol,
// and the second returned string is the description.
// This is different from [GetCodeDesc] and involves allocating and freeing two char array of
// [MAX_STR_LEN] + 1 size.
func GetCodeDescSimple(resCode uint32) (uint32, string, string) {
	symbol := (*C.char)(C.calloc(MAX_STR_LEN+1, C.sizeof_char))
	defer C.free(unsafe.Pointer(symbol))
	des := (*C.char)(C.calloc(MAX_STR_LEN+1, C.sizeof_char))
	defer C.free(unsafe.Pointer(des))

	r := C.MSK_getcodedesc(C.MSKrescodee(resCode), symbol, des)
	return uint32(r), C.GoString(symbol), C.GoString(symbol)
}

// fmtError formats the error string
func fmtError(format string, resCode uint32) error {
	_, symbol, desc := GetCodeDescSimple(resCode)
	return fmt.Errorf(format, symbol, desc)
}

// AppendVars wraps MSK_appendvars, which adds variables
// to the task.
func (task *Task) AppendVars(num int32) uint32 {
	return uint32(C.MSK_appendvars(task.task, C.MSKint32t(num)))
}

// AppendCons wraps MSK_appendcons, which add linear constraints
// to the task
func (task *Task) AppendCons(numcon int32) uint32 {
	return uint32(C.MSK_appendcons(task.task, C.MSKint32t(numcon)))
}

// PutCj wraps MSK_putcj, which set the coefficient in the objective function.
func (task *Task) PutCj(j int32, c_j float64) uint32 {
	return uint32(C.MSK_putcj(task.task, C.MSKint32t(j), C.MSKrealt(c_j)))
}

// PutCSlice wraps MSK_putcslice, which set a slice of coefficients in the objective
func (task *Task) PutCSlice(first, last int32, slice *float64) uint32 {
	return uint32(C.MSK_putcslice(task.task, C.MSKint32t(first), C.MSKint32t(last), (*C.MSKrealt)(slice)))
}

// PutVarType wraps MSK_putvartype and sets the type of the variable
func (task *Task) PutVarType(j int32, vartype VariableType) uint32 {
	return uint32(C.MSK_putvartype(task.task, C.MSKint32t(j), C.MSKvariabletypee(vartype)))
}

// PutVarbound wraps MSK_putvarbound, which set the bound for a variable.
func (task *Task) PutVarbound(j int32, bkx BoundKey, blx, bux float64) uint32 {
	return uint32(C.MSK_putvarbound(task.task, C.MSKint32t(j), C.MSKboundkeye(bkx), C.MSKrealt(blx), C.MSKrealt(bux)))
}

// PutVarboundSliceConst wraps MSK_putvarboundsliceconst, which set the bound for a slice of variables.
func (task *Task) PutVarboundSliceConst(first, last int32, bkx BoundKey, blx, bux float64) uint32 {
	return uint32(C.MSK_putvarboundsliceconst(task.task, C.MSKint32t(first), C.MSKint32t(last), C.MSKboundkeye(bkx), C.MSKrealt(blx), C.MSKrealt(bux)))
}

// PutConBound wraps MSK_putconbound, which set the bound for a contraint
func (task *Task) PutConBound(i int32, bkc BoundKey, blc, buc float64) uint32 {
	return uint32(C.MSK_putconbound(task.task, C.MSKint32t(i), C.MSKboundkeye(bkc), C.MSKrealt(blc), C.MSKrealt(buc)))
}

// PutObjsense wraps MSK_putobjsense set the objective sense - which is either minimize or maximize
func (task *Task) PutObjsense(sense ObjectiveSense) uint32 {
	return uint32(C.MSK_putobjsense(task.task, C.MSKobjsensee(sense)))
}

// PutAij wraps MSK_putaij, which set the value of the linear constraints matrix A[i,j]
func (task *Task) PutAij(i, j int32, aij float64) uint32 {
	return uint32(C.MSK_putaij(task.task, C.MSKint32t(i), C.MSKint32t(j), C.MSKrealt(aij)))
}

// PutACol wraps MSK_putacol, and puts a column of A matrix.
func (task *Task) PutACol(j int32, nzj int32, subj *int32, valj *float64) uint32 {
	return uint32(C.MSK_putacol(task.task, C.MSKint32t(j), C.MSKint32t(nzj), (*C.MSKint32t)(subj), (*C.MSKrealt)(valj)))
}

// AppendAfes wraps MSK_appendafes and adds affine expressions to the task.
func (task *Task) AppendAfes(num int64) uint32 {
	return uint32(C.MSK_appendafes(task.task, C.MSKint64t(num)))
}

// PutAfeFEntry wraps MSK_putafefentry and set an entry in the  affine expression F matrix.
func (task *Task) PutAfeFEntry(afeidx int64, varidx int32, value float64) uint32 {
	return uint32(C.MSK_putafefentry(task.task, C.MSKint64t(afeidx), C.MSKint32t(varidx), C.MSKrealt(value)))
}

// PutAfeFEntryList wraps MSK_putafefentrylist, which set a portion of the affine expression F matrix
func (task *Task) PutAfeFEntryList(numentr int64, afeidx *int64, varidx *int32, val *float64) uint32 {
	return uint32(C.MSK_putafefentrylist(task.task, C.MSKint64t(numentr), (*C.MSKint64t)(afeidx), (*C.MSKint32t)(varidx), (*C.MSKrealt)(val)))
}

// PutAfeFRow wraps MSK_putafefrow and sets a row of affine expression F matrix
func (task *Task) PutAfeFRow(afeidx int64, numnz int32, varidx *int32, val *float64) uint32 {
	return uint32(C.MSK_putafefrow(task.task, C.MSKint64t(afeidx), C.MSKint32t(numnz), (*C.MSKint32t)(varidx), (*C.MSKrealt)(val)))
}

// PutAfeFCol wraps MSK_putafefcol and sets a column of affine expression F matrix
func (task *Task) PutAfeFCol(varidx int32, numnz int64, afeidx *int64, val *float64) uint32 {
	return uint32(C.MSK_putafefcol(task.task, C.MSKint32t(varidx), C.MSKint64t(numnz), (*C.MSKint64t)(afeidx), (*C.MSKrealt)(val)))
}

// PutAfeG wraps MSK_putafeg and sets the value at afeidx to g
func (task *Task) PutAfeG(afeidx int64, g float64) uint32 {
	return uint32(C.MSK_putafeg(task.task, C.MSKint64t(afeidx), C.MSKrealt(g)))
}

// PutAfeGSlice wraps MSK_putafegslice and sets a slice of values in g
func (task *Task) PutAfeGSlice(first, last int64, slice *float64) uint32 {
	return uint32(C.MSK_putafegslice(task.task, C.MSKint64t(first), C.MSKint64t(last), (*C.MSKrealt)(slice)))
}

// AppendRZeroDomain wraps MSK_appendrzerodomain and add a real zero domain of dimension n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendRZeroDomain(n int64) (r uint32, domidx int64) {
	r = uint32(C.MSK_appendrzerodomain(task.task, C.MSKint64t(n), (*C.MSKint64t)(&domidx)))
	return
}

// AppendQuadraticConeDomain wraps MSK_appendquadraticconedomain and adds a new quadratic cone of size n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendQuadraticConeDomain(n int64) (r uint32, domidx int64) {
	r = uint32(C.MSK_appendquadraticconedomain(task.task, C.MSKint64t(n), (*C.MSKint64t)(&domidx)))
	return
}

// AppendRotatedQuadraticConeDomain wraps MSK_appendrquadraticconedomain and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendRotatedQuadraticConeDomain(n int64) (r uint32, domidx int64) {
	r = uint32(C.MSK_appendrquadraticconedomain(task.task, C.MSKint64t(n), (*C.MSKint64t)(&domidx)))
	return
}

// AppendRQuadraticConeDomain wraps MSK_appendrquadraticconedomain and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful - this is same as [AppendRotatedQuadraticConeDomain], but with the word "rotated" fully spelled out.
func (task *Task) AppendRQuadraticConeDomain(n int64) (r uint32, domidx int64) {
	r = uint32(C.MSK_appendrquadraticconedomain(task.task, C.MSKint64t(n), (*C.MSKint64t)(&domidx)))
	return
}

// AppendPrimalPowerConeDomain wraps MSK_appendprimalpowerconedomain and add a primal power cone to the task
func (task *Task) AppendPrimalPowerConeDomain(n, nleft int64, alpha *float64) (r uint32, domidx int64) {
	r = uint32(C.MSK_appendprimalpowerconedomain(task.task, C.MSKint64t(n), C.MSKint64t(nleft), (*C.MSKrealt)(alpha), (*C.MSKint64t)(&domidx)))
	return
}

// AppendAcc wraps MSK_appendacc and adds an affine conic constraint to the task, where the afe idx is provided
// by an array or pointer - if the afe idx is sequential, use [Task.AppendAccSeq] to avoid allocating an array.
func (task *Task) AppendAcc(domidx, numafeidx int64, afeidxlist *int64, b *float64) uint32 {
	return uint32(C.MSK_appendacc(task.task, C.MSKint64t(domidx), C.MSKint64t(numafeidx), (*C.MSKint64t)(afeidxlist), (*C.MSKrealt)(b)))
}

// AppendAccs wraps MSK_appendacc and adds a list of affine conic constraints to the task.
func (task *Task) AppendAccs(numaccs int64, domidxs *int64, numafeidx int64, afeidxlist *int64, b *float64) uint32 {
	return uint32(C.MSK_appendaccs(task.task, C.MSKint64t(numaccs), (*C.MSKint64t)(domidxs), C.MSKint64t(numafeidx), (*C.MSKint64t)(afeidxlist), (*C.MSKrealt)(b)))
}

// AppendAccSeq wraps MSK_appendaccseq and adds an affine conic constraint to the task where
// the affine idx is sequential.
func (task *Task) AppendAccSeq(domidx, numafeidx, afeidxfirst int64, b *float64) uint32 {
	return uint32(C.MSK_appendaccseq(task.task, C.MSKint64t(domidx), C.MSKint64t(numafeidx), C.MSKint64t(afeidxfirst), (*C.MSKrealt)(b)))
}

// AppendAccsSeq wraps MSK_appendaccsseq and append a block of accs to the tas - assuming affine expressions are sequential.
func (task *Task) AppendAccsSeq(numaccs int64, domidxs *int64, numafeidx, afeidxfirst int64, b *float64) uint32 {
	return uint32(C.MSK_appendaccsseq(task.task, C.MSKint64t(numaccs), (*C.MSKint64t)(domidxs), C.MSKint64t(numafeidx), C.MSKint64t(afeidxfirst), (*C.MSKrealt)(b)))
}

// PutVarName wraps MSK_putvarname and sets a name for variable at j.
// Allocate a new C array and copy the data over, then free it - this is a costly function.
func (task *Task) PutVarName(j int32, name string) uint32 {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return uint32(C.MSK_putvarname(task.task, C.MSKint32t(j), cstr))
}

// PutConName wraps MSK_putconname and sets a name for a linear constraint at indext i.
func (task *Task) PutConName(i int32, name string) uint32 {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return uint32(C.MSK_putconname(task.task, C.MSKint32t(i), cstr))
}

// PutAccName wraps MSK_putaccname and sets a name for an affine conic constraint.
func (task *Task) PutAccName(accidx int64, name string) uint32 {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return uint32(C.MSK_putaccname(task.task, C.MSKint64t(accidx), cstr))
}

// OptimizeTerm wraps MSK_optimizeterm, which optimizes the problem.
func (task *Task) OptimizeTerm() (res uint32, trmcode uint32) {
	res = uint32(C.MSK_optimizetrm(task.task, (*C.MSKrescodee)(&trmcode)))
	return
}

// GetNumVar wraps MSK_getnumvar, which obtains the number of variables in task.
func (task *Task) GetNumVar() (res uint32, numVar int32) {
	res = uint32(C.MSK_getnumvar(task.task, (*C.MSKint32t)(&numVar)))
	return
}

// GetSolSta wraps MSK_getsolsta, which returns the solution status
func (task *Task) GetSolSta(whichsol SolType) (res uint32, solSta SolSta) {
	res = uint32(C.MSK_getsolsta(task.task, C.MSKsoltypee(whichsol), (*C.MSKsolstae)(&solSta)))
	return
}

// GetXx wraps MSK_getxx, which gets the solution from the task.
// xx can be nil, in which case the number of variables will be queried from the task and
// a new slice created.
func (task *Task) GetXx(whichsol SolType, xx []float64) (uint32, []float64) {
	var res uint32
	var numVar int32
	if xx == nil {
		res, numVar = task.GetNumVar()
		if res != RES_OK {
			return res, nil
		}
		xx = make([]float64, numVar)
	}

	res = uint32(C.MSK_getxx(task.task, C.MSKsoltypee(whichsol), (*C.MSKrealt)(&xx[0])))

	return res, xx
}

// GetXxSlice wraps MSK_getxxslice, which gets a slice of the solution. xx can be nil, in which case the number of variables
// will be last - first and a new slice will be created.
func (task *Task) GetXxSlice(whichsol SolType, first, last int32, xx []float64) (uint32, []float64) {
	var res uint32
	if xx == nil {
		xx = make([]float64, last-first)
	}

	res = uint32(C.MSK_getxxslice(task.task, C.MSKsoltypee(whichsol), C.MSKint32t(first), C.MSKint32t(last), (*C.MSKrealt)(&xx[0])))

	return res, xx
}

// GetAccN wraps MSK_getaccn and returns the dimension of cone at index accidx.
func (task *Task) GetAccN(accidx int64) (uint32, int64) {
	var accn int64
	res := uint32(C.MSK_getaccn(task.task, C.MSKint64t(accidx), (*C.MSKint64t)(&accn)))
	return res, accn
}

// GetAccDotY wraps MSK_getaccdoty and returns doty dual result of cone at idnex accidx.
// doty can be nil, in which case the dimension of the cone will be queried from the task and
// a new slice will be created.
func (task *Task) GetAccDotY(whichsol SolType, accidx int64, doty []float64) (uint32, []float64) {
	var res uint32

	if doty == nil {
		var numdoty int64
		res, numdoty = task.GetAccN(accidx)
		if res != RES_OK {
			return res, nil
		}
		doty = make([]float64, numdoty)
	}

	res = uint32(C.MSK_getaccdoty(task.task, C.MSKsoltypee(whichsol), C.MSKint64t(accidx), (*C.MSKrealt)(&doty[0])))

	return res, doty
}

// EvaluateAcc gets the activity of the cone at index accidx
func (task *Task) EvaluateAcc(whichsol SolType, accidx int64, activity []float64) (uint32, []float64) {
	var res uint32

	if activity == nil {
		var numdoty int64
		res, numdoty = task.GetAccN(accidx)
		if res != RES_OK {
			return res, nil
		}
		activity = make([]float64, numdoty)
	}

	res = uint32(C.MSK_evaluateacc(task.task, C.MSKsoltypee(whichsol), C.MSKint64t(accidx), (*C.MSKrealt)(&activity[0])))

	return res, activity
}

// PutIntParam wraps MSK_putintparam and sets the param to parvalue.
func (task *Task) PutIntParam(param IParam, parvalue int32) uint32 {
	return uint32(C.MSK_putintparam(task.task, C.MSKiparame(param), C.MSKint32t(parvalue)))
}

// writerHolder holds a writer. This must be passed to C api with a handle.
type writerHolder struct {
	writer io.Writer
}

// writeStreamToWriter is the function for C api's task stream handler,
// and it has a signature of
//
//	void streamfunc(void* handle, char* data)
//
//export writeStreamToWriter
func writeStreamToWriter(p unsafe.Pointer, v *C.char) {
	h := cgo.Handle(p)

	w, ok := h.Value().(writerHolder)
	if !ok {
		return
	}
	if w.writer == nil {
		return
	}

	w.writer.Write([]byte(C.GoString(v)))
}

// LinkFuncToTaskStream wraps MSK_linkfuctotaskstream using [io.Writer] instead of callbacks.
func (task *Task) LinkFuncToTaskStream(whichstream StreamType, w io.Writer) uint32 {
	writer := writerHolder{
		writer: w,
	}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return uint32(C.MSK_linkfunctotaskstream(task.task, C.MSKstreamtypee(whichstream), C.MSKuserhandle_t(ptr), (*[0]byte)(C.writeStreamToWriter)))
}

// SolutionSummary wraps MSK_solutionsummary, which prints the summary of the solution to the given stream.
func (task *Task) SolutionSummary(whichstream StreamType) uint32 {
	return uint32(C.MSK_solutionsummary(task.task, C.MSKstreamtypee(whichstream)))
}

// WriteData wraps MSK_writedata and write data to a file.
func (task *Task) WriteData(filename string) uint32 {
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	return uint32(C.MSK_writedata(task.task, filenameC))
}

// ReadData wraps MSK_readdata and read data from a file
func (task *Task) ReadData(filename string) uint32 {
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	return uint32(C.MSK_readdata(task.task, filenameC))
}

// writeFuncToWriter is the function for C api's MSKhwritefunc, which has a signature of
//
//	void writefunchandle(void* handle, void* src, size_t count)
//
//export writeFuncToWriter
func writeFuncToWriter(handle unsafe.Pointer, src unsafe.Pointer, count C.size_t) C.size_t {
	h := cgo.Handle(handle)
	w, ok := h.Value().(writerHolder)
	if !ok {
		return 0
	}
	if w.writer == nil {
		return 0
	}

	bytePtr := (*byte)(src)

	byteSlice := unsafe.Slice(bytePtr, count)

	n, _ := w.writer.Write(byteSlice)
	return C.size_t(n)
}

// WriteDataHandle wraps MSK_writedatahandle using [io.Writer] instead of using callbacks.
func (task *Task) WriteDataHandle(handle io.Writer, format DataFormat, compress CompressType) uint32 {
	writer := writerHolder{writer: handle}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return uint32(C.MSK_writedatahandle(task.task, (*[0]byte)(C.writeFuncToWriter), C.MSKuserhandle_t(ptr), C.MSKdataformate(format), C.MSKcompresstypee(compress)))
}

// Potrf wraps MSK_potrf and performs Cholesky decomposition of symmetric
// square matrix a.
func POTRF(env *Env, uplo UpLo, n int32, a *float64) uint32 {
	return uint32(C.MSK_potrf(env.getEnv(), C.MSKuploe(uplo), C.MSKint32t(n), (*C.MSKrealt)(a)))
}

// GEMM wraps MSK_gemm and performs a general matrix multiplication
func GEMM(env *Env, transa, transb Transpose, m, n, k int32, alpha float64, a, b *float64, beta float64, c *float64) uint32 {
	return uint32(C.MSK_gemm(
		env.getEnv(),
		C.MSKtransposee(transa),
		C.MSKtransposee(transb),
		C.MSKint32t(m),
		C.MSKint32t(n),
		C.MSKint32t(k),
		C.MSKrealt(alpha),
		(*C.MSKrealt)(a),
		(*C.MSKrealt)(b),
		C.MSKrealt(beta),
		(*C.MSKrealt)(c)))
}
