// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <mosek.h>
import "C"

import (
	"unsafe"

	"github.com/fardream/gmsk/res"
)

// We don't know if unsafe will be used or not, so
var _ any = unsafe.Pointer(nil)

// Callbackcodetostr is wrapping MSK_callbackcodetostr
//
// [MSK_callbackcodetostr] has following parameters
//   - code: MSKcallbackcodee
//   - callbackcodestr: char *
//
// [MSK_callbackcodetostr]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func Callbackcodetostr(
	code CallbackCode,
	callbackcodestr *byte,
) res.Code {
	return res.Code(
		C.MSK_callbackcodetostr(
			C.MSKcallbackcodee(code),
			(*C.char)(unsafe.Pointer(callbackcodestr)),
		),
	)
}

// GetBuildinfo is wrapping MSK_getbuildinfo
//
// [MSK_getbuildinfo] has following parameters
//   - buildstate: char *
//   - builddate: char *
//
// [MSK_getbuildinfo]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func GetBuildinfo(
	buildstate *byte,
	builddate *byte,
) res.Code {
	return res.Code(
		C.MSK_getbuildinfo(
			(*C.char)(unsafe.Pointer(buildstate)),
			(*C.char)(unsafe.Pointer(builddate)),
		),
	)
}

// GetCodedesc is wrapping MSK_getcodedesc
//
// [MSK_getcodedesc] has following parameters
//   - code: MSKrescodee
//   - symname: char *
//   - str: char *
//
// [MSK_getcodedesc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func GetCodedesc(
	code res.Code,
	symname *byte,
	str *byte,
) res.Code {
	return res.Code(
		C.MSK_getcodedesc(
			C.MSKrescodee(code),
			(*C.char)(unsafe.Pointer(symname)),
			(*C.char)(unsafe.Pointer(str)),
		),
	)
}

// GetResponseclass is wrapping MSK_getresponseclass
//
// [MSK_getresponseclass] has following parameters
//   - r: MSKrescodee
//   - rc: MSKrescodetypee *
//
// [MSK_getresponseclass]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func GetResponseclass(
	r res.Code,
	rc *ResCodeType,
) res.Code {
	return res.Code(
		C.MSK_getresponseclass(
			C.MSKrescodee(r),
			(*C.MSKrescodetypee)(rc),
		),
	)
}

// GetVersion is wrapping MSK_getversion
//
// [MSK_getversion] has following parameters
//   - major: MSKint32t *
//   - minor: MSKint32t *
//   - revision: MSKint32t *
//
// [MSK_getversion]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func GetVersion(
	major *int32,
	minor *int32,
	revision *int32,
) res.Code {
	return res.Code(
		C.MSK_getversion(
			(*C.MSKint32t)(major),
			(*C.MSKint32t)(minor),
			(*C.MSKint32t)(revision),
		),
	)
}

// Isinfinity is wrapping MSK_isinfinity
//
// [MSK_isinfinity] has following parameters
//   - value: MSKrealt
//
// [MSK_isinfinity]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func Isinfinity(
	value float64,
) int32 {
	return int32(
		C.MSK_isinfinity(
			C.MSKrealt(value),
		),
	)
}

// Licensecleanup is wrapping MSK_licensecleanup
//
// [MSK_licensecleanup] has following parameters
//
// [MSK_licensecleanup]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func Licensecleanup() res.Code {
	return res.Code(
		C.MSK_licensecleanup(),
	)
}

// Symnamtovalue is wrapping MSK_symnamtovalue
//
// [MSK_symnamtovalue] has following parameters
//   - name: const char *
//   - value: char *
//
// [MSK_symnamtovalue]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func Symnamtovalue(
	name *byte,
	value *byte,
) int32 {
	return int32(
		C.MSK_symnamtovalue(
			(*C.char)(unsafe.Pointer(name)),
			(*C.char)(unsafe.Pointer(value)),
		),
	)
}

// Utf8towchar is wrapping MSK_utf8towchar
//
// [MSK_utf8towchar] has following parameters
//   - outputlen: size_t
//   - len: size_t *
//   - conv: size_t *
//   - output: MSKwchart *
//   - input: const char *
//
// [MSK_utf8towchar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func Utf8towchar(
	outputlen uint64,
	len *uint64,
	conv *uint64,
	output *int32,
	input *byte,
) res.Code {
	return res.Code(
		C.MSK_utf8towchar(
			C.size_t(outputlen),
			(*C.size_t)(len),
			(*C.size_t)(conv),
			(*C.MSKwchart)(output),
			(*C.char)(unsafe.Pointer(input)),
		),
	)
}

// Wchartoutf8 is wrapping MSK_wchartoutf8
//
// [MSK_wchartoutf8] has following parameters
//   - outputlen: size_t
//   - len: size_t *
//   - conv: size_t *
//   - output: char *
//   - input: const MSKwchart *
//
// [MSK_wchartoutf8]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html
func Wchartoutf8(
	outputlen uint64,
	len *uint64,
	conv *uint64,
	output *byte,
	input *int32,
) res.Code {
	return res.Code(
		C.MSK_wchartoutf8(
			C.size_t(outputlen),
			(*C.size_t)(len),
			(*C.size_t)(conv),
			(*C.char)(unsafe.Pointer(output)),
			(*C.MSKwchart)(input),
		),
	)
}
