package gmsk_test

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fardream/gmsk"
)

// Semidefinite optimization, reproduced from sdo2.c in MOSEK C api.
//
//	min   <C1,X1> + <C2,X2>
//	st.   <A1,X1> + <A2,X2> = b
//	            (X2)_{1,2} <= k
//
//	where X1, X2 are symmetric positive semidefinite,
//
//	C1, C2, A1, A2 are assumed to be constant symmetric matrices,
//	and b, k are constants.
func Example_semidefiniteOptimization_sdo2() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	/* Input data */
	const numbarvar int32 = 2
	dimbarvar := []int32{3, 4} /* Dimension of semidefinite variables */

	/* Objective coefficients concatenated */
	Cj := []int32{0, 0, 1, 1, 1, 1} /* Which symmetric variable (j) */
	Ck := []int32{0, 2, 0, 1, 1, 2} /* Which entry (k,l)->v */
	Cl := []int32{0, 2, 0, 0, 1, 2}
	Cv := []float64{1.0, 6.0, 1.0, -3.0, 2.0, 1.0}

	/* Equality constraints coefficients concatenated */
	Ai := []int32{0, 0, 0, 0, 0, 0} /* Which constraint (i = 0) */
	Aj := []int32{0, 0, 0, 1, 1, 1} /* Which symmetric variable (j) */
	Ak := []int32{0, 2, 2, 1, 1, 3} /* Which entry (k,l)->v */
	Al := []int32{0, 0, 2, 0, 1, 3}
	Av := []float64{1.0, 1.0, 2.0, 1.0, -1.0, -3.0}

	/* The second constraint - one-term inequality */
	var A2i int32 = 1 /* Which constraint (i = 1) */
	var A2j int32 = 1 /* Which symmetric variable (j = 1) */
	var A2k int32 = 1 /* Which entry A(1,0) = A(0,1) = 0.5 */
	var A2l int32 = 0
	var A2v float64 = 0.5

	/* Constraint bounds and values */
	const numcon int32 = 2
	bkc := []gmsk.BoundKey{gmsk.BK_FX, gmsk.BK_UP}
	blc := []float64{23.0, -gmsk.INFINITY}
	buc := []float64{23.0, -3.0}

	var i, j, dim int32
	// var barx *float64

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

	/* Append empty constraints.
	   The constraints will initially have no bounds. */
	checkOk(task.AppendCons(numcon))

	/* Append semidefinite variables. */
	checkOk(task.AppendBarvars(numbarvar, dimbarvar))

	/* Set objective (6 nonzeros).*/
	checkOk(task.PutBarcBlockTriplet(6, Cj, Ck, Cl, Cv))

	/* Set the equality constraint (6 nonzeros).*/
	checkOk(task.PutBaraBlockTriplet(6, Ai, Aj, Ak, Al, Av))

	/* Set the inequality constraint (1 nonzero).*/
	checkOk(task.PutBaraBlockTriplet(1, []int32{A2i}, []int32{A2j}, []int32{A2k}, []int32{A2l}, []float64{A2v}))

	/* Set constraint bounds */
	checkOk(task.PutConBoundSlice(0, 2, bkc, blc, buc))

	/* Run optimizer */
	trmcode, r := task.OptimizeTrm()

	task.SolutionSummary(gmsk.STREAM_LOG)

	checkOk(r)

	solsta, r := task.GetSolSta(gmsk.SOL_ITR)
	checkOk(r)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		/* Retrieve the soution for all symmetric variables */
		fmt.Printf("Solution (lower triangular part vectorized):\n")
		for i = 0; i < numbarvar; i++ {
			var b strings.Builder

			dim = dimbarvar[i] * (dimbarvar[i] + 1) / 2
			barx := make([]float64, dim)
			barx, r = task.GetBarXj(gmsk.SOL_ITR, i, barx)
			checkOk(r)
			fmt.Fprintf(&b, "X%d: ", i+1)
			for j = 0; j < dim; j++ {
				fmt.Fprintf(&b, "%.3f ", barx[j])
			}

			fmt.Println(strings.TrimRight(b.String(), " "))
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
	// Solution (lower triangular part vectorized):
	// X1: 21.047 0.000 4.077 5.534 0.000 0.790
	// X2: 5.054 -3.000 0.000 0.000 1.781 0.000 0.000 0.000 0.000 0.000
}
