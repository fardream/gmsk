// gmsk is an unofficial wrapper for MOSEK, the conic optimizer from [MOSEK ApS]
//
// [MOSEK ApS]: https://www.mosek.com
package gmsk

// #cgo LDFLAGS: -lmosek64
//
// #include <stdlib.h> // stdlib.h is required for calloc and free
// #include <mosek.h>
//
// extern void writeStreamToWriter(void*, char*);
import "C"

import (
	"fmt"
	"io"
	"runtime/cgo"
	"unsafe"
)

// INFINITY is MSK_INFINITY (which is different from the double's infinity)
const INFINITY = C.MSK_INFINITY

// Int32t is the int type in MOSEK, which is int32t/int32
type Int32t = C.MSKint32t

// Realt is the double type in MOSEK, which is double/float64
type Realt = C.MSKrealt

// ResCode is the enumerator of MSKrescodee
type ResCode = C.MSKrescodee

const (
	RES_OK        ResCode = C.MSK_RES_OK        // OK return code from mosek functions
	RES_ERR_SPACE ResCode = C.MSK_RES_ERR_SPACE // Err
)

// MSK_MAX_STR_LEN is the max length of strings in mosek
const MAX_STR_LEN = C.MSK_MAX_STR_LEN

// ObjectiveSense is the MSKobjsense type
type ObjectiveSense = C.MSKobjsensee

const (
	OBJECTIVE_SENSE_MINIMIZE ObjectiveSense = C.MSK_OBJECTIVE_SENSE_MINIMIZE // Objective is to maximize
	OBJECTIVE_SENSE_MAXIMIZE ObjectiveSense = C.MSK_OBJECTIVE_SENSE_MAXIMIZE // Objective is to minimize
)

// SolType is the solution type
type SolType = C.MSKsoltypee

const (
	SOL_ITR SolType = C.MSK_SOL_ITR // Iterior Point Solution.
	SOL_BAS SolType = C.MSK_SOL_BAS // Basic Solution.
	SOL_ITG SolType = C.MSK_SOL_ITG // Integer Solution.
)

// SolSta is the solution status
type SolSta = C.MSKsolstae

const (
	SOL_STA_UNKNOWN            SolSta = C.MSK_SOL_STA_UNKNOWN
	SOL_STA_OPTIMAL            SolSta = C.MSK_SOL_STA_OPTIMAL
	SOL_STA_PRIM_FEAS          SolSta = C.MSK_SOL_STA_PRIM_FEAS
	SOL_STA_DUAL_FEAS          SolSta = C.MSK_SOL_STA_DUAL_FEAS
	SOL_STA_PRIM_AND_DUAL_FEAS SolSta = C.MSK_SOL_STA_PRIM_AND_DUAL_FEAS
	SOL_STA_PRIM_INFEAS_CER    SolSta = C.MSK_SOL_STA_PRIM_INFEAS_CER
	SOL_STA_DUAL_INFEAS_CER    SolSta = C.MSK_SOL_STA_DUAL_INFEAS_CER
	SOL_STA_PRIM_ILLPOSED_CER  SolSta = C.MSK_SOL_STA_PRIM_ILLPOSED_CER
	SOL_STA_DUAL_ILLPOSED_CER  SolSta = C.MSK_SOL_STA_DUAL_ILLPOSED_CER
	SOL_STA_INTEGER_OPTIMAL    SolSta = C.MSK_SOL_STA_INTEGER_OPTIMAL
)

// BoundKey is MSKboundkey enum, indicate the type of the bound
type BoundKey = C.MSKboundkeye

const (
	BK_LO BoundKey = C.MSK_BK_LO // Lower bound
	BK_UP BoundKey = C.MSK_BK_UP // Upper bound
	BK_FX BoundKey = C.MSK_BK_FX // Fixed bound
	BK_FR BoundKey = C.MSK_BK_FR // Free
	BK_RA BoundKey = C.MSK_BK_RA // Range bound
)

type StreamType = C.MSKstreamtypee

const (
	STREAM_LOG StreamType = C.MSK_STREAM_LOG
	STREAM_MSG StreamType = C.MSK_STREAM_MSG
	STREAM_ERR StreamType = C.MSK_STREAM_ERR
	STREAM_WRN StreamType = C.MSK_STREAM_WRN
)

// Env wraps mosek environment
type Env struct {
	env C.MSKenv_t
}

// MakeEnv creates a new mosek environment
func MakeEnv() (*Env, error) {
	var env C.MSKenv_t = nil
	r := C.MSK_makeenv(&env, nil)
	if r == RES_OK {
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

// PutVarbound wraps MSK_putvarbound, which set the bound for a variable.
func (task *Task) PutVarbound(j Int32t, bkx BoundKey, blx, bux Realt) ResCode {
	return C.MSK_putvarbound(task.task, j, bkx, blx, bux)
}

// PutConBound wraps MSK_putconbound, which set the bound for a contraint
func (task *Task) PutConBound(i Int32t, bkc BoundKey, blc, buc Realt) ResCode {
	return C.MSK_putconbound(task.task, i, bkc, blc, buc)
}

// PutObjsense wraps MSK_putobjsense set the objective sense - which is either minimize or maximize
func (task *Task) PutObjsense(sense ObjectiveSense) ResCode {
	return C.MSK_putobjsense(task.task, sense)
}

// PutACol wraps MSK_putacol, and puts a column of A matrix.
func (task *Task) PutACol(j Int32t, nzj Int32t, subj *Int32t, valj *Realt) ResCode {
	return C.MSK_putacol(task.task, j, nzj, subj, valj)
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

// getXx wraps MSK_getxx, which gets the solution from the task.
func (task *Task) getXx(solType SolType, result *Realt) ResCode {
	return C.MSK_getxx(task.task, solType, result)
}

// getXx wraps MSK_getxx, which gets the solution from the task.
func (task *Task) GetXx(solType SolType, result []Realt) (ResCode, []Realt) {
	var res ResCode
	var numVar Int32t
	if result == nil {
		res, numVar = task.GetNumVar()
		if res != RES_OK {
			return res, nil
		}
		result = make([]Realt, numVar)
	}

	res = task.getXx(solType, &result[0])

	return res, result
}

// writerHolder holds a writer. This must be passed to C api with a handle.
type writerHolder struct {
	writer io.Writer
}

// writeStreamToWriter is the function for C api's task stream handler,
// and it has a signature of
//
//	void streamfunc(void* handle, char * data)
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

// LinkFuncToTaskStream wraps MSK_linkfuctotaskstream. Instead of using call back function,
// pass in a writer that will take the stream of tasks.
func (task *Task) LinkFuncToTaskStream(whichstream StreamType, w io.Writer) ResCode {
	writer := writerHolder{
		writer: w,
	}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return C.MSK_linkfunctotaskstream(task.task, whichstream, C.MSKuserhandle_t(ptr), (*[0]byte)(C.writeStreamToWriter))
}

func (task *Task) SolutionSummary(whichstream StreamType) ResCode {
	return C.MSK_solutionsummary(task.task, whichstream)
}
