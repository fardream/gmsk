package gmsk

import (
	"errors"

	"github.com/fardream/gmsk/res"
)

// MskError is [res.Error] and wraps the response codes [res.Code].
type MskError = res.Error

// NewError creates an error from [res.Code]. The returned error will be nil for [res.OK].
// The underlying type of the error is a [MskError].
func NewError(code res.Code) error {
	return code.ToError()
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
