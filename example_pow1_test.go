package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Power cone example, reproduced from pow1.c in MOSEK C api.

// Purpose: Demonstrates how to solve the problem
//
//	maximize x^0.2*y^0.8 + z^0.4 - x
//	      st x + y + 0.5z = 2
//	         x,y,z >= 0
func Example_powerCone_pow1() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	const numvar, numcon int32 = 5, 1

	const numafe, numacc, f_nnz int64 = 6, 2, 5

	bkx := [5]gmsk.BoundKey{}
	blx, bux := [5]float64{}, [5]float64{}

	val := [3]float64{1, 1, -1}
	sub := [3]int32{3, 4, 0}

	aval := [3]float64{1, 1, 0.5}
	asub := [3]int32{0, 1, 2}

	afeidx := []int64{0, 1, 2, 3, 5}
	varidx := []int32{0, 1, 3, 2, 4}
	f_val := []float64{1, 1, 1, 1, 1}
	var g float64 = 1.0

	alpha_1 := []float64{0.2, 0.8}
	alpha_2 := []float64{0.4, 0.6}

	domidx := []int64{0, 0}

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

	/* Append 'numafe' affine expressions.
	   The affine expressions will initially be empty. */
	checkOk(task.AppendAfes(numafe))

	/* Set up the linear part */
	checkOk(task.PutCList(3, sub[:], val[:]))
	checkOk(task.PutARow(0, 3, asub[:], aval[:]))
	checkOk(task.PutConBound(0, gmsk.BK_FX, 2, 2))
	for i := 0; i < 5; i++ {
		bkx[i], blx[i], bux[i] = gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY
	}
	checkOk(task.PutVarBoundSlice(0, numvar, bkx[:], blx[:], bux[:]))

	/* Set the non-zero entries of the F matrix */
	checkOk(task.PutAfeFEntryList(f_nnz, afeidx, varidx, f_val))
	/* Set the non-zero element of the g vector */
	checkOk(task.PutAfeG(4, g))

	/* Append the primal power cone domains with their respective parameter values. */
	domidx[0], r = task.AppendPrimalPowerConeDomain(3, 2, alpha_1)
	checkOk(r)
	domidx[1], r = task.AppendPrimalPowerConeDomain(3, 2, alpha_2)
	checkOk(r)

	/* Append two ACCs made up of the AFEs and the domains defined above. */
	checkOk(task.AppendAccsSeq(numacc, domidx, numafe, afeidx[0], nil))

	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

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
			checkOk(gmsk.NewError(gmsk.RES_ERR_SPACE))
		}

		fmt.Printf("Optimal primal solution\n")
		for j := 0; j < 3; j++ {
			fmt.Printf("x[%d]: %.3e\n", j, xx[j])
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
	// x[0]: 6.394e-02
	// x[1]: 7.833e-01
	// x[2]: 2.306e+00
}
