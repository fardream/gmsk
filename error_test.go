package gmsk_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fardream/gmsk"
	"github.com/fardream/gmsk/res"
)

func TestMosekError_Ok(t *testing.T) {
	e := fmt.Errorf("%w", gmsk.NewError(res.ERR_API_INTERNAL))
	ev := &gmsk.MskError{}
	if !errors.As(e, &ev) {
		t.Errorf("%#v is not a mosek error", e)
	}
}
