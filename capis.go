// gmsk is an unofficial wrapper for MOSEK, the conic optimizer from [MOSEK ApS].
//
// gmsk is based on mosek's C api, which must be installed and configured before
// using gmsk.
//
// Most routines of mosek's C api returns a [response code] [ResCode] (which is an enum/uint32)
// to indicate if the routine is successful or not (MSK_RES_OK or [RES_OK] or 0
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
// The input to val can be `&val[0]`. That being said, gmsk takes care of this, and
// most of the pointer inputs in gmsk are slices.
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

	r := ResCode(C.MSK_makeenv(&env, cdbgfile))
	if r != RES_OK {
		return nil, NewError(r)
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
func DeleteEnv(env *Env) error {
	if env.env != nil {
		return NewError(ResCode(C.MSK_deleteenv(&env.env)))
	}

	return nil
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
	r := ResCode(C.MSK_maketask(e, mi32(maxnumcon), mi32(maxnumvar), &task))
	if r != RES_OK {
		return nil, NewError(r)
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
func fmtError(format string, resCode ResCode) error {
	symbol, desc, _ := GetCodedesc(resCode)
	return fmt.Errorf(format, symbol, desc)
}

// AppendRotatedQuadraticConeDomain wraps MSK_appendrquadraticconedomain
// and adds a new *rotated* quadratic cone of size n to the task.
// returns the index of the domain if successful.
// This is same as [Task.AppendRQuadraticConeDomain],
// but with the word "rotated" fully spelled out.
func (task *Task) AppendRotatedQuadraticConeDomain(n int64) (domidx int64, err error) {
	err = NewErrorFromInt(C.MSK_appendrquadraticconedomain(task.task, mi64(n), pi64(&domidx)))

	return
}

// GetXx wraps MSK_getxx, which gets the solution from the task.
// xx can be nil, in which case the number of variables will be queried from the task and
// a new slice created.
func (task *Task) GetXx(whichsol SolType, xx []float64) ([]float64, error) {
	var err error
	var numVar int32
	if xx == nil {
		numVar, err = task.GetNumVar()
		if err != nil {
			return nil, err
		}
		xx = make([]float64, numVar)
	}

	err = ResCode(C.MSK_getxx(task.task, C.MSKsoltypee(whichsol), prl(&xx[0]))).ToError()

	return xx, err
}

// GetXxSlice wraps MSK_getxxslice, which gets a slice of the solution. xx can be nil, in which case the number of variables
// will be last - first and a new slice will be created.
func (task *Task) GetXxSlice(whichsol SolType, first, last int32, xx []float64) ([]float64, error) {
	var r ResCode
	if xx == nil {
		xx = make([]float64, last-first)
	}

	r = ResCode(
		C.MSK_getxxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			mi32(first),
			mi32(last),
			prl(&xx[0])))

	return xx, r.ToError()
}

// GetBarxj wraps MSK_getbarxj and retrieves the semi definite matrix at j.
func (task *Task) GetBarXj(whichsol SolType, j int32, barxj []float64) ([]float64, error) {
	var err error

	if barxj == nil {
		var lenbarvarj int64
		lenbarvarj, err = task.GetLenBarvarJ(j)
		if err != nil {
			return barxj, err
		}
		barxj = make([]float64, lenbarvarj)
	}

	err = NewErrorFromInt(
		C.MSK_getbarxj(task.task, C.MSKsoltypee(whichsol), mi32(j), prl(&barxj[0])),
	)

	return barxj, err
}

// GetAccN wraps MSK_getaccn and returns the dimension of cone at index accidx.
func (task *Task) GetAccN(accidx int64) (int64, error) {
	var accn int64
	err := NewErrorFromInt(C.MSK_getaccn(task.task, mi64(accidx), pi64(&accn)))
	return accn, err
}

// GetAccDotY wraps MSK_getaccdoty and returns doty dual result of cone at idnex accidx.
// doty can be nil, in which case the dimension of the cone will be queried from the task and
// a new slice will be created.
func (task *Task) GetAccDotY(whichsol SolType, accidx int64, doty []float64) ([]float64, error) {
	var err error

	if doty == nil {
		var numdoty int64
		numdoty, err = task.GetAccN(accidx)
		if err != nil {
			return nil, err
		}
		doty = make([]float64, numdoty)
	}

	err = NewErrorFromInt(C.MSK_getaccdoty(task.task, C.MSKsoltypee(whichsol), mi64(accidx), prl(&doty[0])))

	return doty, err
}

// EvaluateAcc gets the activity of the cone at index accidx
func (task *Task) EvaluateAcc(whichsol SolType, accidx int64, activity []float64) ([]float64, error) {
	var err error

	if activity == nil {
		var numdoty int64
		numdoty, err = task.GetAccN(accidx)
		if err != nil {
			return nil, err
		}
		activity = make([]float64, numdoty)
	}

	err = NewErrorFromInt(C.MSK_evaluateacc(task.task, C.MSKsoltypee(whichsol), mi64(accidx), prl(&activity[0])))

	return activity, err
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

	_, _ = w.writer.Write(unsafe.Slice((*byte)(unsafe.Pointer(v)), n))
}

// LinkFuncToTaskStream wraps MSK_linkfuctotaskstream using [io.Writer] instead of callbacks.
func (task *Task) LinkFuncToTaskStream(whichstream StreamType, w io.Writer) error {
	writer := writerHolder{
		writer: w,
	}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return ResCode(
		C.MSK_linkfunctotaskstream(
			task.task,
			C.MSKstreamtypee(whichstream),
			C.MSKuserhandle_t(ptr), // staticcheck will complain, but this is fine.
			(*[0]byte)(C.writeStreamToWriter))).ToError()
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
func (task *Task) WriteDataHandle(handle io.Writer, format DataFormat, compress CompressType) error {
	writer := writerHolder{writer: handle}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return ResCode(
		C.MSK_writedatahandle(
			task.task,
			(*[0]byte)(C.writeFuncToWriter),
			C.MSKuserhandle_t(ptr), // staticcheck will complain, but this is fine.
			C.MSKdataformate(format),
			C.MSKcompresstypee(compress))).ToError()
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

func getPtrToFirst[T any](v []T) *T {
	if v == nil || len(v) == 0 {
		return nil
	} else {
		return &v[0]
	}
}

// RescodeToStr is wrapping [MSK_rescodetostr]
//
// [MSK_rescodetostr]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.rescodetostr
func RescodeToStr(
	c ResCode,
) (str string, r error) {
	// function template: prepare for output of booleans
	c_str := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_str))

	r = ResCode(
		C.MSK_rescodetostr(
			C.MSKrescodee(c),
			c_str,
		),
	).ToError()

	if r == nil {
		str = C.GoString(c_str)
	}

	return
}

// WriteBSolutionHandle wraps [MSK_writebsolutionhandle) and uses [io.Writer] instead of using callbacks.
//
// [MSK_writebsolutionhandle]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.writebsolution
func (task *Task) WriteBSolutionHandle(handle io.Writer, compress CompressType) error {
	writer := writerHolder{writer: handle}

	ptr := cgo.NewHandle(writer)
	task.writerHandles = append(task.writerHandles, ptr)

	return ResCode(C.MSK_writebsolutionhandle(
		task.task,
		(*[0]byte)(C.writeFuncToWriter),
		C.MSKuserhandle_t(ptr), // staticcheck will complain, but this is fine.
		C.MSKcompresstypee(compress),
	)).ToError()
}
