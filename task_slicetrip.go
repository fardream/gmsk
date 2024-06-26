// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

// GetAColSliceTrip is wrapping [MSK_getacolslicetrip],
// Obtains a sequence of columns from the coefficient matrix in triplet format.
//
// Arguments:
//
//   - `first` Index of the first column in the sequence.
//   - `last` Index of the last column in the sequence plus one.
//   - `subi` Constraint subscripts.
//   - `subj` Column subscripts.
//   - `val` Values.
//
// [MSK_getacolslicetrip]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getacolslicetrip
func (task *Task) GetAColSliceTrip(
	first int32,
	last int32,
	maxnumnz int64,
	subi []int32,
	subj []int32,
	val []float64,
) error {
	return ResCode(
		C.MSK_getacolslicetrip(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			C.MSKint64t(maxnumnz),
			(*C.MSKint32t)(getPtrToFirst(subi)),
			(*C.MSKint32t)(getPtrToFirst(subj)),
			(*C.MSKrealt)(getPtrToFirst(val)),
		),
	).ToError()
}

// GetARowSliceTrip is wrapping [MSK_getarowslicetrip],
// Obtains a sequence of rows from the coefficient matrix in sparse triplet format.
//
// Arguments:
//
//   - `first` Index of the first row in the sequence.
//   - `last` Index of the last row in the sequence plus one.
//   - `subi` Constraint subscripts.
//   - `subj` Column subscripts.
//   - `val` Values.
//
// [MSK_getarowslicetrip]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getarowslicetrip
func (task *Task) GetARowSliceTrip(
	first int32,
	last int32,
	maxnumnz int64,
	subi []int32,
	subj []int32,
	val []float64,
) error {
	return ResCode(
		C.MSK_getarowslicetrip(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			C.MSKint64t(maxnumnz),
			(*C.MSKint32t)(getPtrToFirst(subi)),
			(*C.MSKint32t)(getPtrToFirst(subj)),
			(*C.MSKrealt)(getPtrToFirst(val)),
		),
	).ToError()
}
