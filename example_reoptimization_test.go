package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Reoptimization example, reproduced from reoptimization.c in MOSEK C api.
func Example_reoptimization() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	printres := func(n int32, x []float64) {
		for i, v := range x {
			if int32(i) >= n-1 {
				fmt.Printf("%.3f\n", v)
			} else {
				fmt.Printf("%.3f ", v)
			}
		}
	}

	var numvar int32 = 3
	var numcon int32 = 3
	var i, j int32
	c := []float64{1.5, 2.5, 3.0}
	ptrb := []int32{0, 3, 6}
	ptre := []int32{3, 6, 9}
	asub := []int32{
		0, 1, 2,
		0, 1, 2,
		0, 1, 2,
	}

	aval := []float64{
		2.0, 3.0, 2.0,
		4.0, 2.0, 3.0,
		3.0, 3.0, 2.0,
	}

	bkc := []gmsk.BoundKey{gmsk.BK_UP, gmsk.BK_UP, gmsk.BK_UP}
	blc := []float64{-gmsk.INFINITY, -gmsk.INFINITY, -gmsk.INFINITY}
	buc := []float64{100000, 50000, 60000}

	bkx := []gmsk.BoundKey{gmsk.BK_LO, gmsk.BK_LO, gmsk.BK_LO}
	blx := []float64{0.0, 0.0, 0.0}
	bux := []float64{+gmsk.INFINITY, +gmsk.INFINITY, +gmsk.INFINITY}

	var xx []float64
	var varidx, conidx int32

	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := env.MakeTask(numcon, numvar)
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	/* Append the constraints. */
	checkOk(task.AppendCons(numcon))

	/* Append the variables. */
	checkOk(task.AppendVars(numvar))

	/* Put C. */
	checkOk(task.PutCfix(0))

	for j = 0; j < numvar; j++ {
		checkOk(task.PutCJ(j, c[j]))
	}

	/* Put constraint bounds. */
	for i = 0; i < numcon; i++ {
		checkOk(task.PutConBound(i, bkc[i], blc[i], buc[i]))
	}

	/* Put variable bounds. */
	for j = 0; j < numvar; j++ {
		checkOk(task.PutVarBound(j, bkx[j], blx[j], bux[j]))
	}

	/* Put A. */
	if numcon > 0 {
		for j = 0; j < numvar; j++ {
			checkOk(task.PutACol(
				j,
				ptre[j]-ptrb[j],
				asub[ptrb[j]:ptre[j]],
				aval[ptrb[j]:ptre[j]]))
		}
	}

	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	_, r = task.OptimizeTrm()
	checkOk(r)

	xx = make([]float64, numvar)

	xx, r = task.GetXx(
		gmsk.SOL_BAS, /* Basic solution.       */
		xx)
	checkOk(r)

	printres(numvar, xx)

	/******************** Make a change to the A matrix **********/
	checkOk(task.PutAij(0, 0, 3))
	_, r = task.OptimizeTrm()
	checkOk(r)

	xx, r = task.GetXx(
		gmsk.SOL_BAS, /* Basic solution.       */
		xx)
	checkOk(r)

	printres(numvar, xx)

	/*********************** Add a new variable ******************/
	/* Get index of new variable, this should be 3 */
	varidx, r = task.GetNumVar()
	/* Append a new variable x_3 to the problem */
	checkOk(task.AppendVars(1))
	numvar++
	/* Set bounds on new variable */
	task.PutVarBound(
		varidx,
		gmsk.BK_LO,
		0,
		gmsk.INFINITY)

	/* Change objective */
	checkOk(task.PutCJ(varidx, 1))

	/* Put new values in the A matrix */
	{
		acolsub := []int32{0, 2}
		acolval := []float64{4.0, 1.0}

		checkOk(task.PutACol(
			varidx, /* column index */
			2,      /* num nz in column*/
			acolsub,
			acolval))
	}

	/* Change optimizer to free simplex and reoptimize */
	checkOk(task.PutIntParam(gmsk.IPAR_OPTIMIZER, gmsk.OPTIMIZER_FREE_SIMPLEX))

	_, r = task.OptimizeTrm()
	checkOk(r)
	xx = make([]float64, numvar)
	xx, r = task.GetXx(
		gmsk.SOL_BAS, /* Basic solution.       */
		xx)
	checkOk(r)

	printres(numvar, xx)

	/* **************** Add a new constraint ******************* */
	/* Get index of new constraint*/
	conidx, r = task.GetNumCon()
	checkOk(r)

	/* Append a new constraint */
	checkOk(task.AppendCons(1))
	numcon++

	/* Set bounds on new constraint */
	checkOk(
		task.PutConBound(
			conidx,
			gmsk.BK_UP,
			-gmsk.INFINITY,
			30000))

	/* Put new values in the A matrix */
	{
		arowsub := []int32{0, 1, 2, 3}
		arowval := []float64{1.0, 2.0, 1.0, 1.0}

		checkOk(
			task.PutARow(
				conidx, /* row index */
				4,      /* num nz in row*/
				arowsub,
				arowval))
	}

	_, r = task.OptimizeTrm()
	checkOk(r)

	xx, r = task.GetXx(
		gmsk.SOL_BAS, /* Basic solution.       */
		xx)
	checkOk(r)

	printres(numvar, xx)

	/* **************** Change constraint bounds ******************* */
	{
		newbkc := []gmsk.BoundKey{gmsk.BK_UP, gmsk.BK_UP, gmsk.BK_UP, gmsk.BK_UP}
		newblc := []float64{-gmsk.INFINITY, -gmsk.INFINITY, -gmsk.INFINITY, -gmsk.INFINITY}
		newbuc := []float64{80000, 40000, 50000, 22000}

		checkOk(task.PutConBoundSlice(0, numcon, newbkc, newblc, newbuc))
	}

	_, r = task.OptimizeTrm()
	checkOk(r)

	xx, r = task.GetXx(
		gmsk.SOL_BAS, /* Basic solution.       */
		xx)
	checkOk(r)

	printres(numvar, xx)
	// Output:
	// 0.000 16000.000 6000.000
	// 0.000 16000.000 6000.000
	// 0.000 12142.857 8571.429 6428.571
	// 0.000 1000.000 16000.000 12000.000
	// 0.000 0.000 13333.333 8666.667
}
