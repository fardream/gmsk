package gmsk

import (
	"errors"

	"github.com/fardream/gmsk/res"
)

// MskError wraps the response codes [res.Code].
type MskError struct {
	code res.Code
}

var _ error = (*MskError)(nil)

// NewError creates an error from [res.Code]. The returned error will be nil for [res.OK].
func NewError(code res.Code) *MskError {
	if code.IsOk() {
		return nil
	}

	return &MskError{code: code}
}

func (err MskError) Error() string {
	return err.code.String()
}

// Ok checks if the error is nil or if the contained [res.Code] is [res.OK].
func (err MskError) Ok() bool {
	return err.code.IsOk()
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
	if is && target != nil && target.Ok() {
		target = nil
	}

	return target, is
}
