// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKsimseltype_enum/SimSelType

package gmsk

import "strconv"

// SimSelType is MSKsimseltype_enum.
//
// Simplex selection strategy
type SimSelType uint32

const (
	SIM_SELECTION_FREE    SimSelType = 0 // The optimizer chooses the pricing strategy.
	SIM_SELECTION_FULL    SimSelType = 1 // The optimizer uses full pricing.
	SIM_SELECTION_ASE     SimSelType = 2 // The optimizer uses approximate steepest-edge pricing.
	SIM_SELECTION_DEVEX   SimSelType = 3 // The optimizer uses devex steepest-edge pricing.
	SIM_SELECTION_SE      SimSelType = 4 // The optimizer uses steepest-edge selection.
	SIM_SELECTION_PARTIAL SimSelType = 5 // The optimizer uses a partial selection approach.
)

var _SimSelType_map = map[SimSelType]string{
	SIM_SELECTION_FREE:    "SIM_SELECTION_FREE",
	SIM_SELECTION_FULL:    "SIM_SELECTION_FULL",
	SIM_SELECTION_ASE:     "SIM_SELECTION_ASE",
	SIM_SELECTION_DEVEX:   "SIM_SELECTION_DEVEX",
	SIM_SELECTION_SE:      "SIM_SELECTION_SE",
	SIM_SELECTION_PARTIAL: "SIM_SELECTION_PARTIAL",
}

func (e SimSelType) String() string {
	if v, ok := _SimSelType_map[e]; ok {
		return v
	}
	return "SimSelType(" + strconv.FormatInt(int64(e), 10) + ")"
}