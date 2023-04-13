// gmsk is an unofficial wrapper for MOSEK, the conic optimizer from [MOSEK ApS].
//
// gmsk is based on mosek's C api, which must be installed and configured before
// using gmsk.
//
// Most routines of mosek's C api returns a [response code] [ResCode] (whic is an enum/uint32)
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
// Besides [response code], MOSEK sometimes needs to return some other information
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
// [response code]: https://docs.mosek.com/latest/capi/response-codes.html
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

// fmtError formats the error string
func fmtError(format string, resCode res.Code) error {
	_, symbol, desc := GetCodedesc(resCode)
	return fmt.Errorf(format, symbol, desc)
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

// GetBarxj wraps MSK_getbarxj and retrieves the semi definite matrix at j.
func (task *Task) GetBarXj(whichsol SolType, j int32, barxj []float64) (res.Code, []float64) {
	var r res.Code

	if barxj == nil {
		var lenbarvarj int64
		r, lenbarvarj = task.GetLenBarvarJ(j)
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

// intToBool converts an integer to a bool.
// c doesn't have a bool type and MOSEK uses int32 as bool.
func intToBool(i C.MSKbooleant) bool {
	return i != 0
}

func boolToInt(i bool) C.MSKbooleant {
	if i {
		return 1
	} else {
		return 0
	}
}
