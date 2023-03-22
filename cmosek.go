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

// MakeEnv creates a new mosek environment
func MakeEnv() (*Env, error) {
	var env C.MSKenv_t = nil
	r := C.MSK_makeenv(&env, nil)
	if r != RES_OK {
		return nil, fmtError("failed to create environment: %s %s", r)
	}

	return &Env{env: env}, nil
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
func MakeTask(env *Env, maxnumcon, maxnumvar Int32t) (*Task, error) {
	var task C.MSKtask_t = nil
	var e C.MSKenv_t = nil
	if env != nil {
		e = env.env
	}
	r := C.MSK_maketask(e, maxnumcon, maxnumvar, &task)
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
func GetCodeDesc(resCode ResCode, sym []byte, desc []byte) (ResCode, bool) {
	if len(sym) <= MAX_STR_LEN || len(desc) <= MAX_STR_LEN {
		return 0, false
	}

	symstart := unsafe.Pointer(&sym[0])
	descstart := unsafe.Pointer(&desc[0])
	r := C.MSK_getcodedesc(resCode, (*C.char)(symstart), (*C.char)(descstart))

	return r, true
}

// GetCodeDescSimple gets the description of the return code.
// Please check ResCode to see if the operation is successful.
// The first returned string is the symbol,
// and the second returned string is the description.
// This is different from [GetCodeDesc] and involves allocating and freeing two char array of
// [MAX_STR_LEN] + 1 size.
func GetCodeDescSimple(resCode ResCode) (ResCode, string, string) {
	symbol := (*C.char)(C.calloc(MAX_STR_LEN+1, C.sizeof_char))
	defer C.free(unsafe.Pointer(symbol))
	des := (*C.char)(C.calloc(MAX_STR_LEN+1, C.sizeof_char))
	defer C.free(unsafe.Pointer(des))

	r := C.MSK_getcodedesc(resCode, symbol, des)
	return r, C.GoString(symbol), C.GoString(symbol)
}

// fmtError formats the error string
func fmtError(format string, resCode ResCode) error {
	_, symbol, desc := GetCodeDescSimple(resCode)
	return fmt.Errorf(format, symbol, desc)
}

// AppendVars wraps MSK_appendvars, which adds variables
// to the task.
func (task *Task) AppendVars(num Int32t) ResCode {
	return C.MSK_appendvars(task.task, num)
}

// AppendCons wraps MSK_appendcons, which add linear constraints
// to the task
func (task *Task) AppendCons(numcon Int32t) ResCode {
	return C.MSK_appendcons(task.task, numcon)
}

// PutCj wraps MSK_putcj, which set the coefficient in the objective function.
func (task *Task) PutCj(j Int32t, c_j Realt) ResCode {
	return C.MSK_putcj(task.task, j, c_j)
}

// PutCSlice wraps MSK_putcslice, which set a slice of coefficients in the objective
func (task *Task) PutCSlice(first, last Int32t, slice *Realt) ResCode {
	return C.MSK_putcslice(task.task, first, last, slice)
}

// PutVarbound wraps MSK_putvarbound, which set the bound for a variable.
func (task *Task) PutVarbound(j Int32t, bkx BoundKey, blx, bux Realt) ResCode {
	return C.MSK_putvarbound(task.task, j, bkx, blx, bux)
}

// PutVarboundSliceConst wraps MSK_putvarboundsliceconst, which set the bound for a slice of variables.
func (task *Task) PutVarboundSliceConst(first, last Int32t, bkx BoundKey, blx, bux Realt) ResCode {
	return C.MSK_putvarboundsliceconst(task.task, first, last, bkx, blx, bux)
}

// PutConBound wraps MSK_putconbound, which set the bound for a contraint
func (task *Task) PutConBound(i Int32t, bkc BoundKey, blc, buc Realt) ResCode {
	return C.MSK_putconbound(task.task, i, bkc, blc, buc)
}

// PutObjsense wraps MSK_putobjsense set the objective sense - which is either minimize or maximize
func (task *Task) PutObjsense(sense ObjectiveSense) ResCode {
	return C.MSK_putobjsense(task.task, sense)
}

// PutAij wraps MSK_putaij, which set the value of the linear constraints matrix A[i,j]
func (task *Task) PutAij(i, j Int32t, aij Realt) ResCode {
	return C.MSK_putaij(task.task, i, j, aij)
}

// PutACol wraps MSK_putacol, and puts a column of A matrix.
func (task *Task) PutACol(j Int32t, nzj Int32t, subj *Int32t, valj *Realt) ResCode {
	return C.MSK_putacol(task.task, j, nzj, subj, valj)
}

// AppendAfes wraps MSK_appendafes and adds affine expressions to the task.
func (task *Task) AppendAfes(num Int64t) ResCode {
	return C.MSK_appendafes(task.task, num)
}

// PutAfeFEntry wraps MSK_putafefentry and set an entry in the  affine expression F matrix.
func (task *Task) PutAfeFEntry(afeidx Int64t, varidx Int32t, value Realt) ResCode {
	return C.MSK_putafefentry(task.task, afeidx, varidx, value)
}

// PutAfeFEntryList wraps MSK_putafefentrylist, which set a portion of the affine expression F matrix
func (task *Task) PutAfeFEntryList(numentr Int64t, afeidx *Int64t, varidx *Int32t, val *Realt) ResCode {
	return C.MSK_putafefentrylist(task.task, numentr, afeidx, varidx, val)
}

// PutAfeFRow wraps MSK_putafefrow and sets a row of affine expression F matrix
func (task *Task) PutAfeFRow(afeidx Int64t, numnz Int32t, varidx *Int32t, val *Realt) ResCode {
	return C.MSK_putafefrow(task.task, afeidx, numnz, varidx, val)
}

// PutAfeG wraps MSK_putafeg and sets the value at afeidx to g
func (task *Task) PutAfeG(afeidx Int64t, g Realt) ResCode {
	return C.MSK_putafeg(task.task, afeidx, g)
}

// PutAfeGSlice wraps MSK_putafegslice and sets a slice of values in g
func (task *Task) PutAfeGSlice(first, last Int64t, slice *Realt) ResCode {
	return C.MSK_putafegslice(task.task, first, last, slice)
}

// AppendRZeroDomain wraps MSK_appendrzerodomain and add a real zero domain of dimension n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendRZeroDomain(n Int64t) (r ResCode, domidx Int64t) {
	r = C.MSK_appendrzerodomain(task.task, n, &domidx)
	return
}

// AppendQuadraticConeDomain wraps MSK_appendquadraticconedomain and adds a new quadratic cone of size n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendQuadraticConeDomain(n Int64t) (r ResCode, domidx Int64t) {
	r = C.MSK_appendquadraticconedomain(task.task, n, &domidx)
	return
}

// AppendRotatedQuadraticConeDomain wraps MSK_appendrquadraticconedomain and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendRotatedQuadraticConeDomain(n Int64t) (r ResCode, domidx Int64t) {
	r = C.MSK_appendrquadraticconedomain(task.task, n, &domidx)
	return
}

// AppendRQuadraticConeDomain wraps MSK_appendrquadraticconedomain and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful - this is same as [AppendRotatedQuadraticConeDomain], but with the word "rotated" fully spelled out.
func (task *Task) AppendRQuadraticConeDomain(n Int64t) (r ResCode, domidx Int64t) {
	r = C.MSK_appendrquadraticconedomain(task.task, n, &domidx)
	return
}

// AppendAcc wraps MSK_appendacc and adds an affine conic constraint to the task, where the afe idx is provided
// by an array or pointer - if the afe idx is sequential, use [Task.AppendAccSeq] to avoid allocating an array.
func (task *Task) AppendAcc(domidx, numafeidx Int64t, afeidxlist *Int64t, b *Realt) ResCode {
	return C.MSK_appendacc(task.task, domidx, numafeidx, afeidxlist, b)
}

// AppendAccSeq wraps MSK_appendaccseq and adds an affine conic constraint to the task where
// the affine idx is sequential.
func (task *Task) AppendAccSeq(domidx, numafeidx, afeidxfirst Int64t, b *Realt) ResCode {
	return C.MSK_appendaccseq(task.task, domidx, numafeidx, afeidxfirst, b)
}

// AppendAccsSeq wraps MSK_appendaccsseq and append a block of accs to the tas - assuming affine expressions are sequential.
func (task *Task) AppendAccsSeq(numaccs Int64t, domidxs *Int64t, numafeidx, afeidxfirst Int64t, b *Realt) ResCode {
	return C.MSK_appendaccsseq(task.task, numaccs, domidxs, numafeidx, afeidxfirst, b)
}

// PutVarName wraps MSK_putvarname and sets a name for variable at j.
// Allocate a new C array and copy the data over, then free it - this is a costly function.
func (task *Task) PutVarName(j Int32t, name string) ResCode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return C.MSK_putvarname(task.task, j, cstr)
}

// PutConName wraps MSK_putconname and sets a name for a linear constraint at indext i.
func (task *Task) PutConName(i Int32t, name string) ResCode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return C.MSK_putconname(task.task, i, cstr)
}

// PutAccName wraps MSK_putaccname and sets a name for an affine conic constraint.
func (task *Task) PutAccName(accidx Int64t, name string) ResCode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return C.MSK_putaccname(task.task, accidx, cstr)
}

// OptimizeTerm wraps MSK_optimizeterm, which optimizes the problem.
func (task *Task) OptimizeTerm() (res ResCode, trmcode ResCode) {
	res = C.MSK_optimizetrm(task.task, &trmcode)
	return
}

// GetNumVar wraps MSK_getnumvar, which obtains the number of variables in task.
func (task *Task) GetNumVar() (res ResCode, numVar Int32t) {
	res = C.MSK_getnumvar(task.task, &numVar)
	return
}

// GetSolSta wraps MSK_getsolsta, which returns the solution status
func (task *Task) GetSolSta(whichsol SolType) (res ResCode, solSta SolSta) {
	res = C.MSK_getsolsta(task.task, whichsol, &solSta)
	return
}

// GetXx wraps MSK_getxx, which gets the solution from the task.
// xx can be nil, in which case the number of variables will be queried from the task and
// a new slice created.
func (task *Task) GetXx(whichsol SolType, xx []Realt) (ResCode, []Realt) {
	var res ResCode
	var numVar Int32t
	if xx == nil {
		res, numVar = task.GetNumVar()
		if res != RES_OK {
			return res, nil
		}
		xx = make([]Realt, numVar)
	}

	res = C.MSK_getxx(task.task, whichsol, &xx[0])

	return res, xx
}

// GetXxSlice wraps MSK_getxxslice, which gets a slice of the solution. xx can be nil, in which case the number of variables
// will be last - first and a new slice will be created.
func (task *Task) GetXxSlice(whichsol SolType, first, last Int32t, xx []Realt) (ResCode, []Realt) {
	var res ResCode
	if xx == nil {
		xx = make([]Realt, last-first)
	}

	res = C.MSK_getxxslice(task.task, whichsol, first, last, &xx[0])

	return res, xx
}

// GetAccN wraps MSK_getaccn and returns the dimension of cone at index accidx.
func (task *Task) GetAccN(accidx Int64t) (ResCode, Int64t) {
	var accn Int64t
	res := C.MSK_getaccn(task.task, accidx, &accn)
	return res, accn
}

// GetAccDotY wraps MSK_getaccdoty and returns doty dual result of cone at idnex accidx.
// doty can be nil, in which case the dimension of the cone will be queried from the task and
// a new slice will be created.
func (task *Task) GetAccDotY(whichsol SolType, accidx Int64t, doty []Realt) (ResCode, []Realt) {
	var res ResCode

	if doty == nil {
		var numdoty Int64t
		res, numdoty = task.GetAccN(accidx)
		if res != RES_OK {
			return res, nil
		}
		doty = make([]Realt, numdoty)
	}

	res = C.MSK_getaccdoty(task.task, whichsol, accidx, &doty[0])

	return res, doty
}

// EvaluateAcc gets the activity of the cone at index accidx
func (task *Task) EvaluateAcc(whichsol SolType, accidx Int64t, activity []Realt) (ResCode, []Realt) {
	var res ResCode

	if activity == nil {
		var numdoty Int64t
		res, numdoty = task.GetAccN(accidx)
		if res != RES_OK {
			return res, nil
		}
		activity = make([]Realt, numdoty)
	}

	res = C.MSK_evaluateacc(task.task, whichsol, accidx, &activity[0])

	return res, activity
}

// PutIntParam wraps MSK_putintparam and sets the param to parvalue.
func (task *Task) PutIntParam(param IParam, parvalue Int32t) ResCode {
	return C.MSK_putintparam(task.task, param, parvalue)
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
func (task *Task) LinkFuncToTaskStream(whichstream StreamType, w io.Writer) ResCode {
	writer := writerHolder{
		writer: w,
	}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return C.MSK_linkfunctotaskstream(task.task, whichstream, C.MSKuserhandle_t(ptr), (*[0]byte)(C.writeStreamToWriter))
}

// SolutionSummary wraps MSK_solutionsummary, which prints the summary of the solution to the given stream.
func (task *Task) SolutionSummary(whichstream StreamType) ResCode {
	return C.MSK_solutionsummary(task.task, whichstream)
}

// WriteData wraps MSK_writedata and write data to a file.
func (task *Task) WriteData(filename string) ResCode {
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	return C.MSK_writedata(task.task, filenameC)
}

// ReadData wraps MSK_readdata and read data from a file
func (task *Task) ReadData(filename string) ResCode {
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	return C.MSK_readdata(task.task, filenameC)
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
func (task *Task) WriteDataHandle(handle io.Writer, format DataFormat, compress CompressType) ResCode {
	writer := writerHolder{writer: handle}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return C.MSK_writedatahandle(task.task, (*[0]byte)(C.writeFuncToWriter), C.MSKuserhandle_t(ptr), format, compress)
}
