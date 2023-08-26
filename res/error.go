package res

// Error wraps the response codes [res.Code].
type Error struct {
	code Code
}

var _ error = (*Error)(nil)

func (err Error) Error() string {
	return err.code.String()
}

// Ok checks if the error is nil or if the contained [res.Code] is [res.OK].
func (err Error) Ok() bool {
	return err.code.IsOk()
}

func (code Code) ToError() error {
	if code.IsOk() {
		return nil
	}

	return &Error{code: code}
}

func NewErrorFromInt[TInt ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int8 | ~uint8 | ~int16 | ~uint16](code TInt) error {
	return Code(code).ToError()
}
