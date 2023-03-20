package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Linear programming example 1, reproduced from mosek c api example lo1.c
func Example_linearOptimization1() {
	const numvar, numcon gmsk.Int32t = 4, 3
	c := []gmsk.Realt{3.0, 1.0, 5.0, 1.0}

	/* Below is the sparse representation of the A
	   matrix stored by column. */
	aptrb := []gmsk.Int32t{0, 2, 5, 7}
	aptre := []gmsk.Int32t{2, 5, 7, 9}
	asub := []gmsk.Int32t{
		0, 1,
		0, 1, 2,
		0, 1,
		1, 2,
	}
	aval := []gmsk.Realt{
		3, 2,
		1, 1, 2,
		2, 3,
		1, 3,
	}
	/* Bounds on constraints. */
	bkc := []gmsk.BoundKey{gmsk.BK_FX, gmsk.BK_LO, gmsk.BK_UP}
	blc := []gmsk.Realt{30, 15, -gmsk.INFINITY}
	buc := []gmsk.Realt{30, gmsk.INFINITY, 25}

	/* Bounds on variables. */
	bkx := []gmsk.BoundKey{gmsk.BK_LO, gmsk.BK_RA, gmsk.BK_LO, gmsk.BK_LO}
	blx := []gmsk.Realt{0, 0, 0, 0}
	bux := []gmsk.Realt{gmsk.INFINITY, 10, gmsk.INFINITY, gmsk.INFINITY}

	checkOk := func(r gmsk.ResCode) {
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Fatalf("failed: %s %s", sym, desc)
		}
	}

	r := gmsk.RES_OK

	/* Create the optimization task. */
	task, err := gmsk.MakeTask(nil, numcon, numvar)
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteTask(task)

	/* Directs the log task stream to the 'printstr' function. */
	// Note here use os.Stderr to prevent the example from failing
	task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr)

	/* Append 'numcon' empty constraints.
	   The constraints will initially have no bounds. */
	checkOk(task.AppendCons(numcon))
	/* Append 'numvar' variables.
	   The variables will initially be fixed at zero (x=0). */
	checkOk(task.AppendVars(numvar))

	for j := gmsk.Int32t(0); j < numvar && r == gmsk.RES_OK; j++ {
		/* Set the linear term c_j in the objective.*/
		r = task.PutCj(j, c[j])
		if r != gmsk.RES_OK {
			break
		}

		r = task.PutVarbound(
			j,      /* Index of variable.*/
			bkx[j], /* Bound key.*/
			blx[j], /* Numerical value of lower bound.*/
			bux[j]) /* Numerical value of upper bound.*/
		if r != gmsk.RES_OK {
			break
		}

		/* Input column j of A */
		r = task.PutACol(
			j,                 /* Variable (column) index.*/
			aptre[j]-aptrb[j], /* Number of non-zeros in column j.*/
			&asub[aptrb[j]],   /* Pointer to row indexes of column j.*/
			&aval[aptrb[j]])   /* Pointer to Values of column j.*/
	}

	checkOk(r)

	/* Set the bounds on constraints.
	   for i=1, ...,numcon : blc[i] <= constraint i <= buc[i] */
	for i := gmsk.Int32t(0); i < numcon && r == gmsk.RES_OK; i++ {
		r = task.PutConBound(
			i,      /* Index of constraint.*/
			bkc[i], /* Bound key.*/
			blc[i], /* Numerical value of lower bound.*/
			buc[i]) /* Numerical value of upper bound.*/
	}

	checkOk(r)

	/* Maximize objective function. */
	checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	var trmcode gmsk.ResCode
	/* Run optimizer */
	r, trmcode = task.OptimizeTerm()
	checkOk(r)

	/* Print a summary containing information
	   about the solution for debugging purposes. */
	task.SolutionSummary(gmsk.STREAM_LOG)

	var solsta gmsk.SolSta

	r, solsta = task.GetSolSta(gmsk.SOL_BAS)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		r, xx := task.GetXx(
			gmsk.SOL_BAS, /* Request the basic solution. */
			nil)
		if r != gmsk.RES_OK {
			r = gmsk.RES_ERR_SPACE
			break
		}
		fmt.Print("Optimal primal solution\n")
		for j := gmsk.Int32t(0); j < numvar; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}
	case gmsk.SOL_STA_DUAL_INFEAS_CER:
		fallthrough
	case gmsk.SOL_STA_PRIM_INFEAS_CER:
		fmt.Printf("Primal or dual infeasibility certificate found.\n")
	case gmsk.SOL_STA_UNKNOWN:
		/* If the solutions status is unknown, print the termination code
		   indicating why the optimizer terminated prematurely. */
		_, symname, _ := gmsk.GetCodeDescSimple(trmcode)
		fmt.Printf("The solution status is unknown.\n")
		fmt.Printf("The optimizer terminitated with code: %s\n", symname)
	default:
		fmt.Printf("Other solution status.\n")
	}

	if r != gmsk.RES_OK {
		/* In case of an error print error code and description. */
		_, symname, desc := gmsk.GetCodeDescSimple(r)
		fmt.Printf("Error %s - '%s'\n", symname, desc)
	}

	/* Delete the task and the associated data. */
	// but we don't need to because of defer

	// Output: Optimal primal solution
	// x[0]: 0.000000e+00
	// x[1]: 0.000000e+00
	// x[2]: 1.500000e+01
	// x[3]: 8.333333e+00
}
