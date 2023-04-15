// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKboundkey_enum/BoundKey

package gmsk

import "strconv"

// BoundKey is MSKboundkey_enum.
//
// Bound keys
type BoundKey uint32

const (
	BK_LO BoundKey = 0 // The constraint or variable has a finite lower bound and an infinite upper bound.
	BK_UP BoundKey = 1 // The constraint or variable has an infinite lower bound and an finite upper bound.
	BK_FX BoundKey = 2 // The constraint or variable is fixed.
	BK_FR BoundKey = 3 // The constraint or variable is free.
	BK_RA BoundKey = 4 // The constraint or variable is ranged.
)

var _BoundKey_map = map[BoundKey]string{
	BK_LO: "BK_LO",
	BK_UP: "BK_UP",
	BK_FX: "BK_FX",
	BK_FR: "BK_FR",
	BK_RA: "BK_RA",
}

func (e BoundKey) String() string {
	if v, ok := _BoundKey_map[e]; ok {
		return v
	}
	return "BoundKey(" + strconv.FormatInt(int64(e), 10) + ")"
}
