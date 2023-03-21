package gmsk_test

import (
	"log"
	"os"

	"github.com/fardream/gmsk"
)

func ExampleTask_WriteDataHandle() {
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

	CheckOk(task.WriteDataHandle(os.Stdout, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE))
	// Output: Task
	//     # Written by MOSEK v10.0.38
	//     # problemtype: Linear Problem
	//     # number of linear variables: 1
	//     # number of linear constraints: 0
	//     # number of  old-style cones: 0
	//     # number of positive semidefinite variables: 0
	//     # number of positive semidefinite matrixes: 0
	//     # number of affine conic constraints: 0
	//     # number of disjunctive constraints: 0
	//     # number of old-style A nonzeros: 0
	// Objective
	//     Minimize + @x0
	// Constraints
	// Variables
	//     @x0[2;3]
}
