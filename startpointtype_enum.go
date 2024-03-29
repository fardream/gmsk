// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKstartpointtype_enum/StarPointType

package gmsk

import "strconv"

// StarPointType is MSKstartpointtype_enum.
//
// Starting point types
type StarPointType uint32

const (
	STARTING_POINT_FREE     StarPointType = 0 // The starting point is chosen automatically.
	STARTING_POINT_GUESS    StarPointType = 1 // The optimizer guesses a starting point.
	STARTING_POINT_CONSTANT StarPointType = 2 // The optimizer constructs a starting point by assigning a constant value to all primal and dual variables. This starting point is normally robust.
)

var _StarPointType_map = map[StarPointType]string{
	STARTING_POINT_FREE:     "STARTING_POINT_FREE",
	STARTING_POINT_GUESS:    "STARTING_POINT_GUESS",
	STARTING_POINT_CONSTANT: "STARTING_POINT_CONSTANT",
}

func (e StarPointType) String() string {
	if v, ok := _StarPointType_map[e]; ok {
		return v
	}
	return "StarPointType(" + strconv.FormatInt(int64(e), 10) + ")"
}
