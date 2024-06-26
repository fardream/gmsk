// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKsolitem_enum/SolItem

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// SolItem is MSKsolitem_enum.
//
// Solution items
type SolItem uint32

const (
	SOL_ITEM_XC  SolItem = C.MSK_SOL_ITEM_XC  // Solution for the constraints.
	SOL_ITEM_XX  SolItem = C.MSK_SOL_ITEM_XX  // Variable solution.
	SOL_ITEM_Y   SolItem = C.MSK_SOL_ITEM_Y   // Lagrange multipliers for equations.
	SOL_ITEM_SLC SolItem = C.MSK_SOL_ITEM_SLC // Lagrange multipliers for lower bounds on the constraints.
	SOL_ITEM_SUC SolItem = C.MSK_SOL_ITEM_SUC // Lagrange multipliers for upper bounds on the constraints.
	SOL_ITEM_SLX SolItem = C.MSK_SOL_ITEM_SLX // Lagrange multipliers for lower bounds on the variables.
	SOL_ITEM_SUX SolItem = C.MSK_SOL_ITEM_SUX // Lagrange multipliers for upper bounds on the variables.
	SOL_ITEM_SNX SolItem = C.MSK_SOL_ITEM_SNX // Lagrange multipliers corresponding to the conic constraints on the variables.
)

var _SolItem_map = map[SolItem]string{
	SOL_ITEM_XC:  "SOL_ITEM_XC",
	SOL_ITEM_XX:  "SOL_ITEM_XX",
	SOL_ITEM_Y:   "SOL_ITEM_Y",
	SOL_ITEM_SLC: "SOL_ITEM_SLC",
	SOL_ITEM_SUC: "SOL_ITEM_SUC",
	SOL_ITEM_SLX: "SOL_ITEM_SLX",
	SOL_ITEM_SUX: "SOL_ITEM_SUX",
	SOL_ITEM_SNX: "SOL_ITEM_SNX",
}

func (e SolItem) String() string {
	if v, ok := _SolItem_map[e]; ok {
		return v
	}
	return "SolItem(" + strconv.FormatInt(int64(e), 10) + ")"
}
