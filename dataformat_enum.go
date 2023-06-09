// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKdataformat_enum/DataFormat

package gmsk

import "strconv"

// DataFormat is MSKdataformat_enum.
//
// Data format types
type DataFormat uint32

const (
	DATA_FORMAT_EXTENSION DataFormat = 0 // The file extension is used to determine the data file format.
	DATA_FORMAT_MPS       DataFormat = 1 // The data file is MPS formatted.
	DATA_FORMAT_LP        DataFormat = 2 // The data file is LP formatted.
	DATA_FORMAT_OP        DataFormat = 3 // The data file is an optimization problem formatted file.
	DATA_FORMAT_FREE_MPS  DataFormat = 4 // The data a free MPS formatted file.
	DATA_FORMAT_TASK      DataFormat = 5 // Generic task dump file.
	DATA_FORMAT_PTF       DataFormat = 6 // (P)retty (T)ext (F)format.
	DATA_FORMAT_CB        DataFormat = 7 // Conic benchmark format,
	DATA_FORMAT_JSON_TASK DataFormat = 8 // JSON based task format.
)

var _DataFormat_map = map[DataFormat]string{
	DATA_FORMAT_EXTENSION: "DATA_FORMAT_EXTENSION",
	DATA_FORMAT_MPS:       "DATA_FORMAT_MPS",
	DATA_FORMAT_LP:        "DATA_FORMAT_LP",
	DATA_FORMAT_OP:        "DATA_FORMAT_OP",
	DATA_FORMAT_FREE_MPS:  "DATA_FORMAT_FREE_MPS",
	DATA_FORMAT_TASK:      "DATA_FORMAT_TASK",
	DATA_FORMAT_PTF:       "DATA_FORMAT_PTF",
	DATA_FORMAT_CB:        "DATA_FORMAT_CB",
	DATA_FORMAT_JSON_TASK: "DATA_FORMAT_JSON_TASK",
}

func (e DataFormat) String() string {
	if v, ok := _DataFormat_map[e]; ok {
		return v
	}
	return "DataFormat(" + strconv.FormatInt(int64(e), 10) + ")"
}
