// Automatically generated by github.com/fardream/gen-gmsk
// enum for MSKmiodatapermmethod_enum/MioDataPermMethod

package gmsk

// #include <mosek.h>
import "C"

import "strconv"

// MioDataPermMethod is MSKmiodatapermmethod_enum.
//
// Specifies the problem data permutation method for mixed-integer problems.
type MioDataPermMethod uint32

const (
	MIO_DATA_PERMUTATION_METHOD_NONE         MioDataPermMethod = C.MSK_MIO_DATA_PERMUTATION_METHOD_NONE         // No problem data permutation is applied.
	MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT MioDataPermMethod = C.MSK_MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT // A random cyclic shift is applied to permute the problem data.
	MIO_DATA_PERMUTATION_METHOD_RANDOM       MioDataPermMethod = C.MSK_MIO_DATA_PERMUTATION_METHOD_RANDOM       // A random permutation is applied to the problem data.
)

var _MioDataPermMethod_map = map[MioDataPermMethod]string{
	MIO_DATA_PERMUTATION_METHOD_NONE:         "MIO_DATA_PERMUTATION_METHOD_NONE",
	MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT: "MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT",
	MIO_DATA_PERMUTATION_METHOD_RANDOM:       "MIO_DATA_PERMUTATION_METHOD_RANDOM",
}

func (e MioDataPermMethod) String() string {
	if v, ok := _MioDataPermMethod_map[e]; ok {
		return v
	}
	return "MioDataPermMethod(" + strconv.FormatInt(int64(e), 10) + ")"
}
