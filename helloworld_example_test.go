package gmsk_test

import (
	"fmt"
	"log"

	"github.com/fardream/gmsk"
)

// This is reproduce of MOSEK C example hellowworld.c
// However, the response code is checked here.
func Example_helloworld() {
	CheckOk := func(r gmsk.ResCode) {
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Fatalf("Failed: %s %s", sym, desc)
		}
	}

	task, err := gmsk.MakeTask(nil, 0, 0)
	defer gmsk.DeleteTask(task)

	if err != nil {
		log.Fatal(err)
	}

	CheckOk(gmsk.AppendVars(task, 1))
	CheckOk(gmsk.PutCj(task, 0, 1.0))
	CheckOk(gmsk.PutVarbound(task, 0, gmsk.BK_RA, 2.0, 3.0))
	CheckOk(gmsk.PutObjsense(task, gmsk.OBJECTIVE_SENSE_MINIMIZE))
	res, _ := gmsk.OptimizeTerm(task)
	CheckOk(res)
	result := make([]float64, 1)
	res, result = gmsk.GetXx(task, gmsk.SOL_ITR, result)

	fmt.Printf("Solution x = %.6f\n", result[0])

	// Output: Solution x = 2.000000
}
