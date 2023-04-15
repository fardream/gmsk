// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKmiocontsoltype_enum/MioContSolType

package gmsk

import "strconv"

// MioContSolType is MSKmiocontsoltype_enum.
//
// Continuous mixed-integer solution type
type MioContSolType uint32

const (
	MIO_CONT_SOL_NONE    MioContSolType = 0 // No interior-point or basic solution.
	MIO_CONT_SOL_ROOT    MioContSolType = 1 // Solutions to the root node problem.
	MIO_CONT_SOL_ITG     MioContSolType = 2 // A feasible primal solution.
	MIO_CONT_SOL_ITG_REL MioContSolType = 3 // A feasible primal solution or a root node solution if the problem is infeasible.
)

var _MioContSolType_map = map[MioContSolType]string{
	MIO_CONT_SOL_NONE:    "MIO_CONT_SOL_NONE",
	MIO_CONT_SOL_ROOT:    "MIO_CONT_SOL_ROOT",
	MIO_CONT_SOL_ITG:     "MIO_CONT_SOL_ITG",
	MIO_CONT_SOL_ITG_REL: "MIO_CONT_SOL_ITG_REL",
}

func (e MioContSolType) String() string {
	if v, ok := _MioContSolType_map[e]; ok {
		return v
	}
	return "MioContSolType(" + strconv.FormatInt(int64(e), 10) + ")"
}