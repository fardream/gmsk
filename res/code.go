// res package contains the MSKrescodee or [response code]
// of MOSEK. This is to prevent the so many codes
// from poluting the name space of gmsk.
//
// [response code]: https://docs.mosek.com/latest/capi/response-codes.html
package res

// IsOk checks if the result is ok.
func (c Code) IsOk() bool {
	return c == OK
}

// NotOk checks if the result is not ok.
func (c Code) NotOk() bool {
	return c != OK
}
