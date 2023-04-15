// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKscalingtype_enum/ScalingType

package gmsk

import "strconv"

// ScalingType is MSKscalingtype_enum.
//
// Scaling type
type ScalingType uint32

const (
	SCALING_FREE ScalingType = 0 // The optimizer chooses the scaling heuristic.
	SCALING_NONE ScalingType = 1 // No scaling is performed.
)

var _ScalingType_map = map[ScalingType]string{
	SCALING_FREE: "SCALING_FREE",
	SCALING_NONE: "SCALING_NONE",
}

func (e ScalingType) String() string {
	if v, ok := _ScalingType_map[e]; ok {
		return v
	}
	return "ScalingType(" + strconv.FormatInt(int64(e), 10) + ")"
}