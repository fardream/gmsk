// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"github.com/fardream/gmsk/res"
)

// AppendAcc is wrapping [MSK_appendacc] and
// adds an affine conic constraint to the task, where the afe idx is provided
// by an array or pointer - if the afe idx is sequential, use [Task.AppendAccSeq]
// to avoid allocating an array.
//
// [MSK_appendacc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - domidx: MSKint64t
//   - numafeidx: MSKint64t
//   - afeidxlist: const MSKint64t *
//   - b: const MSKrealt *
//
// [MSK_appendacc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendacc
func (task *Task) AppendAcc(
	domidx int64,
	numafeidx int64,
	afeidxlist *int64,
	b *float64,
) res.Code {
	return res.Code(
		C.MSK_appendacc(
			task.task,
			C.MSKint64t(domidx),
			C.MSKint64t(numafeidx),
			(*C.MSKint64t)(afeidxlist),
			(*C.MSKrealt)(b),
		),
	)
}

// AppendAccs is wrapping [MSK_appendaccs] and
// adds a list of affine conic constraints to the task.
//
// [MSK_appendaccs] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - numaccs: MSKint64t
//   - domidxs: const MSKint64t *
//   - numafeidx: MSKint64t
//   - afeidxlist: const MSKint64t *
//   - b: const MSKrealt *
//
// [MSK_appendaccs]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendaccs
func (task *Task) AppendAccs(
	numaccs int64,
	domidxs *int64,
	numafeidx int64,
	afeidxlist *int64,
	b *float64,
) res.Code {
	return res.Code(
		C.MSK_appendaccs(
			task.task,
			C.MSKint64t(numaccs),
			(*C.MSKint64t)(domidxs),
			C.MSKint64t(numafeidx),
			(*C.MSKint64t)(afeidxlist),
			(*C.MSKrealt)(b),
		),
	)
}

// AppendAccSeq is wrapping [MSK_appendaccseq] and
// adds an affine conic constraint to the task where the affine idx is sequential.
//
// [MSK_appendaccseq] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - domidx: MSKint64t
//   - numafeidx: MSKint64t
//   - afeidxfirst: MSKint64t
//   - b: const MSKrealt *
//
// [MSK_appendaccseq]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendaccseq
func (task *Task) AppendAccSeq(
	domidx int64,
	numafeidx int64,
	afeidxfirst int64,
	b *float64,
) res.Code {
	return res.Code(
		C.MSK_appendaccseq(
			task.task,
			C.MSKint64t(domidx),
			C.MSKint64t(numafeidx),
			C.MSKint64t(afeidxfirst),
			(*C.MSKrealt)(b),
		),
	)
}

// AppendAccsSeq is wrapping [MSK_appendaccsseq] and
// append a block of accs to the tas - assuming affine expressions are sequential.
//
// [MSK_appendaccsseq] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - numaccs: MSKint64t
//   - domidxs: const MSKint64t *
//   - numafeidx: MSKint64t
//   - afeidxfirst: MSKint64t
//   - b: const MSKrealt *
//
// [MSK_appendaccsseq]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendaccsseq
func (task *Task) AppendAccsSeq(
	numaccs int64,
	domidxs *int64,
	numafeidx int64,
	afeidxfirst int64,
	b *float64,
) res.Code {
	return res.Code(
		C.MSK_appendaccsseq(
			task.task,
			C.MSKint64t(numaccs),
			(*C.MSKint64t)(domidxs),
			C.MSKint64t(numafeidx),
			C.MSKint64t(afeidxfirst),
			(*C.MSKrealt)(b),
		),
	)
}

// AppendAfes is wrapping [MSK_appendafes] and
// adds affine expressions to the task.
//
// [MSK_appendafes] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint64t
//
// [MSK_appendafes]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendafes
func (task *Task) AppendAfes(
	num int64,
) res.Code {
	return res.Code(
		C.MSK_appendafes(
			task.task,
			C.MSKint64t(num),
		),
	)
}

// AppendBarvars is wrapping [MSK_appendbarvars] and
// adds semidefinite matrix variables to the task.
// Barvar because MOSEK uses bar{x} notation to indicate an element
// of a semidefinite matrix.
// The dimension of each of the semidefinite variables are provided
// through the pointer dim.
//
// [MSK_appendbarvars] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//   - dim: const MSKint32t *
//
// [MSK_appendbarvars]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendbarvars
func (task *Task) AppendBarvars(
	num int32,
	dim *int32,
) res.Code {
	return res.Code(
		C.MSK_appendbarvars(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(dim),
		),
	)
}

// AppendCone is wrapping [MSK_appendcone]
//
// [MSK_appendcone] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - ct: MSKconetypee
//   - conepar: MSKrealt
//   - nummem: MSKint32t
//   - submem: const MSKint32t *
//
// [MSK_appendcone]/AppendCone is deprecated by mosek and will be removed in a future release.
//
// [MSK_appendcone]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendcone
func (task *Task) AppendCone(
	ct ConeType,
	conepar float64,
	nummem int32,
	submem *int32,
) res.Code {
	return res.Code(
		C.MSK_appendcone(
			task.task,
			C.MSKconetypee(ct),
			C.MSKrealt(conepar),
			C.MSKint32t(nummem),
			(*C.MSKint32t)(submem),
		),
	)
}

// AppendConeSeq is wrapping [MSK_appendconeseq]
//
// [MSK_appendconeseq] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - ct: MSKconetypee
//   - conepar: MSKrealt
//   - nummem: MSKint32t
//   - j: MSKint32t
//
// [MSK_appendconeseq]/AppendConeSeq is deprecated by mosek and will be removed in a future release.
//
// [MSK_appendconeseq]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendconeseq
func (task *Task) AppendConeSeq(
	ct ConeType,
	conepar float64,
	nummem int32,
	j int32,
) res.Code {
	return res.Code(
		C.MSK_appendconeseq(
			task.task,
			C.MSKconetypee(ct),
			C.MSKrealt(conepar),
			C.MSKint32t(nummem),
			C.MSKint32t(j),
		),
	)
}

// AppendConesSeq is wrapping [MSK_appendconesseq]
//
// [MSK_appendconesseq] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//   - ct: const MSKconetypee *
//   - conepar: const MSKrealt *
//   - nummem: const MSKint32t *
//   - j: MSKint32t
//
// [MSK_appendconesseq]/AppendConesSeq is deprecated by mosek and will be removed in a future release.
//
// [MSK_appendconesseq]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendconesseq
func (task *Task) AppendConesSeq(
	num int32,
	ct *ConeType,
	conepar *float64,
	nummem *int32,
	j int32,
) res.Code {
	return res.Code(
		C.MSK_appendconesseq(
			task.task,
			C.MSKint32t(num),
			(*C.MSKconetypee)(ct),
			(*C.MSKrealt)(conepar),
			(*C.MSKint32t)(nummem),
			C.MSKint32t(j),
		),
	)
}

// AppendCons is wrapping [MSK_appendcons] and
// add vanilla linear constraints to the task.
//
// [MSK_appendcons] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//
// [MSK_appendcons]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendcons
func (task *Task) AppendCons(
	num int32,
) res.Code {
	return res.Code(
		C.MSK_appendcons(
			task.task,
			C.MSKint32t(num),
		),
	)
}

// AppendDjcs is wrapping [MSK_appenddjcs] and
// adds disjunctive constraints to the task.
//
// [MSK_appenddjcs] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint64t
//
// [MSK_appenddjcs]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appenddjcs
func (task *Task) AppendDjcs(
	num int64,
) res.Code {
	return res.Code(
		C.MSK_appenddjcs(
			task.task,
			C.MSKint64t(num),
		),
	)
}

// AppendSparseSymmat is wrapping [MSK_appendsparsesymmat] and
// adds a sparse and symmetric matrix to the task.
// matrix is represented in coordinate format, and only lower triangular portion of the matrix should be
// specified.
// Those matrices can be used as either coefficent in the objective or constraints. The matrix is identified
// by the returned idx.
//
// [MSK_appendsparsesymmat] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - dim: MSKint32t
//   - nz: MSKint64t
//   - subi: const MSKint32t *
//   - subj: const MSKint32t *
//   - valij: const MSKrealt *
//   - idx: MSKint64t *
//
// [MSK_appendsparsesymmat]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendsparsesymmat
func (task *Task) AppendSparseSymmat(
	dim int32,
	nz int64,
	subi *int32,
	subj *int32,
	valij *float64,
) (r res.Code, idx int64) {
	r = res.Code(
		C.MSK_appendsparsesymmat(
			task.task,
			C.MSKint32t(dim),
			C.MSKint64t(nz),
			(*C.MSKint32t)(subi),
			(*C.MSKint32t)(subj),
			(*C.MSKrealt)(valij),
			(*C.MSKint64t)(&idx),
		),
	)

	return
}

// AppendSparseSymmatList is wrapping [MSK_appendsparsesymmatlist]
//
// [MSK_appendsparsesymmatlist] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//   - dims: const MSKint32t *
//   - nz: const MSKint64t *
//   - subi: const MSKint32t *
//   - subj: const MSKint32t *
//   - valij: const MSKrealt *
//   - idx: MSKint64t *
//
// [MSK_appendsparsesymmatlist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendsparsesymmatlist
func (task *Task) AppendSparseSymmatList(
	num int32,
	dims *int32,
	nz *int64,
	subi *int32,
	subj *int32,
	valij *float64,
	idx *int64,
) res.Code {
	return res.Code(
		C.MSK_appendsparsesymmatlist(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(dims),
			(*C.MSKint64t)(nz),
			(*C.MSKint32t)(subi),
			(*C.MSKint32t)(subj),
			(*C.MSKrealt)(valij),
			(*C.MSKint64t)(idx),
		),
	)
}

// AppendVars is wrapping [MSK_appendvars] and
// add variables to the task.
//
// [MSK_appendvars] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//
// [MSK_appendvars]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendvars
func (task *Task) AppendVars(
	num int32,
) res.Code {
	return res.Code(
		C.MSK_appendvars(
			task.task,
			C.MSKint32t(num),
		),
	)
}
