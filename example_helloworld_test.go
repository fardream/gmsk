package gmsk_test

import (
	"fmt"
	"log"

	"github.com/fardream/gmsk"
)

// This is reproduced from MOSEK C example hellowworld.c
// However, the response code is checked here.
func Example_helloworld() {
	CheckOk := func(r gmsk.ResCode) {
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Fatalf("Failed: %s %s", sym, desc)
		}
	}

	task, err := gmsk.MakeTask(nil, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteTask(task)

	CheckOk(task.AppendVars(1))
	CheckOk(task.PutCj(0, 1.0))
	CheckOk(task.PutVarbound(0, gmsk.BK_RA, 2.0, 3.0))
	CheckOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MINIMIZE))
	res, _ := task.OptimizeTerm()
	CheckOk(res)
	result := make([]float64, 1)
	res, result = task.GetXx(gmsk.SOL_ITR, result)

	fmt.Printf("Solution x = %.6f\n", result[0])

	// Output: Solution x = 2.000000
}
