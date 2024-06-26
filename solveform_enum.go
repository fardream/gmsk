// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKsolveform_enum/Solveform

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// Solveform is MSKsolveform_enum.
//
// Solve primal or dual form
type Solveform uint32

const (
	SOLVE_FREE   Solveform = C.MSK_SOLVE_FREE   // The optimizer is free to solve either the primal or the dual problem.
	SOLVE_PRIMAL Solveform = C.MSK_SOLVE_PRIMAL // The optimizer should solve the primal problem.
	SOLVE_DUAL   Solveform = C.MSK_SOLVE_DUAL   // The optimizer should solve the dual problem.
)

var _Solveform_map = map[Solveform]string{
	SOLVE_FREE:   "SOLVE_FREE",
	SOLVE_PRIMAL: "SOLVE_PRIMAL",
	SOLVE_DUAL:   "SOLVE_DUAL",
}

func (e Solveform) String() string {
	if v, ok := _Solveform_map[e]; ok {
		return v
	}
	return "Solveform(" + strconv.FormatInt(int64(e), 10) + ")"
}
