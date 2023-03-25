package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Example of mixed integer linear optimization, reproduced from milo1.c in MOSEK C api.
func Example_mixedIntegeLinearOptimization1() {
	checkOk := func(r gmsk.ResCode) {
		if !r.IsOk() {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	const numvar, numcon int32 = 2, 2

	c := []float64{1, 0.64}
	bkc := []gmsk.BoundKey{gmsk.BK_UP, gmsk.BK_LO}
	blc := []float64{-gmsk.INFINITY, -4}
	buc := []float64{250, gmsk.INFINITY}

	bkx := []gmsk.BoundKey{gmsk.BK_LO, gmsk.BK_LO}
	blx := []float64{0, 0}
	bux := []float64{gmsk.INFINITY, gmsk.INFINITY}

	aptrb := []int32{0, 2}
	aptre := []int32{2, 4}
	asub := []int32{0, 1, 0, 1}
	aval := []float64{50, 3, 31, -2}

	var i, j int32
	r := gmsk.RES_OK

	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := env.MakeTask(0, 0)
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	/* Append 'numcon' empty constraints.
	   The constraints will initially have no bounds. */
	checkOk(task.AppendCons(numcon))

	/* Append 'numvar' variables.
	   The variables will initially be fixed at zero (x=0). */
	checkOk(task.AppendVars(numvar))

	/* Optionally add a constant term to the objective. */
	checkOk(task.PutCFix(0))

	for j = 0; j < numvar && r.IsOk(); j++ {
		/* Set the linear term c_j in the objective.*/
		checkOk(task.PutCj(j, c[j]))

		/* Set the bounds on variable j.
		   blx[j] <= x_j <= bux[j] */

		checkOk(task.PutVarbound(
			j,       /* Index of variable.*/
			bkx[j],  /* Bound key.*/
			blx[j],  /* Numerical value of lower bound.*/
			bux[j])) /* Numerical value of upper bound.*/

		/* Input column j of A */
		if aptre[j]-aptrb[j] > 0 {
			r = task.PutACol(
				j,                 /* Variable (column) index.*/
				aptre[j]-aptrb[j], /* Number of non-zeros in column j.*/
				&asub[aptrb[j]],   /* Pointer to row indexes of column j.*/
				&aval[aptrb[j]])   /* Pointer to Values of column j.*/
		}
	}

	checkOk(r)

	/* Set the bounds on constraints.
	   for i=1, ...,numcon : blc[i] <= constraint i <= buc[i] */
	for i = 0; i < numcon && r.IsOk(); i++ {
		r = task.PutConbound(
			i,      /* Index of constraint.*/
			bkc[i], /* Bound key.*/
			blc[i], /* Numerical value of lower bound.*/
			buc[i]) /* Numerical value of upper bound.*/
	}
	checkOk(r)

	/* Specify integer variables. */
	for j = 0; j < numvar && r.IsOk(); j++ {
		r = task.PutVarType(j, gmsk.VAR_TYPE_INT)
	}
	checkOk(r)

	checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	/* Set max solution time */
	checkOk(task.PutDouParam(gmsk.DPAR_MIO_MAX_TIME, 60))

	/* Run optimizer */
	r, trmcode := task.OptimizeTerm()

	/* Print a summary containing information
	   about the solution for debugging purposes*/
	task.SolutionSummary(gmsk.STREAM_LOG)

	checkOk(r)

	r, solsta := task.GetSolSta(gmsk.SOL_ITG)
	checkOk(r)

	xx := make([]float64, numvar)

	switch solsta {
	case gmsk.SOL_STA_INTEGER_OPTIMAL:
		r, xx = task.GetXx(
			gmsk.SOL_ITG, /* Request the integer solution. */
			xx)
		checkOk(r)
		fmt.Printf("Optimal solution.\n")
		for j = 0; j < numvar; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}

	case gmsk.SOL_STA_PRIM_FEAS:
		/* A feasible but not necessarily optimal solution was located. */
		r, xx = task.GetXx(gmsk.SOL_ITG, xx)
		checkOk(r)
		fmt.Printf("Feasible solution.\n")
		for j = 0; j < numvar; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}

	case gmsk.SOL_STA_UNKNOWN:
		r, prosta := task.GetProSta(gmsk.SOL_ITG)
		checkOk(r)
		switch prosta {
		case gmsk.PRO_STA_PRIM_INFEAS_OR_UNBOUNDED:
			fmt.Printf("Problem status Infeasible or unbounded\n")
		case gmsk.PRO_STA_PRIM_INFEAS:
			fmt.Printf("Problem status Infeasible.\n")
		case gmsk.PRO_STA_UNKNOWN:
			fmt.Printf("Problem status unknown. Termination code %d.\n", trmcode)
		default:
			fmt.Printf("Other problem status.")
		}
	default:
		fmt.Printf("Other solution status.")
	}
	// Output:
	// Optimal solution.
	// x[0]: 5.000000e+00
	// x[1]: 0.000000e+00
}
