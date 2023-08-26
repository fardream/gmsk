package res

import "errors"

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

// IsMskError checks if the err wraps [MskError]
func IsMskError(err error) bool {
	_, is := AsMskError(err)

	return is
}

// AsMskError converts an error into a [MskError] if err wraps a [MskError].
// The second returned bool indicates if the operation is successful.
func AsMskError(err error) (*Error, bool) {
	target := &Error{}
	is := errors.As(err, target)
	if is && target != nil && target.Ok() {
		target = nil
	}

	return target, is
}

// Convert the error into a [Code]
func (err Error) ToResCode() Code {
	return err.code
}
