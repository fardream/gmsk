// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"unsafe"

	"github.com/fardream/gmsk/res"
)

// CallbackcodeToStr is wrapping [MSK_callbackcodetostr]
//
// [MSK_callbackcodetostr]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.callbackcodetostr
func CallbackcodeToStr(
	code CallbackCode,
) (callbackcodestr string, r error) {
	// function template: prepare for output of booleans
	c_callbackcodestr := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_callbackcodestr))

	r = res.Code(
		C.MSK_callbackcodetostr(
			C.MSKcallbackcodee(code),
			c_callbackcodestr,
		),
	).ToError()

	if r == nil {
		callbackcodestr = C.GoString(c_callbackcodestr)
	}

	return
}

// GetBuildInfo is wrapping [MSK_getbuildinfo]
//
// [MSK_getbuildinfo]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.getbuildinfo
func GetBuildInfo() (buildstate, builddate string, r error) {
	// function template: prepare for output of booleans
	c_buildstate := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_buildstate))
	c_builddate := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_builddate))

	r = res.Code(
		C.MSK_getbuildinfo(
			c_buildstate,
			c_builddate,
		),
	).ToError()

	if r == nil {
		buildstate = C.GoString(c_buildstate)
		builddate = C.GoString(c_builddate)
	}

	return
}

// GetCodedesc is wrapping [MSK_getcodedesc]
//
// [MSK_getcodedesc]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.getcodedesc
func GetCodedesc(
	code res.Code,
) (symname, str string, r error) {
	// function template: prepare for output of booleans
	c_symname := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_symname))
	c_str := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_str))

	r = res.Code(
		C.MSK_getcodedesc(
			C.MSKrescodee(code),
			c_symname,
			c_str,
		),
	).ToError()

	if r == nil {
		symname = C.GoString(c_symname)
		str = C.GoString(c_str)
	}

	return
}

// GetResponseclass is wrapping [MSK_getresponseclass]
//
// [MSK_getresponseclass]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.getresponseclass
func GetResponseclass(
	r res.Code,
) (rc ResCodeType, rescode error) {
	rescode = res.Code(
		C.MSK_getresponseclass(
			C.MSKrescodee(r),
			(*C.MSKrescodetypee)(&rc),
		),
	).ToError()

	return
}

// GetVersion is wrapping [MSK_getversion]
//
// [MSK_getversion]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.getversion
func GetVersion() (major, minor, revision int32, r error) {
	r = res.Code(
		C.MSK_getversion(
			(*C.MSKint32t)(&major),
			(*C.MSKint32t)(&minor),
			(*C.MSKint32t)(&revision),
		),
	).ToError()

	return
}

// Isinfinity is wrapping [MSK_isinfinity]
//
// [MSK_isinfinity]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.isinfinity
func Isinfinity(
	value float64,
) bool {
	return intToBool(
		C.MSK_isinfinity(
			C.MSKrealt(value),
		),
	)
}

// Licensecleanup is wrapping [MSK_licensecleanup]
//
// [MSK_licensecleanup]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.licensecleanup
func Licensecleanup() error {
	return res.Code(
		C.MSK_licensecleanup(),
	).ToError()
}

// Symnamtovalue is wrapping [MSK_symnamtovalue]
//
// [MSK_symnamtovalue]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.symnamtovalue
func Symnamtovalue(
	name string,
	value *byte,
) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return intToBool(
		C.MSK_symnamtovalue(
			c_name,
			(*C.char)(unsafe.Pointer(value)),
		),
	)
}

// Utf8towchar is wrapping [MSK_utf8towchar]
//
// [MSK_utf8towchar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.utf8towchar
func Utf8towchar(
	outputlen uint64,
	len []uint64,
	conv []uint64,
	output []int32,
	input string,
) error {
	c_input := C.CString(input)
	defer C.free(unsafe.Pointer(c_input))

	return res.Code(
		C.MSK_utf8towchar(
			C.size_t(outputlen),
			(*C.size_t)(getPtrToFirst(len)),
			(*C.size_t)(getPtrToFirst(conv)),
			(*C.MSKwchart)(getPtrToFirst(output)),
			c_input,
		),
	).ToError()
}

// Wchartoutf8 is wrapping [MSK_wchartoutf8]
//
// [MSK_wchartoutf8]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.wchartoutf8
func Wchartoutf8(
	outputlen uint64,
	len []uint64,
	conv []uint64,
	output *byte,
	input []int32,
) error {
	return res.Code(
		C.MSK_wchartoutf8(
			C.size_t(outputlen),
			(*C.size_t)(getPtrToFirst(len)),
			(*C.size_t)(getPtrToFirst(conv)),
			(*C.char)(unsafe.Pointer(output)),
			(*C.MSKwchart)(getPtrToFirst(input)),
		),
	).ToError()
}
