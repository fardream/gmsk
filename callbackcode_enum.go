// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKcallbackcode_enum/CallbackCode

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// CallbackCode is MSKcallbackcode_enum.
//
// Progress callback codes
type CallbackCode uint32

const (
	CALLBACK_BEGIN_BI                 CallbackCode = C.MSK_CALLBACK_BEGIN_BI                 // The basis identification procedure has been started.
	CALLBACK_BEGIN_CONIC              CallbackCode = C.MSK_CALLBACK_BEGIN_CONIC              // The callback function is called when the conic optimizer is started.
	CALLBACK_BEGIN_DUAL_BI            CallbackCode = C.MSK_CALLBACK_BEGIN_DUAL_BI            // The callback function is called from within the basis identification procedure when the dual phase is started.
	CALLBACK_BEGIN_DUAL_SENSITIVITY   CallbackCode = C.MSK_CALLBACK_BEGIN_DUAL_SENSITIVITY   // Dual sensitivity analysis is started.
	CALLBACK_BEGIN_DUAL_SETUP_BI      CallbackCode = C.MSK_CALLBACK_BEGIN_DUAL_SETUP_BI      // The callback function is called when the dual BI phase is started.
	CALLBACK_BEGIN_DUAL_SIMPLEX       CallbackCode = C.MSK_CALLBACK_BEGIN_DUAL_SIMPLEX       // The callback function is called when the dual simplex optimizer started.
	CALLBACK_BEGIN_DUAL_SIMPLEX_BI    CallbackCode = C.MSK_CALLBACK_BEGIN_DUAL_SIMPLEX_BI    // The callback function is called from within the basis identification procedure when the dual simplex clean-up phase is started.
	CALLBACK_BEGIN_INFEAS_ANA         CallbackCode = C.MSK_CALLBACK_BEGIN_INFEAS_ANA         // The callback function is called when the infeasibility analyzer is started.
	CALLBACK_BEGIN_INTPNT             CallbackCode = C.MSK_CALLBACK_BEGIN_INTPNT             // The callback function is called when the interior-point optimizer is started.
	CALLBACK_BEGIN_LICENSE_WAIT       CallbackCode = C.MSK_CALLBACK_BEGIN_LICENSE_WAIT       // Begin waiting for license.
	CALLBACK_BEGIN_MIO                CallbackCode = C.MSK_CALLBACK_BEGIN_MIO                // The callback function is called when the mixed-integer optimizer is started.
	CALLBACK_BEGIN_OPTIMIZER          CallbackCode = C.MSK_CALLBACK_BEGIN_OPTIMIZER          // The callback function is called when the optimizer is started.
	CALLBACK_BEGIN_PRESOLVE           CallbackCode = C.MSK_CALLBACK_BEGIN_PRESOLVE           // The callback function is called when the presolve is started.
	CALLBACK_BEGIN_PRIMAL_BI          CallbackCode = C.MSK_CALLBACK_BEGIN_PRIMAL_BI          // The callback function is called from within the basis identification procedure when the primal phase is started.
	CALLBACK_BEGIN_PRIMAL_REPAIR      CallbackCode = C.MSK_CALLBACK_BEGIN_PRIMAL_REPAIR      // Begin primal feasibility repair.
	CALLBACK_BEGIN_PRIMAL_SENSITIVITY CallbackCode = C.MSK_CALLBACK_BEGIN_PRIMAL_SENSITIVITY // Primal sensitivity analysis is started.
	CALLBACK_BEGIN_PRIMAL_SETUP_BI    CallbackCode = C.MSK_CALLBACK_BEGIN_PRIMAL_SETUP_BI    // The callback function is called when the primal BI setup is started.
	CALLBACK_BEGIN_PRIMAL_SIMPLEX     CallbackCode = C.MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX     // The callback function is called when the primal simplex optimizer is started.
	CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI  CallbackCode = C.MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI  // The callback function is called from within the basis identification procedure when the primal simplex clean-up phase is started.
	CALLBACK_BEGIN_QCQO_REFORMULATE   CallbackCode = C.MSK_CALLBACK_BEGIN_QCQO_REFORMULATE   // Begin QCQO reformulation.
	CALLBACK_BEGIN_READ               CallbackCode = C.MSK_CALLBACK_BEGIN_READ               // MOSEK has started reading a problem file.
	CALLBACK_BEGIN_ROOT_CUTGEN        CallbackCode = C.MSK_CALLBACK_BEGIN_ROOT_CUTGEN        // The callback function is called when root cut generation is started.
	CALLBACK_BEGIN_SIMPLEX            CallbackCode = C.MSK_CALLBACK_BEGIN_SIMPLEX            // The callback function is called when the simplex optimizer is started.
	CALLBACK_BEGIN_SIMPLEX_BI         CallbackCode = C.MSK_CALLBACK_BEGIN_SIMPLEX_BI         // The callback function is called from within the basis identification procedure when the simplex clean-up phase is started.
	CALLBACK_BEGIN_SOLVE_ROOT_RELAX   CallbackCode = C.MSK_CALLBACK_BEGIN_SOLVE_ROOT_RELAX   // The callback function is called when solution of root relaxation is started.
	CALLBACK_BEGIN_TO_CONIC           CallbackCode = C.MSK_CALLBACK_BEGIN_TO_CONIC           // Begin conic reformulation.
	CALLBACK_BEGIN_WRITE              CallbackCode = C.MSK_CALLBACK_BEGIN_WRITE              // MOSEK has started writing a problem file.
	CALLBACK_CONIC                    CallbackCode = C.MSK_CALLBACK_CONIC                    // The callback function is called from within the conic optimizer after the information database has been updated.
	CALLBACK_DUAL_SIMPLEX             CallbackCode = C.MSK_CALLBACK_DUAL_SIMPLEX             // The callback function is called from within the dual simplex optimizer.
	CALLBACK_END_BI                   CallbackCode = C.MSK_CALLBACK_END_BI                   // The callback function is called when the basis identification procedure is terminated.
	CALLBACK_END_CONIC                CallbackCode = C.MSK_CALLBACK_END_CONIC                // The callback function is called when the conic optimizer is terminated.
	CALLBACK_END_DUAL_BI              CallbackCode = C.MSK_CALLBACK_END_DUAL_BI              // The callback function is called from within the basis identification procedure when the dual phase is terminated.
	CALLBACK_END_DUAL_SENSITIVITY     CallbackCode = C.MSK_CALLBACK_END_DUAL_SENSITIVITY     // Dual sensitivity analysis is terminated.
	CALLBACK_END_DUAL_SETUP_BI        CallbackCode = C.MSK_CALLBACK_END_DUAL_SETUP_BI        // The callback function is called when the dual BI phase is terminated.
	CALLBACK_END_DUAL_SIMPLEX         CallbackCode = C.MSK_CALLBACK_END_DUAL_SIMPLEX         // The callback function is called when the dual simplex optimizer is terminated.
	CALLBACK_END_DUAL_SIMPLEX_BI      CallbackCode = C.MSK_CALLBACK_END_DUAL_SIMPLEX_BI      // The callback function is called from within the basis identification procedure when the dual clean-up phase is terminated.
	CALLBACK_END_INFEAS_ANA           CallbackCode = C.MSK_CALLBACK_END_INFEAS_ANA           // The callback function is called when the infeasibility analyzer is terminated.
	CALLBACK_END_INTPNT               CallbackCode = C.MSK_CALLBACK_END_INTPNT               // The callback function is called when the interior-point optimizer is terminated.
	CALLBACK_END_LICENSE_WAIT         CallbackCode = C.MSK_CALLBACK_END_LICENSE_WAIT         // End waiting for license.
	CALLBACK_END_MIO                  CallbackCode = C.MSK_CALLBACK_END_MIO                  // The callback function is called when the mixed-integer optimizer is terminated.
	CALLBACK_END_OPTIMIZER            CallbackCode = C.MSK_CALLBACK_END_OPTIMIZER            // The callback function is called when the optimizer is terminated.
	CALLBACK_END_PRESOLVE             CallbackCode = C.MSK_CALLBACK_END_PRESOLVE             // The callback function is called when the presolve is completed.
	CALLBACK_END_PRIMAL_BI            CallbackCode = C.MSK_CALLBACK_END_PRIMAL_BI            // The callback function is called from within the basis identification procedure when the primal phase is terminated.
	CALLBACK_END_PRIMAL_REPAIR        CallbackCode = C.MSK_CALLBACK_END_PRIMAL_REPAIR        // End primal feasibility repair.
	CALLBACK_END_PRIMAL_SENSITIVITY   CallbackCode = C.MSK_CALLBACK_END_PRIMAL_SENSITIVITY   // Primal sensitivity analysis is terminated.
	CALLBACK_END_PRIMAL_SETUP_BI      CallbackCode = C.MSK_CALLBACK_END_PRIMAL_SETUP_BI      // The callback function is called when the primal BI setup is terminated.
	CALLBACK_END_PRIMAL_SIMPLEX       CallbackCode = C.MSK_CALLBACK_END_PRIMAL_SIMPLEX       // The callback function is called when the primal simplex optimizer is terminated.
	CALLBACK_END_PRIMAL_SIMPLEX_BI    CallbackCode = C.MSK_CALLBACK_END_PRIMAL_SIMPLEX_BI    // The callback function is called from within the basis identification procedure when the primal clean-up phase is terminated.
	CALLBACK_END_QCQO_REFORMULATE     CallbackCode = C.MSK_CALLBACK_END_QCQO_REFORMULATE     // End QCQO reformulation.
	CALLBACK_END_READ                 CallbackCode = C.MSK_CALLBACK_END_READ                 // MOSEK has finished reading a problem file.
	CALLBACK_END_ROOT_CUTGEN          CallbackCode = C.MSK_CALLBACK_END_ROOT_CUTGEN          // The callback function is called when root cut generation is terminated.
	CALLBACK_END_SIMPLEX              CallbackCode = C.MSK_CALLBACK_END_SIMPLEX              // The callback function is called when the simplex optimizer is terminated.
	CALLBACK_END_SIMPLEX_BI           CallbackCode = C.MSK_CALLBACK_END_SIMPLEX_BI           // The callback function is called from within the basis identification procedure when the simplex clean-up phase is terminated.
	CALLBACK_END_SOLVE_ROOT_RELAX     CallbackCode = C.MSK_CALLBACK_END_SOLVE_ROOT_RELAX     // The callback function is called when solution of root relaxation is terminated.
	CALLBACK_END_TO_CONIC             CallbackCode = C.MSK_CALLBACK_END_TO_CONIC             // End conic reformulation.
	CALLBACK_END_WRITE                CallbackCode = C.MSK_CALLBACK_END_WRITE                // MOSEK has finished writing a problem file.
	CALLBACK_IM_BI                    CallbackCode = C.MSK_CALLBACK_IM_BI                    // The callback function is called from within the basis identification procedure at an intermediate point.
	CALLBACK_IM_CONIC                 CallbackCode = C.MSK_CALLBACK_IM_CONIC                 // The callback function is called at an intermediate stage within the conic optimizer where the information database has not been updated.
	CALLBACK_IM_DUAL_BI               CallbackCode = C.MSK_CALLBACK_IM_DUAL_BI               // The callback function is called from within the basis identification procedure at an intermediate point in the dual phase.
	CALLBACK_IM_DUAL_SENSIVITY        CallbackCode = C.MSK_CALLBACK_IM_DUAL_SENSIVITY        // The callback function is called at an intermediate stage of the dual sensitivity analysis.
	CALLBACK_IM_DUAL_SIMPLEX          CallbackCode = C.MSK_CALLBACK_IM_DUAL_SIMPLEX          // The callback function is called at an intermediate point in the dual simplex optimizer.
	CALLBACK_IM_INTPNT                CallbackCode = C.MSK_CALLBACK_IM_INTPNT                // The callback function is called at an intermediate stage within the interior-point optimizer where the information database has not been updated.
	CALLBACK_IM_LICENSE_WAIT          CallbackCode = C.MSK_CALLBACK_IM_LICENSE_WAIT          // MOSEK is waiting for a license.
	CALLBACK_IM_LU                    CallbackCode = C.MSK_CALLBACK_IM_LU                    // The callback function is called from within the LU factorization procedure at an intermediate point.
	CALLBACK_IM_MIO                   CallbackCode = C.MSK_CALLBACK_IM_MIO                   // The callback function is called at an intermediate point in the mixed-integer optimizer.
	CALLBACK_IM_MIO_DUAL_SIMPLEX      CallbackCode = C.MSK_CALLBACK_IM_MIO_DUAL_SIMPLEX      // The callback function is called at an intermediate point in the mixed-integer optimizer while running the dual simplex optimizer.
	CALLBACK_IM_MIO_INTPNT            CallbackCode = C.MSK_CALLBACK_IM_MIO_INTPNT            // The callback function is called at an intermediate point in the mixed-integer optimizer while running the interior-point optimizer.
	CALLBACK_IM_MIO_PRIMAL_SIMPLEX    CallbackCode = C.MSK_CALLBACK_IM_MIO_PRIMAL_SIMPLEX    // The callback function is called at an intermediate point in the mixed-integer optimizer while running the primal simplex optimizer.
	CALLBACK_IM_ORDER                 CallbackCode = C.MSK_CALLBACK_IM_ORDER                 // The callback function is called from within the matrix ordering procedure at an intermediate point.
	CALLBACK_IM_PRESOLVE              CallbackCode = C.MSK_CALLBACK_IM_PRESOLVE              // The callback function is called from within the presolve procedure at an intermediate stage.
	CALLBACK_IM_PRIMAL_BI             CallbackCode = C.MSK_CALLBACK_IM_PRIMAL_BI             // The callback function is called from within the basis identification procedure at an intermediate point in the primal phase.
	CALLBACK_IM_PRIMAL_SENSIVITY      CallbackCode = C.MSK_CALLBACK_IM_PRIMAL_SENSIVITY      // The callback function is called at an intermediate stage of the primal sensitivity analysis.
	CALLBACK_IM_PRIMAL_SIMPLEX        CallbackCode = C.MSK_CALLBACK_IM_PRIMAL_SIMPLEX        // The callback function is called at an intermediate point in the primal simplex optimizer.
	CALLBACK_IM_QO_REFORMULATE        CallbackCode = C.MSK_CALLBACK_IM_QO_REFORMULATE        // The callback function is called at an intermediate stage of the conic quadratic reformulation.
	CALLBACK_IM_READ                  CallbackCode = C.MSK_CALLBACK_IM_READ                  // Intermediate stage in reading.
	CALLBACK_IM_ROOT_CUTGEN           CallbackCode = C.MSK_CALLBACK_IM_ROOT_CUTGEN           // The callback is called from within root cut generation at an intermediate stage.
	CALLBACK_IM_SIMPLEX               CallbackCode = C.MSK_CALLBACK_IM_SIMPLEX               // The callback function is called from within the simplex optimizer at an intermediate point.
	CALLBACK_IM_SIMPLEX_BI            CallbackCode = C.MSK_CALLBACK_IM_SIMPLEX_BI            // The callback function is called from within the basis identification procedure at an intermediate point in the simplex clean-up phase.
	CALLBACK_INTPNT                   CallbackCode = C.MSK_CALLBACK_INTPNT                   // The callback function is called from within the interior-point optimizer after the information database has been updated.
	CALLBACK_NEW_INT_MIO              CallbackCode = C.MSK_CALLBACK_NEW_INT_MIO              // The callback function is called after a new integer solution has been located by the mixed-integer optimizer.
	CALLBACK_PRIMAL_SIMPLEX           CallbackCode = C.MSK_CALLBACK_PRIMAL_SIMPLEX           // The callback function is called from within the primal simplex optimizer.
	CALLBACK_READ_OPF                 CallbackCode = C.MSK_CALLBACK_READ_OPF                 // The callback function is called from the OPF reader.
	CALLBACK_READ_OPF_SECTION         CallbackCode = C.MSK_CALLBACK_READ_OPF_SECTION         // A chunk of Q non-zeros has been read from a problem file.
	CALLBACK_RESTART_MIO              CallbackCode = C.MSK_CALLBACK_RESTART_MIO              // The callback function is called when the mixed-integer optimizer is restarted.
	CALLBACK_SOLVING_REMOTE           CallbackCode = C.MSK_CALLBACK_SOLVING_REMOTE           // The callback function is called while the task is being solved on a remote server.
	CALLBACK_UPDATE_DUAL_BI           CallbackCode = C.MSK_CALLBACK_UPDATE_DUAL_BI           // The callback function is called from within the basis identification procedure at an intermediate point in the dual phase.
	CALLBACK_UPDATE_DUAL_SIMPLEX      CallbackCode = C.MSK_CALLBACK_UPDATE_DUAL_SIMPLEX      // The callback function is called in the dual simplex optimizer.
	CALLBACK_UPDATE_DUAL_SIMPLEX_BI   CallbackCode = C.MSK_CALLBACK_UPDATE_DUAL_SIMPLEX_BI   // The callback function is called from within the basis identification procedure at an intermediate point in the dual simplex clean-up phase.
	CALLBACK_UPDATE_PRESOLVE          CallbackCode = C.MSK_CALLBACK_UPDATE_PRESOLVE          // The callback function is called from within the presolve procedure.
	CALLBACK_UPDATE_PRIMAL_BI         CallbackCode = C.MSK_CALLBACK_UPDATE_PRIMAL_BI         // The callback function is called from within the basis identification procedure at an intermediate point in the primal phase.
	CALLBACK_UPDATE_PRIMAL_SIMPLEX    CallbackCode = C.MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX    // The callback function is called  in the primal simplex optimizer.
	CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI CallbackCode = C.MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI // The callback function is called from within the basis identification procedure at an intermediate point in the primal simplex clean-up phase.
	CALLBACK_UPDATE_SIMPLEX           CallbackCode = C.MSK_CALLBACK_UPDATE_SIMPLEX           // The callback function is called from simplex optimizer.
	CALLBACK_WRITE_OPF                CallbackCode = C.MSK_CALLBACK_WRITE_OPF                // The callback function is called from the OPF writer.
)

var _CallbackCode_map = map[CallbackCode]string{
	CALLBACK_BEGIN_BI:                 "CALLBACK_BEGIN_BI",
	CALLBACK_BEGIN_CONIC:              "CALLBACK_BEGIN_CONIC",
	CALLBACK_BEGIN_DUAL_BI:            "CALLBACK_BEGIN_DUAL_BI",
	CALLBACK_BEGIN_DUAL_SENSITIVITY:   "CALLBACK_BEGIN_DUAL_SENSITIVITY",
	CALLBACK_BEGIN_DUAL_SETUP_BI:      "CALLBACK_BEGIN_DUAL_SETUP_BI",
	CALLBACK_BEGIN_DUAL_SIMPLEX:       "CALLBACK_BEGIN_DUAL_SIMPLEX",
	CALLBACK_BEGIN_DUAL_SIMPLEX_BI:    "CALLBACK_BEGIN_DUAL_SIMPLEX_BI",
	CALLBACK_BEGIN_INFEAS_ANA:         "CALLBACK_BEGIN_INFEAS_ANA",
	CALLBACK_BEGIN_INTPNT:             "CALLBACK_BEGIN_INTPNT",
	CALLBACK_BEGIN_LICENSE_WAIT:       "CALLBACK_BEGIN_LICENSE_WAIT",
	CALLBACK_BEGIN_MIO:                "CALLBACK_BEGIN_MIO",
	CALLBACK_BEGIN_OPTIMIZER:          "CALLBACK_BEGIN_OPTIMIZER",
	CALLBACK_BEGIN_PRESOLVE:           "CALLBACK_BEGIN_PRESOLVE",
	CALLBACK_BEGIN_PRIMAL_BI:          "CALLBACK_BEGIN_PRIMAL_BI",
	CALLBACK_BEGIN_PRIMAL_REPAIR:      "CALLBACK_BEGIN_PRIMAL_REPAIR",
	CALLBACK_BEGIN_PRIMAL_SENSITIVITY: "CALLBACK_BEGIN_PRIMAL_SENSITIVITY",
	CALLBACK_BEGIN_PRIMAL_SETUP_BI:    "CALLBACK_BEGIN_PRIMAL_SETUP_BI",
	CALLBACK_BEGIN_PRIMAL_SIMPLEX:     "CALLBACK_BEGIN_PRIMAL_SIMPLEX",
	CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI:  "CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI",
	CALLBACK_BEGIN_QCQO_REFORMULATE:   "CALLBACK_BEGIN_QCQO_REFORMULATE",
	CALLBACK_BEGIN_READ:               "CALLBACK_BEGIN_READ",
	CALLBACK_BEGIN_ROOT_CUTGEN:        "CALLBACK_BEGIN_ROOT_CUTGEN",
	CALLBACK_BEGIN_SIMPLEX:            "CALLBACK_BEGIN_SIMPLEX",
	CALLBACK_BEGIN_SIMPLEX_BI:         "CALLBACK_BEGIN_SIMPLEX_BI",
	CALLBACK_BEGIN_SOLVE_ROOT_RELAX:   "CALLBACK_BEGIN_SOLVE_ROOT_RELAX",
	CALLBACK_BEGIN_TO_CONIC:           "CALLBACK_BEGIN_TO_CONIC",
	CALLBACK_BEGIN_WRITE:              "CALLBACK_BEGIN_WRITE",
	CALLBACK_CONIC:                    "CALLBACK_CONIC",
	CALLBACK_DUAL_SIMPLEX:             "CALLBACK_DUAL_SIMPLEX",
	CALLBACK_END_BI:                   "CALLBACK_END_BI",
	CALLBACK_END_CONIC:                "CALLBACK_END_CONIC",
	CALLBACK_END_DUAL_BI:              "CALLBACK_END_DUAL_BI",
	CALLBACK_END_DUAL_SENSITIVITY:     "CALLBACK_END_DUAL_SENSITIVITY",
	CALLBACK_END_DUAL_SETUP_BI:        "CALLBACK_END_DUAL_SETUP_BI",
	CALLBACK_END_DUAL_SIMPLEX:         "CALLBACK_END_DUAL_SIMPLEX",
	CALLBACK_END_DUAL_SIMPLEX_BI:      "CALLBACK_END_DUAL_SIMPLEX_BI",
	CALLBACK_END_INFEAS_ANA:           "CALLBACK_END_INFEAS_ANA",
	CALLBACK_END_INTPNT:               "CALLBACK_END_INTPNT",
	CALLBACK_END_LICENSE_WAIT:         "CALLBACK_END_LICENSE_WAIT",
	CALLBACK_END_MIO:                  "CALLBACK_END_MIO",
	CALLBACK_END_OPTIMIZER:            "CALLBACK_END_OPTIMIZER",
	CALLBACK_END_PRESOLVE:             "CALLBACK_END_PRESOLVE",
	CALLBACK_END_PRIMAL_BI:            "CALLBACK_END_PRIMAL_BI",
	CALLBACK_END_PRIMAL_REPAIR:        "CALLBACK_END_PRIMAL_REPAIR",
	CALLBACK_END_PRIMAL_SENSITIVITY:   "CALLBACK_END_PRIMAL_SENSITIVITY",
	CALLBACK_END_PRIMAL_SETUP_BI:      "CALLBACK_END_PRIMAL_SETUP_BI",
	CALLBACK_END_PRIMAL_SIMPLEX:       "CALLBACK_END_PRIMAL_SIMPLEX",
	CALLBACK_END_PRIMAL_SIMPLEX_BI:    "CALLBACK_END_PRIMAL_SIMPLEX_BI",
	CALLBACK_END_QCQO_REFORMULATE:     "CALLBACK_END_QCQO_REFORMULATE",
	CALLBACK_END_READ:                 "CALLBACK_END_READ",
	CALLBACK_END_ROOT_CUTGEN:          "CALLBACK_END_ROOT_CUTGEN",
	CALLBACK_END_SIMPLEX:              "CALLBACK_END_SIMPLEX",
	CALLBACK_END_SIMPLEX_BI:           "CALLBACK_END_SIMPLEX_BI",
	CALLBACK_END_SOLVE_ROOT_RELAX:     "CALLBACK_END_SOLVE_ROOT_RELAX",
	CALLBACK_END_TO_CONIC:             "CALLBACK_END_TO_CONIC",
	CALLBACK_END_WRITE:                "CALLBACK_END_WRITE",
	CALLBACK_IM_BI:                    "CALLBACK_IM_BI",
	CALLBACK_IM_CONIC:                 "CALLBACK_IM_CONIC",
	CALLBACK_IM_DUAL_BI:               "CALLBACK_IM_DUAL_BI",
	CALLBACK_IM_DUAL_SENSIVITY:        "CALLBACK_IM_DUAL_SENSIVITY",
	CALLBACK_IM_DUAL_SIMPLEX:          "CALLBACK_IM_DUAL_SIMPLEX",
	CALLBACK_IM_INTPNT:                "CALLBACK_IM_INTPNT",
	CALLBACK_IM_LICENSE_WAIT:          "CALLBACK_IM_LICENSE_WAIT",
	CALLBACK_IM_LU:                    "CALLBACK_IM_LU",
	CALLBACK_IM_MIO:                   "CALLBACK_IM_MIO",
	CALLBACK_IM_MIO_DUAL_SIMPLEX:      "CALLBACK_IM_MIO_DUAL_SIMPLEX",
	CALLBACK_IM_MIO_INTPNT:            "CALLBACK_IM_MIO_INTPNT",
	CALLBACK_IM_MIO_PRIMAL_SIMPLEX:    "CALLBACK_IM_MIO_PRIMAL_SIMPLEX",
	CALLBACK_IM_ORDER:                 "CALLBACK_IM_ORDER",
	CALLBACK_IM_PRESOLVE:              "CALLBACK_IM_PRESOLVE",
	CALLBACK_IM_PRIMAL_BI:             "CALLBACK_IM_PRIMAL_BI",
	CALLBACK_IM_PRIMAL_SENSIVITY:      "CALLBACK_IM_PRIMAL_SENSIVITY",
	CALLBACK_IM_PRIMAL_SIMPLEX:        "CALLBACK_IM_PRIMAL_SIMPLEX",
	CALLBACK_IM_QO_REFORMULATE:        "CALLBACK_IM_QO_REFORMULATE",
	CALLBACK_IM_READ:                  "CALLBACK_IM_READ",
	CALLBACK_IM_ROOT_CUTGEN:           "CALLBACK_IM_ROOT_CUTGEN",
	CALLBACK_IM_SIMPLEX:               "CALLBACK_IM_SIMPLEX",
	CALLBACK_IM_SIMPLEX_BI:            "CALLBACK_IM_SIMPLEX_BI",
	CALLBACK_INTPNT:                   "CALLBACK_INTPNT",
	CALLBACK_NEW_INT_MIO:              "CALLBACK_NEW_INT_MIO",
	CALLBACK_PRIMAL_SIMPLEX:           "CALLBACK_PRIMAL_SIMPLEX",
	CALLBACK_READ_OPF:                 "CALLBACK_READ_OPF",
	CALLBACK_READ_OPF_SECTION:         "CALLBACK_READ_OPF_SECTION",
	CALLBACK_RESTART_MIO:              "CALLBACK_RESTART_MIO",
	CALLBACK_SOLVING_REMOTE:           "CALLBACK_SOLVING_REMOTE",
	CALLBACK_UPDATE_DUAL_BI:           "CALLBACK_UPDATE_DUAL_BI",
	CALLBACK_UPDATE_DUAL_SIMPLEX:      "CALLBACK_UPDATE_DUAL_SIMPLEX",
	CALLBACK_UPDATE_DUAL_SIMPLEX_BI:   "CALLBACK_UPDATE_DUAL_SIMPLEX_BI",
	CALLBACK_UPDATE_PRESOLVE:          "CALLBACK_UPDATE_PRESOLVE",
	CALLBACK_UPDATE_PRIMAL_BI:         "CALLBACK_UPDATE_PRIMAL_BI",
	CALLBACK_UPDATE_PRIMAL_SIMPLEX:    "CALLBACK_UPDATE_PRIMAL_SIMPLEX",
	CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI: "CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI",
	CALLBACK_UPDATE_SIMPLEX:           "CALLBACK_UPDATE_SIMPLEX",
	CALLBACK_WRITE_OPF:                "CALLBACK_WRITE_OPF",
}

func (e CallbackCode) String() string {
	if v, ok := _CallbackCode_map[e]; ok {
		return v
	}
	return "CallbackCode(" + strconv.FormatInt(int64(e), 10) + ")"
}
