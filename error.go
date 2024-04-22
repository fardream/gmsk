package gmsk

import "errors"

// Error wraps the response codes [ResCode].
type MskError struct {
	code ResCode
}

var _ error = (*MskError)(nil)

func (err MskError) Error() string {
	return err.code.String()
}

// Ok checks if the error is nil or if the contained [ResCode] is [RES_OK].
func (err MskError) Ok() bool {
	return err.code.IsOk()
}

func (code ResCode) ToError() error {
	if code.IsOk() {
		return nil
	}

	return &MskError{code: code}
}

func NewErrorFromInt[TInt ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int8 | ~uint8 | ~int16 | ~uint16](code TInt) error {
	return ResCode(code).ToError()
}

// IsMskError checks if the err wraps [MskError]
func IsMskError(err error) bool {
	_, is := AsMskError(err)

	return is
}

// AsMskError converts an error into a [MskError] if err wraps a [MskError].
// The second returned bool indicates if the operation is successful.
func AsMskError(err error) (*MskError, bool) {
	target := &MskError{}
	is := errors.As(err, target)
	if is && target.Ok() {
		target = nil
	}

	return target, is
}

// Convert the error into a [ResCode]
func (err MskError) ToResCode() ResCode {
	return err.code
}

// NewError creates an error from [res.Code]. The returned error will be nil for [res.OK].
// The underlying type of the error is a [MskError].
func NewError(code ResCode) error {
	return code.ToError()
}

// IsOk checks if the result is ok.
func (c ResCode) IsOk() bool {
	return c == RES_OK
}

// NotOk checks if the result is not ok.
func (c ResCode) NotOk() bool {
	return c != RES_OK
}
