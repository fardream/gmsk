// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKfeature_enum/Feature

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// Feature is MSKfeature_enum.
//
// License feature
type Feature uint32

const (
	FEATURE_PTS  Feature = C.MSK_FEATURE_PTS  // Base system.
	FEATURE_PTON Feature = C.MSK_FEATURE_PTON // Conic extension.
)

var _Feature_map = map[Feature]string{
	FEATURE_PTS:  "FEATURE_PTS",
	FEATURE_PTON: "FEATURE_PTON",
}

func (e Feature) String() string {
	if v, ok := _Feature_map[e]; ok {
		return v
	}
	return "Feature(" + strconv.FormatInt(int64(e), 10) + ")"
}
