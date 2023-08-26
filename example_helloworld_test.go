package gmsk_test

import (
	"fmt"
	"log"

	"github.com/fardream/gmsk"
)

// This is reproduced from MOSEK C example hellowworld.c
// However, the response code is checked here.
func Example_helloworld() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	task, err := gmsk.MakeTask(nil, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.AppendVars(1))
	checkOk(task.PutCJ(0, 1.0))
	checkOk(task.PutVarBound(0, gmsk.BK_RA, 2.0, 3.0))
	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MINIMIZE))
	_, res := task.OptimizeTrm()
	checkOk(res)
	result := make([]float64, 1)
	result, res = task.GetXx(gmsk.SOL_ITR, result)

	fmt.Printf("Solution x = %.6f\n", result[0])

	// Output: Solution x = 2.000000
}
