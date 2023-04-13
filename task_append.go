// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"github.com/fardream/gmsk/res"
)

// AppendAcc is wrapping [MSK_appendacc],
// Appends an affine conic constraint to the task.
//
// Arguments:
//
//   - `domidx` Domain index.
//   - `afeidxlist` List of affine expression indexes.
//   - `b` The vector of constant terms added to affine expressions. Optional, can be NULL.
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

// AppendAccs is wrapping [MSK_appendaccs],
// Appends a number of affine conic constraint to the task.
//
// Arguments:
//
//   - `domidxs` Domain indices.
//   - `afeidxlist` List of affine expression indexes.
//   - `b` The vector of constant terms added to affine expressions. Optional, can be NULL.
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

// AppendAccSeq is wrapping [MSK_appendaccseq],
// Appends an affine conic constraint to the task.
//
// Arguments:
//
//   - `domidx` Domain index.
//   - `afeidxfirst` Index of the first affine expression.
//   - `b` The vector of constant terms added to affine expressions. Optional, can be NULL.
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

// AppendAccsSeq is wrapping [MSK_appendaccsseq],
// Appends a number of affine conic constraint to the task.
//
// Arguments:
//
//   - `domidxs` Domain indices.
//   - `numafeidx` Number of affine expressions in the affine expression list (must equal the sum of dimensions of the domains).
//   - `afeidxfirst` Index of the first affine expression.
//   - `b` The vector of constant terms added to affine expressions. Optional, can be NULL.
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

// AppendAfes is wrapping [MSK_appendafes],
// Appends a number of empty affine expressions to the optimization task.
//
// Arguments:
//
//   - `num` Number of empty affine expressions which should be appended.
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

// AppendBarvars is wrapping [MSK_appendbarvars],
// Appends semidefinite variables to the problem.
//
// Arguments:
//
//   - `dim` Dimensions of symmetric matrix variables to be added.
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

// AppendCone is wrapping [MSK_appendcone],
// Appends a new conic constraint to the problem.
//
// Arguments:
//
//   - `ct` Specifies the type of the cone.
//   - `conepar` For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
//   - `submem` Variable subscripts of the members in the cone.
//
// Deprecated: [MSK_appendcone]/AppendCone is deprecated by mosek and will be removed in a future release.
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

// AppendConeSeq is wrapping [MSK_appendconeseq],
// Appends a new conic constraint to the problem.
//
// Arguments:
//
//   - `ct` Specifies the type of the cone.
//   - `conepar` For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
//   - `nummem` Number of member variables in the cone.
//   - `j` Index of the first variable in the conic constraint.
//
// Deprecated: [MSK_appendconeseq]/AppendConeSeq is deprecated by mosek and will be removed in a future release.
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

// AppendConesSeq is wrapping [MSK_appendconesseq],
// Appends multiple conic constraints to the problem.
//
// Arguments:
//
//   - `ct` Specifies the type of the cone.
//   - `conepar` For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
//   - `nummem` Numbers of member variables in the cones.
//   - `j` Index of the first variable in the first cone to be appended.
//
// Deprecated: [MSK_appendconesseq]/AppendConesSeq is deprecated by mosek and will be removed in a future release.
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

// AppendCons is wrapping [MSK_appendcons],
// Appends a number of constraints to the optimization task.
//
// Arguments:
//
//   - `num` Number of constraints which should be appended.
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

// AppendDjcs is wrapping [MSK_appenddjcs],
// Appends a number of empty disjunctive constraints to the task.
//
// Arguments:
//
//   - `num` Number of empty disjunctive constraints which should be appended.
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

// AppendSparseSymMat is wrapping [MSK_appendsparsesymmat],
// Appends a general sparse symmetric matrix to the storage of symmetric matrices.
//
// Arguments:
//
//   - `dim` Dimension of the symmetric matrix that is appended.
//   - `subi` Row subscript in the triplets.
//   - `subj` Column subscripts in the triplets.
//   - `valij` Values of each triplet.
//
// Returns:
//
//   - `idx` Unique index assigned to the inputted matrix.
//
// [MSK_appendsparsesymmat]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendsparsesymmat
func (task *Task) AppendSparseSymMat(
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

// AppendSparseSymMatList is wrapping [MSK_appendsparsesymmatlist],
// Appends a general sparse symmetric matrix to the storage of symmetric matrices.
//
// Arguments:
//
//   - `dims` Dimensions of the symmetric matrixes.
//   - `nz` Number of nonzeros for each matrix.
//   - `subi` Row subscript in the triplets.
//   - `subj` Column subscripts in the triplets.
//   - `valij` Values of each triplet.
//   - `idx` Unique index assigned to the inputted matrix.
//
// [MSK_appendsparsesymmatlist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.appendsparsesymmatlist
func (task *Task) AppendSparseSymMatList(
	num int32,
	dims *int32,
	nz *int64,
	subi *int32,
	subj *int32,
	valij *float64,
) (r res.Code, idx int64) {
	r = res.Code(
		C.MSK_appendsparsesymmatlist(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(dims),
			(*C.MSKint64t)(nz),
			(*C.MSKint32t)(subi),
			(*C.MSKint32t)(subj),
			(*C.MSKrealt)(valij),
			(*C.MSKint64t)(&idx),
		),
	)

	return
}

// AppendVars is wrapping [MSK_appendvars],
// Appends a number of variables to the optimization task.
//
// Arguments:
//
//   - `num` Number of variables which should be appended.
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
