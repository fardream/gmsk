package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Mixed integer conic optimization example, reproduced from mico1.c from MOSEK C api.
//
//	Purpose :   Demonstrates how to solve a small mixed
//	            integer conic optimization problem.
//
//	            minimize    x^2 + y^2
//	            subject to  x >= e^y + 3.8
//	                        x, y - integer
func Example_mixedIntegerConicOptimization_mico1() {
	checkOk := func(r gmsk.ResCode) {
		if !r.IsOk() {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	var numvar int32 = 3 /* x, y, t */

	vart := []gmsk.VariableType{gmsk.VAR_TYPE_INT, gmsk.VAR_TYPE_INT}
	intsub := []int32{0, 1}

	var r gmsk.ResCode

	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	task, err := env.MakeTask(0, 0)
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	checkOk(task.AppendVars(numvar))

	checkOk(task.PutVarboundSliceConst(0, numvar, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	/* Integrality constraints */
	checkOk(task.PutVarTypeList(2, &intsub[0], &vart[0]))

	/* Objective */
	checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MINIMIZE))
	checkOk(task.PutCj(2, 1.0)) /* Minimize t */

	/* Set up the affine expressions */
	/* x, x-3.8, y, t, 1.0 */
	afeidx := []int64{0, 1, 2, 3}
	varidx := []int32{0, 0, 1, 2}
	val := []float64{1.0, 1.0, 1.0, 1.0}
	g := []float64{0.0, -3.8, 0.0, 0.0, 1.0}
	afeidxExp := []int64{1, 4, 2}
	afeidxQuad := []int64{3, 0, 2}

	checkOk(task.AppendAfes(5))

	checkOk(task.PutAfeFEntryList(4, &afeidx[0], &varidx[0], &val[0]))

	checkOk(task.PutAfeGSlice(0, 5, &g[0]))

	// Add constraint (x-3.8, 1, y) \in \EXP
	r, domExp := task.AppendPrimalExpConeDomain()
	checkOk(r)
	checkOk(task.AppendAcc(domExp, 3, &afeidxExp[0], nil))

	// Add constraint (t, x, y) \in \QUAD
	r, domQuad := task.AppendQuadraticConeDomain(3)
	checkOk(r)
	checkOk(task.AppendAcc(domQuad, 3, &afeidxQuad[0], nil))

	r, _ = task.OptimizeTerm()
	checkOk(r)

	xx := make([]float64, 2)

	r, xx = task.GetXxSlice(gmsk.SOL_ITG, 0, 2, xx)
	checkOk(r)

	fmt.Printf("x = %.2f, y = %.2f\n", xx[0], xx[1])
	// Output:
	// x = 4.00, y = -2.00
}
