// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKoptimizertype_enum/OptimizerType

package gmsk

// #include <mosek.h>
import "C"

// OptimizerType is MSKoptimizertype_enum.
//
// can be set for the integer parameter [IPAR_OPTIMIZER]
type OptimizerType = int32

const (
	OPTIMIZER_CONIC          OptimizerType = C.MSK_OPTIMIZER_CONIC          // The optimizer for problems having conic constraints.
	OPTIMIZER_DUAL_SIMPLEX   OptimizerType = C.MSK_OPTIMIZER_DUAL_SIMPLEX   // The dual simplex optimizer is used.
	OPTIMIZER_FREE           OptimizerType = C.MSK_OPTIMIZER_FREE           // The optimizer is chosen automatically.
	OPTIMIZER_FREE_SIMPLEX   OptimizerType = C.MSK_OPTIMIZER_FREE_SIMPLEX   // One of the simplex optimizers is used.
	OPTIMIZER_INTPNT         OptimizerType = C.MSK_OPTIMIZER_INTPNT         // The interior-point optimizer is used.
	OPTIMIZER_MIXED_INT      OptimizerType = C.MSK_OPTIMIZER_MIXED_INT      // The mixed-integer optimizer.
	OPTIMIZER_PRIMAL_SIMPLEX OptimizerType = C.MSK_OPTIMIZER_PRIMAL_SIMPLEX // The primal simplex optimizer is used.
)
