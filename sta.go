package gmsk

// SolSta is the solution status
type SolSta uint32

const (
	SOL_STA_UNKNOWN            SolSta = 0
	SOL_STA_OPTIMAL            SolSta = 1 // Optimal solution
	SOL_STA_PRIM_FEAS          SolSta = 2 // Primal feasible, but may not be optimal
	SOL_STA_DUAL_FEAS          SolSta = 3
	SOL_STA_PRIM_AND_DUAL_FEAS SolSta = 4
	SOL_STA_PRIM_INFEAS_CER    SolSta = 5 // Primal infeasible
	SOL_STA_DUAL_INFEAS_CER    SolSta = 6 // Dual infeasible
	SOL_STA_PRIM_ILLPOSED_CER  SolSta = 7
	SOL_STA_DUAL_ILLPOSED_CER  SolSta = 8
	SOL_STA_INTEGER_OPTIMAL    SolSta = 9 // Optimal for integer programming
)

// ProSta is the problem status
type ProSta uint32

const (
	PRO_STA_UNKNOWN                  ProSta = 0
	PRO_STA_PRIM_AND_DUAL_FEAS       ProSta = 1
	PRO_STA_PRIM_FEAS                ProSta = 2
	PRO_STA_DUAL_FEAS                ProSta = 3
	PRO_STA_PRIM_INFEAS              ProSta = 4
	PRO_STA_DUAL_INFEAS              ProSta = 5
	PRO_STA_PRIM_AND_DUAL_INFEAS     ProSta = 6
	PRO_STA_ILL_POSED                ProSta = 7
	PRO_STA_PRIM_INFEAS_OR_UNBOUNDED ProSta = 8
)
