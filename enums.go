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

// UpLo indicates if the matrix is upper triangular (up) or lower triangular (lo)
type UpLo uint32

const (
	UPLO_LO UpLo = 0 // Lower triangular
	UPLO_UP UpLo = 1 // Upper triangular
)

// Transpose indicates if the matrix input should be transposed or not.
type Transpose uint32

const (
	TRANSPOSE_NO  Transpose = 0 // No transpose
	TRANSPOSE_YES Transpose = 1 // Transpose
)
