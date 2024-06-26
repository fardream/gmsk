// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKsolsta_enum/SolSta

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// SolSta is MSKsolsta_enum.
//
// Solution status keys
type SolSta uint32

const (
	SOL_STA_UNKNOWN            SolSta = C.MSK_SOL_STA_UNKNOWN            // Status of the solution is unknown.
	SOL_STA_OPTIMAL            SolSta = C.MSK_SOL_STA_OPTIMAL            // The solution is optimal.
	SOL_STA_PRIM_FEAS          SolSta = C.MSK_SOL_STA_PRIM_FEAS          // The solution is primal feasible.
	SOL_STA_DUAL_FEAS          SolSta = C.MSK_SOL_STA_DUAL_FEAS          // The solution is dual feasible.
	SOL_STA_PRIM_AND_DUAL_FEAS SolSta = C.MSK_SOL_STA_PRIM_AND_DUAL_FEAS // The solution is both primal and dual feasible.
	SOL_STA_PRIM_INFEAS_CER    SolSta = C.MSK_SOL_STA_PRIM_INFEAS_CER    // The solution is a certificate of primal infeasibility.
	SOL_STA_DUAL_INFEAS_CER    SolSta = C.MSK_SOL_STA_DUAL_INFEAS_CER    // The solution is a certificate of dual infeasibility.
	SOL_STA_PRIM_ILLPOSED_CER  SolSta = C.MSK_SOL_STA_PRIM_ILLPOSED_CER  // The solution is a certificate that the primal problem is illposed.
	SOL_STA_DUAL_ILLPOSED_CER  SolSta = C.MSK_SOL_STA_DUAL_ILLPOSED_CER  // The solution is a certificate that the dual problem is illposed.
	SOL_STA_INTEGER_OPTIMAL    SolSta = C.MSK_SOL_STA_INTEGER_OPTIMAL    // The primal solution is integer optimal.
)

var _SolSta_map = map[SolSta]string{
	SOL_STA_UNKNOWN:            "SOL_STA_UNKNOWN",
	SOL_STA_OPTIMAL:            "SOL_STA_OPTIMAL",
	SOL_STA_PRIM_FEAS:          "SOL_STA_PRIM_FEAS",
	SOL_STA_DUAL_FEAS:          "SOL_STA_DUAL_FEAS",
	SOL_STA_PRIM_AND_DUAL_FEAS: "SOL_STA_PRIM_AND_DUAL_FEAS",
	SOL_STA_PRIM_INFEAS_CER:    "SOL_STA_PRIM_INFEAS_CER",
	SOL_STA_DUAL_INFEAS_CER:    "SOL_STA_DUAL_INFEAS_CER",
	SOL_STA_PRIM_ILLPOSED_CER:  "SOL_STA_PRIM_ILLPOSED_CER",
	SOL_STA_DUAL_ILLPOSED_CER:  "SOL_STA_DUAL_ILLPOSED_CER",
	SOL_STA_INTEGER_OPTIMAL:    "SOL_STA_INTEGER_OPTIMAL",
}

func (e SolSta) String() string {
	if v, ok := _SolSta_map[e]; ok {
		return v
	}
	return "SolSta(" + strconv.FormatInt(int64(e), 10) + ")"
}
