// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKobjsense_enum/ObjectiveSense

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// ObjectiveSense is MSKobjsense_enum.
//
// Objective sense types
type ObjectiveSense uint32

const (
	OBJECTIVE_SENSE_MINIMIZE ObjectiveSense = C.MSK_OBJECTIVE_SENSE_MINIMIZE // The problem should be minimized.
	OBJECTIVE_SENSE_MAXIMIZE ObjectiveSense = C.MSK_OBJECTIVE_SENSE_MAXIMIZE // The problem should be maximized.
)

var _ObjectiveSense_map = map[ObjectiveSense]string{
	OBJECTIVE_SENSE_MINIMIZE: "OBJECTIVE_SENSE_MINIMIZE",
	OBJECTIVE_SENSE_MAXIMIZE: "OBJECTIVE_SENSE_MAXIMIZE",
}

func (e ObjectiveSense) String() string {
	if v, ok := _ObjectiveSense_map[e]; ok {
		return v
	}
	return "ObjectiveSense(" + strconv.FormatInt(int64(e), 10) + ")"
}
