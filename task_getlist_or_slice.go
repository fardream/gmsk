// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

// GetAccAfeIdxList is wrapping [MSK_getaccafeidxlist],
// Obtains the list of affine expressions appearing in the affine conic constraint.
//
// Arguments:
//
//   - `accidx` Index of the affine conic constraint.
//   - `afeidxlist` List of indexes of affine expressions appearing in the constraint.
//
// [MSK_getaccafeidxlist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getaccafeidxlist
func (task *Task) GetAccAfeIdxList(
	accidx int64,
	afeidxlist []int64,
) error {
	return ResCode(
		C.MSK_getaccafeidxlist(
			task.task,
			C.MSKint64t(accidx),
			(*C.MSKint64t)(getPtrToFirst(afeidxlist)),
		),
	).ToError()
}

// GetAColSlice is wrapping [MSK_getacolslice],
// Obtains a sequence of columns from the coefficient matrix.
//
// Arguments:
//
//   - `first` Index of the first column in the sequence.
//   - `last` Index of the last column in the sequence plus one.
//   - `ptrb` Column start pointers.
//   - `ptre` Column end pointers.
//   - `sub` Contains the row subscripts.
//   - `val` Contains the coefficient values.
//
// [MSK_getacolslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getacolslice
func (task *Task) GetAColSlice(
	first int32,
	last int32,
	maxnumnz int32,
	ptrb []int32,
	ptre []int32,
	sub []int32,
	val []float64,
) error {
	return ResCode(
		C.MSK_getacolslice(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			C.MSKint32t(maxnumnz),
			(*C.MSKint32t)(getPtrToFirst(ptrb)),
			(*C.MSKint32t)(getPtrToFirst(ptre)),
			(*C.MSKint32t)(getPtrToFirst(sub)),
			(*C.MSKrealt)(getPtrToFirst(val)),
		),
	).ToError()
}

// GetAfeGSlice is wrapping [MSK_getafegslice],
// Obtains a sequence of coefficients from the vector g.
//
// Arguments:
//
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `g` The slice of g as a dense vector.
//
// [MSK_getafegslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getafegslice
func (task *Task) GetAfeGSlice(
	first int64,
	last int64,
	g []float64,
) error {
	return ResCode(
		C.MSK_getafegslice(
			task.task,
			C.MSKint64t(first),
			C.MSKint64t(last),
			(*C.MSKrealt)(getPtrToFirst(g)),
		),
	).ToError()
}

// GetARowSlice is wrapping [MSK_getarowslice],
// Obtains a sequence of rows from the coefficient matrix.
//
// Arguments:
//
//   - `first` Index of the first row in the sequence.
//   - `last` Index of the last row in the sequence plus one.
//   - `ptrb` Row start pointers.
//   - `ptre` Row end pointers.
//   - `sub` Contains the column subscripts.
//   - `val` Contains the coefficient values.
//
// [MSK_getarowslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getarowslice
func (task *Task) GetARowSlice(
	first int32,
	last int32,
	maxnumnz int32,
	ptrb []int32,
	ptre []int32,
	sub []int32,
	val []float64,
) error {
	return ResCode(
		C.MSK_getarowslice(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			C.MSKint32t(maxnumnz),
			(*C.MSKint32t)(getPtrToFirst(ptrb)),
			(*C.MSKint32t)(getPtrToFirst(ptre)),
			(*C.MSKint32t)(getPtrToFirst(sub)),
			(*C.MSKrealt)(getPtrToFirst(val)),
		),
	).ToError()
}

// GetBarsSlice is wrapping [MSK_getbarsslice],
// Obtains the dual solution for a sequence of semidefinite variables.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` Index of the first semidefinite variable in the slice.
//   - `last` Index of the last semidefinite variable in the slice plus one.
//   - `slicesize` Denotes the length of the array barsslice.
//   - `barsslice` Dual solution values of symmetric matrix variables in the slice, stored sequentially.
//
// [MSK_getbarsslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getbarsslice
func (task *Task) GetBarsSlice(
	whichsol SolType,
	first int32,
	last int32,
	slicesize int64,
	barsslice []float64,
) error {
	return ResCode(
		C.MSK_getbarsslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			C.MSKint64t(slicesize),
			(*C.MSKrealt)(getPtrToFirst(barsslice)),
		),
	).ToError()
}

// GetBarxSlice is wrapping [MSK_getbarxslice],
// Obtains the primal solution for a sequence of semidefinite variables.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` Index of the first semidefinite variable in the slice.
//   - `last` Index of the last semidefinite variable in the slice plus one.
//   - `slicesize` Denotes the length of the array barxslice.
//   - `barxslice` Solution values of symmetric matrix variables in the slice, stored sequentially.
//
// [MSK_getbarxslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getbarxslice
func (task *Task) GetBarxSlice(
	whichsol SolType,
	first int32,
	last int32,
	slicesize int64,
	barxslice []float64,
) error {
	return ResCode(
		C.MSK_getbarxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			C.MSKint64t(slicesize),
			(*C.MSKrealt)(getPtrToFirst(barxslice)),
		),
	).ToError()
}

// GetCList is wrapping [MSK_getclist],
// Obtains a sequence of coefficients from the objective.
//
// Arguments:
//
//   - `subj` A list of variable indexes.
//   - `c` Linear terms of the requested list of the objective as a dense vector.
//
// [MSK_getclist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getclist
func (task *Task) GetCList(
	num int32,
	subj []int32,
	c []float64,
) error {
	return ResCode(
		C.MSK_getclist(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(getPtrToFirst(subj)),
			(*C.MSKrealt)(getPtrToFirst(c)),
		),
	).ToError()
}

// GetConBoundSlice is wrapping [MSK_getconboundslice],
// Obtains bounds information for a slice of the constraints.
//
// Arguments:
//
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `bk` Bound keys.
//   - `bl` Values for lower bounds.
//   - `bu` Values for upper bounds.
//
// [MSK_getconboundslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getconboundslice
func (task *Task) GetConBoundSlice(
	first int32,
	last int32,
	bk []BoundKey,
	bl []float64,
	bu []float64,
) error {
	return ResCode(
		C.MSK_getconboundslice(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKboundkeye)(getPtrToFirst(bk)),
			(*C.MSKrealt)(getPtrToFirst(bl)),
			(*C.MSKrealt)(getPtrToFirst(bu)),
		),
	).ToError()
}

// GetCSlice is wrapping [MSK_getcslice],
// Obtains a sequence of coefficients from the objective.
//
// Arguments:
//
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `c` Linear terms of the requested slice of the objective as a dense vector.
//
// [MSK_getcslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getcslice
func (task *Task) GetCSlice(
	first int32,
	last int32,
	c []float64,
) error {
	return ResCode(
		C.MSK_getcslice(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(c)),
		),
	).ToError()
}

// GetDjcAfeIdxList is wrapping [MSK_getdjcafeidxlist],
// Obtains the list of affine expression indexes in a disjunctive constraint.
//
// Arguments:
//
//   - `djcidx` Index of the disjunctive constraint.
//   - `afeidxlist` List of affine expression indexes.
//
// [MSK_getdjcafeidxlist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdjcafeidxlist
func (task *Task) GetDjcAfeIdxList(
	djcidx int64,
	afeidxlist []int64,
) error {
	return ResCode(
		C.MSK_getdjcafeidxlist(
			task.task,
			C.MSKint64t(djcidx),
			(*C.MSKint64t)(getPtrToFirst(afeidxlist)),
		),
	).ToError()
}

// GetDjcDomainIdxList is wrapping [MSK_getdjcdomainidxlist],
// Obtains the list of domain indexes in a disjunctive constraint.
//
// Arguments:
//
//   - `djcidx` Index of the disjunctive constraint.
//   - `domidxlist` List of term sizes.
//
// [MSK_getdjcdomainidxlist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdjcdomainidxlist
func (task *Task) GetDjcDomainIdxList(
	djcidx int64,
	domidxlist []int64,
) error {
	return ResCode(
		C.MSK_getdjcdomainidxlist(
			task.task,
			C.MSKint64t(djcidx),
			(*C.MSKint64t)(getPtrToFirst(domidxlist)),
		),
	).ToError()
}

// GetDjcTermSizeList is wrapping [MSK_getdjctermsizelist],
// Obtains the list of term sizes in a disjunctive constraint.
//
// Arguments:
//
//   - `djcidx` Index of the disjunctive constraint.
//   - `termsizelist` List of term sizes.
//
// [MSK_getdjctermsizelist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdjctermsizelist
func (task *Task) GetDjcTermSizeList(
	djcidx int64,
	termsizelist []int64,
) error {
	return ResCode(
		C.MSK_getdjctermsizelist(
			task.task,
			C.MSKint64t(djcidx),
			(*C.MSKint64t)(getPtrToFirst(termsizelist)),
		),
	).ToError()
}

// GetSkcSlice is wrapping [MSK_getskcslice],
// Obtains the status keys for a slice of the constraints.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `skc` Status keys for the constraints.
//
// [MSK_getskcslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getskcslice
func (task *Task) GetSkcSlice(
	whichsol SolType,
	first int32,
	last int32,
	skc []StaKey,
) error {
	return ResCode(
		C.MSK_getskcslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKstakeye)(getPtrToFirst(skc)),
		),
	).ToError()
}

// GetSkxSlice is wrapping [MSK_getskxslice],
// Obtains the status keys for a slice of the scalar variables.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `skx` Status keys for the variables.
//
// [MSK_getskxslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getskxslice
func (task *Task) GetSkxSlice(
	whichsol SolType,
	first int32,
	last int32,
	skx []StaKey,
) error {
	return ResCode(
		C.MSK_getskxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKstakeye)(getPtrToFirst(skx)),
		),
	).ToError()
}

// GetSlcSlice is wrapping [MSK_getslcslice],
// Obtains a slice of the slc vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `slc` Dual variables corresponding to the lower bounds on the constraints.
//
// [MSK_getslcslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getslcslice
func (task *Task) GetSlcSlice(
	whichsol SolType,
	first int32,
	last int32,
	slc []float64,
) error {
	return ResCode(
		C.MSK_getslcslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(slc)),
		),
	).ToError()
}

// GetSlxSlice is wrapping [MSK_getslxslice],
// Obtains a slice of the slx vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `slx` Dual variables corresponding to the lower bounds on the variables.
//
// [MSK_getslxslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getslxslice
func (task *Task) GetSlxSlice(
	whichsol SolType,
	first int32,
	last int32,
	slx []float64,
) error {
	return ResCode(
		C.MSK_getslxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(slx)),
		),
	).ToError()
}

// GetSnxSlice is wrapping [MSK_getsnxslice],
// Obtains a slice of the snx vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `snx` Dual variables corresponding to the conic constraints on the variables.
//
// [MSK_getsnxslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getsnxslice
func (task *Task) GetSnxSlice(
	whichsol SolType,
	first int32,
	last int32,
	snx []float64,
) error {
	return ResCode(
		C.MSK_getsnxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(snx)),
		),
	).ToError()
}

// GetSolutionSlice is wrapping [MSK_getsolutionslice],
// Obtains a slice of the solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `solitem` Which part of the solution is required.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `values` The values of the requested solution elements.
//
// [MSK_getsolutionslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getsolutionslice
func (task *Task) GetSolutionSlice(
	whichsol SolType,
	solitem SolItem,
	first int32,
	last int32,
	values []float64,
) error {
	return ResCode(
		C.MSK_getsolutionslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKsoliteme(solitem),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(values)),
		),
	).ToError()
}

// GetSucSlice is wrapping [MSK_getsucslice],
// Obtains a slice of the suc vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `suc` Dual variables corresponding to the upper bounds on the constraints.
//
// [MSK_getsucslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getsucslice
func (task *Task) GetSucSlice(
	whichsol SolType,
	first int32,
	last int32,
	suc []float64,
) error {
	return ResCode(
		C.MSK_getsucslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(suc)),
		),
	).ToError()
}

// GetSuxSlice is wrapping [MSK_getsuxslice],
// Obtains a slice of the sux vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `sux` Dual variables corresponding to the upper bounds on the variables.
//
// [MSK_getsuxslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getsuxslice
func (task *Task) GetSuxSlice(
	whichsol SolType,
	first int32,
	last int32,
	sux []float64,
) error {
	return ResCode(
		C.MSK_getsuxslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(sux)),
		),
	).ToError()
}

// GetVarBoundSlice is wrapping [MSK_getvarboundslice],
// Obtains bounds information for a slice of the variables.
//
// Arguments:
//
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `bk` Bound keys.
//   - `bl` Values for lower bounds.
//   - `bu` Values for upper bounds.
//
// [MSK_getvarboundslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getvarboundslice
func (task *Task) GetVarBoundSlice(
	first int32,
	last int32,
	bk []BoundKey,
	bl []float64,
	bu []float64,
) error {
	return ResCode(
		C.MSK_getvarboundslice(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKboundkeye)(getPtrToFirst(bk)),
			(*C.MSKrealt)(getPtrToFirst(bl)),
			(*C.MSKrealt)(getPtrToFirst(bu)),
		),
	).ToError()
}

// GetVarTypeList is wrapping [MSK_getvartypelist],
// Obtains the variable type for one or more variables.
//
// Arguments:
//
//   - `subj` A list of variable indexes.
//   - `vartype` Returns the variables types corresponding the variable indexes requested.
//
// [MSK_getvartypelist]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getvartypelist
func (task *Task) GetVarTypeList(
	num int32,
	subj []int32,
	vartype []VariableType,
) error {
	return ResCode(
		C.MSK_getvartypelist(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(getPtrToFirst(subj)),
			(*C.MSKvariabletypee)(getPtrToFirst(vartype)),
		),
	).ToError()
}

// GetXcSlice is wrapping [MSK_getxcslice],
// Obtains a slice of the xc vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `xc` Primal constraint solution.
//
// [MSK_getxcslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getxcslice
func (task *Task) GetXcSlice(
	whichsol SolType,
	first int32,
	last int32,
	xc []float64,
) error {
	return ResCode(
		C.MSK_getxcslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(xc)),
		),
	).ToError()
}

// GetYSlice is wrapping [MSK_getyslice],
// Obtains a slice of the y vector for a solution.
//
// Arguments:
//
//   - `whichsol` Selects a solution.
//   - `first` First index in the sequence.
//   - `last` Last index plus 1 in the sequence.
//   - `y` Vector of dual variables corresponding to the constraints.
//
// [MSK_getyslice]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getyslice
func (task *Task) GetYSlice(
	whichsol SolType,
	first int32,
	last int32,
	y []float64,
) error {
	return ResCode(
		C.MSK_getyslice(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKrealt)(getPtrToFirst(y)),
		),
	).ToError()
}
