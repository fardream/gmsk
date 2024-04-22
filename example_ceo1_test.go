package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Conic exponential optimization, reproduced from ceo1.c in MOSEK C api.
func Example_conicExponentialOptimization1_ceo1() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	const numvar, numcon int32 = 3, 1
	const numafe, f_nnz int64 = 3, 3

	bkc := gmsk.BK_FX
	blc := 1.0
	buc := 1.0

	bkx := []gmsk.BoundKey{
		gmsk.BK_FR,
		gmsk.BK_FR,
		gmsk.BK_FR,
	}
	blx := []float64{
		-gmsk.INFINITY,
		-gmsk.INFINITY,
		-gmsk.INFINITY,
	}
	bux := []float64{
		gmsk.INFINITY,
		gmsk.INFINITY,
		gmsk.INFINITY,
	}

	c := []float64{
		1,
		1,
		0,
	}

	a := []float64{1, 1, 1}
	asub := []int32{0, 1, 2}

	afeidx := []int64{0, 1, 2}
	varidx := []int32{0, 1, 2}
	f_val := []float64{1, 1, 1}

	var domidx int64 = 0

	/* Create the mosek environment. */
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	/* Create the optimization task. */
	task, err := env.MakeTask(numcon, numvar)
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

	/* Append 'numafe' affine expressions.
	   The affine expressions will initially be empty. */
	checkOk(task.AppendAfes(numafe))

	/* Set up the linear part */
	checkOk(task.PutCSlice(0, numvar, c))
	checkOk(task.PutARow(0, numvar, asub, a))
	checkOk(task.PutConBound(0, bkc, blc, buc))
	checkOk(task.PutVarBoundSlice(0, numvar, bkx, blx, bux))

	checkOk(task.PutAfeFEntryList(f_nnz, afeidx, varidx, f_val))
	domidx, r = task.AppendPrimalExpConeDomain()
	checkOk(r)
	checkOk(task.AppendAccSeq(domidx, numafe, 0, nil))

	/* Run optimizer */
	trmcode, r := task.OptimizeTrm()

	/* Print a summary containing information
	   about the solution for debugging purposes*/
	task.SolutionSummary(gmsk.STREAM_LOG)

	checkOk(r)

	solsta, r := task.GetSolSta(gmsk.SOL_ITR)
	checkOk(r)

	switch solsta {
	case gmsk.SOL_STA_OPTIMAL:
		xx, r := task.GetXx(gmsk.SOL_ITR, nil)
		if r != nil {
			checkOk(gmsk.NewErrorFromInt(gmsk.RES_ERR_SPACE))
		}

		fmt.Printf("Optimal primal solution\n")
		for j := 0; j < 3; j++ {
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
	// Output:
	// Optimal primal solution
	// x[0]: 6.117883e-01
	// x[1]: 1.704000e-01
	// x[2]: 2.178117e-01
}
