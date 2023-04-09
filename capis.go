// gmsk is an unofficial wrapper for MOSEK, the conic optimizer from [MOSEK ApS].
//
// gmsk is based on mosek's C api, which must be installed and configured before
// using gmsk.
//
// Most routines of mosek's C api returns a result code [ResCode] (whic is an enum/uint32)
// to indicate if the routine is successful or not (MSK_RES_OK or [RES_OK]
// indicates success). Check mosek documentation for more information.
//
// Almost all float point numbers in MOSEK are IEEE-754 64-bit float point number,
// or double in C/C++/float64 in go.
//
// Most indices associated with variables and A (constraints) matrix are
// signed 32-bit integers, int32_t in C/C++ and int32 in go. However, indices
// associated with affine expression (afe) are 64-bit, or int64_t in C/C++
// and int64 in go.
//
// Besides result code, MOSEK sometimes needs to return some other information
// (for example, index of newly created domain). This is achieved in C/C++ by
// passing in a "destination" pointer. This is not done in go because go supports
// multiple return values.
//
// For majority of pointer inputs to MOSEK calls (such as below),
// the pointers can be obtained by take the address of the desired element in a
// slice/array in golang
//
//	MSKrescodee MSK_putafefcol(
//	 task MSKtask_t,
//	 MSKint32t varidx,
//	 MSKint64t numz,
//	 MSKint64t *afeidx,
//	 MSKrealt *val);
//
// The input to val can be `&val[0]`.
//
// Wrappers for routines accepting C string (or "const char *" type)
// instead accept go's string. However, this will create a new C.CString
// (destroy it once it is copied over by mosek).
//
// Note that there is always overhead of calling c functions from go.
//
// [MOSEK ApS]: https://www.mosek.com
package gmsk

// #cgo LDFLAGS: -lmosek64
//
// #include <stdlib.h> // required for calloc and free
// #include <string.h> // required for strnlen
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

	"github.com/fardream/gmsk/res"
)

// INFINITY is MSK_INFINITY (which is different from the double's infinity)
const INFINITY float64 = C.MSK_INFINITY

// MSK_MAX_STR_LEN is the max length of strings in mosek
const MAX_STR_LEN = C.MSK_MAX_STR_LEN

// some local types to shorten the types from MOSEK.
type (
	mi32 = C.MSKint32t
	mi64 = C.MSKint64t
	pi32 = *C.MSKint32t
	pi64 = *C.MSKint64t
	mrl  = C.MSKrealt
	prl  = *C.MSKrealt
	mbk  = C.MSKboundkeye
	pbk  = *C.MSKboundkeye
)

// Env wraps mosek environment
type Env struct {
	env C.MSKenv_t
}

// MakeEnv creates a new mosek environment.
// dbgfile can be zero or one, error when more than two dbgfile are provided.
//
// Remember the env needs to be deleted.
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

	r := res.Code(C.MSK_makeenv(&env, cdbgfile))
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

// MakeTask creates a new task in this environment.
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
	r := res.Code(C.MSK_maketask(e, mi32(maxnumcon), mi32(maxnumvar), &task))
	if r != RES_OK {
		return nil, fmtError("failed to create task: %s %s", r)
	}

	return &Task{task: task}, nil
}

// DeleteTask deletes the task
func DeleteTask(task *Task) {
	if task == nil {
		return
	}
	if task.task != nil {
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
func GetCodeDesc(resCode res.Code, sym []byte, desc []byte) (res.Code, bool) {
	if len(sym) <= MAX_STR_LEN || len(desc) <= MAX_STR_LEN {
		return 0, false
	}

	symstart := unsafe.Pointer(&sym[0])
	descstart := unsafe.Pointer(&desc[0])

	r := C.MSK_getcodedesc(C.MSKrescodee(resCode), (*C.char)(symstart), (*C.char)(descstart))

	return res.Code(r), true
}

// GetCodeDescSimple gets the description of the return code.
// Please check [res.Code] to see if the operation is successful.
// The first returned string is the symbol,
// and the second returned string is the description.
// This is different from [GetCodeDesc] and involves allocating and freeing two char array of
// [MAX_STR_LEN] + 1 size.
func GetCodeDescSimple(resCode res.Code) (res.Code, string, string) {
	// use calloc, which will zero out the memory location.
	symbol := (*C.char)(C.calloc(MAX_STR_LEN+1, C.sizeof_char))
	defer C.free(unsafe.Pointer(symbol))

	des := (*C.char)(C.calloc(MAX_STR_LEN+1, C.sizeof_char))
	defer C.free(unsafe.Pointer(des))

	r := C.MSK_getcodedesc(C.MSKrescodee(resCode), symbol, des)

	return res.Code(r), C.GoString(symbol), C.GoString(symbol)
}

// fmtError formats the error string
func fmtError(format string, resCode res.Code) error {
	_, symbol, desc := GetCodeDescSimple(resCode)
	return fmt.Errorf(format, symbol, desc)
}

// AppendBarVars adds semidefinite matrix variables to the task.
// Barvar because MOSEK uses bar{x} notation to indicate an element
// of a semidefinite matrix.
// The dimension of each of the semidefinite variables are provided
// through the pointer dim.
func (task *Task) AppendBarVars(num int32, dim *int32) res.Code {
	return res.Code(
		C.MSK_appendbarvars(
			task.task,
			mi32(num),
			pi32(dim)))
}

// PutCFix wraps MSK_putcfix and adds a constant term to the objective.
func (task *Task) PutCFix(cfix float64) res.Code {
	return res.Code(C.MSK_putcfix(task.task, mrl(cfix)))
}

// PutBarCj wraps MSK_putbarcj and adds a positive semidefinite matrix to the objective.
// j is the index of the matrix variable, and num is number of weight matrices. Those
// weight matrices are identified by sub (the idx when they are added to the task
// by [Task.AppendSparseSymmat].
func (task *Task) PutBarCj(j int32, num int64, sub *int64, weights *float64) res.Code {
	return res.Code(
		C.MSK_putbarcj(
			task.task,
			mi32(j),
			mi64(num),
			pi64(sub),
			prl(weights),
		),
	)
}

// PutBarCBlockTriplet wraps MSK_putbarcblocktriplet and set the coefficients for
// matrix variable. suj is the index of the matrix variable, subk/subl are row/colmn index
// of the cofficients, and valjkl is the coefficients value.
func (task *Task) PutBarCBlockTriplet(num int64, subj, subk, subl *int32, valjkl *float64) res.Code {
	return res.Code(
		C.MSK_putbarcblocktriplet(
			task.task,
			mi64(num),
			pi32(subj),
			pi32(subk),
			pi32(subl),
			prl(valjkl),
		),
	)
}

// PutQObj wraps MSK_putqobj and set the cofficient for quadratic objective term.
func (task *Task) PutQObj(numqonz int32, qosubi, qosubj *int32, qoval *float64) res.Code {
	return res.Code(
		C.MSK_putqobj(
			task.task,
			mi32(numqonz),
			pi32(qosubi),
			pi32(qosubj),
			prl(qoval),
		),
	)
}

// PutVarType wraps MSK_putvartype and sets the type of the variable
func (task *Task) PutVarType(j int32, vartype VariableType) res.Code {
	return res.Code(
		C.MSK_putvartype(
			task.task,
			mi32(j),
			C.MSKvariabletypee(vartype)))
}

// PutVarTypeList wraps MSK_putvartypelist and sets the type of a list of variables
func (task *Task) PutVarTypeList(num int32, subj *int32, vartype *VariableType) res.Code {
	return res.Code(
		C.MSK_putvartypelist(
			task.task,
			mi32(num),
			pi32(subj),
			(*C.MSKvariabletypee)(vartype),
		),
	)
}

// PutACol wraps MSK_putacol, and puts a column of A matrix.
func (task *Task) PutACol(j int32, nzj int32, subj *int32, valj *float64) res.Code {
	return res.Code(C.MSK_putacol(task.task, mi32(j), mi32(nzj), pi32(subj), prl(valj)))
}

// PutARow wraps MSK_putarow and sets a row of the A matrix.
func (task *Task) PutARow(i int32, nzi int32, subi *int32, vali *float64) res.Code {
	return res.Code(
		C.MSK_putarow(
			task.task,
			mi32(i),
			mi32(nzi),
			pi32(subi),
			prl(vali),
		),
	)
}

// PutBarAij wraps MSK_putbaraij and a semidefinite matrix to constraint.
// i is the index of the constraint,and j is the index of the semidefinite matrix variable.
// num is the number of coefficients matrices, and sub/weights are the coefficient matrices'
// idx (when they are added by [Task.AppendSparseSymmat]) and weights.
func (task *Task) PutBarAij(i, j int32, num int64, sub *int64, weights *float64) res.Code {
	return res.Code(
		C.MSK_putbaraij(
			task.task,
			mi32(i),
			mi32(j),
			mi64(num),
			pi64(sub),
			prl(weights)),
	)
}

// PutBarABlockTriplet wraps MSK_putbarablocktriplet and sets constraints for matrix variable.
// subi is the index of the constraint, subj is the index of the matrix variable,
// subk and subl are the indices of the coefficients and valijkl are the coefficients value.
func (task *Task) PutBarABlockTriplet(num int64, subi, subj, subk, subl *int32, valijkl *float64) res.Code {
	return res.Code(
		C.MSK_putbarablocktriplet(
			task.task,
			mi64(num),
			pi32(subi),
			pi32(subj),
			pi32(subk),
			pi32(subl),
			prl(valijkl),
		),
	)
}

// PutQConK wraps MSK_putqconk and sets the quandratic constraint's coefficients.
func (task *Task) PutQConK(k int32, numqcnz int32, qcsubi, qcsubj *int32, qcval *float64) res.Code {
	return res.Code(
		C.MSK_putqconk(
			task.task,
			mi32(k),
			mi32(numqcnz),
			pi32(qcsubi),
			pi32(qcsubj),
			prl(qcval),
		),
	)
}

// PutAfeFEntry wraps MSK_putafefentry and set an entry in the  affine expression F matrix.
func (task *Task) PutAfeFEntry(afeidx int64, varidx int32, value float64) res.Code {
	return res.Code(C.MSK_putafefentry(task.task, mi64(afeidx), mi32(varidx), mrl(value)))
}

// PutAfeFEntryList wraps MSK_putafefentrylist, which set a portion of the affine expression F matrix
func (task *Task) PutAfeFEntryList(numentr int64, afeidx *int64, varidx *int32, val *float64) res.Code {
	return res.Code(
		C.MSK_putafefentrylist(
			task.task,
			mi64(numentr),
			pi64(afeidx),
			pi32(varidx),
			prl(val),
		),
	)
}

// PutAfeFRow wraps MSK_putafefrow and sets a row of affine expression F matrix
func (task *Task) PutAfeFRow(afeidx int64, numnz int32, varidx *int32, val *float64) res.Code {
	return res.Code(
		C.MSK_putafefrow(
			task.task,
			mi64(afeidx),
			mi32(numnz),
			pi32(varidx),
			prl(val),
		),
	)
}

// PutAfeFCol wraps MSK_putafefcol and sets a column of affine expression F matrix
func (task *Task) PutAfeFCol(varidx int32, numnz int64, afeidx *int64, val *float64) res.Code {
	return res.Code(
		C.MSK_putafefcol(
			task.task,
			mi32(varidx),
			mi64(numnz),
			pi64(afeidx),
			prl(val),
		),
	)
}

// PutAfeBarFBlockTriplet wraps MSK_putafebarfblocktriplet and sets a matrix variable to the F matrix of affine expression.
// the rows of afe idx is provided in afeidx, the matrix variables are indexed by barvaridx, and subk, subl are the indices
// of the coefficients for those matrix variables, and valkl are the corresponding coefficients.
func (task *Task) PutAfeBarFBlockTriplet(numtrip int64, afeidx *int64, barvaridx, subk, subl *int32, valkl *float64) res.Code {
	return res.Code(
		C.MSK_putafebarfblocktriplet(
			task.task,
			mi64(numtrip),
			pi64(afeidx),
			pi32(barvaridx),
			pi32(subk),
			pi32(subl),
			prl(valkl),
		),
	)
}

// PutAfeG wraps MSK_putafeg and sets the value at afeidx to g
func (task *Task) PutAfeG(afeidx int64, g float64) res.Code {
	return res.Code(C.MSK_putafeg(task.task, mi64(afeidx), mrl(g)))
}

// PutAfeGSlice wraps MSK_putafegslice and sets a slice of values in g
func (task *Task) PutAfeGSlice(first, last int64, slice *float64) res.Code {
	return res.Code(
		C.MSK_putafegslice(
			task.task,
			mi64(first),
			mi64(last),
			prl(slice),
		),
	)
}

// AppendRZeroDomain wraps MSK_appendrzerodomain and add a real zero domain of dimension n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendRZeroDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendrzerodomain(
			task.task,
			mi64(n),
			pi64(&domidx),
		),
	)
	return
}

// AppendRPlusDomain wraps MSK_appendrplusdomain and adds a R-Plus domain to the task, which is x >= 0.
func (task *Task) AppendRPlusDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendrplusdomain(
			task.task,
			mi64(n),
			pi64(&domidx),
		),
	)
	return
}

// AppendRMinusDomain wraps MSK_appendrminusdomain and adds a R-minus domain to the task, which is x <= 0.
func (task *Task) AppendRMinusDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendrminusdomain(
			task.task,
			mi64(n),
			pi64(&domidx),
		),
	)

	return
}

// AppendRDomain wraps MSK_appendrdomain and adds the whole n dimension R space domain to the task.
func (task *Task) AppendRDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendrdomain(
			task.task,
			mi64(n),
			pi64(&domidx),
		),
	)

	return
}

// AppendQuadraticConeDomain wraps MSK_appendquadraticconedomain and adds a new quadratic cone of size n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendQuadraticConeDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(C.MSK_appendquadraticconedomain(task.task, mi64(n), pi64(&domidx)))
	return
}

// AppendRotatedQuadraticConeDomain wraps MSK_appendrquadraticconedomain
// and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful.
// This is same as [Task.AppendRQuadraticConeDomain],
// but with the word "rotated" fully spelled out.
func (task *Task) AppendRotatedQuadraticConeDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(C.MSK_appendrquadraticconedomain(task.task, mi64(n), pi64(&domidx)))
	return
}

// AppendRQuadraticConeDomain wraps MSK_appendrquadraticconedomain and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful.
func (task *Task) AppendRQuadraticConeDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(C.MSK_appendrquadraticconedomain(task.task, mi64(n), pi64(&domidx)))
	return
}

// AppendPrimalPowerConeDomain wraps MSK_appendprimalpowerconedomain and add a primal power cone to the task
func (task *Task) AppendPrimalPowerConeDomain(n, nleft int64, alpha *float64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendprimalpowerconedomain(
			task.task,
			mi64(n),
			mi64(nleft),
			prl(alpha),
			pi64(&domidx)))

	return
}

// AppendDualPowerConeDomain wraps MSK_appenddualpowerconedomain and adds a dual power cone to the task.
func (task *Task) AppendDualPowerConeDomain(n, nleft int64, alpha *float64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appenddualpowerconedomain(
			task.task,
			mi64(n),
			mi64(nleft),
			prl(alpha),
			pi64(&domidx),
		),
	)

	return
}

// AppendPrimalExpConeDomain wraps MSK_appendprimalexpconedomain and appends a primal exponential cone to the task.
func (task *Task) AppendPrimalExpConeDomain() (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendprimalexpconedomain(
			task.task,
			pi64(&domidx),
		),
	)
	return
}

// AppendDualExpConeDomain wraps MSK_appenddualexpconedomain and appends a dual exponential cone to the task.
func (task *Task) AppendDualExpConeDomain() (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appenddualexpconedomain(
			task.task,
			pi64(&domidx),
		),
	)

	return
}

// AppendPrimalGeoMeanConeDomain wraps MSK_appendprimalgeomeanconedomain and adds a geometric mean cone domain to the task.
func (task *Task) AppendPrimalGeoMeanConeDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendprimalgeomeanconedomain(
			task.task,
			mi64(n),
			pi64(&domidx),
		),
	)

	return
}

// AppendDualGeoMeanConeDomain wraps MSK_appenddualgeomeanconedomain and adds a geometric mean cone domain to the task.
func (task *Task) AppendDualGeoMeanConeDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appenddualgeomeanconedomain(
			task.task,
			mi64(n),
			pi64(&domidx),
		),
	)

	return
}

// AppendSvecPsdConeDomain wraps MSK_appendsvecpsdconedomain and adds a SVEC_PSD domain,
// or vectorized postive semidefinite matrix. n must be k(k+1)/2.
func (task *Task) AppendSvecPsdConeDomain(n int64) (r res.Code, domidx int64) {
	r = res.Code(
		C.MSK_appendsvecpsdconedomain(task.task, mi64(n), pi64(&domidx)),
	)
	return
}

// AppendSparseSymmat wraps MSK_appendsparsesymmat and adds a sparse and symmetric matrix to the task.
// matrix is represented in coordinate format, and only lower triangular portion of the matrix should be
// specified.
// Those matrices can be used as either coefficent in the objective or constraints. The matrix is identified
// by the returned idx.
func (task *Task) AppendSparseSymmat(dim int32, nz int64, subi, subj *int32, valij *float64) (r res.Code, idx int64) {
	r = res.Code(
		C.MSK_appendsparsesymmat(
			task.task,
			mi32(dim),
			mi64(nz),
			pi32(subi),
			pi32(subj),
			prl(valij),
			pi64(&idx),
		),
	)

	return
}

// GetSolSta wraps MSK_getsolsta, which returns the solution status
func (task *Task) GetSolSta(whichsol SolType) (r res.Code, solSta SolSta) {
	r = res.Code(C.MSK_getsolsta(task.task, C.MSKsoltypee(whichsol), (*C.MSKsolstae)(&solSta)))
	return
}

// GetProSta wraps MSK_getprosta and gets the problem status
func (task *Task) GetProSta(whichsol SolType) (r res.Code, problemsta ProSta) {
	r = res.Code(C.MSK_getprosta(task.task, C.MSKsoltypee(whichsol), (*C.MSKprostae)(&problemsta)))
	return
}

// GetXx wraps MSK_getxx, which gets the solution from the task.
// xx can be nil, in which case the number of variables will be queried from the task and
// a new slice created.
func (task *Task) GetXx(whichsol SolType, xx []float64) (res.Code, []float64) {
	var r res.Code
	var numVar int32
	if xx == nil {
		r, numVar = task.GetNumVar()
		if r != RES_OK {
			return r, nil
		}
		xx = make([]float64, numVar)
	}

	r = res.Code(C.MSK_getxx(task.task, C.MSKsoltypee(whichsol), prl(&xx[0])))

	return r, xx
}

// GetXxSlice wraps MSK_getxxslice, which gets a slice of the solution. xx can be nil, in which case the number of variables
// will be last - first and a new slice will be created.
func (task *Task) GetXxSlice(whichsol SolType, first, last int32, xx []float64) (res.Code, []float64) {
	var r res.Code
	if xx == nil {
		xx = make([]float64, last-first)
	}

	r = res.Code(
		C.MSK_getxxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			mi32(first),
			mi32(last),
			prl(&xx[0])))

	return r, xx
}

// GetLenBarVarJ wraps MSK_getlenbarvarj and returns the length of semidefinite matrix variable's length at j
func (task *Task) GetLenBarVarJ(j int32) (r res.Code, lengthbarvarj int64) {
	r = res.Code(
		C.MSK_getlenbarvarj(
			task.task,
			mi32(j),
			pi64(&lengthbarvarj)),
	)
	return
}

// GetBarxj wraps MSK_getbarxj and retrieves the semi definite matrix at j.
func (task *Task) GetBarXj(whichsol SolType, j int32, barxj []float64) (res.Code, []float64) {
	var r res.Code

	if barxj == nil {
		var lenbarvarj int64
		r, lenbarvarj = task.GetLenBarVarJ(j)
		if r != RES_OK {
			return r, barxj
		}
		barxj = make([]float64, lenbarvarj)
	}

	r = res.Code(
		C.MSK_getbarxj(task.task, C.MSKsoltypee(whichsol), mi32(j), prl(&barxj[0])),
	)

	return r, barxj
}

// GetAccN wraps MSK_getaccn and returns the dimension of cone at index accidx.
func (task *Task) GetAccN(accidx int64) (res.Code, int64) {
	var accn int64
	res := res.Code(C.MSK_getaccn(task.task, mi64(accidx), pi64(&accn)))
	return res, accn
}

// GetAccDotY wraps MSK_getaccdoty and returns doty dual result of cone at idnex accidx.
// doty can be nil, in which case the dimension of the cone will be queried from the task and
// a new slice will be created.
func (task *Task) GetAccDotY(whichsol SolType, accidx int64, doty []float64) (res.Code, []float64) {
	var r res.Code

	if doty == nil {
		var numdoty int64
		r, numdoty = task.GetAccN(accidx)
		if r != RES_OK {
			return r, nil
		}
		doty = make([]float64, numdoty)
	}

	r = res.Code(C.MSK_getaccdoty(task.task, C.MSKsoltypee(whichsol), mi64(accidx), prl(&doty[0])))

	return r, doty
}

// EvaluateAcc gets the activity of the cone at index accidx
func (task *Task) EvaluateAcc(whichsol SolType, accidx int64, activity []float64) (res.Code, []float64) {
	var r res.Code

	if activity == nil {
		var numdoty int64
		r, numdoty = task.GetAccN(accidx)
		if r != RES_OK {
			return r, nil
		}
		activity = make([]float64, numdoty)
	}

	r = res.Code(C.MSK_evaluateacc(task.task, C.MSKsoltypee(whichsol), mi64(accidx), prl(&activity[0])))

	return r, activity
}

// GetIntInf wraps MSK_getintinf and retrieve integer information from the task.
func (task *Task) GetIntInf(whichiinf IInfItem) (r res.Code, ivalue int32) {
	r = res.Code(
		C.MSK_getintinf(task.task, C.MSKiinfiteme(whichiinf), pi32(&ivalue)),
	)
	return
}

// GetDouInf wraps MSK_getdouinf and retrieves double information from the task.
func (task *Task) GetDouInf(whichdinf DInfItem) (r res.Code, dvalue float64) {
	r = res.Code(
		C.MSK_getdouinf(
			task.task,
			C.MSKdinfiteme(whichdinf),
			prl(&dvalue),
		),
	)
	return
}

// PutIntParam wraps MSK_putintparam and sets the param to parvalue.
func (task *Task) PutIntParam(param IParam, parvalue int32) res.Code {
	return res.Code(C.MSK_putintparam(task.task, C.MSKiparame(param), mi32(parvalue)))
}

// PutDouParam wraps MSK_putdouparam and sets the param to the parvalue
func (task *Task) PutDouParam(param DParam, parvalue float64) res.Code {
	return res.Code(C.MSK_putdouparam(task.task, C.MSKdparame(param), mrl(parvalue)))
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

	const MAXLEN = 16777216 // or 16MB
	n := C.strnlen(v, MAXLEN)

	w.writer.Write(unsafe.Slice((*byte)(unsafe.Pointer(v)), n))
}

// LinkFuncToTaskStream wraps MSK_linkfuctotaskstream using [io.Writer] instead of callbacks.
func (task *Task) LinkFuncToTaskStream(whichstream StreamType, w io.Writer) res.Code {
	writer := writerHolder{
		writer: w,
	}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return res.Code(
		C.MSK_linkfunctotaskstream(
			task.task,
			C.MSKstreamtypee(whichstream),
			C.MSKuserhandle_t(ptr), // staticcheck will complain, but this is fine.
			(*[0]byte)(C.writeStreamToWriter)))
}

// SolutionSummary wraps MSK_solutionsummary, which prints the summary of the solution to the given stream.
func (task *Task) SolutionSummary(whichstream StreamType) res.Code {
	return res.Code(C.MSK_solutionsummary(task.task, C.MSKstreamtypee(whichstream)))
}

// ReadData wraps MSK_readdata and read data from a file
func (task *Task) ReadData(filename string) res.Code {
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	return res.Code(C.MSK_readdata(task.task, filenameC))
}

// writeFuncToWriter is the function for C api's MSKhwritefunc, which has a signature of
//
//	size_t  (MSKAPI * MSKhwritefunc) (
//	  MSKuserhandle_t handle,
//	  const void * src,
//	  const size_t count)
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

	byteSlice := unsafe.Slice((*byte)(src), count)

	n, _ := w.writer.Write(byteSlice)

	return C.size_t(n)
}

// WriteDataHandle wraps MSK_writedatahandle using [io.Writer] instead of using callbacks.
func (task *Task) WriteDataHandle(handle io.Writer, format DataFormat, compress CompressType) res.Code {
	writer := writerHolder{writer: handle}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return res.Code(
		C.MSK_writedatahandle(
			task.task,
			(*[0]byte)(C.writeFuncToWriter),
			C.MSKuserhandle_t(ptr), // staticcheck will complain, but this is fine.
			C.MSKdataformate(format),
			C.MSKcompresstypee(compress)))
}
