package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Semidefinite optimization example, reproduced from sdo1.c in MOSEK C api.
//
//	minimize    Tr [2, 1, 0; 1, 2, 1; 0, 1, 2]*X + x0
//
//	subject to  Tr [1, 0, 0; 0, 1, 0; 0, 0, 1]*X + x0           = 1
//	            Tr [1, 1, 1; 1, 1, 1; 1, 1, 1]*X      + x1 + x2 = 0.5
//	            (x0,x1,x2) \in Q,  X \in PSD
func Example_semidefiniteOptimization_sdo1() {
	const NUMCON = 2 /* Number of constraints.              */
	const NUMVAR = 3 /* Number of conic quadratic variables */
	// const NUMANZ = 3    /* Number of non-zeros in A            */
	const NUMAFE = 3    /* Number of affine expressions        */
	const NUMFNZ = 3    /* Number of non-zeros in F            */
	const NUMBARVAR = 1 /* Number of semidefinite variables    */

	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	DIMBARVAR := []int32{3}               /* Dimension of semidefinite cone */
	LENBARVAR := []int32{3 * (3 + 1) / 2} /* Number of scalar SD variables  */

	bkc := []gmsk.BoundKey{gmsk.BK_FX, gmsk.BK_FX}
	blc := []float64{1.0, 0.5}
	buc := []float64{1.0, 0.5}

	barc_i := []int32{0, 1, 1, 2, 2}
	barc_j := []int32{0, 0, 1, 1, 2}
	barc_v := []float64{2.0, 1.0, 2.0, 1.0, 2.0}

	aptrb := []int32{0, 1}
	aptre := []int32{1, 3}
	asub := []int32{0, 1, 2} /* column subscripts of A */
	aval := []float64{1, 1, 1}

	bara_i := []int32{0, 1, 2, 0, 1, 2, 1, 2, 2}
	bara_j := []int32{0, 1, 2, 0, 0, 0, 1, 1, 2}
	bara_v := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
	// conesub := []int32{0, 1, 2}

	afeidx := []int64{0, 1, 2}
	varidx := []int32{0, 1, 2}
	f_val := []float64{1, 1, 1}

	var i, j int32
	var idx int64
	var falpha float64 = 1

	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := env.MakeTask(NUMCON, 0)
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

	/* Append 'NUMAFE' affine expressions.*/
	checkOk(task.AppendAfes(NUMAFE))

	/* Append 'NUMBARVAR' semidefinite variables. */
	checkOk(task.AppendBarvars(NUMBARVAR, DIMBARVAR))

	/* Optionally add a constant term to the objective. */
	checkOk(task.PutCfix(0))

	/* Set the linear term c_j in the objective.*/
	checkOk(task.PutCJ(0, 1))

	for j = 0; j < NUMVAR && r == nil; j++ {
		r = task.PutVarBound(j, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY)
	}
	checkOk(r)

	/* Set the linear term barc_j in the objective.*/
	idx, r = task.AppendSparseSymMat(DIMBARVAR[0], 5, barc_i, barc_j, barc_v)
	checkOk(r)
	checkOk(task.PutBarcJ(0, 1, []int64{idx}, []float64{falpha}))

	/* Set the bounds on constraints.
	   for i=1, ...,NUMCON : blc[i] <= constraint i <= buc[i] */
	for i = 0; i < NUMCON && r == nil; i++ {
		r = task.PutConBound(
			i,      /* Index of constraint.*/
			bkc[i], /* Bound key.*/
			blc[i], /* Numerical value of lower bound.*/
			buc[i]) /* Numerical value of upper bound.*/
	}
	checkOk(r)

	/* Input A row by row */
	for i = 0; i < NUMCON && r == nil; i++ {
		ni := aptre[i] - aptrb[i] // need to check zero since go checks
		if ni <= 0 {
			continue
		}
		r = task.PutARow(i, ni, asub[aptrb[i]:aptre[i]], aval[aptrb[i]:aptre[i]])
	}

	/* Append the affine conic constraint with quadratic cone */
	checkOk(task.PutAfeFEntryList(NUMFNZ, afeidx, varidx, f_val))
	qdomidx, r := task.AppendQuadraticConeDomain(3)
	checkOk(r)
	checkOk(task.AppendAcc(qdomidx, 3, afeidx, nil))

	/* Add the first row of barA */
	idx, r = task.AppendSparseSymMat(DIMBARVAR[0], 3, bara_i, bara_j, bara_v)
	checkOk(r)
	checkOk(task.PutBaraIj(0, 0, 1, []int64{idx}, []float64{falpha}))

	/* Add the second row of barA */
	idx, r = task.AppendSparseSymMat(DIMBARVAR[0], 6, bara_i[3:], bara_j[3:], bara_v[3:])
	checkOk(r)
	checkOk(task.PutBaraIj(1, 0, 1, []int64{idx}, []float64{falpha}))

	trmcode, r := task.OptimizeTrm()

	task.SolutionSummary(gmsk.STREAM_LOG)

	checkOk(r)

	solsta, r := task.GetSolSta(gmsk.SOL_ITR)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		xx := make([]float64, NUMVAR)
		barx := make([]float64, LENBARVAR[0])

		xx, r = task.GetXx(gmsk.SOL_ITR, xx)
		checkOk(r)
		barx, r = task.GetBarXj(gmsk.SOL_ITR, 0, barx)
		checkOk(r)

		fmt.Printf("Optimal primal solution\n")
		for i = 0; i < NUMVAR; i++ {
			fmt.Printf("x[%d]   : % e\n", i, xx[i])
		}
		for i = 0; i < LENBARVAR[0]; i++ {
			fmt.Printf("barx[%d]: % e\n", i, barx[i])
		}

	case gmsk.SOL_STA_DUAL_INFEAS_CER:
		fallthrough
	case gmsk.SOL_STA_PRIM_INFEAS_CER:
		fmt.Printf("Primal or dual infeasibility certificate found.\n")

	case gmsk.SOL_STA_UNKNOWN:
		fmt.Printf("The status of the solution could not be determined. Termination code: %d.\n", trmcode)

	default:
		fmt.Printf("Other solution status.")
	}
	// Output:
	// Optimal primal solution
	// x[0]   :  2.544049e-01
	// x[1]   :  1.798914e-01
	// x[2]   :  1.798914e-01
	// barx[0]:  2.172534e-01
	// barx[1]: -2.599712e-01
	// barx[2]:  2.172534e-01
	// barx[3]:  3.110884e-01
	// barx[4]: -2.599712e-01
	// barx[5]:  2.172534e-01
}
