package gmsk

// #cgo LDFLAGS: -lmosek64
//
// #include <stdlib.h> // stdlib.h is required for calloc and free
// #include <mosek.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// ResCode is the enumerator of MSKrescodee
type ResCode = C.MSKrescodee

// MSK_RES_OK is OK return code from mosek functions
const RES_OK ResCode = C.MSK_RES_OK

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

// BoundKey is MSKboundkey enum, indicate the type of the bound
type BoundKey = C.MSKboundkeye

const (
	BK_LO BoundKey = C.MSK_BK_LO // Lower bound
	BK_UP BoundKey = C.MSK_BK_UP // Upper bound
	BK_FX BoundKey = C.MSK_BK_FX // Fixed bound
	BK_FR BoundKey = C.MSK_BK_FR // Free
	BK_RA BoundKey = C.MSK_BK_RA // Range bound
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
}

// MakeTask is the equivalent of MSK_maketask.
// To use the global environment, use nil for env
func MakeTask(env *Env, maxnumcon, maxnumvar int) (*Task, error) {
	var task C.MSKtask_t = nil
	var e C.MSKenv_t = nil
	if env != nil {
		e = env.env
	}
	r := C.MSK_maketask(e, C.int(maxnumcon), C.int(maxnumvar), &task)
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
func AppendVars(task *Task, num int) ResCode {
	return C.MSK_appendvars(task.task, C.int(num))
}

// PutCj wraps MSK_putcj, which set the coefficient in the objective function.
func PutCj(task *Task, j int, c_j float64) ResCode {
	return C.MSK_putcj(task.task, C.int(j), C.double(c_j))
}

// PutVarbound wraps MSK_putvarbound, which set the bound for a variable.
func PutVarbound(task *Task, j int, boundType BoundKey, l, u float64) ResCode {
	return C.MSK_putvarbound(task.task, C.int(j), boundType, C.double(l), C.double(u))
}

// PutObjsense wraps MSK_putobjsense set the objective sense - which is either minimize or maximize
func PutObjsense(task *Task, sense ObjectiveSense) ResCode {
	return C.MSK_putobjsense(task.task, sense)
}

// OptimizeTerm wraps MSK_optimizeterm, which optimizes the problem.
func OptimizeTerm(task *Task) (res ResCode, trmcode ResCode) {
	res = C.MSK_optimizetrm(task.task, &trmcode)
	return
}

// GetNumVar wraps MSK_getnumvar, which obtains the number of variables in task.
func GetNumVar(task *Task) (res ResCode, numVar int) {
	var result C.MSKint32t
	res = C.MSK_getnumvar(task.task, &result)

	numVar = int(result)

	return
}

// getXx wraps MSK_getxx, which gets the solution from the task.
func getXx(task *Task, solType SolType, result *float64) ResCode {
	return C.MSK_getxx(task.task, solType, (*C.double)(result))
}

// getXx wraps MSK_getxx, which gets the solution from the task.
func GetXx(task *Task, solType SolType, result []float64) (ResCode, []float64) {
	var res ResCode
	var numVar int
	if result == nil {
		res, numVar = GetNumVar(task)
		if res != RES_OK {
			return res, nil
		}
		result = make([]float64, numVar)
	}

	res = getXx(task, solType, &result[0])

	return res, result
}
