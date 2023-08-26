package gmsk_test

import (
	"testing"

	"github.com/fardream/gmsk"
)

func TestGetCodedesc(t *testing.T) {
	a, b, _ := gmsk.GetCodedesc(gmsk.RES_OK)
	if a != "MSK_RES_OK" || b != "No error occurred." {
		t.Fatalf("description for RES_OK is %s %s", a, b)
	}
}
