// There are many enums in MOSEK, this consolidate everything here.

package gmsk

// #include <mosek.h>
import "C"

// INFINITY is MSK_INFINITY (which is different from the double's infinity)
const INFINITY float64 = C.MSK_INFINITY

// MSK_MAX_STR_LEN is the max length of strings in mosek
const MAX_STR_LEN = C.MSK_MAX_STR_LEN

// ObjectiveSense is the MSKobjsense type
type ObjectiveSense uint32

const (
	OBJECTIVE_SENSE_MINIMIZE ObjectiveSense = C.MSK_OBJECTIVE_SENSE_MINIMIZE // Objective is to maximize
	OBJECTIVE_SENSE_MAXIMIZE ObjectiveSense = C.MSK_OBJECTIVE_SENSE_MAXIMIZE // Objective is to minimize
)

// VarType is the variable type of mosek
type VariableType uint32

const (
	VAR_TYPE_CONT VariableType = C.MSK_VAR_TYPE_CONT // Continuous variable
	VAR_TYPE_INT  VariableType = C.MSK_VAR_TYPE_INT  // Integer variable
)

// SolType is the solution type
type SolType uint32

const (
	SOL_ITR SolType = C.MSK_SOL_ITR // Iterior Point Solution.
	SOL_BAS SolType = C.MSK_SOL_BAS // Basic Solution.
	SOL_ITG SolType = C.MSK_SOL_ITG // Integer Solution.
)

// SolSta is the solution status
type SolSta = uint32

const (
	SOL_STA_UNKNOWN            SolSta = C.MSK_SOL_STA_UNKNOWN
	SOL_STA_OPTIMAL            SolSta = C.MSK_SOL_STA_OPTIMAL
	SOL_STA_PRIM_FEAS          SolSta = C.MSK_SOL_STA_PRIM_FEAS
	SOL_STA_DUAL_FEAS          SolSta = C.MSK_SOL_STA_DUAL_FEAS
	SOL_STA_PRIM_AND_DUAL_FEAS SolSta = C.MSK_SOL_STA_PRIM_AND_DUAL_FEAS
	SOL_STA_PRIM_INFEAS_CER    SolSta = C.MSK_SOL_STA_PRIM_INFEAS_CER
	SOL_STA_DUAL_INFEAS_CER    SolSta = C.MSK_SOL_STA_DUAL_INFEAS_CER
	SOL_STA_PRIM_ILLPOSED_CER  SolSta = C.MSK_SOL_STA_PRIM_ILLPOSED_CER
	SOL_STA_DUAL_ILLPOSED_CER  SolSta = C.MSK_SOL_STA_DUAL_ILLPOSED_CER
	SOL_STA_INTEGER_OPTIMAL    SolSta = C.MSK_SOL_STA_INTEGER_OPTIMAL
)

// BoundKey is MSKboundkey enum, indicate the type of the bound
type BoundKey uint32

const (
	BK_LO BoundKey = C.MSK_BK_LO // Lower bound
	BK_UP BoundKey = C.MSK_BK_UP // Upper bound
	BK_FX BoundKey = C.MSK_BK_FX // Fixed bound
	BK_FR BoundKey = C.MSK_BK_FR // Free
	BK_RA BoundKey = C.MSK_BK_RA // Range bound
)

// StreamType is MSKstreamtypee, the type of the stream.
type StreamType uint32

const (
	STREAM_LOG StreamType = C.MSK_STREAM_LOG
	STREAM_MSG StreamType = C.MSK_STREAM_MSG
	STREAM_ERR StreamType = C.MSK_STREAM_ERR
	STREAM_WRN StreamType = C.MSK_STREAM_WRN
)

// DataFormat is MSKdataformate and format of the data file.
type DataFormat uint32

const (
	DATA_FORMAT_EXTENSION DataFormat = C.MSK_DATA_FORMAT_EXTENSION
	DATA_FORMAT_MPS       DataFormat = C.MSK_DATA_FORMAT_MPS
	DATA_FORMAT_LP        DataFormat = C.MSK_DATA_FORMAT_LP
	DATA_FORMAT_OP        DataFormat = C.MSK_DATA_FORMAT_OP
	DATA_FORMAT_FREE_MPS  DataFormat = C.MSK_DATA_FORMAT_FREE_MPS
	DATA_FORMAT_TASK      DataFormat = C.MSK_DATA_FORMAT_TASK
	DATA_FORMAT_PTF       DataFormat = C.MSK_DATA_FORMAT_PTF
	DATA_FORMAT_CB        DataFormat = C.MSK_DATA_FORMAT_CB
	DATA_FORMAT_JSON_TASK DataFormat = C.MSK_DATA_FORMAT_JSON_TASK
)

// CompressType is the compression type for data file
type CompressType uint32

const (
	COMPRESS_NONE CompressType = C.MSK_COMPRESS_NONE
	COMPRESS_FREE CompressType = C.MSK_COMPRESS_FREE
	COMPRESS_GZIP CompressType = C.MSK_COMPRESS_GZIP
	COMPRESS_ZSTD CompressType = C.MSK_COMPRESS_ZSTD
)

// IParam is the integer parameter enum (MSKiparam), which
// tells what paramete the integer parameter is set for
// in MSK_putintparam or [Task.PutIntParam].
type IParam uint32

const (
	IPAR_ANA_SOL_BASIS                      IParam = C.MSK_IPAR_ANA_SOL_BASIS
	IPAR_ANA_SOL_PRINT_VIOLATED             IParam = C.MSK_IPAR_ANA_SOL_PRINT_VIOLATED
	IPAR_AUTO_SORT_A_BEFORE_OPT             IParam = C.MSK_IPAR_AUTO_SORT_A_BEFORE_OPT
	IPAR_AUTO_UPDATE_SOL_INFO               IParam = C.MSK_IPAR_AUTO_UPDATE_SOL_INFO
	IPAR_BASIS_SOLVE_USE_PLUS_ONE           IParam = C.MSK_IPAR_BASIS_SOLVE_USE_PLUS_ONE
	IPAR_BI_CLEAN_OPTIMIZER                 IParam = C.MSK_IPAR_BI_CLEAN_OPTIMIZER
	IPAR_BI_IGNORE_MAX_ITER                 IParam = C.MSK_IPAR_BI_IGNORE_MAX_ITER
	IPAR_BI_IGNORE_NUM_ERROR                IParam = C.MSK_IPAR_BI_IGNORE_NUM_ERROR
	IPAR_BI_MAX_ITERATIONS                  IParam = C.MSK_IPAR_BI_MAX_ITERATIONS
	IPAR_CACHE_LICENSE                      IParam = C.MSK_IPAR_CACHE_LICENSE
	IPAR_CHECK_CONVEXITY                    IParam = C.MSK_IPAR_CHECK_CONVEXITY
	IPAR_COMPRESS_STATFILE                  IParam = C.MSK_IPAR_COMPRESS_STATFILE
	IPAR_INFEAS_GENERIC_NAMES               IParam = C.MSK_IPAR_INFEAS_GENERIC_NAMES
	IPAR_INFEAS_PREFER_PRIMAL               IParam = C.MSK_IPAR_INFEAS_PREFER_PRIMAL
	IPAR_INFEAS_REPORT_AUTO                 IParam = C.MSK_IPAR_INFEAS_REPORT_AUTO
	IPAR_INFEAS_REPORT_LEVEL                IParam = C.MSK_IPAR_INFEAS_REPORT_LEVEL
	IPAR_INTPNT_BASIS                       IParam = C.MSK_IPAR_INTPNT_BASIS
	IPAR_INTPNT_DIFF_STEP                   IParam = C.MSK_IPAR_INTPNT_DIFF_STEP
	IPAR_INTPNT_HOTSTART                    IParam = C.MSK_IPAR_INTPNT_HOTSTART
	IPAR_INTPNT_MAX_ITERATIONS              IParam = C.MSK_IPAR_INTPNT_MAX_ITERATIONS
	IPAR_INTPNT_MAX_NUM_COR                 IParam = C.MSK_IPAR_INTPNT_MAX_NUM_COR
	IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS    IParam = C.MSK_IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS
	IPAR_INTPNT_OFF_COL_TRH                 IParam = C.MSK_IPAR_INTPNT_OFF_COL_TRH
	IPAR_INTPNT_ORDER_GP_NUM_SEEDS          IParam = C.MSK_IPAR_INTPNT_ORDER_GP_NUM_SEEDS
	IPAR_INTPNT_ORDER_METHOD                IParam = C.MSK_IPAR_INTPNT_ORDER_METHOD
	IPAR_INTPNT_PURIFY                      IParam = C.MSK_IPAR_INTPNT_PURIFY
	IPAR_INTPNT_REGULARIZATION_USE          IParam = C.MSK_IPAR_INTPNT_REGULARIZATION_USE
	IPAR_INTPNT_SCALING                     IParam = C.MSK_IPAR_INTPNT_SCALING
	IPAR_INTPNT_SOLVE_FORM                  IParam = C.MSK_IPAR_INTPNT_SOLVE_FORM
	IPAR_INTPNT_STARTING_POINT              IParam = C.MSK_IPAR_INTPNT_STARTING_POINT
	IPAR_LICENSE_DEBUG                      IParam = C.MSK_IPAR_LICENSE_DEBUG
	IPAR_LICENSE_PAUSE_TIME                 IParam = C.MSK_IPAR_LICENSE_PAUSE_TIME
	IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS       IParam = C.MSK_IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS
	IPAR_LICENSE_TRH_EXPIRY_WRN             IParam = C.MSK_IPAR_LICENSE_TRH_EXPIRY_WRN
	IPAR_LICENSE_WAIT                       IParam = C.MSK_IPAR_LICENSE_WAIT
	IPAR_LOG                                IParam = C.MSK_IPAR_LOG
	IPAR_LOG_ANA_PRO                        IParam = C.MSK_IPAR_LOG_ANA_PRO
	IPAR_LOG_BI                             IParam = C.MSK_IPAR_LOG_BI
	IPAR_LOG_BI_FREQ                        IParam = C.MSK_IPAR_LOG_BI_FREQ
	IPAR_LOG_CHECK_CONVEXITY                IParam = C.MSK_IPAR_LOG_CHECK_CONVEXITY
	IPAR_LOG_CUT_SECOND_OPT                 IParam = C.MSK_IPAR_LOG_CUT_SECOND_OPT
	IPAR_LOG_EXPAND                         IParam = C.MSK_IPAR_LOG_EXPAND
	IPAR_LOG_FEAS_REPAIR                    IParam = C.MSK_IPAR_LOG_FEAS_REPAIR
	IPAR_LOG_FILE                           IParam = C.MSK_IPAR_LOG_FILE
	IPAR_LOG_INCLUDE_SUMMARY                IParam = C.MSK_IPAR_LOG_INCLUDE_SUMMARY
	IPAR_LOG_INFEAS_ANA                     IParam = C.MSK_IPAR_LOG_INFEAS_ANA
	IPAR_LOG_INTPNT                         IParam = C.MSK_IPAR_LOG_INTPNT
	IPAR_LOG_LOCAL_INFO                     IParam = C.MSK_IPAR_LOG_LOCAL_INFO
	IPAR_LOG_MIO                            IParam = C.MSK_IPAR_LOG_MIO
	IPAR_LOG_MIO_FREQ                       IParam = C.MSK_IPAR_LOG_MIO_FREQ
	IPAR_LOG_ORDER                          IParam = C.MSK_IPAR_LOG_ORDER
	IPAR_LOG_PRESOLVE                       IParam = C.MSK_IPAR_LOG_PRESOLVE
	IPAR_LOG_RESPONSE                       IParam = C.MSK_IPAR_LOG_RESPONSE
	IPAR_LOG_SENSITIVITY                    IParam = C.MSK_IPAR_LOG_SENSITIVITY
	IPAR_LOG_SENSITIVITY_OPT                IParam = C.MSK_IPAR_LOG_SENSITIVITY_OPT
	IPAR_LOG_SIM                            IParam = C.MSK_IPAR_LOG_SIM
	IPAR_LOG_SIM_FREQ                       IParam = C.MSK_IPAR_LOG_SIM_FREQ
	IPAR_LOG_SIM_MINOR                      IParam = C.MSK_IPAR_LOG_SIM_MINOR
	IPAR_LOG_STORAGE                        IParam = C.MSK_IPAR_LOG_STORAGE
	IPAR_MAX_NUM_WARNINGS                   IParam = C.MSK_IPAR_MAX_NUM_WARNINGS
	IPAR_MIO_BRANCH_DIR                     IParam = C.MSK_IPAR_MIO_BRANCH_DIR
	IPAR_MIO_CONIC_OUTER_APPROXIMATION      IParam = C.MSK_IPAR_MIO_CONIC_OUTER_APPROXIMATION
	IPAR_MIO_CONSTRUCT_SOL                  IParam = C.MSK_IPAR_MIO_CONSTRUCT_SOL
	IPAR_MIO_CUT_CLIQUE                     IParam = C.MSK_IPAR_MIO_CUT_CLIQUE
	IPAR_MIO_CUT_CMIR                       IParam = C.MSK_IPAR_MIO_CUT_CMIR
	IPAR_MIO_CUT_GMI                        IParam = C.MSK_IPAR_MIO_CUT_GMI
	IPAR_MIO_CUT_IMPLIED_BOUND              IParam = C.MSK_IPAR_MIO_CUT_IMPLIED_BOUND
	IPAR_MIO_CUT_KNAPSACK_COVER             IParam = C.MSK_IPAR_MIO_CUT_KNAPSACK_COVER
	IPAR_MIO_CUT_LIPRO                      IParam = C.MSK_IPAR_MIO_CUT_LIPRO
	IPAR_MIO_CUT_SELECTION_LEVEL            IParam = C.MSK_IPAR_MIO_CUT_SELECTION_LEVEL
	IPAR_MIO_DATA_PERMUTATION_METHOD        IParam = C.MSK_IPAR_MIO_DATA_PERMUTATION_METHOD
	IPAR_MIO_FEASPUMP_LEVEL                 IParam = C.MSK_IPAR_MIO_FEASPUMP_LEVEL
	IPAR_MIO_HEURISTIC_LEVEL                IParam = C.MSK_IPAR_MIO_HEURISTIC_LEVEL
	IPAR_MIO_MAX_NUM_BRANCHES               IParam = C.MSK_IPAR_MIO_MAX_NUM_BRANCHES
	IPAR_MIO_MAX_NUM_RELAXS                 IParam = C.MSK_IPAR_MIO_MAX_NUM_RELAXS
	IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS        IParam = C.MSK_IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS
	IPAR_MIO_MAX_NUM_SOLUTIONS              IParam = C.MSK_IPAR_MIO_MAX_NUM_SOLUTIONS
	IPAR_MIO_MEMORY_EMPHASIS_LEVEL          IParam = C.MSK_IPAR_MIO_MEMORY_EMPHASIS_LEVEL
	IPAR_MIO_MODE                           IParam = C.MSK_IPAR_MIO_MODE
	IPAR_MIO_NODE_OPTIMIZER                 IParam = C.MSK_IPAR_MIO_NODE_OPTIMIZER
	IPAR_MIO_NODE_SELECTION                 IParam = C.MSK_IPAR_MIO_NODE_SELECTION
	IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL       IParam = C.MSK_IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL
	IPAR_MIO_PERSPECTIVE_REFORMULATE        IParam = C.MSK_IPAR_MIO_PERSPECTIVE_REFORMULATE
	IPAR_MIO_PRESOLVE_AGGREGATOR_USE        IParam = C.MSK_IPAR_MIO_PRESOLVE_AGGREGATOR_USE
	IPAR_MIO_PROBING_LEVEL                  IParam = C.MSK_IPAR_MIO_PROBING_LEVEL
	IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT IParam = C.MSK_IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT
	IPAR_MIO_QCQO_REFORMULATION_METHOD      IParam = C.MSK_IPAR_MIO_QCQO_REFORMULATION_METHOD
	IPAR_MIO_RINS_MAX_NODES                 IParam = C.MSK_IPAR_MIO_RINS_MAX_NODES
	IPAR_MIO_ROOT_OPTIMIZER                 IParam = C.MSK_IPAR_MIO_ROOT_OPTIMIZER
	IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL     IParam = C.MSK_IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL
	IPAR_MIO_SEED                           IParam = C.MSK_IPAR_MIO_SEED
	IPAR_MIO_SYMMETRY_LEVEL                 IParam = C.MSK_IPAR_MIO_SYMMETRY_LEVEL
	IPAR_MIO_VB_DETECTION_LEVEL             IParam = C.MSK_IPAR_MIO_VB_DETECTION_LEVEL
	IPAR_MT_SPINCOUNT                       IParam = C.MSK_IPAR_MT_SPINCOUNT
	IPAR_NG                                 IParam = C.MSK_IPAR_NG
	IPAR_NUM_THREADS                        IParam = C.MSK_IPAR_NUM_THREADS
	IPAR_OPF_WRITE_HEADER                   IParam = C.MSK_IPAR_OPF_WRITE_HEADER
	IPAR_OPF_WRITE_HINTS                    IParam = C.MSK_IPAR_OPF_WRITE_HINTS
	IPAR_OPF_WRITE_LINE_LENGTH              IParam = C.MSK_IPAR_OPF_WRITE_LINE_LENGTH
	IPAR_OPF_WRITE_PARAMETERS               IParam = C.MSK_IPAR_OPF_WRITE_PARAMETERS
	IPAR_OPF_WRITE_PROBLEM                  IParam = C.MSK_IPAR_OPF_WRITE_PROBLEM
	IPAR_OPF_WRITE_SOL_BAS                  IParam = C.MSK_IPAR_OPF_WRITE_SOL_BAS
	IPAR_OPF_WRITE_SOL_ITG                  IParam = C.MSK_IPAR_OPF_WRITE_SOL_ITG
	IPAR_OPF_WRITE_SOL_ITR                  IParam = C.MSK_IPAR_OPF_WRITE_SOL_ITR
	IPAR_OPF_WRITE_SOLUTIONS                IParam = C.MSK_IPAR_OPF_WRITE_SOLUTIONS
	IPAR_OPTIMIZER                          IParam = C.MSK_IPAR_OPTIMIZER
	IPAR_PARAM_READ_CASE_NAME               IParam = C.MSK_IPAR_PARAM_READ_CASE_NAME
	IPAR_PARAM_READ_IGN_ERROR               IParam = C.MSK_IPAR_PARAM_READ_IGN_ERROR
	IPAR_PRESOLVE_ELIMINATOR_MAX_FILL       IParam = C.MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_FILL
	IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES  IParam = C.MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES
	IPAR_PRESOLVE_LEVEL                     IParam = C.MSK_IPAR_PRESOLVE_LEVEL
	IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH       IParam = C.MSK_IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH
	IPAR_PRESOLVE_LINDEP_REL_WORK_TRH       IParam = C.MSK_IPAR_PRESOLVE_LINDEP_REL_WORK_TRH
	IPAR_PRESOLVE_LINDEP_USE                IParam = C.MSK_IPAR_PRESOLVE_LINDEP_USE
	IPAR_PRESOLVE_MAX_NUM_PASS              IParam = C.MSK_IPAR_PRESOLVE_MAX_NUM_PASS
	IPAR_PRESOLVE_MAX_NUM_REDUCTIONS        IParam = C.MSK_IPAR_PRESOLVE_MAX_NUM_REDUCTIONS
	IPAR_PRESOLVE_USE                       IParam = C.MSK_IPAR_PRESOLVE_USE
	IPAR_PRIMAL_REPAIR_OPTIMIZER            IParam = C.MSK_IPAR_PRIMAL_REPAIR_OPTIMIZER
	IPAR_PTF_WRITE_PARAMETERS               IParam = C.MSK_IPAR_PTF_WRITE_PARAMETERS
	IPAR_PTF_WRITE_SOLUTIONS                IParam = C.MSK_IPAR_PTF_WRITE_SOLUTIONS
	IPAR_PTF_WRITE_TRANSFORM                IParam = C.MSK_IPAR_PTF_WRITE_TRANSFORM
	IPAR_READ_DEBUG                         IParam = C.MSK_IPAR_READ_DEBUG
	IPAR_READ_KEEP_FREE_CON                 IParam = C.MSK_IPAR_READ_KEEP_FREE_CON
	IPAR_READ_MPS_FORMAT                    IParam = C.MSK_IPAR_READ_MPS_FORMAT
	IPAR_READ_MPS_WIDTH                     IParam = C.MSK_IPAR_READ_MPS_WIDTH
	IPAR_READ_TASK_IGNORE_PARAM             IParam = C.MSK_IPAR_READ_TASK_IGNORE_PARAM
	IPAR_REMOTE_USE_COMPRESSION             IParam = C.MSK_IPAR_REMOTE_USE_COMPRESSION
	IPAR_REMOVE_UNUSED_SOLUTIONS            IParam = C.MSK_IPAR_REMOVE_UNUSED_SOLUTIONS
	IPAR_SENSITIVITY_ALL                    IParam = C.MSK_IPAR_SENSITIVITY_ALL
	IPAR_SENSITIVITY_OPTIMIZER              IParam = C.MSK_IPAR_SENSITIVITY_OPTIMIZER
	IPAR_SENSITIVITY_TYPE                   IParam = C.MSK_IPAR_SENSITIVITY_TYPE
	IPAR_SIM_BASIS_FACTOR_USE               IParam = C.MSK_IPAR_SIM_BASIS_FACTOR_USE
	IPAR_SIM_DEGEN                          IParam = C.MSK_IPAR_SIM_DEGEN
	IPAR_SIM_DETECT_PWL                     IParam = C.MSK_IPAR_SIM_DETECT_PWL
	IPAR_SIM_DUAL_CRASH                     IParam = C.MSK_IPAR_SIM_DUAL_CRASH
	IPAR_SIM_DUAL_PHASEONE_METHOD           IParam = C.MSK_IPAR_SIM_DUAL_PHASEONE_METHOD
	IPAR_SIM_DUAL_RESTRICT_SELECTION        IParam = C.MSK_IPAR_SIM_DUAL_RESTRICT_SELECTION
	IPAR_SIM_DUAL_SELECTION                 IParam = C.MSK_IPAR_SIM_DUAL_SELECTION
	IPAR_SIM_EXPLOIT_DUPVEC                 IParam = C.MSK_IPAR_SIM_EXPLOIT_DUPVEC
	IPAR_SIM_HOTSTART                       IParam = C.MSK_IPAR_SIM_HOTSTART
	IPAR_SIM_HOTSTART_LU                    IParam = C.MSK_IPAR_SIM_HOTSTART_LU
	IPAR_SIM_MAX_ITERATIONS                 IParam = C.MSK_IPAR_SIM_MAX_ITERATIONS
	IPAR_SIM_MAX_NUM_SETBACKS               IParam = C.MSK_IPAR_SIM_MAX_NUM_SETBACKS
	IPAR_SIM_NON_SINGULAR                   IParam = C.MSK_IPAR_SIM_NON_SINGULAR
	IPAR_SIM_PRIMAL_CRASH                   IParam = C.MSK_IPAR_SIM_PRIMAL_CRASH
	IPAR_SIM_PRIMAL_PHASEONE_METHOD         IParam = C.MSK_IPAR_SIM_PRIMAL_PHASEONE_METHOD
	IPAR_SIM_PRIMAL_RESTRICT_SELECTION      IParam = C.MSK_IPAR_SIM_PRIMAL_RESTRICT_SELECTION
	IPAR_SIM_PRIMAL_SELECTION               IParam = C.MSK_IPAR_SIM_PRIMAL_SELECTION
	IPAR_SIM_REFACTOR_FREQ                  IParam = C.MSK_IPAR_SIM_REFACTOR_FREQ
	IPAR_SIM_REFORMULATION                  IParam = C.MSK_IPAR_SIM_REFORMULATION
	IPAR_SIM_SAVE_LU                        IParam = C.MSK_IPAR_SIM_SAVE_LU
	IPAR_SIM_SCALING                        IParam = C.MSK_IPAR_SIM_SCALING
	IPAR_SIM_SCALING_METHOD                 IParam = C.MSK_IPAR_SIM_SCALING_METHOD
	IPAR_SIM_SEED                           IParam = C.MSK_IPAR_SIM_SEED
	IPAR_SIM_SOLVE_FORM                     IParam = C.MSK_IPAR_SIM_SOLVE_FORM
	IPAR_SIM_STABILITY_PRIORITY             IParam = C.MSK_IPAR_SIM_STABILITY_PRIORITY
	IPAR_SIM_SWITCH_OPTIMIZER               IParam = C.MSK_IPAR_SIM_SWITCH_OPTIMIZER
	IPAR_SOL_FILTER_KEEP_BASIC              IParam = C.MSK_IPAR_SOL_FILTER_KEEP_BASIC
	IPAR_SOL_FILTER_KEEP_RANGED             IParam = C.MSK_IPAR_SOL_FILTER_KEEP_RANGED
	IPAR_SOL_READ_NAME_WIDTH                IParam = C.MSK_IPAR_SOL_READ_NAME_WIDTH
	IPAR_SOL_READ_WIDTH                     IParam = C.MSK_IPAR_SOL_READ_WIDTH
	IPAR_SOLUTION_CALLBACK                  IParam = C.MSK_IPAR_SOLUTION_CALLBACK
	IPAR_TIMING_LEVEL                       IParam = C.MSK_IPAR_TIMING_LEVEL
	IPAR_WRITE_BAS_CONSTRAINTS              IParam = C.MSK_IPAR_WRITE_BAS_CONSTRAINTS
	IPAR_WRITE_BAS_HEAD                     IParam = C.MSK_IPAR_WRITE_BAS_HEAD
	IPAR_WRITE_BAS_VARIABLES                IParam = C.MSK_IPAR_WRITE_BAS_VARIABLES
	IPAR_WRITE_COMPRESSION                  IParam = C.MSK_IPAR_WRITE_COMPRESSION
	IPAR_WRITE_DATA_PARAM                   IParam = C.MSK_IPAR_WRITE_DATA_PARAM
	IPAR_WRITE_FREE_CON                     IParam = C.MSK_IPAR_WRITE_FREE_CON
	IPAR_WRITE_GENERIC_NAMES                IParam = C.MSK_IPAR_WRITE_GENERIC_NAMES
	IPAR_WRITE_GENERIC_NAMES_IO             IParam = C.MSK_IPAR_WRITE_GENERIC_NAMES_IO
	IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS    IParam = C.MSK_IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS
	IPAR_WRITE_INT_CONSTRAINTS              IParam = C.MSK_IPAR_WRITE_INT_CONSTRAINTS
	IPAR_WRITE_INT_HEAD                     IParam = C.MSK_IPAR_WRITE_INT_HEAD
	IPAR_WRITE_INT_VARIABLES                IParam = C.MSK_IPAR_WRITE_INT_VARIABLES
	IPAR_WRITE_JSON_INDENTATION             IParam = C.MSK_IPAR_WRITE_JSON_INDENTATION
	IPAR_WRITE_LP_FULL_OBJ                  IParam = C.MSK_IPAR_WRITE_LP_FULL_OBJ
	IPAR_WRITE_LP_LINE_WIDTH                IParam = C.MSK_IPAR_WRITE_LP_LINE_WIDTH
	IPAR_WRITE_MPS_FORMAT                   IParam = C.MSK_IPAR_WRITE_MPS_FORMAT
	IPAR_WRITE_MPS_INT                      IParam = C.MSK_IPAR_WRITE_MPS_INT
	IPAR_WRITE_SOL_BARVARIABLES             IParam = C.MSK_IPAR_WRITE_SOL_BARVARIABLES
	IPAR_WRITE_SOL_CONSTRAINTS              IParam = C.MSK_IPAR_WRITE_SOL_CONSTRAINTS
	IPAR_WRITE_SOL_HEAD                     IParam = C.MSK_IPAR_WRITE_SOL_HEAD
	IPAR_WRITE_SOL_IGNORE_INVALID_NAMES     IParam = C.MSK_IPAR_WRITE_SOL_IGNORE_INVALID_NAMES
	IPAR_WRITE_SOL_VARIABLES                IParam = C.MSK_IPAR_WRITE_SOL_VARIABLES
	IPAR_WRITE_TASK_INC_SOL                 IParam = C.MSK_IPAR_WRITE_TASK_INC_SOL
	IPAR_WRITE_XML_MODE                     IParam = C.MSK_IPAR_WRITE_XML_MODE
)

// UpLo indicates if the matrix is upper triangular (up) or lower triangular (lo)
type UpLo uint32

const (
	UPLO_LO UpLo = C.MSK_UPLO_LO // Lower triangular
	UPLO_UP UpLo = C.MSK_UPLO_UP // Upper triangular
)

type Transpose uint32

const (
	TRANSPOSE_NO  Transpose = C.MSK_TRANSPOSE_NO  // No transpose
	TRANSPOSE_YES Transpose = C.MSK_TRANSPOSE_YES // Transpose
)
