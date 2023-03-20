package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

func Example_AffineConicConstraints() {
	/* Input data dimensions */
	var n gmsk.Int32t = 3
	var k gmsk.Int64t = 2

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

	checkOk := func(r gmsk.ResCode) {
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
		c := []gmsk.Realt{2.0, 3.0, -1.0}
		checkOk(task.PutCSlice(0, n, &c[0]))
		checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))
	}

	/* One linear constraint sum(x) == 1 */
	checkOk(task.AppendCons(1))
	checkOk(task.PutConBound(0, gmsk.BK_FX, 1, 1))
	for i := gmsk.Int32t(0); i < n; i++ {
		checkOk(task.PutAij(0, i, 1))
	}

	/* Append empty AFE rows for affine expression storage */
	checkOk(task.AppendAfes(k + 1))

	{
		/* Fill in the affine expression storage with data */
		/* F matrix in sparse form */
		Fsubi := []gmsk.Int64t{1, 1, 2, 2} /* G is placed from row 1 of F */
		Fsubj := []gmsk.Int32t{0, 1, 0, 2}
		Fval := []gmsk.Realt{1.5, 0.1, 0.3, 2.1}
		var numEntries gmsk.Int64t = 4

		h := []gmsk.Realt{0, 0.1}
		var gamma gmsk.Realt = 0.03

		/* Fill in F storage */
		checkOk(task.PutAfeFEntryList(numEntries, &Fsubi[0], &Fsubj[0], &Fval[0]))

		/* Fill in g storage */
		checkOk(task.PutAfeG(0, gamma))
		checkOk(task.PutAfeGSlice(1, k+1, &h[0]))
	}

	/* Define a conic quadratic domain */
	r, quadDom := task.AppendQuadraticConeDomain(k + 1)
	checkOk(r)

	{
		/* Create the ACC */
		afeidx := []gmsk.Int64t{0, 1, 2}
		checkOk(task.AppendAcc(quadDom, k+1, &afeidx[0], nil))
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
		/* Fetch the solution */
		xx := make([]gmsk.Realt, n)
		r, xx = task.GetXx(
			gmsk.SOL_ITR, /* Request the interior solution. */
			xx)
		checkOk(r)
		fmt.Println("Optimal primal solution")
		for j := gmsk.Int32t(0); j < n; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}

		/* Fetch the doty dual of the ACC */
		doty := make([]gmsk.Realt, k+1)
		r, doty = task.GetAccDotY(
			gmsk.SOL_ITR, /* Request the interior solution. */
			0,            /* ACC index. */
			doty)
		checkOk(r)

		fmt.Println("Dual doty of the ACC")
		for j := gmsk.Int64t(0); j < k+1; j++ {
			fmt.Printf("doty[%d]: %e\n", j, doty[j])
		}

		/* Fetch the activity of the ACC */
		activity := make([]gmsk.Realt, k+1)
		r, activity = task.EvaluateAcc(
			gmsk.SOL_ITR, /* Request the interior solution. */
			0,            /* ACC index. */
			activity)
		checkOk(r)
		fmt.Println("Activity of the ACC")
		for j := gmsk.Int64t(0); j < k+1; j++ {
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
