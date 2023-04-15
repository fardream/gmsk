// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKparametertype_enum/ParameterType

package gmsk

import "strconv"

// ParameterType is MSKparametertype_enum.
//
// Parameter type
type ParameterType uint32

const (
	PAR_INVALID_TYPE ParameterType = 0 // Not a valid parameter.
	PAR_DOU_TYPE     ParameterType = 1 // Is a double parameter.
	PAR_INT_TYPE     ParameterType = 2 // Is an integer parameter.
	PAR_STR_TYPE     ParameterType = 3 // Is a string parameter.
)

var _ParameterType_map = map[ParameterType]string{
	PAR_INVALID_TYPE: "PAR_INVALID_TYPE",
	PAR_DOU_TYPE:     "PAR_DOU_TYPE",
	PAR_INT_TYPE:     "PAR_INT_TYPE",
	PAR_STR_TYPE:     "PAR_STR_TYPE",
}

func (e ParameterType) String() string {
	if v, ok := _ParameterType_map[e]; ok {
		return v
	}
	return "ParameterType(" + strconv.FormatInt(int64(e), 10) + ")"
}