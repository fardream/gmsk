package gmsk_test

import (
	"testing"

	"github.com/fardream/gmsk"
)

func TestGetCodeDesc(t *testing.T) {
	_, a, b := gmsk.GetCodeDescSimple(gmsk.RES_OK)
	if a != "MSK_RES_OK" || b != "MSK_RES_OK" {
		t.Fatalf("description for RES_OK is %s %s", a, b)
	}
}
