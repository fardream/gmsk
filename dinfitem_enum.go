// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKdinfitem_enum/DInfItem

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// DInfItem is MSKdinfitem_enum.
//
// Double information items
type DInfItem uint32

const (
	DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY   DInfItem = C.MSK_DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY   // Density percentage of the scalarized constraint matrix.
	DINF_BI_CLEAN_DUAL_TIME                             DInfItem = C.MSK_DINF_BI_CLEAN_DUAL_TIME                             // Time  spent within the dual clean-up optimizer of the basis identification procedure since its invocation.
	DINF_BI_CLEAN_PRIMAL_TIME                           DInfItem = C.MSK_DINF_BI_CLEAN_PRIMAL_TIME                           // Time spent within the primal clean-up optimizer of the basis identification procedure since its invocation.
	DINF_BI_CLEAN_TIME                                  DInfItem = C.MSK_DINF_BI_CLEAN_TIME                                  // Time spent within the clean-up phase of the basis identification procedure since its invocation.
	DINF_BI_DUAL_TIME                                   DInfItem = C.MSK_DINF_BI_DUAL_TIME                                   // Time spent within the dual phase basis identification procedure since its invocation.
	DINF_BI_PRIMAL_TIME                                 DInfItem = C.MSK_DINF_BI_PRIMAL_TIME                                 // Time  spent within the primal phase of the basis identification procedure since its invocation.
	DINF_BI_TIME                                        DInfItem = C.MSK_DINF_BI_TIME                                        // Time spent within the basis identification procedure since its invocation.
	DINF_INTPNT_DUAL_FEAS                               DInfItem = C.MSK_DINF_INTPNT_DUAL_FEAS                               // Dual feasibility measure reported by the interior-point optimizer.
	DINF_INTPNT_DUAL_OBJ                                DInfItem = C.MSK_DINF_INTPNT_DUAL_OBJ                                // Dual objective value reported by the interior-point optimizer.
	DINF_INTPNT_FACTOR_NUM_FLOPS                        DInfItem = C.MSK_DINF_INTPNT_FACTOR_NUM_FLOPS                        // An estimate of the number of flops used in the factorization.
	DINF_INTPNT_OPT_STATUS                              DInfItem = C.MSK_DINF_INTPNT_OPT_STATUS                              // A measure of optimality of the solution.
	DINF_INTPNT_ORDER_TIME                              DInfItem = C.MSK_DINF_INTPNT_ORDER_TIME                              // Order time (in seconds).
	DINF_INTPNT_PRIMAL_FEAS                             DInfItem = C.MSK_DINF_INTPNT_PRIMAL_FEAS                             // Primal feasibility measure reported by the interior-point optimizer.
	DINF_INTPNT_PRIMAL_OBJ                              DInfItem = C.MSK_DINF_INTPNT_PRIMAL_OBJ                              // Primal objective value reported by the interior-point optimizer.
	DINF_INTPNT_TIME                                    DInfItem = C.MSK_DINF_INTPNT_TIME                                    // Time spent within the interior-point optimizer since its invocation.
	DINF_MIO_CLIQUE_SELECTION_TIME                      DInfItem = C.MSK_DINF_MIO_CLIQUE_SELECTION_TIME                      // Selection time for clique cuts.
	DINF_MIO_CLIQUE_SEPARATION_TIME                     DInfItem = C.MSK_DINF_MIO_CLIQUE_SEPARATION_TIME                     // Separation time for clique cuts.
	DINF_MIO_CMIR_SELECTION_TIME                        DInfItem = C.MSK_DINF_MIO_CMIR_SELECTION_TIME                        // Selection time for CMIR cuts.
	DINF_MIO_CMIR_SEPARATION_TIME                       DInfItem = C.MSK_DINF_MIO_CMIR_SEPARATION_TIME                       // Separation time for CMIR cuts.
	DINF_MIO_CONSTRUCT_SOLUTION_OBJ                     DInfItem = C.MSK_DINF_MIO_CONSTRUCT_SOLUTION_OBJ                     // Optimal objective value corresponding to the feasible solution.
	DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE                  DInfItem = C.MSK_DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE                  // Value of the dual bound after presolve but before cut generation.
	DINF_MIO_GMI_SELECTION_TIME                         DInfItem = C.MSK_DINF_MIO_GMI_SELECTION_TIME                         // Selection time for GMI cuts.
	DINF_MIO_GMI_SEPARATION_TIME                        DInfItem = C.MSK_DINF_MIO_GMI_SEPARATION_TIME                        // Separation time for GMI cuts.
	DINF_MIO_IMPLIED_BOUND_SELECTION_TIME               DInfItem = C.MSK_DINF_MIO_IMPLIED_BOUND_SELECTION_TIME               // Selection time for implied bound cuts.
	DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME              DInfItem = C.MSK_DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME              // Separation time for implied bound cuts.
	DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ              DInfItem = C.MSK_DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ              // Optimal objective value corresponding to the user provided initial solution.
	DINF_MIO_KNAPSACK_COVER_SELECTION_TIME              DInfItem = C.MSK_DINF_MIO_KNAPSACK_COVER_SELECTION_TIME              // Selection time for knapsack cover.
	DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME             DInfItem = C.MSK_DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME             // Separation time for knapsack cover.
	DINF_MIO_LIPRO_SELECTION_TIME                       DInfItem = C.MSK_DINF_MIO_LIPRO_SELECTION_TIME                       // Selection time for lift-and-project cuts.
	DINF_MIO_LIPRO_SEPARATION_TIME                      DInfItem = C.MSK_DINF_MIO_LIPRO_SEPARATION_TIME                      // Separation time for lift-and-project cuts.
	DINF_MIO_OBJ_ABS_GAP                                DInfItem = C.MSK_DINF_MIO_OBJ_ABS_GAP                                // If the mixed-integer optimizer has computed a feasible solution and a bound, this contains the absolute gap.
	DINF_MIO_OBJ_BOUND                                  DInfItem = C.MSK_DINF_MIO_OBJ_BOUND                                  // The best bound on the objective value known.
	DINF_MIO_OBJ_INT                                    DInfItem = C.MSK_DINF_MIO_OBJ_INT                                    // The primal objective value corresponding to the best integer feasible solution.
	DINF_MIO_OBJ_REL_GAP                                DInfItem = C.MSK_DINF_MIO_OBJ_REL_GAP                                // If the mixed-integer optimizer has computed a feasible solution and a bound, this contains the relative gap.
	DINF_MIO_PROBING_TIME                               DInfItem = C.MSK_DINF_MIO_PROBING_TIME                               // Total time for probing.
	DINF_MIO_ROOT_CUT_SELECTION_TIME                    DInfItem = C.MSK_DINF_MIO_ROOT_CUT_SELECTION_TIME                    // Total time for cut selection.
	DINF_MIO_ROOT_CUT_SEPARATION_TIME                   DInfItem = C.MSK_DINF_MIO_ROOT_CUT_SEPARATION_TIME                   // Total time for cut separation.
	DINF_MIO_ROOT_OPTIMIZER_TIME                        DInfItem = C.MSK_DINF_MIO_ROOT_OPTIMIZER_TIME                        // Time spent in the contiuous optimizer while processing the root node relaxation.
	DINF_MIO_ROOT_PRESOLVE_TIME                         DInfItem = C.MSK_DINF_MIO_ROOT_PRESOLVE_TIME                         // Time spent presolving the problem at the root node.
	DINF_MIO_ROOT_TIME                                  DInfItem = C.MSK_DINF_MIO_ROOT_TIME                                  // Time spent processing the root node.
	DINF_MIO_SYMMETRY_DETECTION_TIME                    DInfItem = C.MSK_DINF_MIO_SYMMETRY_DETECTION_TIME                    // Total time for symmetry detection.
	DINF_MIO_SYMMETRY_FACTOR                            DInfItem = C.MSK_DINF_MIO_SYMMETRY_FACTOR                            // Degree to which the problem is affected by detected symmetry.
	DINF_MIO_TIME                                       DInfItem = C.MSK_DINF_MIO_TIME                                       // Time spent in the mixed-integer optimizer.
	DINF_MIO_USER_OBJ_CUT                               DInfItem = C.MSK_DINF_MIO_USER_OBJ_CUT                               // If the objective cut is used, then this information item has the value of the cut.
	DINF_OPTIMIZER_TICKS                                DInfItem = C.MSK_DINF_OPTIMIZER_TICKS                                // Total number of ticks spent in the optimizer since it was invoked. It is strictly negative if it is not available.
	DINF_OPTIMIZER_TIME                                 DInfItem = C.MSK_DINF_OPTIMIZER_TIME                                 // Total time spent in the optimizer since it was invoked.
	DINF_PRESOLVE_ELI_TIME                              DInfItem = C.MSK_DINF_PRESOLVE_ELI_TIME                              // Total time spent in the eliminator since the presolve was invoked.
	DINF_PRESOLVE_LINDEP_TIME                           DInfItem = C.MSK_DINF_PRESOLVE_LINDEP_TIME                           // Total time spent  in the linear dependency checker since the presolve was invoked.
	DINF_PRESOLVE_TIME                                  DInfItem = C.MSK_DINF_PRESOLVE_TIME                                  // Total time (in seconds) spent in the presolve since it was invoked.
	DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION             DInfItem = C.MSK_DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION             // Total perturbation of the bounds of the primal problem.
	DINF_PRIMAL_REPAIR_PENALTY_OBJ                      DInfItem = C.MSK_DINF_PRIMAL_REPAIR_PENALTY_OBJ                      // The optimal objective value of the penalty function.
	DINF_QCQO_REFORMULATE_MAX_PERTURBATION              DInfItem = C.MSK_DINF_QCQO_REFORMULATE_MAX_PERTURBATION              // Maximum absolute diagonal perturbation occurring during the QCQO reformulation.
	DINF_QCQO_REFORMULATE_TIME                          DInfItem = C.MSK_DINF_QCQO_REFORMULATE_TIME                          // Time spent with conic quadratic reformulation.
	DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING DInfItem = C.MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING // Worst Cholesky column scaling.
	DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING   DInfItem = C.MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING   // Worst Cholesky diagonal scaling.
	DINF_READ_DATA_TIME                                 DInfItem = C.MSK_DINF_READ_DATA_TIME                                 // Time spent reading the data file.
	DINF_REMOTE_TIME                                    DInfItem = C.MSK_DINF_REMOTE_TIME                                    // The total real time in seconds spent when optimizing on a server by the process performing the optimization on the server
	DINF_SIM_DUAL_TIME                                  DInfItem = C.MSK_DINF_SIM_DUAL_TIME                                  // Time spent in the dual simplex optimizer since invoking it.
	DINF_SIM_FEAS                                       DInfItem = C.MSK_DINF_SIM_FEAS                                       // Feasibility measure reported by the simplex optimizer.
	DINF_SIM_OBJ                                        DInfItem = C.MSK_DINF_SIM_OBJ                                        // Objective value reported by the simplex optimizer.
	DINF_SIM_PRIMAL_TIME                                DInfItem = C.MSK_DINF_SIM_PRIMAL_TIME                                // Time spent in the primal simplex optimizer since invoking it.
	DINF_SIM_TIME                                       DInfItem = C.MSK_DINF_SIM_TIME                                       // Time spent in the simplex optimizer since invoking it.
	DINF_SOL_BAS_DUAL_OBJ                               DInfItem = C.MSK_DINF_SOL_BAS_DUAL_OBJ                               // Dual objective value of the basic solution. Updated by the function updatesolutioninfo.
	DINF_SOL_BAS_DVIOLCON                               DInfItem = C.MSK_DINF_SOL_BAS_DVIOLCON                               // Maximal dual bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
	DINF_SOL_BAS_DVIOLVAR                               DInfItem = C.MSK_DINF_SOL_BAS_DVIOLVAR                               // Maximal dual bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
	DINF_SOL_BAS_NRM_BARX                               DInfItem = C.MSK_DINF_SOL_BAS_NRM_BARX                               // Infinity norm of barx in the basic solution.
	DINF_SOL_BAS_NRM_SLC                                DInfItem = C.MSK_DINF_SOL_BAS_NRM_SLC                                // Infinity norm of slc in the basic solution.
	DINF_SOL_BAS_NRM_SLX                                DInfItem = C.MSK_DINF_SOL_BAS_NRM_SLX                                // Infinity norm of slx in the basic solution.
	DINF_SOL_BAS_NRM_SUC                                DInfItem = C.MSK_DINF_SOL_BAS_NRM_SUC                                // Infinity norm of suc in the basic solution.
	DINF_SOL_BAS_NRM_SUX                                DInfItem = C.MSK_DINF_SOL_BAS_NRM_SUX                                // Infinity norm of sux in the basic solution.
	DINF_SOL_BAS_NRM_XC                                 DInfItem = C.MSK_DINF_SOL_BAS_NRM_XC                                 // Infinity norm of xc in the basic solution.
	DINF_SOL_BAS_NRM_XX                                 DInfItem = C.MSK_DINF_SOL_BAS_NRM_XX                                 // Infinity norm of xx in the basic solution.
	DINF_SOL_BAS_NRM_Y                                  DInfItem = C.MSK_DINF_SOL_BAS_NRM_Y                                  // Infinity norm of Y in the basic solution.
	DINF_SOL_BAS_PRIMAL_OBJ                             DInfItem = C.MSK_DINF_SOL_BAS_PRIMAL_OBJ                             // Primal objective value of the basic solution. Updated by the function updatesolutioninfo.
	DINF_SOL_BAS_PVIOLCON                               DInfItem = C.MSK_DINF_SOL_BAS_PVIOLCON                               // Maximal primal bound violation for xc in the basic solution. Updated by the function updatesolutioninfo.
	DINF_SOL_BAS_PVIOLVAR                               DInfItem = C.MSK_DINF_SOL_BAS_PVIOLVAR                               // Maximal primal bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_NRM_BARX                               DInfItem = C.MSK_DINF_SOL_ITG_NRM_BARX                               // Infinity norm of barx in the integer solution.
	DINF_SOL_ITG_NRM_XC                                 DInfItem = C.MSK_DINF_SOL_ITG_NRM_XC                                 // Infinity norm of xc in the integer solution.
	DINF_SOL_ITG_NRM_XX                                 DInfItem = C.MSK_DINF_SOL_ITG_NRM_XX                                 // Infinity norm of xx in the integer solution.
	DINF_SOL_ITG_PRIMAL_OBJ                             DInfItem = C.MSK_DINF_SOL_ITG_PRIMAL_OBJ                             // Primal objective value of the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLACC                               DInfItem = C.MSK_DINF_SOL_ITG_PVIOLACC                               // Maximal primal violation for affine conic constraints in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLBARVAR                            DInfItem = C.MSK_DINF_SOL_ITG_PVIOLBARVAR                            // Maximal primal bound violation for barx in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLCON                               DInfItem = C.MSK_DINF_SOL_ITG_PVIOLCON                               // Maximal primal bound violation for xc in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLCONES                             DInfItem = C.MSK_DINF_SOL_ITG_PVIOLCONES                             // Maximal primal violation for primal conic constraints in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLDJC                               DInfItem = C.MSK_DINF_SOL_ITG_PVIOLDJC                               // Maximal primal violation for disjunctive constraints in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLITG                               DInfItem = C.MSK_DINF_SOL_ITG_PVIOLITG                               // Maximal violation for the integer constraints in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITG_PVIOLVAR                               DInfItem = C.MSK_DINF_SOL_ITG_PVIOLVAR                               // Maximal primal bound violation for xx in the integer solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_DUAL_OBJ                               DInfItem = C.MSK_DINF_SOL_ITR_DUAL_OBJ                               // Dual objective value of the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_DVIOLACC                               DInfItem = C.MSK_DINF_SOL_ITR_DVIOLACC                               // Maximal dual violation for affine conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_DVIOLBARVAR                            DInfItem = C.MSK_DINF_SOL_ITR_DVIOLBARVAR                            // Maximal dual bound violation for barx in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_DVIOLCON                               DInfItem = C.MSK_DINF_SOL_ITR_DVIOLCON                               // Maximal dual bound violation for xc in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_DVIOLCONES                             DInfItem = C.MSK_DINF_SOL_ITR_DVIOLCONES                             // Maximal dual violation for conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_DVIOLVAR                               DInfItem = C.MSK_DINF_SOL_ITR_DVIOLVAR                               // Maximal dual bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_NRM_BARS                               DInfItem = C.MSK_DINF_SOL_ITR_NRM_BARS                               // Infinity norm of bars in the interior-point solution.
	DINF_SOL_ITR_NRM_BARX                               DInfItem = C.MSK_DINF_SOL_ITR_NRM_BARX                               // Infinity norm of barx in the interior-point solution.
	DINF_SOL_ITR_NRM_SLC                                DInfItem = C.MSK_DINF_SOL_ITR_NRM_SLC                                // Infinity norm of slc in the interior-point solution.
	DINF_SOL_ITR_NRM_SLX                                DInfItem = C.MSK_DINF_SOL_ITR_NRM_SLX                                // Infinity norm of slx in the interior-point solution.
	DINF_SOL_ITR_NRM_SNX                                DInfItem = C.MSK_DINF_SOL_ITR_NRM_SNX                                // Infinity norm of snx in the interior-point solution.
	DINF_SOL_ITR_NRM_SUC                                DInfItem = C.MSK_DINF_SOL_ITR_NRM_SUC                                // Infinity norm of suc in the interior-point solution.
	DINF_SOL_ITR_NRM_SUX                                DInfItem = C.MSK_DINF_SOL_ITR_NRM_SUX                                // Infinity norm of sux in the interior-point solution.
	DINF_SOL_ITR_NRM_XC                                 DInfItem = C.MSK_DINF_SOL_ITR_NRM_XC                                 // Infinity norm of xc in the interior-point solution.
	DINF_SOL_ITR_NRM_XX                                 DInfItem = C.MSK_DINF_SOL_ITR_NRM_XX                                 // Infinity norm of xx in the interior-point solution.
	DINF_SOL_ITR_NRM_Y                                  DInfItem = C.MSK_DINF_SOL_ITR_NRM_Y                                  // Infinity norm of Y in the interior-point solution.
	DINF_SOL_ITR_PRIMAL_OBJ                             DInfItem = C.MSK_DINF_SOL_ITR_PRIMAL_OBJ                             // Primal objective value of the interior-point solution.
	DINF_SOL_ITR_PVIOLACC                               DInfItem = C.MSK_DINF_SOL_ITR_PVIOLACC                               // Maximal primal violation for affine conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_PVIOLBARVAR                            DInfItem = C.MSK_DINF_SOL_ITR_PVIOLBARVAR                            // Maximal primal bound violation for barx in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_PVIOLCON                               DInfItem = C.MSK_DINF_SOL_ITR_PVIOLCON                               // Maximal primal bound violation for xc in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_PVIOLCONES                             DInfItem = C.MSK_DINF_SOL_ITR_PVIOLCONES                             // Maximal primal violation for conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_SOL_ITR_PVIOLVAR                               DInfItem = C.MSK_DINF_SOL_ITR_PVIOLVAR                               // Maximal primal bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
	DINF_TO_CONIC_TIME                                  DInfItem = C.MSK_DINF_TO_CONIC_TIME                                  // Time spent in the last to conic reformulation.
	DINF_WRITE_DATA_TIME                                DInfItem = C.MSK_DINF_WRITE_DATA_TIME                                // Time spent writing the data file.
)

var _DInfItem_map = map[DInfItem]string{
	DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY:   "DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY",
	DINF_BI_CLEAN_DUAL_TIME:                             "DINF_BI_CLEAN_DUAL_TIME",
	DINF_BI_CLEAN_PRIMAL_TIME:                           "DINF_BI_CLEAN_PRIMAL_TIME",
	DINF_BI_CLEAN_TIME:                                  "DINF_BI_CLEAN_TIME",
	DINF_BI_DUAL_TIME:                                   "DINF_BI_DUAL_TIME",
	DINF_BI_PRIMAL_TIME:                                 "DINF_BI_PRIMAL_TIME",
	DINF_BI_TIME:                                        "DINF_BI_TIME",
	DINF_INTPNT_DUAL_FEAS:                               "DINF_INTPNT_DUAL_FEAS",
	DINF_INTPNT_DUAL_OBJ:                                "DINF_INTPNT_DUAL_OBJ",
	DINF_INTPNT_FACTOR_NUM_FLOPS:                        "DINF_INTPNT_FACTOR_NUM_FLOPS",
	DINF_INTPNT_OPT_STATUS:                              "DINF_INTPNT_OPT_STATUS",
	DINF_INTPNT_ORDER_TIME:                              "DINF_INTPNT_ORDER_TIME",
	DINF_INTPNT_PRIMAL_FEAS:                             "DINF_INTPNT_PRIMAL_FEAS",
	DINF_INTPNT_PRIMAL_OBJ:                              "DINF_INTPNT_PRIMAL_OBJ",
	DINF_INTPNT_TIME:                                    "DINF_INTPNT_TIME",
	DINF_MIO_CLIQUE_SELECTION_TIME:                      "DINF_MIO_CLIQUE_SELECTION_TIME",
	DINF_MIO_CLIQUE_SEPARATION_TIME:                     "DINF_MIO_CLIQUE_SEPARATION_TIME",
	DINF_MIO_CMIR_SELECTION_TIME:                        "DINF_MIO_CMIR_SELECTION_TIME",
	DINF_MIO_CMIR_SEPARATION_TIME:                       "DINF_MIO_CMIR_SEPARATION_TIME",
	DINF_MIO_CONSTRUCT_SOLUTION_OBJ:                     "DINF_MIO_CONSTRUCT_SOLUTION_OBJ",
	DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE:                  "DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE",
	DINF_MIO_GMI_SELECTION_TIME:                         "DINF_MIO_GMI_SELECTION_TIME",
	DINF_MIO_GMI_SEPARATION_TIME:                        "DINF_MIO_GMI_SEPARATION_TIME",
	DINF_MIO_IMPLIED_BOUND_SELECTION_TIME:               "DINF_MIO_IMPLIED_BOUND_SELECTION_TIME",
	DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME:              "DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME",
	DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ:              "DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ",
	DINF_MIO_KNAPSACK_COVER_SELECTION_TIME:              "DINF_MIO_KNAPSACK_COVER_SELECTION_TIME",
	DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME:             "DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME",
	DINF_MIO_LIPRO_SELECTION_TIME:                       "DINF_MIO_LIPRO_SELECTION_TIME",
	DINF_MIO_LIPRO_SEPARATION_TIME:                      "DINF_MIO_LIPRO_SEPARATION_TIME",
	DINF_MIO_OBJ_ABS_GAP:                                "DINF_MIO_OBJ_ABS_GAP",
	DINF_MIO_OBJ_BOUND:                                  "DINF_MIO_OBJ_BOUND",
	DINF_MIO_OBJ_INT:                                    "DINF_MIO_OBJ_INT",
	DINF_MIO_OBJ_REL_GAP:                                "DINF_MIO_OBJ_REL_GAP",
	DINF_MIO_PROBING_TIME:                               "DINF_MIO_PROBING_TIME",
	DINF_MIO_ROOT_CUT_SELECTION_TIME:                    "DINF_MIO_ROOT_CUT_SELECTION_TIME",
	DINF_MIO_ROOT_CUT_SEPARATION_TIME:                   "DINF_MIO_ROOT_CUT_SEPARATION_TIME",
	DINF_MIO_ROOT_OPTIMIZER_TIME:                        "DINF_MIO_ROOT_OPTIMIZER_TIME",
	DINF_MIO_ROOT_PRESOLVE_TIME:                         "DINF_MIO_ROOT_PRESOLVE_TIME",
	DINF_MIO_ROOT_TIME:                                  "DINF_MIO_ROOT_TIME",
	DINF_MIO_SYMMETRY_DETECTION_TIME:                    "DINF_MIO_SYMMETRY_DETECTION_TIME",
	DINF_MIO_SYMMETRY_FACTOR:                            "DINF_MIO_SYMMETRY_FACTOR",
	DINF_MIO_TIME:                                       "DINF_MIO_TIME",
	DINF_MIO_USER_OBJ_CUT:                               "DINF_MIO_USER_OBJ_CUT",
	DINF_OPTIMIZER_TICKS:                                "DINF_OPTIMIZER_TICKS",
	DINF_OPTIMIZER_TIME:                                 "DINF_OPTIMIZER_TIME",
	DINF_PRESOLVE_ELI_TIME:                              "DINF_PRESOLVE_ELI_TIME",
	DINF_PRESOLVE_LINDEP_TIME:                           "DINF_PRESOLVE_LINDEP_TIME",
	DINF_PRESOLVE_TIME:                                  "DINF_PRESOLVE_TIME",
	DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION:             "DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION",
	DINF_PRIMAL_REPAIR_PENALTY_OBJ:                      "DINF_PRIMAL_REPAIR_PENALTY_OBJ",
	DINF_QCQO_REFORMULATE_MAX_PERTURBATION:              "DINF_QCQO_REFORMULATE_MAX_PERTURBATION",
	DINF_QCQO_REFORMULATE_TIME:                          "DINF_QCQO_REFORMULATE_TIME",
	DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING: "DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING",
	DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING:   "DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING",
	DINF_READ_DATA_TIME:                                 "DINF_READ_DATA_TIME",
	DINF_REMOTE_TIME:                                    "DINF_REMOTE_TIME",
	DINF_SIM_DUAL_TIME:                                  "DINF_SIM_DUAL_TIME",
	DINF_SIM_FEAS:                                       "DINF_SIM_FEAS",
	DINF_SIM_OBJ:                                        "DINF_SIM_OBJ",
	DINF_SIM_PRIMAL_TIME:                                "DINF_SIM_PRIMAL_TIME",
	DINF_SIM_TIME:                                       "DINF_SIM_TIME",
	DINF_SOL_BAS_DUAL_OBJ:                               "DINF_SOL_BAS_DUAL_OBJ",
	DINF_SOL_BAS_DVIOLCON:                               "DINF_SOL_BAS_DVIOLCON",
	DINF_SOL_BAS_DVIOLVAR:                               "DINF_SOL_BAS_DVIOLVAR",
	DINF_SOL_BAS_NRM_BARX:                               "DINF_SOL_BAS_NRM_BARX",
	DINF_SOL_BAS_NRM_SLC:                                "DINF_SOL_BAS_NRM_SLC",
	DINF_SOL_BAS_NRM_SLX:                                "DINF_SOL_BAS_NRM_SLX",
	DINF_SOL_BAS_NRM_SUC:                                "DINF_SOL_BAS_NRM_SUC",
	DINF_SOL_BAS_NRM_SUX:                                "DINF_SOL_BAS_NRM_SUX",
	DINF_SOL_BAS_NRM_XC:                                 "DINF_SOL_BAS_NRM_XC",
	DINF_SOL_BAS_NRM_XX:                                 "DINF_SOL_BAS_NRM_XX",
	DINF_SOL_BAS_NRM_Y:                                  "DINF_SOL_BAS_NRM_Y",
	DINF_SOL_BAS_PRIMAL_OBJ:                             "DINF_SOL_BAS_PRIMAL_OBJ",
	DINF_SOL_BAS_PVIOLCON:                               "DINF_SOL_BAS_PVIOLCON",
	DINF_SOL_BAS_PVIOLVAR:                               "DINF_SOL_BAS_PVIOLVAR",
	DINF_SOL_ITG_NRM_BARX:                               "DINF_SOL_ITG_NRM_BARX",
	DINF_SOL_ITG_NRM_XC:                                 "DINF_SOL_ITG_NRM_XC",
	DINF_SOL_ITG_NRM_XX:                                 "DINF_SOL_ITG_NRM_XX",
	DINF_SOL_ITG_PRIMAL_OBJ:                             "DINF_SOL_ITG_PRIMAL_OBJ",
	DINF_SOL_ITG_PVIOLACC:                               "DINF_SOL_ITG_PVIOLACC",
	DINF_SOL_ITG_PVIOLBARVAR:                            "DINF_SOL_ITG_PVIOLBARVAR",
	DINF_SOL_ITG_PVIOLCON:                               "DINF_SOL_ITG_PVIOLCON",
	DINF_SOL_ITG_PVIOLCONES:                             "DINF_SOL_ITG_PVIOLCONES",
	DINF_SOL_ITG_PVIOLDJC:                               "DINF_SOL_ITG_PVIOLDJC",
	DINF_SOL_ITG_PVIOLITG:                               "DINF_SOL_ITG_PVIOLITG",
	DINF_SOL_ITG_PVIOLVAR:                               "DINF_SOL_ITG_PVIOLVAR",
	DINF_SOL_ITR_DUAL_OBJ:                               "DINF_SOL_ITR_DUAL_OBJ",
	DINF_SOL_ITR_DVIOLACC:                               "DINF_SOL_ITR_DVIOLACC",
	DINF_SOL_ITR_DVIOLBARVAR:                            "DINF_SOL_ITR_DVIOLBARVAR",
	DINF_SOL_ITR_DVIOLCON:                               "DINF_SOL_ITR_DVIOLCON",
	DINF_SOL_ITR_DVIOLCONES:                             "DINF_SOL_ITR_DVIOLCONES",
	DINF_SOL_ITR_DVIOLVAR:                               "DINF_SOL_ITR_DVIOLVAR",
	DINF_SOL_ITR_NRM_BARS:                               "DINF_SOL_ITR_NRM_BARS",
	DINF_SOL_ITR_NRM_BARX:                               "DINF_SOL_ITR_NRM_BARX",
	DINF_SOL_ITR_NRM_SLC:                                "DINF_SOL_ITR_NRM_SLC",
	DINF_SOL_ITR_NRM_SLX:                                "DINF_SOL_ITR_NRM_SLX",
	DINF_SOL_ITR_NRM_SNX:                                "DINF_SOL_ITR_NRM_SNX",
	DINF_SOL_ITR_NRM_SUC:                                "DINF_SOL_ITR_NRM_SUC",
	DINF_SOL_ITR_NRM_SUX:                                "DINF_SOL_ITR_NRM_SUX",
	DINF_SOL_ITR_NRM_XC:                                 "DINF_SOL_ITR_NRM_XC",
	DINF_SOL_ITR_NRM_XX:                                 "DINF_SOL_ITR_NRM_XX",
	DINF_SOL_ITR_NRM_Y:                                  "DINF_SOL_ITR_NRM_Y",
	DINF_SOL_ITR_PRIMAL_OBJ:                             "DINF_SOL_ITR_PRIMAL_OBJ",
	DINF_SOL_ITR_PVIOLACC:                               "DINF_SOL_ITR_PVIOLACC",
	DINF_SOL_ITR_PVIOLBARVAR:                            "DINF_SOL_ITR_PVIOLBARVAR",
	DINF_SOL_ITR_PVIOLCON:                               "DINF_SOL_ITR_PVIOLCON",
	DINF_SOL_ITR_PVIOLCONES:                             "DINF_SOL_ITR_PVIOLCONES",
	DINF_SOL_ITR_PVIOLVAR:                               "DINF_SOL_ITR_PVIOLVAR",
	DINF_TO_CONIC_TIME:                                  "DINF_TO_CONIC_TIME",
	DINF_WRITE_DATA_TIME:                                "DINF_WRITE_DATA_TIME",
}

func (e DInfItem) String() string {
	if v, ok := _DInfItem_map[e]; ok {
		return v
	}
	return "DInfItem(" + strconv.FormatInt(int64(e), 10) + ")"
}
