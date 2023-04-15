// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKproblemtype_enum/ProblemType

package gmsk

import "strconv"

// ProblemType is MSKproblemtype_enum.
//
// Problem types
type ProblemType uint32

const (
	PROBTYPE_LO    ProblemType = 0 // The problem is a linear optimization problem.
	PROBTYPE_QO    ProblemType = 1 // The problem is a quadratic optimization problem.
	PROBTYPE_QCQO  ProblemType = 2 // The problem is a quadratically constrained optimization problem.
	PROBTYPE_CONIC ProblemType = 3 // A conic optimization.
	PROBTYPE_MIXED ProblemType = 4 // General nonlinear constraints and conic constraints. This combination can not be solved by MOSEK.
)

var _ProblemType_map = map[ProblemType]string{
	PROBTYPE_LO:    "PROBTYPE_LO",
	PROBTYPE_QO:    "PROBTYPE_QO",
	PROBTYPE_QCQO:  "PROBTYPE_QCQO",
	PROBTYPE_CONIC: "PROBTYPE_CONIC",
	PROBTYPE_MIXED: "PROBTYPE_MIXED",
}

func (e ProblemType) String() string {
	if v, ok := _ProblemType_map[e]; ok {
		return v
	}
	return "ProblemType(" + strconv.FormatInt(int64(e), 10) + ")"
}