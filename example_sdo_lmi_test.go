package gmsk_test

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/fardream/gmsk"
)

// Example of semidefinite matrix with linear matrix inequality, reproduced from sdo_lmi.c in
// MOSEK C api.
//
//	Purpose :   To solve a problem with an LMI and an affine conic constrained problem with a PSD term
//
//	             minimize    Tr [1, 0; 0, 1]*X + x(1) + x(2) + 1
//
//	             subject to  Tr [0, 1; 1, 0]*X - x(1) - x(2) >= 0
//	                         x(1) [0, 1; 1, 3] + x(2) [3, 1; 1, 0] - [1, 0; 0, 1] >> 0
//	                         X >> 0
func Example_semidefiniteOptimization_sdo_lmi() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	const NUMVAR = 2    /* Number of scalar variables */
	const NUMAFE = 4    /* Number of affine expressions        */
	const NUMFNZ = 6    /* Number of non-zeros in F            */
	const NUMBARVAR = 1 /* Number of semidefinite variables    */

	DIMBARVAR := []int32{2}               /* Dimension of semidefinite cone */
	LENBARVAR := []int64{2 * (2 + 1) / 2} /* Number of scalar SD variables  */

	barc_j := []int32{0, 0}
	barc_k := []int32{0, 1}
	barc_l := []int32{0, 1}
	barc_v := []float64{1, 1}

	barf_i := []int64{0, 0}
	barf_j := []int32{0, 0}
	barf_k := []int32{0, 1}
	barf_l := []int32{0, 0}

	barf_v := []float64{0, 1}

	afeidx := []int64{0, 0, 1, 2, 2, 3}
	varidx := []int32{0, 1, 1, 0, 1, 0}
	f_val := []float64{-1, -1, 3, math.Sqrt(2), math.Sqrt(2), 3}
	g := []float64{0, -1, 0, -1}

	var i, j int32

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

	/* Append 'NUMAFE' empty affine expressions. */
	checkOk(task.AppendAfes(NUMAFE))

	/* Append 'NUMVAR' scalar variables.
	   The variables will initially be fixed at zero (x=0). */
	checkOk(task.AppendVars(NUMVAR))

	/* Append 'NUMBARVAR' semidefinite variables. */
	checkOk(task.AppendBarvars(NUMBARVAR, &DIMBARVAR[0]))

	/* Set the constant term in the objective. */
	checkOk(task.PutCfix(1))

	/* Set c_j and the bounds for each scalar variable*/
	for j = 0; j < NUMVAR && r == nil; j++ {
		r = task.PutCJ(j, 1.0)
		checkOk(r)
		r = task.PutVarBound(j, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY)
	}
	checkOk(r)

	/* Set the linear term barc_j in the objective.*/
	checkOk(task.PutBarcBlockTriplet(2, &barc_j[0], &barc_k[0], &barc_l[0], &barc_v[0]))

	/* Set the F matrix */
	checkOk(task.PutAfeFEntryList(NUMFNZ, &afeidx[0], &varidx[0], &f_val[0]))
	/* Set the g vector */
	checkOk(task.PutAfeGSlice(0, NUMAFE, &g[0]))
	/* Set the barF matrix */
	checkOk(task.PutAfeBarfBlockTriplet(2, &barf_i[0], &barf_j[0], &barf_k[0], &barf_l[0], &barf_v[0]))

	/* Append R+ domain and the corresponding ACC */
	acc1_afeidx := []int64{0}
	rplusdom, r := task.AppendRplusDomain(1)
	checkOk(r)
	checkOk(task.AppendAcc(rplusdom, 1, &acc1_afeidx[0], nil))

	/* Append the SVEC_PSD domain and the corresponding ACC */
	acc2_afeidx := []int64{1, 2, 3}
	svecpsddom, r := task.AppendSvecPsdConeDomain(3)
	checkOk(r)
	checkOk(task.AppendAcc(svecpsddom, 3, &acc2_afeidx[0], nil))

	/* Run optimizer */
	trmcode, r := task.OptimizeTrm()

	/* Print a summary containing information
	   about the solution for debugging purposes*/
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
		for i = 0; i < int32(LENBARVAR[0]); i++ {
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
	// x[0]   :  1.000000e+00
	// x[1]   :  1.000000e+00
	// barx[0]:  1.000000e+00
	// barx[1]:  1.000000e+00
	// barx[2]:  1.000000e+00
}
