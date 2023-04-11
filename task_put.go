// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"unsafe"

	"github.com/fardream/gmsk/res"
)

// PutAcc is wrapping [MSK_putacc]
//
// [MSK_putacc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - accidx: MSKint64t
//   - domidx: MSKint64t
//   - numafeidx: MSKint64t
//   - afeidxlist: const MSKint64t *
//   - b: const MSKrealt *
//
// [MSK_putacc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putacc
func (task *Task) PutAcc(
	accidx int64,
	domidx int64,
	numafeidx int64,
	afeidxlist *int64,
	b *float64,
) res.Code {
	return res.Code(
		C.MSK_putacc(
			task.task,
			C.MSKint64t(accidx),
			C.MSKint64t(domidx),
			C.MSKint64t(numafeidx),
			(*C.MSKint64t)(afeidxlist),
			(*C.MSKrealt)(b),
		),
	)
}

// PutAccb is wrapping [MSK_putaccb]
//
// [MSK_putaccb] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - accidx: MSKint64t
//   - lengthb: MSKint64t
//   - b: const MSKrealt *
//
// [MSK_putaccb]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putaccb
func (task *Task) PutAccb(
	accidx int64,
	lengthb int64,
	b *float64,
) res.Code {
	return res.Code(
		C.MSK_putaccb(
			task.task,
			C.MSKint64t(accidx),
			C.MSKint64t(lengthb),
			(*C.MSKrealt)(b),
		),
	)
}

// PutAccbj is wrapping [MSK_putaccbj]
//
// [MSK_putaccbj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - accidx: MSKint64t
//   - j: MSKint64t
//   - bj: MSKrealt
//
// [MSK_putaccbj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putaccbj
func (task *Task) PutAccbj(
	accidx int64,
	j int64,
	bj float64,
) res.Code {
	return res.Code(
		C.MSK_putaccbj(
			task.task,
			C.MSKint64t(accidx),
			C.MSKint64t(j),
			C.MSKrealt(bj),
		),
	)
}

// PutAccDotY is wrapping [MSK_putaccdoty]
//
// [MSK_putaccdoty] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - accidx: MSKint64t
//   - doty: MSKrealt *
//
// [MSK_putaccdoty]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putaccdoty
func (task *Task) PutAccDotY(
	whichsol SolType,
	accidx int64,
	doty *float64,
) res.Code {
	return res.Code(
		C.MSK_putaccdoty(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint64t(accidx),
			(*C.MSKrealt)(doty),
		),
	)
}

// PutACol is wrapping [MSK_putacol] and
// puts a column of A matrix.
//
// [MSK_putacol] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - j: MSKint32t
//   - nzj: MSKint32t
//   - subj: const MSKint32t *
//   - valj: const MSKrealt *
//
// [MSK_putacol]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putacol
func (task *Task) PutACol(
	j int32,
	nzj int32,
	subj *int32,
	valj *float64,
) res.Code {
	return res.Code(
		C.MSK_putacol(
			task.task,
			C.MSKint32t(j),
			C.MSKint32t(nzj),
			(*C.MSKint32t)(subj),
			(*C.MSKrealt)(valj),
		),
	)
}

// PutAColSlice64 is wrapping [MSK_putacolslice64]
//
// [MSK_putacolslice64] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - first: MSKint32t
//   - last: MSKint32t
//   - ptrb: const MSKint64t *
//   - ptre: const MSKint64t *
//   - asub: const MSKint32t *
//   - aval: const MSKrealt *
//
// [MSK_putacolslice64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putacolslice64
func (task *Task) PutAColSlice64(
	first int32,
	last int32,
	ptrb *int64,
	ptre *int64,
	asub *int32,
	aval *float64,
) res.Code {
	return res.Code(
		C.MSK_putacolslice64(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKint64t)(ptrb),
			(*C.MSKint64t)(ptre),
			(*C.MSKint32t)(asub),
			(*C.MSKrealt)(aval),
		),
	)
}

// PutAfeBarFBlockTriplet is wrapping [MSK_putafebarfblocktriplet] and
// sets a matrix variable to the F matrix of affine expression.
// the rows of afe idx is provided in afeidx, the matrix variables are indexed by barvaridx, and subk, subl are the indices
// of the coefficients for those matrix variables, and valkl are the corresponding coefficients.
//
// [MSK_putafebarfblocktriplet] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - numtrip: MSKint64t
//   - afeidx: const MSKint64t *
//   - barvaridx: const MSKint32t *
//   - subk: const MSKint32t *
//   - subl: const MSKint32t *
//   - valkl: const MSKrealt *
//
// [MSK_putafebarfblocktriplet]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafebarfblocktriplet
func (task *Task) PutAfeBarFBlockTriplet(
	numtrip int64,
	afeidx *int64,
	barvaridx *int32,
	subk *int32,
	subl *int32,
	valkl *float64,
) res.Code {
	return res.Code(
		C.MSK_putafebarfblocktriplet(
			task.task,
			C.MSKint64t(numtrip),
			(*C.MSKint64t)(afeidx),
			(*C.MSKint32t)(barvaridx),
			(*C.MSKint32t)(subk),
			(*C.MSKint32t)(subl),
			(*C.MSKrealt)(valkl),
		),
	)
}

// PutAfeBarFEntry is wrapping [MSK_putafebarfentry]
//
// [MSK_putafebarfentry] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - afeidx: MSKint64t
//   - barvaridx: MSKint32t
//   - numterm: MSKint64t
//   - termidx: const MSKint64t *
//   - termweight: const MSKrealt *
//
// [MSK_putafebarfentry]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafebarfentry
func (task *Task) PutAfeBarFEntry(
	afeidx int64,
	barvaridx int32,
	numterm int64,
	termidx *int64,
	termweight *float64,
) res.Code {
	return res.Code(
		C.MSK_putafebarfentry(
			task.task,
			C.MSKint64t(afeidx),
			C.MSKint32t(barvaridx),
			C.MSKint64t(numterm),
			(*C.MSKint64t)(termidx),
			(*C.MSKrealt)(termweight),
		),
	)
}

// PutAfeBarFRow is wrapping [MSK_putafebarfrow]
//
// [MSK_putafebarfrow] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - afeidx: MSKint64t
//   - numentr: MSKint32t
//   - barvaridx: const MSKint32t *
//   - numterm: const MSKint64t *
//   - ptrterm: const MSKint64t *
//   - lenterm: MSKint64t
//   - termidx: const MSKint64t *
//   - termweight: const MSKrealt *
//
// [MSK_putafebarfrow]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafebarfrow
func (task *Task) PutAfeBarFRow(
	afeidx int64,
	numentr int32,
	barvaridx *int32,
	numterm *int64,
	ptrterm *int64,
	lenterm int64,
	termidx *int64,
	termweight *float64,
) res.Code {
	return res.Code(
		C.MSK_putafebarfrow(
			task.task,
			C.MSKint64t(afeidx),
			C.MSKint32t(numentr),
			(*C.MSKint32t)(barvaridx),
			(*C.MSKint64t)(numterm),
			(*C.MSKint64t)(ptrterm),
			C.MSKint64t(lenterm),
			(*C.MSKint64t)(termidx),
			(*C.MSKrealt)(termweight),
		),
	)
}

// PutAfeFCol is wrapping [MSK_putafefcol] and
// sets a column of affine expression F matrix.
//
// [MSK_putafefcol] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - varidx: MSKint32t
//   - numnz: MSKint64t
//   - afeidx: const MSKint64t *
//   - val: const MSKrealt *
//
// [MSK_putafefcol]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafefcol
func (task *Task) PutAfeFCol(
	varidx int32,
	numnz int64,
	afeidx *int64,
	val *float64,
) res.Code {
	return res.Code(
		C.MSK_putafefcol(
			task.task,
			C.MSKint32t(varidx),
			C.MSKint64t(numnz),
			(*C.MSKint64t)(afeidx),
			(*C.MSKrealt)(val),
		),
	)
}

// PutAfeFEntry is wrapping [MSK_putafefentry] and
// sets an entry in the  affine expression F matrix.
//
// [MSK_putafefentry] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - afeidx: MSKint64t
//   - varidx: MSKint32t
//   - value: MSKrealt
//
// [MSK_putafefentry]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafefentry
func (task *Task) PutAfeFEntry(
	afeidx int64,
	varidx int32,
	value float64,
) res.Code {
	return res.Code(
		C.MSK_putafefentry(
			task.task,
			C.MSKint64t(afeidx),
			C.MSKint32t(varidx),
			C.MSKrealt(value),
		),
	)
}

// PutAfeFRow is wrapping [MSK_putafefrow] and
// sets a row of affine expression F matrix.
//
// [MSK_putafefrow] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - afeidx: MSKint64t
//   - numnz: MSKint32t
//   - varidx: const MSKint32t *
//   - val: const MSKrealt *
//
// [MSK_putafefrow]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafefrow
func (task *Task) PutAfeFRow(
	afeidx int64,
	numnz int32,
	varidx *int32,
	val *float64,
) res.Code {
	return res.Code(
		C.MSK_putafefrow(
			task.task,
			C.MSKint64t(afeidx),
			C.MSKint32t(numnz),
			(*C.MSKint32t)(varidx),
			(*C.MSKrealt)(val),
		),
	)
}

// PutAfeG is wrapping [MSK_putafeg] and
// sets the value at afeidx to g.
//
// [MSK_putafeg] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - afeidx: MSKint64t
//   - g: MSKrealt
//
// [MSK_putafeg]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putafeg
func (task *Task) PutAfeG(
	afeidx int64,
	g float64,
) res.Code {
	return res.Code(
		C.MSK_putafeg(
			task.task,
			C.MSKint64t(afeidx),
			C.MSKrealt(g),
		),
	)
}

// PutAij is wrapping [MSK_putaij] and
// set the value of the constraints matrix A[i,j]
//
// [MSK_putaij] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - j: MSKint32t
//   - aij: MSKrealt
//
// [MSK_putaij]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putaij
func (task *Task) PutAij(
	i int32,
	j int32,
	aij float64,
) res.Code {
	return res.Code(
		C.MSK_putaij(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(j),
			C.MSKrealt(aij),
		),
	)
}

// PutARow is wrapping [MSK_putarow] and
// puts a row of A matrix.
//
// [MSK_putarow] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - nzi: MSKint32t
//   - subi: const MSKint32t *
//   - vali: const MSKrealt *
//
// [MSK_putarow]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putarow
func (task *Task) PutARow(
	i int32,
	nzi int32,
	subi *int32,
	vali *float64,
) res.Code {
	return res.Code(
		C.MSK_putarow(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(nzi),
			(*C.MSKint32t)(subi),
			(*C.MSKrealt)(vali),
		),
	)
}

// PutARowSlice64 is wrapping [MSK_putarowslice64]
//
// [MSK_putarowslice64] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - first: MSKint32t
//   - last: MSKint32t
//   - ptrb: const MSKint64t *
//   - ptre: const MSKint64t *
//   - asub: const MSKint32t *
//   - aval: const MSKrealt *
//
// [MSK_putarowslice64]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putarowslice64
func (task *Task) PutARowSlice64(
	first int32,
	last int32,
	ptrb *int64,
	ptre *int64,
	asub *int32,
	aval *float64,
) res.Code {
	return res.Code(
		C.MSK_putarowslice64(
			task.task,
			C.MSKint32t(first),
			C.MSKint32t(last),
			(*C.MSKint64t)(ptrb),
			(*C.MSKint64t)(ptre),
			(*C.MSKint32t)(asub),
			(*C.MSKrealt)(aval),
		),
	)
}

// PutAtruncatetol is wrapping [MSK_putatruncatetol]
//
// [MSK_putatruncatetol] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - tolzero: MSKrealt
//
// [MSK_putatruncatetol]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putatruncatetol
func (task *Task) PutAtruncatetol(
	tolzero float64,
) res.Code {
	return res.Code(
		C.MSK_putatruncatetol(
			task.task,
			C.MSKrealt(tolzero),
		),
	)
}

// PutBarABlockTriplet is wrapping [MSK_putbarablocktriplet] and
// sets constraints for matrix variable.
// subi is the index of the constraint, subj is the index of the matrix variable,
// subk and subl are the indices of the coefficients and valijkl are the coefficients value.
//
// [MSK_putbarablocktriplet] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint64t
//   - subi: const MSKint32t *
//   - subj: const MSKint32t *
//   - subk: const MSKint32t *
//   - subl: const MSKint32t *
//   - valijkl: const MSKrealt *
//
// [MSK_putbarablocktriplet]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbarablocktriplet
func (task *Task) PutBarABlockTriplet(
	num int64,
	subi *int32,
	subj *int32,
	subk *int32,
	subl *int32,
	valijkl *float64,
) res.Code {
	return res.Code(
		C.MSK_putbarablocktriplet(
			task.task,
			C.MSKint64t(num),
			(*C.MSKint32t)(subi),
			(*C.MSKint32t)(subj),
			(*C.MSKint32t)(subk),
			(*C.MSKint32t)(subl),
			(*C.MSKrealt)(valijkl),
		),
	)
}

// PutBarAij is wrapping [MSK_putbaraij] and
// a semidefinite matrix to constraint.
// i is the index of the constraint,and j is the index of the semidefinite matrix variable.
// num is the number of coefficients matrices, and sub/weights are the coefficient matrices'
// idx (when they are added by [Task.AppendSparseSymmat]) and weights.
//
// [MSK_putbaraij] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - j: MSKint32t
//   - num: MSKint64t
//   - sub: const MSKint64t *
//   - weights: const MSKrealt *
//
// [MSK_putbaraij]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbaraij
func (task *Task) PutBarAij(
	i int32,
	j int32,
	num int64,
	sub *int64,
	weights *float64,
) res.Code {
	return res.Code(
		C.MSK_putbaraij(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(j),
			C.MSKint64t(num),
			(*C.MSKint64t)(sub),
			(*C.MSKrealt)(weights),
		),
	)
}

// PutBarCBlockTriplet is wrapping [MSK_putbarcblocktriplet] and
// sets the coefficients for
// matrix variable. suj is the index of the matrix variable, subk/subl are row/colmn index
// of the cofficients, and valjkl is the coefficients value.
//
// [MSK_putbarcblocktriplet] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint64t
//   - subj: const MSKint32t *
//   - subk: const MSKint32t *
//   - subl: const MSKint32t *
//   - valjkl: const MSKrealt *
//
// [MSK_putbarcblocktriplet]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbarcblocktriplet
func (task *Task) PutBarCBlockTriplet(
	num int64,
	subj *int32,
	subk *int32,
	subl *int32,
	valjkl *float64,
) res.Code {
	return res.Code(
		C.MSK_putbarcblocktriplet(
			task.task,
			C.MSKint64t(num),
			(*C.MSKint32t)(subj),
			(*C.MSKint32t)(subk),
			(*C.MSKint32t)(subl),
			(*C.MSKrealt)(valjkl),
		),
	)
}

// PutBarCj is wrapping [MSK_putbarcj] and
// adds a positive semidefinite matrix to the objective.
// j is the index of the matrix variable, and num is number of weight matrices. Those
// weight matrices are identified by sub (the idx when they are added to the task
// by [Task.AppendSparseSymmat].
//
// [MSK_putbarcj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - j: MSKint32t
//   - num: MSKint64t
//   - sub: const MSKint64t *
//   - weights: const MSKrealt *
//
// [MSK_putbarcj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbarcj
func (task *Task) PutBarCj(
	j int32,
	num int64,
	sub *int64,
	weights *float64,
) res.Code {
	return res.Code(
		C.MSK_putbarcj(
			task.task,
			C.MSKint32t(j),
			C.MSKint64t(num),
			(*C.MSKint64t)(sub),
			(*C.MSKrealt)(weights),
		),
	)
}

// PutBarsj is wrapping [MSK_putbarsj]
//
// [MSK_putbarsj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - j: MSKint32t
//   - barsj: const MSKrealt *
//
// [MSK_putbarsj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbarsj
func (task *Task) PutBarsj(
	whichsol SolType,
	j int32,
	barsj *float64,
) res.Code {
	return res.Code(
		C.MSK_putbarsj(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(j),
			(*C.MSKrealt)(barsj),
		),
	)
}

// PutBarxj is wrapping [MSK_putbarxj]
//
// [MSK_putbarxj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - j: MSKint32t
//   - barxj: const MSKrealt *
//
// [MSK_putbarxj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbarxj
func (task *Task) PutBarxj(
	whichsol SolType,
	j int32,
	barxj *float64,
) res.Code {
	return res.Code(
		C.MSK_putbarxj(
			task.task,
			C.MSKsoltypee(whichsol),
			C.MSKint32t(j),
			(*C.MSKrealt)(barxj),
		),
	)
}

// PutCFix is wrapping [MSK_putcfix] and
// sets the constant term in the objective.
//
// [MSK_putcfix] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - cfix: MSKrealt
//
// [MSK_putcfix]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putcfix
func (task *Task) PutCFix(
	cfix float64,
) res.Code {
	return res.Code(
		C.MSK_putcfix(
			task.task,
			C.MSKrealt(cfix),
		),
	)
}

// PutCj is wrapping [MSK_putcj] and
// set the coefficient in the objective function.
//
// [MSK_putcj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - j: MSKint32t
//   - cj: MSKrealt
//
// [MSK_putcj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putcj
func (task *Task) PutCj(
	j int32,
	cj float64,
) res.Code {
	return res.Code(
		C.MSK_putcj(
			task.task,
			C.MSKint32t(j),
			C.MSKrealt(cj),
		),
	)
}

// PutConbound is wrapping [MSK_putconbound] and
// set the bound for a contraint
//
// [MSK_putconbound] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - bkc: MSKboundkeye
//   - blc: MSKrealt
//   - buc: MSKrealt
//
// [MSK_putconbound]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putconbound
func (task *Task) PutConbound(
	i int32,
	bkc BoundKey,
	blc float64,
	buc float64,
) res.Code {
	return res.Code(
		C.MSK_putconbound(
			task.task,
			C.MSKint32t(i),
			C.MSKboundkeye(bkc),
			C.MSKrealt(blc),
			C.MSKrealt(buc),
		),
	)
}

// PutConboundListConst is wrapping [MSK_putconboundlistconst]
//
// [MSK_putconboundlistconst] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//   - sub: const MSKint32t *
//   - bkc: MSKboundkeye
//   - blc: MSKrealt
//   - buc: MSKrealt
//
// [MSK_putconboundlistconst]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putconboundlistconst
func (task *Task) PutConboundListConst(
	num int32,
	sub *int32,
	bkc BoundKey,
	blc float64,
	buc float64,
) res.Code {
	return res.Code(
		C.MSK_putconboundlistconst(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(sub),
			C.MSKboundkeye(bkc),
			C.MSKrealt(blc),
			C.MSKrealt(buc),
		),
	)
}

// PutCone is wrapping [MSK_putcone]
//
// [MSK_putcone] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - k: MSKint32t
//   - ct: MSKconetypee
//   - conepar: MSKrealt
//   - nummem: MSKint32t
//   - submem: const MSKint32t *
//
// Deprecated: [MSK_putcone]/PutCone is deprecated by mosek and will be removed in a future release.
//
// [MSK_putcone]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putcone
func (task *Task) PutCone(
	k int32,
	ct ConeType,
	conepar float64,
	nummem int32,
	submem *int32,
) res.Code {
	return res.Code(
		C.MSK_putcone(
			task.task,
			C.MSKint32t(k),
			C.MSKconetypee(ct),
			C.MSKrealt(conepar),
			C.MSKint32t(nummem),
			(*C.MSKint32t)(submem),
		),
	)
}

// PutConsolutioni is wrapping [MSK_putconsolutioni]
//
// [MSK_putconsolutioni] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - whichsol: MSKsoltypee
//   - sk: MSKstakeye
//   - x: MSKrealt
//   - sl: MSKrealt
//   - su: MSKrealt
//
// [MSK_putconsolutioni]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putconsolutioni
func (task *Task) PutConsolutioni(
	i int32,
	whichsol SolType,
	sk StaKey,
	x float64,
	sl float64,
	su float64,
) res.Code {
	return res.Code(
		C.MSK_putconsolutioni(
			task.task,
			C.MSKint32t(i),
			C.MSKsoltypee(whichsol),
			C.MSKstakeye(sk),
			C.MSKrealt(x),
			C.MSKrealt(sl),
			C.MSKrealt(su),
		),
	)
}

// PutDjc is wrapping [MSK_putdjc] and
// sets the disjunctive constraint.
//
// [MSK_putdjc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - djcidx: MSKint64t
//   - numdomidx: MSKint64t
//   - domidxlist: const MSKint64t *
//   - numafeidx: MSKint64t
//   - afeidxlist: const MSKint64t *
//   - b: const MSKrealt *
//   - numterms: MSKint64t
//   - termsizelist: const MSKint64t *
//
// [MSK_putdjc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putdjc
func (task *Task) PutDjc(
	djcidx int64,
	numdomidx int64,
	domidxlist *int64,
	numafeidx int64,
	afeidxlist *int64,
	b *float64,
	numterms int64,
	termsizelist *int64,
) res.Code {
	return res.Code(
		C.MSK_putdjc(
			task.task,
			C.MSKint64t(djcidx),
			C.MSKint64t(numdomidx),
			(*C.MSKint64t)(domidxlist),
			C.MSKint64t(numafeidx),
			(*C.MSKint64t)(afeidxlist),
			(*C.MSKrealt)(b),
			C.MSKint64t(numterms),
			(*C.MSKint64t)(termsizelist),
		),
	)
}

// PutDouParam is wrapping [MSK_putdouparam] and
// sets a float point parameter.
//
// [MSK_putdouparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - param: MSKdparame
//   - parvalue: MSKrealt
//
// [MSK_putdouparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putdouparam
func (task *Task) PutDouParam(
	param DParam,
	parvalue float64,
) res.Code {
	return res.Code(
		C.MSK_putdouparam(
			task.task,
			C.MSKdparame(param),
			C.MSKrealt(parvalue),
		),
	)
}

// PutIntParam is wrapping [MSK_putintparam] and
// sets an integer parameter.
//
// [MSK_putintparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - param: MSKiparame
//   - parvalue: MSKint32t
//
// [MSK_putintparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putintparam
func (task *Task) PutIntParam(
	param IParam,
	parvalue int32,
) res.Code {
	return res.Code(
		C.MSK_putintparam(
			task.task,
			C.MSKiparame(param),
			C.MSKint32t(parvalue),
		),
	)
}

// PutNadouparam is wrapping [MSK_putnadouparam]
//
// [MSK_putnadouparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - paramname: const char *
//   - parvalue: MSKrealt
//
// [MSK_putnadouparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putnadouparam
func (task *Task) PutNadouparam(
	paramname string,
	parvalue float64,
) res.Code {
	c_paramname := C.CString(paramname)
	defer C.free(unsafe.Pointer(c_paramname))

	return res.Code(
		C.MSK_putnadouparam(
			task.task,
			c_paramname,
			C.MSKrealt(parvalue),
		),
	)
}

// PutNaintparam is wrapping [MSK_putnaintparam]
//
// [MSK_putnaintparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - paramname: const char *
//   - parvalue: MSKint32t
//
// [MSK_putnaintparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putnaintparam
func (task *Task) PutNaintparam(
	paramname string,
	parvalue int32,
) res.Code {
	c_paramname := C.CString(paramname)
	defer C.free(unsafe.Pointer(c_paramname))

	return res.Code(
		C.MSK_putnaintparam(
			task.task,
			c_paramname,
			C.MSKint32t(parvalue),
		),
	)
}

// PutNastrparam is wrapping [MSK_putnastrparam]
//
// [MSK_putnastrparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - paramname: const char *
//   - parvalue: const char *
//
// [MSK_putnastrparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putnastrparam
func (task *Task) PutNastrparam(
	paramname string,
	parvalue string,
) res.Code {
	c_paramname := C.CString(paramname)
	defer C.free(unsafe.Pointer(c_paramname))

	c_parvalue := C.CString(parvalue)
	defer C.free(unsafe.Pointer(c_parvalue))

	return res.Code(
		C.MSK_putnastrparam(
			task.task,
			c_paramname,
			c_parvalue,
		),
	)
}

// PutObjsense is wrapping [MSK_putobjsense] and
// set the objective sense - which is either minimize or maximize
//
// [MSK_putobjsense] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - sense: MSKobjsensee
//
// [MSK_putobjsense]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putobjsense
func (task *Task) PutObjsense(
	sense ObjectiveSense,
) res.Code {
	return res.Code(
		C.MSK_putobjsense(
			task.task,
			C.MSKobjsensee(sense),
		),
	)
}

// PutOptserverhost is wrapping [MSK_putoptserverhost]
//
// [MSK_putoptserverhost] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - host: const char *
//
// [MSK_putoptserverhost]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putoptserverhost
func (task *Task) PutOptserverhost(
	host string,
) res.Code {
	c_host := C.CString(host)
	defer C.free(unsafe.Pointer(c_host))

	return res.Code(
		C.MSK_putoptserverhost(
			task.task,
			c_host,
		),
	)
}

// PutParam is wrapping [MSK_putparam]
//
// [MSK_putparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - parname: const char *
//   - parvalue: const char *
//
// [MSK_putparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putparam
func (task *Task) PutParam(
	parname string,
	parvalue string,
) res.Code {
	c_parname := C.CString(parname)
	defer C.free(unsafe.Pointer(c_parname))

	c_parvalue := C.CString(parvalue)
	defer C.free(unsafe.Pointer(c_parvalue))

	return res.Code(
		C.MSK_putparam(
			task.task,
			c_parname,
			c_parvalue,
		),
	)
}

// PutQcon is wrapping [MSK_putqcon]
//
// [MSK_putqcon] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - numqcnz: MSKint32t
//   - qcsubk: const MSKint32t *
//   - qcsubi: const MSKint32t *
//   - qcsubj: const MSKint32t *
//   - qcval: const MSKrealt *
//
// [MSK_putqcon]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putqcon
func (task *Task) PutQcon(
	numqcnz int32,
	qcsubk *int32,
	qcsubi *int32,
	qcsubj *int32,
	qcval *float64,
) res.Code {
	return res.Code(
		C.MSK_putqcon(
			task.task,
			C.MSKint32t(numqcnz),
			(*C.MSKint32t)(qcsubk),
			(*C.MSKint32t)(qcsubi),
			(*C.MSKint32t)(qcsubj),
			(*C.MSKrealt)(qcval),
		),
	)
}

// PutQConK is wrapping [MSK_putqconk] and
// sets the quandratic constraint's coefficients.
//
// [MSK_putqconk] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - k: MSKint32t
//   - numqcnz: MSKint32t
//   - qcsubi: const MSKint32t *
//   - qcsubj: const MSKint32t *
//   - qcval: const MSKrealt *
//
// [MSK_putqconk]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putqconk
func (task *Task) PutQConK(
	k int32,
	numqcnz int32,
	qcsubi *int32,
	qcsubj *int32,
	qcval *float64,
) res.Code {
	return res.Code(
		C.MSK_putqconk(
			task.task,
			C.MSKint32t(k),
			C.MSKint32t(numqcnz),
			(*C.MSKint32t)(qcsubi),
			(*C.MSKint32t)(qcsubj),
			(*C.MSKrealt)(qcval),
		),
	)
}

// PutQObj is wrapping [MSK_putqobj] and
// sets the cofficient for quadratic objective term.
//
// [MSK_putqobj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - numqonz: MSKint32t
//   - qosubi: const MSKint32t *
//   - qosubj: const MSKint32t *
//   - qoval: const MSKrealt *
//
// [MSK_putqobj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putqobj
func (task *Task) PutQObj(
	numqonz int32,
	qosubi *int32,
	qosubj *int32,
	qoval *float64,
) res.Code {
	return res.Code(
		C.MSK_putqobj(
			task.task,
			C.MSKint32t(numqonz),
			(*C.MSKint32t)(qosubi),
			(*C.MSKint32t)(qosubj),
			(*C.MSKrealt)(qoval),
		),
	)
}

// PutQObjIj is wrapping [MSK_putqobjij]
//
// [MSK_putqobjij] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - j: MSKint32t
//   - qoij: MSKrealt
//
// [MSK_putqobjij]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putqobjij
func (task *Task) PutQObjIj(
	i int32,
	j int32,
	qoij float64,
) res.Code {
	return res.Code(
		C.MSK_putqobjij(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(j),
			C.MSKrealt(qoij),
		),
	)
}

// PutSkc is wrapping [MSK_putskc]
//
// [MSK_putskc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - skc: const MSKstakeye *
//
// [MSK_putskc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putskc
func (task *Task) PutSkc(
	whichsol SolType,
	skc *StaKey,
) res.Code {
	return res.Code(
		C.MSK_putskc(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKstakeye)(skc),
		),
	)
}

// PutSkx is wrapping [MSK_putskx]
//
// [MSK_putskx] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - skx: const MSKstakeye *
//
// [MSK_putskx]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putskx
func (task *Task) PutSkx(
	whichsol SolType,
	skx *StaKey,
) res.Code {
	return res.Code(
		C.MSK_putskx(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKstakeye)(skx),
		),
	)
}

// PutSlc is wrapping [MSK_putslc]
//
// [MSK_putslc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - slc: const MSKrealt *
//
// [MSK_putslc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putslc
func (task *Task) PutSlc(
	whichsol SolType,
	slc *float64,
) res.Code {
	return res.Code(
		C.MSK_putslc(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(slc),
		),
	)
}

// PutSlx is wrapping [MSK_putslx]
//
// [MSK_putslx] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - slx: const MSKrealt *
//
// [MSK_putslx]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putslx
func (task *Task) PutSlx(
	whichsol SolType,
	slx *float64,
) res.Code {
	return res.Code(
		C.MSK_putslx(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(slx),
		),
	)
}

// PutSnx is wrapping [MSK_putsnx]
//
// [MSK_putsnx] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - sux: const MSKrealt *
//
// [MSK_putsnx]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putsnx
func (task *Task) PutSnx(
	whichsol SolType,
	sux *float64,
) res.Code {
	return res.Code(
		C.MSK_putsnx(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(sux),
		),
	)
}

// PutSolution is wrapping [MSK_putsolution]
//
// [MSK_putsolution] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - skc: const MSKstakeye *
//   - skx: const MSKstakeye *
//   - skn: const MSKstakeye *
//   - xc: const MSKrealt *
//   - xx: const MSKrealt *
//   - y: const MSKrealt *
//   - slc: const MSKrealt *
//   - suc: const MSKrealt *
//   - slx: const MSKrealt *
//   - sux: const MSKrealt *
//   - snx: const MSKrealt *
//
// [MSK_putsolution]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putsolution
func (task *Task) PutSolution(
	whichsol SolType,
	skc *StaKey,
	skx *StaKey,
	skn *StaKey,
	xc *float64,
	xx *float64,
	y *float64,
	slc *float64,
	suc *float64,
	slx *float64,
	sux *float64,
	snx *float64,
) res.Code {
	return res.Code(
		C.MSK_putsolution(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKstakeye)(skc),
			(*C.MSKstakeye)(skx),
			(*C.MSKstakeye)(skn),
			(*C.MSKrealt)(xc),
			(*C.MSKrealt)(xx),
			(*C.MSKrealt)(y),
			(*C.MSKrealt)(slc),
			(*C.MSKrealt)(suc),
			(*C.MSKrealt)(slx),
			(*C.MSKrealt)(sux),
			(*C.MSKrealt)(snx),
		),
	)
}

// PutSolutionNew is wrapping [MSK_putsolutionnew]
//
// [MSK_putsolutionnew] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - skc: const MSKstakeye *
//   - skx: const MSKstakeye *
//   - skn: const MSKstakeye *
//   - xc: const MSKrealt *
//   - xx: const MSKrealt *
//   - y: const MSKrealt *
//   - slc: const MSKrealt *
//   - suc: const MSKrealt *
//   - slx: const MSKrealt *
//   - sux: const MSKrealt *
//   - snx: const MSKrealt *
//   - doty: const MSKrealt *
//
// [MSK_putsolutionnew]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putsolutionnew
func (task *Task) PutSolutionNew(
	whichsol SolType,
	skc *StaKey,
	skx *StaKey,
	skn *StaKey,
	xc *float64,
	xx *float64,
	y *float64,
	slc *float64,
	suc *float64,
	slx *float64,
	sux *float64,
	snx *float64,
	doty *float64,
) res.Code {
	return res.Code(
		C.MSK_putsolutionnew(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKstakeye)(skc),
			(*C.MSKstakeye)(skx),
			(*C.MSKstakeye)(skn),
			(*C.MSKrealt)(xc),
			(*C.MSKrealt)(xx),
			(*C.MSKrealt)(y),
			(*C.MSKrealt)(slc),
			(*C.MSKrealt)(suc),
			(*C.MSKrealt)(slx),
			(*C.MSKrealt)(sux),
			(*C.MSKrealt)(snx),
			(*C.MSKrealt)(doty),
		),
	)
}

// PutSolutionyi is wrapping [MSK_putsolutionyi]
//
// [MSK_putsolutionyi] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - i: MSKint32t
//   - whichsol: MSKsoltypee
//   - y: MSKrealt
//
// [MSK_putsolutionyi]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putsolutionyi
func (task *Task) PutSolutionyi(
	i int32,
	whichsol SolType,
	y float64,
) res.Code {
	return res.Code(
		C.MSK_putsolutionyi(
			task.task,
			C.MSKint32t(i),
			C.MSKsoltypee(whichsol),
			C.MSKrealt(y),
		),
	)
}

// PutStrparam is wrapping [MSK_putstrparam]
//
// [MSK_putstrparam] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - param: MSKsparame
//   - parvalue: const char *
//
// [MSK_putstrparam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putstrparam
func (task *Task) PutStrparam(
	param SParam,
	parvalue string,
) res.Code {
	c_parvalue := C.CString(parvalue)
	defer C.free(unsafe.Pointer(c_parvalue))

	return res.Code(
		C.MSK_putstrparam(
			task.task,
			C.MSKsparame(param),
			c_parvalue,
		),
	)
}

// PutSuc is wrapping [MSK_putsuc]
//
// [MSK_putsuc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - suc: const MSKrealt *
//
// [MSK_putsuc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putsuc
func (task *Task) PutSuc(
	whichsol SolType,
	suc *float64,
) res.Code {
	return res.Code(
		C.MSK_putsuc(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(suc),
		),
	)
}

// PutSux is wrapping [MSK_putsux]
//
// [MSK_putsux] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - sux: const MSKrealt *
//
// [MSK_putsux]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putsux
func (task *Task) PutSux(
	whichsol SolType,
	sux *float64,
) res.Code {
	return res.Code(
		C.MSK_putsux(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(sux),
		),
	)
}

// PutVarbound is wrapping [MSK_putvarbound] and
// set the bound for a variable.
//
// [MSK_putvarbound] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - j: MSKint32t
//   - bkx: MSKboundkeye
//   - blx: MSKrealt
//   - bux: MSKrealt
//
// [MSK_putvarbound]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putvarbound
func (task *Task) PutVarbound(
	j int32,
	bkx BoundKey,
	blx float64,
	bux float64,
) res.Code {
	return res.Code(
		C.MSK_putvarbound(
			task.task,
			C.MSKint32t(j),
			C.MSKboundkeye(bkx),
			C.MSKrealt(blx),
			C.MSKrealt(bux),
		),
	)
}

// PutVarboundListConst is wrapping [MSK_putvarboundlistconst]
//
// [MSK_putvarboundlistconst] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - num: MSKint32t
//   - sub: const MSKint32t *
//   - bkx: MSKboundkeye
//   - blx: MSKrealt
//   - bux: MSKrealt
//
// [MSK_putvarboundlistconst]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putvarboundlistconst
func (task *Task) PutVarboundListConst(
	num int32,
	sub *int32,
	bkx BoundKey,
	blx float64,
	bux float64,
) res.Code {
	return res.Code(
		C.MSK_putvarboundlistconst(
			task.task,
			C.MSKint32t(num),
			(*C.MSKint32t)(sub),
			C.MSKboundkeye(bkx),
			C.MSKrealt(blx),
			C.MSKrealt(bux),
		),
	)
}

// PutVarsolutionj is wrapping [MSK_putvarsolutionj]
//
// [MSK_putvarsolutionj] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - j: MSKint32t
//   - whichsol: MSKsoltypee
//   - sk: MSKstakeye
//   - x: MSKrealt
//   - sl: MSKrealt
//   - su: MSKrealt
//   - sn: MSKrealt
//
// [MSK_putvarsolutionj]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putvarsolutionj
func (task *Task) PutVarsolutionj(
	j int32,
	whichsol SolType,
	sk StaKey,
	x float64,
	sl float64,
	su float64,
	sn float64,
) res.Code {
	return res.Code(
		C.MSK_putvarsolutionj(
			task.task,
			C.MSKint32t(j),
			C.MSKsoltypee(whichsol),
			C.MSKstakeye(sk),
			C.MSKrealt(x),
			C.MSKrealt(sl),
			C.MSKrealt(su),
			C.MSKrealt(sn),
		),
	)
}

// PutVarType is wrapping [MSK_putvartype] and
// sets the type of the variable.
//
// [MSK_putvartype] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - j: MSKint32t
//   - vartype: MSKvariabletypee
//
// [MSK_putvartype]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putvartype
func (task *Task) PutVarType(
	j int32,
	vartype VariableType,
) res.Code {
	return res.Code(
		C.MSK_putvartype(
			task.task,
			C.MSKint32t(j),
			C.MSKvariabletypee(vartype),
		),
	)
}

// PutXc is wrapping [MSK_putxc]
//
// [MSK_putxc] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - xc: MSKrealt *
//
// [MSK_putxc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putxc
func (task *Task) PutXc(
	whichsol SolType,
	xc *float64,
) res.Code {
	return res.Code(
		C.MSK_putxc(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(xc),
		),
	)
}

// PutXx is wrapping [MSK_putxx]
//
// [MSK_putxx] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - xx: const MSKrealt *
//
// [MSK_putxx]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putxx
func (task *Task) PutXx(
	whichsol SolType,
	xx *float64,
) res.Code {
	return res.Code(
		C.MSK_putxx(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(xx),
		),
	)
}

// PutY is wrapping [MSK_puty]
//
// [MSK_puty] returns MSKrescodee and has following parameters
//   - task: MSKtask_t
//   - whichsol: MSKsoltypee
//   - y: const MSKrealt *
//
// [MSK_puty]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.puty
func (task *Task) PutY(
	whichsol SolType,
	y *float64,
) res.Code {
	return res.Code(
		C.MSK_puty(
			task.task,
			C.MSKsoltypee(whichsol),
			(*C.MSKrealt)(y),
		),
	)
}
