// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKmionodeseltype_enum/MioNodeSelType

package gmsk

import "strconv"

// MioNodeSelType is MSKmionodeseltype_enum.
//
// Mixed-integer node selection types
type MioNodeSelType uint32

const (
	MIO_NODE_SELECTION_FREE   MioNodeSelType = 0 // The optimizer decides the node selection strategy.
	MIO_NODE_SELECTION_FIRST  MioNodeSelType = 1 // The optimizer employs a depth first node selection strategy.
	MIO_NODE_SELECTION_BEST   MioNodeSelType = 2 // The optimizer employs a best bound node selection strategy.
	MIO_NODE_SELECTION_PSEUDO MioNodeSelType = 3 // The optimizer employs selects the node based on a pseudo cost estimate.
)

var _MioNodeSelType_map = map[MioNodeSelType]string{
	MIO_NODE_SELECTION_FREE:   "MIO_NODE_SELECTION_FREE",
	MIO_NODE_SELECTION_FIRST:  "MIO_NODE_SELECTION_FIRST",
	MIO_NODE_SELECTION_BEST:   "MIO_NODE_SELECTION_BEST",
	MIO_NODE_SELECTION_PSEUDO: "MIO_NODE_SELECTION_PSEUDO",
}

func (e MioNodeSelType) String() string {
	if v, ok := _MioNodeSelType_map[e]; ok {
		return v
	}
	return "MioNodeSelType(" + strconv.FormatInt(int64(e), 10) + ")"
}
