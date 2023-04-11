package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Optimization with disjunctive constraints, reproduced from djc1.c in MOSEK C api.
//
//	Purpose: Demonstrates how to solve the problem with two disjunctions:
//
//	   minimize    2x0 + x1 + 3x2 + x3
//	   subject to   x0 + x1 + x2 + x3 >= -10
//	               (x0-2x1<=-1 and x2=x3=0) or (x2-3x3<=-2 and x1=x2=0)
//	               x0=2.5 or x1=2.5 or x2=2.5 or x3=2.5
func Example_disjunctiveConstraint_djc1() {
	checkOk := func(r gmsk.ResCode) {
		if !r.IsOk() {
			_, sym, desc := gmsk.GetCodeDesc(r)
			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	r := gmsk.RES_OK

	var j, numvar int32
	var numafe, numdjc int64
	var zero1, zero2, rminus1 int64 // Domain indices

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

	// Append free variables
	numvar = 4
	checkOk(task.AppendVars(numvar))
	checkOk(task.PutVarboundSliceConst(0, numvar, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	{
		// The linear part: the linear constraint
		idx := []int32{0, 1, 2, 3}
		val := []float64{1, 1, 1, 1}

		checkOk(task.AppendCons(1))
		checkOk(task.PutARow(0, 4, &idx[0], &val[0]))
		checkOk(task.PutConbound(0, gmsk.BK_LO, -10, -10))
	}
	{
		// The linear part: objective
		idx := []int32{0, 1, 2, 3}
		val := []float64{2, 1, 3, 1}
		checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MINIMIZE))
		checkOk(task.PutCList(4, &idx[0], &val[0]))
	}

	// Fill in the affine expression storage F, g
	numafe = 10
	checkOk(task.AppendAfes(numafe))

	fafeidx := []int64{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fvaridx := []int32{0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3}
	fval := []float64{1.0, -2.0, 1.0, -3.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
	g := []float64{1.0, 2.0, 0.0, 0.0, 0.0, 0.0, -2.5, -2.5, -2.5, -2.5}

	checkOk(task.PutAfeFEntryList(12, &fafeidx[0], &fvaridx[0], &fval[0]))
	checkOk(task.PutAfeGSlice(0, numafe, &g[0]))

	// Create domains
	r, zero1 = task.AppendRZeroDomain(1)
	checkOk(r)
	r, zero2 = task.AppendRZeroDomain(2)
	checkOk(r)
	r, rminus1 = task.AppendRMinusDomain(1)
	checkOk(r)

	// Append disjunctive constraints
	numdjc = 2
	checkOk(task.AppendDjcs(numdjc))

	{
		// First disjunctive constraint
		domidxlist := []int64{rminus1, zero2, rminus1, zero2}
		afeidxlist := []int64{0, 4, 5, 1, 2, 3}
		termsizelist := []int64{2, 2}

		checkOk(task.PutDjc(
			0, // DJC index
			4, &domidxlist[0],
			6, &afeidxlist[0],
			nil, // Unused
			2, &termsizelist[0]))
	}

	{
		// Second disjunctive constraint
		domidxlist := []int64{zero1, zero1, zero1, zero1}
		afeidxlist := []int64{6, 7, 8, 9}
		termsizelist := []int64{1, 1, 1, 1}
		checkOk(task.PutDjc(
			1, // DJC index
			4, &domidxlist[0],
			4, &afeidxlist[0],
			nil, // Unused
			4, &termsizelist[0]))
	}

	// Useful for debugging
	{
		// Write a human-readable file
		checkOk(task.WriteDataHandle(os.Stderr, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE))
		// Directs the log task stream to the 'printstr' function.
		checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))
	}

	// Solve the problem
	r, _ = task.OptimizeTrm()
	checkOk(r)

	/* Print a summary containing information
	   about the solution for debugging purposes. */
	task.SolutionSummary(gmsk.STREAM_LOG)

	r, solsta := task.GetSolSta(gmsk.SOL_ITG)
	checkOk(r)

	switch solsta {
	case gmsk.SOL_STA_INTEGER_OPTIMAL:
		xx := make([]float64, numvar)
		r, xx = task.GetXx(gmsk.SOL_ITG, xx)
		fmt.Printf("Optimal primal solution\n")
		for j = 0; j < numvar; j++ {
			fmt.Printf("x[%d]: %e\n", j, xx[j])
		}

	default:
		fmt.Printf("Another solution status.\n")
	}
	// Output:
	// Optimal primal solution
	// x[0]: 0.000000e+00
	// x[1]: 0.000000e+00
	// x[2]: -1.250000e+01
	// x[3]: 2.500000e+00
}
