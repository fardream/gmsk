package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Conic Quadratic Optimization, reproduced from cqo1.c in MOSEK example.
func Example_conicQuadraticOptimization1_cqo1() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	const (
		numvar = int32(6)
		numcon = int32(1)
		numafe = int64(6)
		numacc = int64(2)
		f_nnz  = int64(6)
	)

	bkc := []gmsk.BoundKey{gmsk.BK_FX}
	blc := []float64{1}
	buc := []float64{1}

	bkx := []gmsk.BoundKey{
		gmsk.BK_LO,
		gmsk.BK_LO,
		gmsk.BK_LO,
		gmsk.BK_FR,
		gmsk.BK_FR,
		gmsk.BK_FR,
	}
	blx := []float64{
		0.0,
		0.0,
		0.0,
		-gmsk.INFINITY,
		-gmsk.INFINITY,
		-gmsk.INFINITY,
	}
	bux := []float64{
		+gmsk.INFINITY,
		+gmsk.INFINITY,
		+gmsk.INFINITY,
		+gmsk.INFINITY,
		+gmsk.INFINITY,
		+gmsk.INFINITY,
	}

	c := []float64{
		0.0,
		0.0,
		0.0,
		1.0,
		1.0,
		1.0,
	}
	var (
		aptrb = []int32{0, 1, 2, 3, 3, 3}
		aptre = []int32{1, 2, 3, 3, 3, 3}
		asub  = []int32{0, 0, 0, 0}
		aval  = []float64{1, 1, 2}
	)
	var (
		afeidx = []int64{0, 1, 2, 3, 4, 5}
		varidx = []int32{3, 0, 1, 4, 5, 2}
		f_val  = []float64{1, 1, 1, 1, 1, 1}
	)

	domidx := []int64{0, 0}

	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := gmsk.MakeTask(env, numcon, numvar)
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	/* Append 'numcon' empty constraints.
	   The constraints will initially have no bounds. */
	checkOk(task.AppendCons(numcon))

	/* Append 'numvar' variables.
	   The variables will initially be fixed at zero (x=0). */
	checkOk(task.AppendVars(numvar))

	/* Append 'numafe' affine expressions.
	   The affine expressions will initially be empty. */
	checkOk(task.AppendAfes(numafe))

	for j := int32(0); j < numvar && r == nil; j++ {
		/* Set the linear term c_j in the objective.*/
		r = task.PutCJ(j, c[j])
		checkOk(r)

		/* Set the bounds on variable j.
		   blx[j] <= x_j <= bux[j] */
		r = task.PutVarBound(
			j,      /* Index of variable.*/
			bkx[j], /* Bound key.*/
			blx[j], /* Numerical value of lower bound.*/
			bux[j]) /* Numerical value of upper bound.*/
		checkOk(r)

		if aptre[j] > aptrb[j] { // looks like go will check if the index is out of range.
			/* Input column j of A */
			r = task.PutACol(
				j,                       /* Variable (column) index.*/
				aptre[j]-aptrb[j],       /* Number of non-zeros in column j.*/
				asub[aptrb[j]:aptre[j]], /* Pointer to row indexes of column j.*/
				aval[aptrb[j]:aptre[j]]) /* Pointer to Values of column j.*/
		}
	}

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

	/* Set the non-zero entries of the F matrix */
	checkOk(task.PutAfeFEntryList(f_nnz, afeidx, varidx, f_val))

	/* Append quadratic cone domain */
	domidx[0], r = task.AppendQuadraticConeDomain(3)
	checkOk(r)
	/* Append rotated quadratic cone domain */
	domidx[1], r = task.AppendRQuadraticConeDomain(3)
	checkOk(r)
	/* Append two ACCs made up of the AFEs and the domains defined above. */
	checkOk(task.AppendAccsSeq(numacc, domidx, numafe, afeidx[0], nil))

	/* Run optimizer */
	trmcode, r := task.OptimizeTrm()

	task.SolutionSummary(gmsk.STREAM_LOG)

	checkOk(r)

	solsta, r := task.GetSolSta(gmsk.SOL_ITR)
	checkOk(r)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		xx, r := task.GetXx(
			gmsk.SOL_ITR, /* Request the interior solution. */
			nil)
		if r != nil {
			r = gmsk.NewErrorFromInt(gmsk.RES_ERR_SPACE)
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
		fmt.Printf("The status of the solution could not be determined. Termination code: %d.\n", trmcode)
	default:
		fmt.Printf("Other solution status.\n")
	}

	// Output: Optimal primal solution
	// x[0]: 2.609204e-01
	// x[1]: 2.609204e-01
	// x[2]: 2.390796e-01
	// x[3]: 3.689972e-01
	// x[4]: 1.690548e-01
	// x[5]: 1.690548e-01
}
