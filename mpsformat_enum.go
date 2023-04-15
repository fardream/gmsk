// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKmpsformat_enum/MPSFormat

package gmsk

import "strconv"

// MPSFormat is MSKmpsformat_enum.
//
// MPS file format type
type MPSFormat uint32

const (
	MPS_FORMAT_STRICT  MPSFormat = 0 // It is assumed that the input file satisfies the MPS format strictly.
	MPS_FORMAT_RELAXED MPSFormat = 1 // It is assumed that the input file satisfies a slightly relaxed version of the MPS format.
	MPS_FORMAT_FREE    MPSFormat = 2 // It is assumed that the input file satisfies the free MPS format. This implies that spaces are not allowed in names. Otherwise the format is free.
	MPS_FORMAT_CPLEX   MPSFormat = 3 // The CPLEX compatible version of the MPS format is employed.
)

var _MPSFormat_map = map[MPSFormat]string{
	MPS_FORMAT_STRICT:  "MPS_FORMAT_STRICT",
	MPS_FORMAT_RELAXED: "MPS_FORMAT_RELAXED",
	MPS_FORMAT_FREE:    "MPS_FORMAT_FREE",
	MPS_FORMAT_CPLEX:   "MPS_FORMAT_CPLEX",
}

func (e MPSFormat) String() string {
	if v, ok := _MPSFormat_map[e]; ok {
		return v
	}
	return "MPSFormat(" + strconv.FormatInt(int64(e), 10) + ")"
}
