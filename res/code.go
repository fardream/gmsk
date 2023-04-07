// res package contains the MSKrescodee
// of MOSEK. This is to prevent the so many codes
// from poluting the name space of gmsk.
package res

// Code is the return code from MOSEK routines
type Code uint32

// IsOk checks if the result is ok.
func (c Code) IsOk() bool {
	return c == OK
}

// NotOk checks if the result is not ok.
func (c Code) NotOk() bool {
	return c != OK
}

func (c Code) GetMsg() (string, bool) {
	s, f := resCodeMsg[c]
	return s, f
}
