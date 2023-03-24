// There are many enums in MOSEK, this consolidate everything here.

package gmsk

import "github.com/fardream/gmsk/res"

// ResCode is return code from mosek.
// This is a reexport to prevent polluting the namespace of gmsk
type ResCode = res.Code

const (
	RES_OK        res.Code = res.OK // RES_OK indicates success return code.
	RES_ERR_SPACE res.Code = res.ERR_SPACE
)

// ObjectiveSense is the MSKobjsense type
type ObjectiveSense uint32

const (
	OBJECTIVE_SENSE_MINIMIZE ObjectiveSense = 0 // Objective is to maximize
	OBJECTIVE_SENSE_MAXIMIZE ObjectiveSense = 1 // Objective is to minimize
)

// VarType is the variable type of mosek
type VariableType uint32

const (
	VAR_TYPE_CONT VariableType = 0 // Continuous variable
	VAR_TYPE_INT  VariableType = 1 // Integer variable
)

// SolType is the solution type
type SolType uint32

const (
	SOL_ITR SolType = 0 // Iterior Point Solution.
	SOL_BAS SolType = 1 // Basic Solution.
	SOL_ITG SolType = 2 // Integer Solution.
)

// SolSta is the solution status
type SolSta uint32

const (
	SOL_STA_UNKNOWN            SolSta = 0
	SOL_STA_OPTIMAL            SolSta = 1
	SOL_STA_PRIM_FEAS          SolSta = 2
	SOL_STA_DUAL_FEAS          SolSta = 3
	SOL_STA_PRIM_AND_DUAL_FEAS SolSta = 4
	SOL_STA_PRIM_INFEAS_CER    SolSta = 5
	SOL_STA_DUAL_INFEAS_CER    SolSta = 6
	SOL_STA_PRIM_ILLPOSED_CER  SolSta = 7
	SOL_STA_DUAL_ILLPOSED_CER  SolSta = 8
	SOL_STA_INTEGER_OPTIMAL    SolSta = 9
)

// BoundKey is MSKboundkey enum, indicate the type of the bound
type BoundKey uint32

const (
	BK_LO BoundKey = 0 // Lower bound
	BK_UP BoundKey = 1 // Upper bound
	BK_FX BoundKey = 2 // Fixed bound
	BK_FR BoundKey = 3 // Free
	BK_RA BoundKey = 4 // Range bound
)

// StreamType is MSKstreamtypee, the type of the stream.
type StreamType uint32

const (
	STREAM_LOG StreamType = 0
	STREAM_MSG StreamType = 1
	STREAM_ERR StreamType = 2
	STREAM_WRN StreamType = 3
)

// DataFormat is MSKdataformate and format of the data file.
type DataFormat uint32

const (
	DATA_FORMAT_EXTENSION DataFormat = 0
	DATA_FORMAT_MPS       DataFormat = 1
	DATA_FORMAT_LP        DataFormat = 2
	DATA_FORMAT_OP        DataFormat = 3
	DATA_FORMAT_FREE_MPS  DataFormat = 4
	DATA_FORMAT_TASK      DataFormat = 5
	DATA_FORMAT_PTF       DataFormat = 6
	DATA_FORMAT_CB        DataFormat = 7
	DATA_FORMAT_JSON_TASK DataFormat = 8
)

// CompressType is the compression type for data file
type CompressType uint32

const (
	COMPRESS_NONE CompressType = 0
	COMPRESS_FREE CompressType = 1
	COMPRESS_GZIP CompressType = 2
	COMPRESS_ZSTD CompressType = 3
)

// IParam is the integer parameter enum (MSKiparam), which
// tells what paramete the integer parameter is set for
// in MSK_putintparam or [Task.PutIntParam].
type IParam uint32

const (
	IPAR_ANA_SOL_BASIS                      IParam = 0
	IPAR_ANA_SOL_PRINT_VIOLATED             IParam = 1
	IPAR_AUTO_SORT_A_BEFORE_OPT             IParam = 2
	IPAR_AUTO_UPDATE_SOL_INFO               IParam = 3
	IPAR_BASIS_SOLVE_USE_PLUS_ONE           IParam = 4
	IPAR_BI_CLEAN_OPTIMIZER                 IParam = 5
	IPAR_BI_IGNORE_MAX_ITER                 IParam = 6
	IPAR_BI_IGNORE_NUM_ERROR                IParam = 7
	IPAR_BI_MAX_ITERATIONS                  IParam = 8
	IPAR_CACHE_LICENSE                      IParam = 9
	IPAR_CHECK_CONVEXITY                    IParam = 10
	IPAR_COMPRESS_STATFILE                  IParam = 11
	IPAR_INFEAS_GENERIC_NAMES               IParam = 12
	IPAR_INFEAS_PREFER_PRIMAL               IParam = 13
	IPAR_INFEAS_REPORT_AUTO                 IParam = 14
	IPAR_INFEAS_REPORT_LEVEL                IParam = 15
	IPAR_INTPNT_BASIS                       IParam = 16
	IPAR_INTPNT_DIFF_STEP                   IParam = 17
	IPAR_INTPNT_HOTSTART                    IParam = 18
	IPAR_INTPNT_MAX_ITERATIONS              IParam = 19
	IPAR_INTPNT_MAX_NUM_COR                 IParam = 20
	IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS    IParam = 21
	IPAR_INTPNT_OFF_COL_TRH                 IParam = 22
	IPAR_INTPNT_ORDER_GP_NUM_SEEDS          IParam = 23
	IPAR_INTPNT_ORDER_METHOD                IParam = 24
	IPAR_INTPNT_PURIFY                      IParam = 25
	IPAR_INTPNT_REGULARIZATION_USE          IParam = 26
	IPAR_INTPNT_SCALING                     IParam = 27
	IPAR_INTPNT_SOLVE_FORM                  IParam = 28
	IPAR_INTPNT_STARTING_POINT              IParam = 29
	IPAR_LICENSE_DEBUG                      IParam = 30
	IPAR_LICENSE_PAUSE_TIME                 IParam = 31
	IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS       IParam = 32
	IPAR_LICENSE_TRH_EXPIRY_WRN             IParam = 33
	IPAR_LICENSE_WAIT                       IParam = 34
	IPAR_LOG                                IParam = 35
	IPAR_LOG_ANA_PRO                        IParam = 36
	IPAR_LOG_BI                             IParam = 37
	IPAR_LOG_BI_FREQ                        IParam = 38
	IPAR_LOG_CHECK_CONVEXITY                IParam = 39
	IPAR_LOG_CUT_SECOND_OPT                 IParam = 40
	IPAR_LOG_EXPAND                         IParam = 41
	IPAR_LOG_FEAS_REPAIR                    IParam = 42
	IPAR_LOG_FILE                           IParam = 43
	IPAR_LOG_INCLUDE_SUMMARY                IParam = 44
	IPAR_LOG_INFEAS_ANA                     IParam = 45
	IPAR_LOG_INTPNT                         IParam = 46
	IPAR_LOG_LOCAL_INFO                     IParam = 47
	IPAR_LOG_MIO                            IParam = 48
	IPAR_LOG_MIO_FREQ                       IParam = 49
	IPAR_LOG_ORDER                          IParam = 50
	IPAR_LOG_PRESOLVE                       IParam = 51
	IPAR_LOG_RESPONSE                       IParam = 52
	IPAR_LOG_SENSITIVITY                    IParam = 53
	IPAR_LOG_SENSITIVITY_OPT                IParam = 54
	IPAR_LOG_SIM                            IParam = 55
	IPAR_LOG_SIM_FREQ                       IParam = 56
	IPAR_LOG_SIM_MINOR                      IParam = 57
	IPAR_LOG_STORAGE                        IParam = 58
	IPAR_MAX_NUM_WARNINGS                   IParam = 59
	IPAR_MIO_BRANCH_DIR                     IParam = 60
	IPAR_MIO_CONIC_OUTER_APPROXIMATION      IParam = 61
	IPAR_MIO_CONSTRUCT_SOL                  IParam = 62
	IPAR_MIO_CUT_CLIQUE                     IParam = 63
	IPAR_MIO_CUT_CMIR                       IParam = 64
	IPAR_MIO_CUT_GMI                        IParam = 65
	IPAR_MIO_CUT_IMPLIED_BOUND              IParam = 66
	IPAR_MIO_CUT_KNAPSACK_COVER             IParam = 67
	IPAR_MIO_CUT_LIPRO                      IParam = 68
	IPAR_MIO_CUT_SELECTION_LEVEL            IParam = 69
	IPAR_MIO_DATA_PERMUTATION_METHOD        IParam = 70
	IPAR_MIO_FEASPUMP_LEVEL                 IParam = 71
	IPAR_MIO_HEURISTIC_LEVEL                IParam = 72
	IPAR_MIO_MAX_NUM_BRANCHES               IParam = 73
	IPAR_MIO_MAX_NUM_RELAXS                 IParam = 74
	IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS        IParam = 75
	IPAR_MIO_MAX_NUM_SOLUTIONS              IParam = 76
	IPAR_MIO_MEMORY_EMPHASIS_LEVEL          IParam = 77
	IPAR_MIO_MODE                           IParam = 78
	IPAR_MIO_NODE_OPTIMIZER                 IParam = 79
	IPAR_MIO_NODE_SELECTION                 IParam = 80
	IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL       IParam = 81
	IPAR_MIO_PERSPECTIVE_REFORMULATE        IParam = 82
	IPAR_MIO_PRESOLVE_AGGREGATOR_USE        IParam = 83
	IPAR_MIO_PROBING_LEVEL                  IParam = 84
	IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT IParam = 85
	IPAR_MIO_QCQO_REFORMULATION_METHOD      IParam = 86
	IPAR_MIO_RINS_MAX_NODES                 IParam = 87
	IPAR_MIO_ROOT_OPTIMIZER                 IParam = 88
	IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL     IParam = 89
	IPAR_MIO_SEED                           IParam = 90
	IPAR_MIO_SYMMETRY_LEVEL                 IParam = 91
	IPAR_MIO_VB_DETECTION_LEVEL             IParam = 92
	IPAR_MT_SPINCOUNT                       IParam = 93
	IPAR_NG                                 IParam = 94
	IPAR_NUM_THREADS                        IParam = 95
	IPAR_OPF_WRITE_HEADER                   IParam = 96
	IPAR_OPF_WRITE_HINTS                    IParam = 97
	IPAR_OPF_WRITE_LINE_LENGTH              IParam = 98
	IPAR_OPF_WRITE_PARAMETERS               IParam = 99
	IPAR_OPF_WRITE_PROBLEM                  IParam = 100
	IPAR_OPF_WRITE_SOL_BAS                  IParam = 101
	IPAR_OPF_WRITE_SOL_ITG                  IParam = 102
	IPAR_OPF_WRITE_SOL_ITR                  IParam = 103
	IPAR_OPF_WRITE_SOLUTIONS                IParam = 104
	IPAR_OPTIMIZER                          IParam = 105
	IPAR_PARAM_READ_CASE_NAME               IParam = 106
	IPAR_PARAM_READ_IGN_ERROR               IParam = 107
	IPAR_PRESOLVE_ELIMINATOR_MAX_FILL       IParam = 108
	IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES  IParam = 109
	IPAR_PRESOLVE_LEVEL                     IParam = 110
	IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH       IParam = 111
	IPAR_PRESOLVE_LINDEP_REL_WORK_TRH       IParam = 112
	IPAR_PRESOLVE_LINDEP_USE                IParam = 113
	IPAR_PRESOLVE_MAX_NUM_PASS              IParam = 114
	IPAR_PRESOLVE_MAX_NUM_REDUCTIONS        IParam = 115
	IPAR_PRESOLVE_USE                       IParam = 116
	IPAR_PRIMAL_REPAIR_OPTIMIZER            IParam = 117
	IPAR_PTF_WRITE_PARAMETERS               IParam = 118
	IPAR_PTF_WRITE_SOLUTIONS                IParam = 119
	IPAR_PTF_WRITE_TRANSFORM                IParam = 120
	IPAR_READ_DEBUG                         IParam = 121
	IPAR_READ_KEEP_FREE_CON                 IParam = 122
	IPAR_READ_MPS_FORMAT                    IParam = 123
	IPAR_READ_MPS_WIDTH                     IParam = 124
	IPAR_READ_TASK_IGNORE_PARAM             IParam = 125
	IPAR_REMOTE_USE_COMPRESSION             IParam = 126
	IPAR_REMOVE_UNUSED_SOLUTIONS            IParam = 127
	IPAR_SENSITIVITY_ALL                    IParam = 128
	IPAR_SENSITIVITY_OPTIMIZER              IParam = 129
	IPAR_SENSITIVITY_TYPE                   IParam = 130
	IPAR_SIM_BASIS_FACTOR_USE               IParam = 131
	IPAR_SIM_DEGEN                          IParam = 132
	IPAR_SIM_DETECT_PWL                     IParam = 133
	IPAR_SIM_DUAL_CRASH                     IParam = 134
	IPAR_SIM_DUAL_PHASEONE_METHOD           IParam = 135
	IPAR_SIM_DUAL_RESTRICT_SELECTION        IParam = 136
	IPAR_SIM_DUAL_SELECTION                 IParam = 137
	IPAR_SIM_EXPLOIT_DUPVEC                 IParam = 138
	IPAR_SIM_HOTSTART                       IParam = 139
	IPAR_SIM_HOTSTART_LU                    IParam = 140
	IPAR_SIM_MAX_ITERATIONS                 IParam = 141
	IPAR_SIM_MAX_NUM_SETBACKS               IParam = 142
	IPAR_SIM_NON_SINGULAR                   IParam = 143
	IPAR_SIM_PRIMAL_CRASH                   IParam = 144
	IPAR_SIM_PRIMAL_PHASEONE_METHOD         IParam = 145
	IPAR_SIM_PRIMAL_RESTRICT_SELECTION      IParam = 146
	IPAR_SIM_PRIMAL_SELECTION               IParam = 147
	IPAR_SIM_REFACTOR_FREQ                  IParam = 148
	IPAR_SIM_REFORMULATION                  IParam = 149
	IPAR_SIM_SAVE_LU                        IParam = 150
	IPAR_SIM_SCALING                        IParam = 151
	IPAR_SIM_SCALING_METHOD                 IParam = 152
	IPAR_SIM_SEED                           IParam = 153
	IPAR_SIM_SOLVE_FORM                     IParam = 154
	IPAR_SIM_STABILITY_PRIORITY             IParam = 155
	IPAR_SIM_SWITCH_OPTIMIZER               IParam = 156
	IPAR_SOL_FILTER_KEEP_BASIC              IParam = 157
	IPAR_SOL_FILTER_KEEP_RANGED             IParam = 158
	IPAR_SOL_READ_NAME_WIDTH                IParam = 159
	IPAR_SOL_READ_WIDTH                     IParam = 160
	IPAR_SOLUTION_CALLBACK                  IParam = 161
	IPAR_TIMING_LEVEL                       IParam = 162
	IPAR_WRITE_BAS_CONSTRAINTS              IParam = 163
	IPAR_WRITE_BAS_HEAD                     IParam = 164
	IPAR_WRITE_BAS_VARIABLES                IParam = 165
	IPAR_WRITE_COMPRESSION                  IParam = 166
	IPAR_WRITE_DATA_PARAM                   IParam = 167
	IPAR_WRITE_FREE_CON                     IParam = 168
	IPAR_WRITE_GENERIC_NAMES                IParam = 169
	IPAR_WRITE_GENERIC_NAMES_IO             IParam = 170
	IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS    IParam = 171
	IPAR_WRITE_INT_CONSTRAINTS              IParam = 172
	IPAR_WRITE_INT_HEAD                     IParam = 173
	IPAR_WRITE_INT_VARIABLES                IParam = 174
	IPAR_WRITE_JSON_INDENTATION             IParam = 175
	IPAR_WRITE_LP_FULL_OBJ                  IParam = 176
	IPAR_WRITE_LP_LINE_WIDTH                IParam = 177
	IPAR_WRITE_MPS_FORMAT                   IParam = 178
	IPAR_WRITE_MPS_INT                      IParam = 179
	IPAR_WRITE_SOL_BARVARIABLES             IParam = 180
	IPAR_WRITE_SOL_CONSTRAINTS              IParam = 181
	IPAR_WRITE_SOL_HEAD                     IParam = 182
	IPAR_WRITE_SOL_IGNORE_INVALID_NAMES     IParam = 183
	IPAR_WRITE_SOL_VARIABLES                IParam = 184
	IPAR_WRITE_TASK_INC_SOL                 IParam = 185
	IPAR_WRITE_XML_MODE                     IParam = 186
)

// UpLo indicates if the matrix is upper triangular (up) or lower triangular (lo)
type UpLo uint32

const (
	UPLO_LO UpLo = 0 // Lower triangular
	UPLO_UP UpLo = 1 // Upper triangular
)

type Transpose uint32

const (
	TRANSPOSE_NO  Transpose = 0 // No transpose
	TRANSPOSE_YES Transpose = 1 // Transpose
)
