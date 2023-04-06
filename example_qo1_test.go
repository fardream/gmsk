package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Quadratic optimization example, reproduced from qo1.c in MOSEK C api.
func Example_quadraticOptimization_qo1() {
	checkOk := func(r gmsk.ResCode) {
		if !r.IsOk() {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	const NUMCON = 1 /* Number of constraints.             */
	const NUMVAR = 3 /* Number of variables.               */
	const NUMQNZ = 4 /* Number of non-zeros in Q.          */

	c := []float64{0.0, -1.0, 0.0}

	bkc := []gmsk.BoundKey{gmsk.BK_LO}
	blc := []float64{1.0}
	buc := []float64{+gmsk.INFINITY}

	bkx := []gmsk.BoundKey{
		gmsk.BK_LO,
		gmsk.BK_LO,
		gmsk.BK_LO,
	}
	blx := []float64{
		0.0,
		0.0,
		0.0,
	}
	bux := []float64{
		+gmsk.INFINITY,
		+gmsk.INFINITY,
		+gmsk.INFINITY,
	}

	aptrb := []int32{0, 1, 2}
	aptre := []int32{1, 2, 3}
	asub := []int32{0, 0, 0}
	aval := []float64{1.0, 1.0, 1.0}

	qsubi := [NUMQNZ]int32{}
	qsubj := [NUMQNZ]int32{}
	qval := [NUMQNZ]float64{}

	var i, j int32
	xx := make([]float64, NUMVAR)

	r := gmsk.RES_OK
	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := env.MakeTask(NUMCON, NUMVAR)
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	/* Append 'NUMCON' empty constraints.
	   The constraints will initially have no bounds. */
	checkOk(task.AppendCons(NUMCON))

	/* Append 'NUMVAR' variables.
	   The variables will initially be fixed at zero (x=0). */
	checkOk(task.AppendVars(NUMVAR))

	/* Optionally add a constant term to the objective. */
	checkOk(task.PutCFix(0))

	for j = 0; j < NUMVAR && r.IsOk(); j++ {
		/* Set the linear term c_j in the objective.*/
		checkOk(task.PutCj(j, c[j]))

		/* Set the bounds on variable j.
		   blx[j] <= x_j <= bux[j] */
		checkOk(
			task.PutVarbound(
				j,      /* Index of variable.*/
				bkx[j], /* Bound key.*/
				blx[j], /* Numerical value of lower bound.*/
				bux[i], /* Numerical value of upper bound.*/
			))
		/* Input column j of A */
		r = task.PutACol(
			j,                 /* Variable (column) index.*/
			aptre[j]-aptrb[j], /* Number of non-zeros in column j.*/
			&asub[aptrb[j]],   /* Pointer to row indexes of column j.*/
			&aval[aptrb[j]])   /* Pointer to Values of column j.*/
	}

	checkOk(r)
	/* Set the bounds on constraints.
	   for i=1, ...,NUMCON : blc[i] <= constraint i <= buc[i] */
	for i = 0; i < NUMCON && r.IsOk(); i++ {
		r = task.PutConbound(
			i,      /* Index of constraint.*/
			bkc[i], /* Bound key.*/
			blc[i], /* Numerical value of lower bound.*/
			buc[i]) /* Numerical value of upper bound.*/
	}

	checkOk(r)

	/*
	 * The lower triangular part of the Q
	 * matrix in the objective is specified.
	 */

	qsubi[0] = 0
	qsubj[0] = 0
	qval[0] = 2.0
	qsubi[1] = 1
	qsubj[1] = 1
	qval[1] = 0.2
	qsubi[2] = 2
	qsubj[2] = 0
	qval[2] = -1.0
	qsubi[3] = 2
	qsubj[3] = 2
	qval[3] = 2.0

	/* Input the Q for the objective. */
	checkOk(task.PutQObj(NUMQNZ, &qsubi[0], &qsubj[0], &qval[0]))

	/* Run optimizer */
	r, trmcode := task.OptimizeTerm()

	/* Print a summary containing information
	   about the solution for debugging purposes*/
	task.SolutionSummary(gmsk.STREAM_LOG)

	r, solsta := task.GetSolSta(gmsk.SOL_ITR)
	checkOk(r)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		r, xx = task.GetXx(
			gmsk.SOL_ITR, /* Request the interior solution. */
			xx)
		if r != gmsk.RES_OK {
			r = gmsk.RES_ERR_SPACE
			break
		}
		fmt.Print("Optimal primal solution\n")
		for j = 0; j < NUMVAR; j++ {
			fmt.Printf("x[%d]: %.2e\n", j, xx[j])
		}
	case gmsk.SOL_STA_DUAL_INFEAS_CER:
		fallthrough
	case gmsk.SOL_STA_PRIM_INFEAS_CER:
		fmt.Printf("Primal or dual infeasibility certificate found.\n")
	case gmsk.SOL_STA_UNKNOWN:
		fmt.Printf("The status of the solution could not be determined. Termination code: %d.\n", trmcode)
	default:
		fmt.Printf("Other solution status.\n")
	}
	// Output:
	// Optimal primal solution
	// x[0]: 1.58e-04
	// x[1]: 5.00e+00
	// x[2]: 1.58e-04
}
