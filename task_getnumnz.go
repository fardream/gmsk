// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"github.com/fardream/gmsk/res"
)

// GetAccfNumNz is wrapping [MSK_getaccfnumnz]
//
// [MSK_getaccfnumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - accfnnz: MSKint64t *
//
// [MSK_getaccfnumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAccfNumNz() (r res.Code, accfnnz int64) {
	r = res.Code(
		C.MSK_getaccfnumnz(
			task.task,
			(*C.MSKint64t)(&accfnnz),
		),
	)

	return
}

// GetAColNumNz is wrapping [MSK_getacolnumnz]
//
// [MSK_getacolnumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - nzj: MSKint32t *
//
// [MSK_getacolnumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAColNumNz(
	i int32,
) (r res.Code, nzj int32) {
	r = res.Code(
		C.MSK_getacolnumnz(
			task.task,
			C.MSKint32t(i),
			(*C.MSKint32t)(&nzj),
		),
	)

	return
}

// GetAColSliceNumNz is wrapping [MSK_getacolslicenumnz]
//
// [MSK_getacolslicenumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - first: MSKint32t
//   - last: MSKint32t
//   - numnz: MSKint32t *
//
// [MSK_getacolslicenumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAColSliceNumNz(
	first int32,
	last int32,
) (r res.Code, numnz int32) {
	r = res.Code(
		C.MSK_getacolslicenumnz(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKint32t)(&numnz),
		),
	)

	return
}

// GetAColSliceNumNz64 is wrapping [MSK_getacolslicenumnz64]
//
// [MSK_getacolslicenumnz64] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - first: MSKint32t
//   - last: MSKint32t
//   - numnz: MSKint64t *
//
// [MSK_getacolslicenumnz64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAColSliceNumNz64(
	first int32,
	last int32,
) (r res.Code, numnz int64) {
	r = res.Code(
		C.MSK_getacolslicenumnz64(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKint64t)(&numnz),
		),
	)

	return
}

// GetAfeFNumNz is wrapping [MSK_getafefnumnz]
//
// [MSK_getafefnumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - numnz: MSKint64t *
//
// [MSK_getafefnumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAfeFNumNz() (r res.Code, numnz int64) {
	r = res.Code(
		C.MSK_getafefnumnz(
			task.task,
			(*C.MSKint64t)(&numnz),
		),
	)

	return
}

// GetAfeFRowNumNz is wrapping [MSK_getafefrownumnz]
//
// [MSK_getafefrownumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - afeidx: MSKint64t
//   - numnz: MSKint32t *
//
// [MSK_getafefrownumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAfeFRowNumNz(
	afeidx int64,
) (r res.Code, numnz int32) {
	r = res.Code(
		C.MSK_getafefrownumnz(
			task.task,
			C.MSKint64t(afeidx),
			(*C.MSKint32t)(&numnz),
		),
	)

	return
}

// GetAPieceNumNz is wrapping [MSK_getapiecenumnz]
//
// [MSK_getapiecenumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - firsti: MSKint32t
//   - lasti: MSKint32t
//   - firstj: MSKint32t
//   - lastj: MSKint32t
//   - numnz: MSKint32t *
//
// [MSK_getapiecenumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetAPieceNumNz(
	firsti int32,
	lasti int32,
	firstj int32,
	lastj int32,
) (r res.Code, numnz int32) {
	r = res.Code(
		C.MSK_getapiecenumnz(
			task.task,
			C.MSKint32t(firsti),
			C.MSKint32t(lasti),
			C.MSKint32t(firstj),
			C.MSKint32t(lastj),
			(*C.MSKint32t)(&numnz),
		),
	)

	return
}

// GetARowNumNz is wrapping [MSK_getarownumnz]
//
// [MSK_getarownumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - nzi: MSKint32t *
//
// [MSK_getarownumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetARowNumNz(
	i int32,
) (r res.Code, nzi int32) {
	r = res.Code(
		C.MSK_getarownumnz(
			task.task,
			C.MSKint32t(i),
			(*C.MSKint32t)(&nzi),
		),
	)

	return
}

// GetARowSliceNumNz is wrapping [MSK_getarowslicenumnz]
//
// [MSK_getarowslicenumnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - first: MSKint32t
//   - last: MSKint32t
//   - numnz: MSKint32t *
//
// [MSK_getarowslicenumnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetARowSliceNumNz(
	first int32,
	last int32,
) (r res.Code, numnz int32) {
	r = res.Code(
		C.MSK_getarowslicenumnz(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKint32t)(&numnz),
		),
	)

	return
}

// GetARowSliceNumNz64 is wrapping [MSK_getarowslicenumnz64]
//
// [MSK_getarowslicenumnz64] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - first: MSKint32t
//   - last: MSKint32t
//   - numnz: MSKint64t *
//
// [MSK_getarowslicenumnz64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func (task *Task) GetARowSliceNumNz64(
	first int32,
	last int32,
) (r res.Code, numnz int64) {
	r = res.Code(
		C.MSK_getarowslicenumnz64(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKint64t)(&numnz),
		),
	)

	return
}
