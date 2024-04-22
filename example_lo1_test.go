package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Linear programming example 1, reproduced from mosek c api example lo1.c
func Example_linearOptimization1_lo1() {
	const numvar, numcon int32 = 4, 3
	c := []float64{3.0, 1.0, 5.0, 1.0}

	/* Below is the sparse representation of the A
	   matrix stored by column. */
	aptrb := []int32{0, 2, 5, 7}
	aptre := []int32{2, 5, 7, 9}
	asub := []int32{
		0, 1,
		0, 1, 2,
		0, 1,
		1, 2,
	}
	aval := []float64{
		3, 2,
		1, 1, 2,
		2, 3,
		1, 3,
	}
	/* Bounds on constraints. */
	bkc := []gmsk.BoundKey{gmsk.BK_FX, gmsk.BK_LO, gmsk.BK_UP}
	blc := []float64{30, 15, -gmsk.INFINITY}
	buc := []float64{30, gmsk.INFINITY, 25}

	/* Bounds on variables. */
	bkx := []gmsk.BoundKey{gmsk.BK_LO, gmsk.BK_RA, gmsk.BK_LO, gmsk.BK_LO}
	blx := []float64{0, 0, 0, 0}
	bux := []float64{gmsk.INFINITY, 10, gmsk.INFINITY, gmsk.INFINITY}

	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

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

	for j := int32(0); j < numvar && r == nil; j++ {
		/* Set the linear term c_j in the objective.*/
		r = task.PutCJ(j, c[j])
		if r != nil {
			break
		}

		r = task.PutVarBound(
			j,      /* Index of variable.*/
			bkx[j], /* Bound key.*/
			blx[j], /* Numerical value of lower bound.*/
			bux[j]) /* Numerical value of upper bound.*/
		if r != nil {
			break
		}

		/* Input column j of A */
		r = task.PutACol(
			j,                       /* Variable (column) index.*/
			aptre[j]-aptrb[j],       /* Number of non-zeros in column j.*/
			asub[aptrb[j]:aptre[j]], /* Pointer to row indexes of column j.*/
			aval[aptrb[j]:aptre[j]]) /* Pointer to Values of column j.*/
	}

	checkOk(r)

	/* Set the bounds on constraints.
	   for i=1, ...,numcon : blc[i] <= constraint i <= buc[i] */
	for i := int32(0); i < numcon && r == nil; i++ {
		r = task.PutConBound(
			i,      /* Index of constraint.*/
			bkc[i], /* Bound key.*/
			blc[i], /* Numerical value of lower bound.*/
			buc[i]) /* Numerical value of upper bound.*/
	}

	checkOk(r)

	/* Maximize objective function. */
	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	var trmcode gmsk.ResCode
	/* Run optimizer */
	trmcode, r = task.OptimizeTrm()
	checkOk(r)

	/* Print a summary containing information
	   about the solution for debugging purposes. */
	task.SolutionSummary(gmsk.STREAM_LOG)

	var solsta gmsk.SolSta

	solsta, r = task.GetSolSta(gmsk.SOL_BAS)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		xx, r := task.GetXx(
			gmsk.SOL_BAS, /* Request the basic solution. */
			nil)
		if r != nil {
			r = gmsk.NewError(gmsk.RES_ERR_SPACE)
			break
		}
		fmt.Print("Optimal primal solution\n")
		for j := int32(0); j < numvar; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}
	case gmsk.SOL_STA_DUAL_INFEAS_CER:
		fallthrough
	case gmsk.SOL_STA_PRIM_INFEAS_CER:
		fmt.Printf("Primal or dual infeasibility certificate found.\n")
	case gmsk.SOL_STA_UNKNOWN:
		/* If the solutions status is unknown, print the termination code
		   indicating why the optimizer terminated prematurely. */
		_, symname, _ := gmsk.GetCodedesc(trmcode)
		fmt.Printf("The solution status is unknown.\n")
		fmt.Printf("The optimizer terminitated with code: %s\n", symname)
	default:
		fmt.Printf("Other solution status.\n")
	}

	if r != nil && gmsk.IsMskError(r) {
		e, _ := gmsk.AsMskError(r)
		/* In case of an error print error code and description. */
		_, symname, desc := gmsk.GetCodedesc(e.ToResCode())
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
