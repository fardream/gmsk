// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

// GetNumAcc is wrapping [MSK_getnumacc],
// Obtains the number of affine conic constraints.
//
// Returns:
//
//   - `num` The number of affine conic constraints.
//
// [MSK_getnumacc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumacc
func (task *Task) GetNumAcc() (num int64, r error) {
	r = ResCode(
		C.MSK_getnumacc(
			task.task,
			(*C.MSKint64t)(&num),
		),
	).ToError()

	return
}

// GetNumAfe is wrapping [MSK_getnumafe],
// Obtains the number of affine expressions.
//
// Returns:
//
//   - `numafe` Number of affine expressions.
//
// [MSK_getnumafe]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumafe
func (task *Task) GetNumAfe() (numafe int64, r error) {
	r = ResCode(
		C.MSK_getnumafe(
			task.task,
			(*C.MSKint64t)(&numafe),
		),
	).ToError()

	return
}

// GetNumANz is wrapping [MSK_getnumanz],
// Obtains the number of non-zeros in the coefficient matrix.
//
// Returns:
//
//   - `numanz` Number of non-zero elements in the linear constraint matrix.
//
// [MSK_getnumanz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumanz
func (task *Task) GetNumANz() (numanz int32, r error) {
	r = ResCode(
		C.MSK_getnumanz(
			task.task,
			(*C.MSKint32t)(&numanz),
		),
	).ToError()

	return
}

// GetNumANz64 is wrapping [MSK_getnumanz64],
// Obtains the number of non-zeros in the coefficient matrix.
//
// Returns:
//
//   - `numanz` Number of non-zero elements in the linear constraint matrix.
//
// [MSK_getnumanz64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumanz64
func (task *Task) GetNumANz64() (numanz int64, r error) {
	r = ResCode(
		C.MSK_getnumanz64(
			task.task,
			(*C.MSKint64t)(&numanz),
		),
	).ToError()

	return
}

// GetNumBaraBlockTriplets is wrapping [MSK_getnumbarablocktriplets],
// Obtains an upper bound on the number of scalar elements in the block triplet form of bara.
//
// Returns:
//
//   - `num` An upper bound on the number of elements in the block triplet form of bara.
//
// [MSK_getnumbarablocktriplets]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumbarablocktriplets
func (task *Task) GetNumBaraBlockTriplets() (num int64, r error) {
	r = ResCode(
		C.MSK_getnumbarablocktriplets(
			task.task,
			(*C.MSKint64t)(&num),
		),
	).ToError()

	return
}

// GetNumBaraNz is wrapping [MSK_getnumbaranz],
// Get the number of nonzero elements in barA.
//
// Returns:
//
//   - `nz` The number of nonzero block elements in barA.
//
// [MSK_getnumbaranz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumbaranz
func (task *Task) GetNumBaraNz() (nz int64, r error) {
	r = ResCode(
		C.MSK_getnumbaranz(
			task.task,
			(*C.MSKint64t)(&nz),
		),
	).ToError()

	return
}

// GetNumBarcBlockTriplets is wrapping [MSK_getnumbarcblocktriplets],
// Obtains an upper bound on the number of elements in the block triplet form of barc.
//
// Returns:
//
//   - `num` An upper bound on the number of elements in the block triplet form of barc.
//
// [MSK_getnumbarcblocktriplets]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumbarcblocktriplets
func (task *Task) GetNumBarcBlockTriplets() (num int64, r error) {
	r = ResCode(
		C.MSK_getnumbarcblocktriplets(
			task.task,
			(*C.MSKint64t)(&num),
		),
	).ToError()

	return
}

// GetNumBarcNz is wrapping [MSK_getnumbarcnz],
// Obtains the number of nonzero elements in barc.
//
// Returns:
//
//   - `nz` The number of nonzero elements in barc.
//
// [MSK_getnumbarcnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumbarcnz
func (task *Task) GetNumBarcNz() (nz int64, r error) {
	r = ResCode(
		C.MSK_getnumbarcnz(
			task.task,
			(*C.MSKint64t)(&nz),
		),
	).ToError()

	return
}

// GetNumBarvar is wrapping [MSK_getnumbarvar],
// Obtains the number of semidefinite variables.
//
// Returns:
//
//   - `numbarvar` Number of semidefinite variables in the problem.
//
// [MSK_getnumbarvar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumbarvar
func (task *Task) GetNumBarvar() (numbarvar int32, r error) {
	r = ResCode(
		C.MSK_getnumbarvar(
			task.task,
			(*C.MSKint32t)(&numbarvar),
		),
	).ToError()

	return
}

// GetNumCon is wrapping [MSK_getnumcon],
// Obtains the number of constraints.
//
// Returns:
//
//   - `numcon` Number of constraints.
//
// [MSK_getnumcon]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumcon
func (task *Task) GetNumCon() (numcon int32, r error) {
	r = ResCode(
		C.MSK_getnumcon(
			task.task,
			(*C.MSKint32t)(&numcon),
		),
	).ToError()

	return
}

// GetNumCone is wrapping [MSK_getnumcone],
// Obtains the number of cones.
//
// Returns:
//
//   - `numcone` Number of conic constraints.
//
// Deprecated: [MSK_getnumcone]/GetNumCone is deprecated by mosek and will be removed in a future release.
//
// [MSK_getnumcone]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumcone
func (task *Task) GetNumCone() (numcone int32, r error) {
	r = ResCode(
		C.MSK_getnumcone(
			task.task,
			(*C.MSKint32t)(&numcone),
		),
	).ToError()

	return
}

// GetNumConeMem is wrapping [MSK_getnumconemem],
// Obtains the number of members in a cone.
//
// Arguments:
//
//   - `k` Index of the cone.
//   - `nummem` Number of member variables in the cone.
//
// Deprecated: [MSK_getnumconemem]/GetNumConeMem is deprecated by mosek and will be removed in a future release.
//
// [MSK_getnumconemem]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumconemem
func (task *Task) GetNumConeMem(
	k int32,
) (nummem int32, r error) {
	r = ResCode(
		C.MSK_getnumconemem(
			task.task,
			C.MSKint32t(k),
			(*C.MSKint32t)(&nummem),
		),
	).ToError()

	return
}

// GetNumDjc is wrapping [MSK_getnumdjc],
// Obtains the number of disjunctive constraints.
//
// Returns:
//
//   - `num` The number of disjunctive constraints.
//
// [MSK_getnumdjc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumdjc
func (task *Task) GetNumDjc() (num int64, r error) {
	r = ResCode(
		C.MSK_getnumdjc(
			task.task,
			(*C.MSKint64t)(&num),
		),
	).ToError()

	return
}

// GetNumDomain is wrapping [MSK_getnumdomain],
// Obtain the number of domains defined.
//
// Returns:
//
//   - `numdomain` Number of domains in the task.
//
// [MSK_getnumdomain]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumdomain
func (task *Task) GetNumDomain() (numdomain int64, r error) {
	r = ResCode(
		C.MSK_getnumdomain(
			task.task,
			(*C.MSKint64t)(&numdomain),
		),
	).ToError()

	return
}

// GetNumIntVar is wrapping [MSK_getnumintvar],
// Obtains the number of integer-constrained variables.
//
// Returns:
//
//   - `numintvar` Number of integer variables.
//
// [MSK_getnumintvar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumintvar
func (task *Task) GetNumIntVar() (numintvar int32, r error) {
	r = ResCode(
		C.MSK_getnumintvar(
			task.task,
			(*C.MSKint32t)(&numintvar),
		),
	).ToError()

	return
}

// GetNumParam is wrapping [MSK_getnumparam],
// Obtains the number of parameters of a given type.
//
// Arguments:
//
//   - `partype` Parameter type.
//   - `numparam` Returns the number of parameters of the requested type.
//
// [MSK_getnumparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumparam
func (task *Task) GetNumParam(
	partype ParameterType,
) (numparam int32, r error) {
	r = ResCode(
		C.MSK_getnumparam(
			task.task,
			C.MSKparametertypee(partype),
			(*C.MSKint32t)(&numparam),
		),
	).ToError()

	return
}

// GetNumQConKNz is wrapping [MSK_getnumqconknz],
// Obtains the number of non-zero quadratic terms in a constraint.
//
// Arguments:
//
//   - `k` Index of the constraint for which the number quadratic terms should be obtained.
//
// Returns:
//
//   - `numqcnz` Number of quadratic terms.
//
// [MSK_getnumqconknz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumqconknz
func (task *Task) GetNumQConKNz(
	k int32,
) (numqcnz int32, r error) {
	r = ResCode(
		C.MSK_getnumqconknz(
			task.task,
			C.MSKint32t(k),
			(*C.MSKint32t)(&numqcnz),
		),
	).ToError()

	return
}

// GetNumQConKNz64 is wrapping [MSK_getnumqconknz64]
//
// [MSK_getnumqconknz64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumqconknz64
func (task *Task) GetNumQConKNz64(
	k int32,
) (numqcnz int64, r error) {
	r = ResCode(
		C.MSK_getnumqconknz64(
			task.task,
			C.MSKint32t(k),
			(*C.MSKint64t)(&numqcnz),
		),
	).ToError()

	return
}

// GetNumQObjNz is wrapping [MSK_getnumqobjnz],
// Obtains the number of non-zero quadratic terms in the objective.
//
// Returns:
//
//   - `numqonz` Number of non-zero elements in the quadratic objective terms.
//
// [MSK_getnumqobjnz]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumqobjnz
func (task *Task) GetNumQObjNz() (numqonz int32, r error) {
	r = ResCode(
		C.MSK_getnumqobjnz(
			task.task,
			(*C.MSKint32t)(&numqonz),
		),
	).ToError()

	return
}

// GetNumQObjNz64 is wrapping [MSK_getnumqobjnz64]
//
// [MSK_getnumqobjnz64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumqobjnz64
func (task *Task) GetNumQObjNz64() (numqonz int64, r error) {
	r = ResCode(
		C.MSK_getnumqobjnz64(
			task.task,
			(*C.MSKint64t)(&numqonz),
		),
	).ToError()

	return
}

// GetNumSymMat is wrapping [MSK_getnumsymmat],
// Obtains the number of symmetric matrices stored.
//
// Arguments:
//
//   - `num` The number of symmetric sparse matrices.
//
// [MSK_getnumsymmat]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumsymmat
func (task *Task) GetNumSymMat() (num int64, r error) {
	r = ResCode(
		C.MSK_getnumsymmat(
			task.task,
			(*C.MSKint64t)(&num),
		),
	).ToError()

	return
}

// GetNumVar is wrapping [MSK_getnumvar],
// Obtains the number of variables.
//
// Returns:
//
//   - `numvar` Number of variables.
//
// [MSK_getnumvar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getnumvar
func (task *Task) GetNumVar() (numvar int32, r error) {
	r = ResCode(
		C.MSK_getnumvar(
			task.task,
			(*C.MSKint32t)(&numvar),
		),
	).ToError()

	return
}
