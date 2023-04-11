// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"github.com/fardream/gmsk/res"
)

// PutMaxNumAcc is wrapping [MSK_putmaxnumacc]
//
// [MSK_putmaxnumacc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumacc: MSKint64t
//
// [MSK_putmaxnumacc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumacc
func (task *Task) PutMaxNumAcc(
	maxnumacc int64,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumacc(
			task.task,
			C.MSKint64t(maxnumacc),
		),
	)
}

// PutMaxNumAfe is wrapping [MSK_putmaxnumafe]
//
// [MSK_putmaxnumafe] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumafe: MSKint64t
//
// [MSK_putmaxnumafe]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumafe
func (task *Task) PutMaxNumAfe(
	maxnumafe int64,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumafe(
			task.task,
			C.MSKint64t(maxnumafe),
		),
	)
}

// PutMaxNumAnz is wrapping [MSK_putmaxnumanz]
//
// [MSK_putmaxnumanz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumanz: MSKint64t
//
// [MSK_putmaxnumanz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumanz
func (task *Task) PutMaxNumAnz(
	maxnumanz int64,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumanz(
			task.task,
			C.MSKint64t(maxnumanz),
		),
	)
}

// PutMaxNumBarvar is wrapping [MSK_putmaxnumbarvar]
//
// [MSK_putmaxnumbarvar] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumbarvar: MSKint32t
//
// [MSK_putmaxnumbarvar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumbarvar
func (task *Task) PutMaxNumBarvar(
	maxnumbarvar int32,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumbarvar(
			task.task,
			C.MSKint32t(maxnumbarvar),
		),
	)
}

// PutMaxNumCon is wrapping [MSK_putmaxnumcon]
//
// [MSK_putmaxnumcon] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumcon: MSKint32t
//
// [MSK_putmaxnumcon]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumcon
func (task *Task) PutMaxNumCon(
	maxnumcon int32,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumcon(
			task.task,
			C.MSKint32t(maxnumcon),
		),
	)
}

// PutMaxNumCone is wrapping [MSK_putmaxnumcone]
//
// [MSK_putmaxnumcone] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumcone: MSKint32t
//
// [MSK_putmaxnumcone]/PutMaxNumCone is deprecated by mosek and will be removed in a future release.
//
// [MSK_putmaxnumcone]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumcone
func (task *Task) PutMaxNumCone(
	maxnumcone int32,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumcone(
			task.task,
			C.MSKint32t(maxnumcone),
		),
	)
}

// PutMaxNumDjc is wrapping [MSK_putmaxnumdjc]
//
// [MSK_putmaxnumdjc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumdjc: MSKint64t
//
// [MSK_putmaxnumdjc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumdjc
func (task *Task) PutMaxNumDjc(
	maxnumdjc int64,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumdjc(
			task.task,
			C.MSKint64t(maxnumdjc),
		),
	)
}

// PutMaxNumDomain is wrapping [MSK_putmaxnumdomain]
//
// [MSK_putmaxnumdomain] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumdomain: MSKint64t
//
// [MSK_putmaxnumdomain]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumdomain
func (task *Task) PutMaxNumDomain(
	maxnumdomain int64,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumdomain(
			task.task,
			C.MSKint64t(maxnumdomain),
		),
	)
}

// PutMaxNumQnz is wrapping [MSK_putmaxnumqnz]
//
// [MSK_putmaxnumqnz] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumqnz: MSKint64t
//
// [MSK_putmaxnumqnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumqnz
func (task *Task) PutMaxNumQnz(
	maxnumqnz int64,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumqnz(
			task.task,
			C.MSKint64t(maxnumqnz),
		),
	)
}

// PutMaxNumVar is wrapping [MSK_putmaxnumvar]
//
// [MSK_putmaxnumvar] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - maxnumvar: MSKint32t
//
// [MSK_putmaxnumvar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putmaxnumvar
func (task *Task) PutMaxNumVar(
	maxnumvar int32,
) res.Code {
	return res.Code(
		C.MSK_putmaxnumvar(
			task.task,
			C.MSKint32t(maxnumvar),
		),
	)
}
