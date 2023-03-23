package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Affine conic constraints example 2, reproduced from acc2.c in MOSEK C Api.
//
// Purpose :   Tutorial example for affine conic constraints.
//
// Models the problem:
//
// maximize c^T x
//
// subject to  sum(x) = 1
//
// gamma >= |Gx+h|_2
//
// This version inputs the linear constraint as an affine conic constraint.
func Example_affineConicConstraints2() {
	r := gmsk.RES_OK
	/* Input data dimensions */
	var n int32 = 3
	var k int64 = 2

	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := gmsk.MakeTask(env, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk := func(r uint32) {
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodeDescSimple(r)

			log.Fatalf("failed: %s %s", sym, desc)
		}
	}

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	/* Create n free variables */
	checkOk(task.AppendVars(n))
	checkOk(task.PutVarboundSliceConst(0, n, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	/* Set up the objective */
	{
		c := []float64{2.0, 3.0, -1.0}
		checkOk(task.PutCSlice(0, n, &c[0]))
		checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))
	}

	{
		/* Set AFE rows representing the linear constraint */
		checkOk(task.AppendAfes(1))
		for i := int32(0); i < n && r == gmsk.RES_OK; i++ {
			r = task.PutAfeFEntry(0, i, 1)
		}
		checkOk(r)
		checkOk(task.PutAfeG(0, -1))
	}

	{
		/* Set AFE rows representing the quadratic constraint */
		/* F matrix in sparse form */
		Fsubi := []int64{2, 2, 3, 3} /* G is placed from row 2 of F */
		Fsubj := []int32{0, 1, 0, 2}
		Fval := []float64{1.5, 0.1, 0.3, 2.1}
		var numEntries int64 = 4
		/* Other data */
		h := []float64{0, 0.1}
		var gamma float64 = 0.03

		checkOk(task.AppendAfes(k + 1))
		checkOk(task.PutAfeFEntryList(numEntries, &Fsubi[0], &Fsubj[0], &Fval[0]))
		checkOk(task.PutAfeG(1, gamma))
		checkOk(task.PutAfeGSlice(2, k+2, &h[0]))
	}

	r, zeroDom := task.AppendRZeroDomain(1)
	checkOk(r)
	/* Define a conic quadratic domain */
	r, quadDom := task.AppendQuadraticConeDomain(k + 1)
	checkOk(r)

	/* Append affine conic constraints */
	{
		/* Linear constraint */
		afeidx := []int64{0}

		checkOk(task.AppendAcc(
			zeroDom,    /* Domain index */
			1,          /* Dimension */
			&afeidx[0], /* Indices of AFE rows */
			nil),       /* Ignored */
		)
	}

	{
		/* Quadratic constraint */
		afeidx := []int64{1, 2, 3}

		checkOk(task.AppendAcc(
			quadDom,    /* Domain index */
			k+1,        /* Dimension */
			&afeidx[0], /* Indices of AFE rows */
			nil),       /* Ignored */
		)
	}

	/* Begin optimization and fetching the solution */
	r, trmcode := task.OptimizeTerm()
	checkOk(r)

	/* Print a summary containing information
	   about the solution for debugging purposes*/
	task.SolutionSummary(gmsk.STREAM_LOG) // use stream log and direct it to stderr

	r, solsta := task.GetSolSta(gmsk.SOL_ITR)
	checkOk(r)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		/* Fetch the primal solution */
		xx := make([]float64, n)
		r, xx = task.GetXx(
			gmsk.SOL_ITR, /* Request the interior solution. */
			xx)
		checkOk(r)
		fmt.Println("Optimal primal solution")
		for j := int32(0); j < n; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}

		/* Fetch the doty dual of the ACC */
		doty := make([]float64, k+1)
		r, doty = task.GetAccDotY(
			gmsk.SOL_ITR, /* Request the interior solution. */
			1,            /* ACC index of quadratic ACC. */
			doty)
		checkOk(r)

		fmt.Println("Dual doty of the ACC")
		for j := int64(0); j < k+1; j++ {
			fmt.Printf("doty[%d]: %e\n", j, doty[j])
		}

		/* Fetch the activity of the ACC */
		activity := make([]float64, k+1)
		r, activity = task.EvaluateAcc(
			gmsk.SOL_ITR, /* Request the interior solution. */
			1,            /* ACC index. */
			activity)
		checkOk(r)
		fmt.Println("Activity of the ACC")
		for j := int64(0); j < k+1; j++ {
			fmt.Printf("activity[%d]: %e\n", j, activity[j])
		}
	case gmsk.SOL_STA_DUAL_INFEAS_CER:
		fallthrough
	case gmsk.SOL_STA_PRIM_INFEAS_CER:
		fmt.Println("Primal or dual infeasibility certificate found.")
	case gmsk.SOL_STA_UNKNOWN:
		fmt.Printf("The status of the solution could not be determined. Termination code: %d.\n", trmcode)
	default:
		fmt.Println("Other solution status.")
	}

	// Output: Optimal primal solution
	// x[0]: -7.838011e-02
	// x[1]: 1.128913e+00
	// x[2]: -5.053279e-02
	// Dual doty of the ACC
	// doty[0]: -1.942968e+00
	// doty[1]: -3.030303e-01
	// doty[2]: -1.919192e+00
	// Activity of the ACC
	// activity[0]: 3.000000e-02
	// activity[1]: -4.678877e-03
	// activity[2]: -2.963289e-02
}
