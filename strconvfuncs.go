package gmsk

// #include <mosek.h>
import "C"

import "unsafe"

// MSKWchar is the char type of mosek's wide chars, or MSKwchart
type MSKWchar C.MSKwchart

// Utf8towchar is wrapping [MSK_utf8towchar]
//
// [MSK_utf8towchar]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.utf8towchar
func Utf8towchar(
	outputlen uint64,
	len []uint64,
	conv []uint64,
	output []MSKWchar,
	input string,
) error {
	c_input := C.CString(input)
	defer C.free(unsafe.Pointer(c_input))

	return ResCode(
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
	input []MSKWchar,
) error {
	return ResCode(
		C.MSK_wchartoutf8(
			C.size_t(outputlen),
			(*C.size_t)(getPtrToFirst(len)),
			(*C.size_t)(getPtrToFirst(conv)),
			(*C.char)(unsafe.Pointer(output)),
			(*C.MSKwchart)(getPtrToFirst(input)),
		),
	).ToError()
}
