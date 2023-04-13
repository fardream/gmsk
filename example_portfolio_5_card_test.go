package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Portfolio optimization with cardinality constraints on the number of assets traded,
// reproduced from portfolio_5_card.c in MOSEK C api.
func Example_portfolio_5_Card() {
	checkOk := func(r gmsk.ResCode) {
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodedesc(r)

			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	const n int32 = 8
	mu := []float64{0.07197, 0.15518, 0.17535, 0.08981, 0.42896, 0.39292, 0.32171, 0.18379}
	// GT must have size n rows
	GT := [...][8]float64{
		{0.30758, 0.12146, 0.11341, 0.11327, 0.17625, 0.11973, 0.10435, 0.10638},
		{0.00000, 0.25042, 0.09946, 0.09164, 0.06692, 0.08706, 0.09173, 0.08506},
		{0.00000, 0.00000, 0.19914, 0.05867, 0.06453, 0.07367, 0.06468, 0.01914},
		{0.00000, 0.00000, 0.00000, 0.20876, 0.04933, 0.03651, 0.09381, 0.07742},
		{0.00000, 0.00000, 0.00000, 0.00000, 0.36096, 0.12574, 0.10157, 0.0571},
		{0.00000, 0.00000, 0.00000, 0.00000, 0.00000, 0.21552, 0.05663, 0.06187},
		{0.00000, 0.00000, 0.00000, 0.00000, 0.00000, 0.00000, 0.22514, 0.03327},
		{0.00000, 0.00000, 0.00000, 0.00000, 0.00000, 0.00000, 0.00000, 0.2202},
	}

	const k int64 = 8 // this is const MSKint32t k       = sizeof(GT) / (n * sizeof(MSKrealt));
	x0 := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	const w float64 = 1
	const gamma float64 = 0.25

	xx := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	markowitz_with_card := func(K int32) gmsk.ResCode {
		// Offset of variables
		const numvar int32 = 3 * n
		const voff_x int32 = 0
		const voff_z int32 = n
		const voff_y int32 = 2 * n

		// Offset of constraints.
		const numcon int32 = 3*n + 2
		_ = numcon // this is not used
		const coff_bud int32 = 0
		const coff_abs1 int32 = 1
		const coff_abs2 int32 = 1 + n
		const coff_swi int32 = 1 + 2*n
		const coff_card int32 = 1 + 3*n

		/* Initial setup. */
		env, err := gmsk.MakeEnv()
		if err != nil {
			log.Fatal(err)
		}
		defer gmsk.DeleteEnv(env)

		task, err := gmsk.MakeTask(env, 0, 0)
		if err != nil {
			log.Fatal(err)
		}
		defer gmsk.DeleteTask(task)

		checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

		// Variables (vector of x, z, y)
		checkOk(task.AppendVars(numvar))
		for j := int32(0); j < n; j++ {
			/* Optionally we can give the variables names */
			checkOk(task.PutVarName(voff_x+j, fmt.Sprintf("x[%d]", 1+j)))
			checkOk(task.PutVarName(voff_z+j, fmt.Sprintf("z[%d]", 1+j)))
			checkOk(task.PutVarName(voff_y+j, fmt.Sprintf("y[%d]", 1+j)))
			/* Apply variable bounds (x >= 0, z free, y binary) */
			checkOk(task.PutVarBound(voff_x+j, gmsk.BK_LO, 0, gmsk.INFINITY))
			checkOk(task.PutVarBound(voff_z+j, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))
			checkOk(task.PutVarBound(voff_y+j, gmsk.BK_RA, 0, 1))
			checkOk(task.PutVarType(voff_y+j, gmsk.VAR_TYPE_INT))
		}

		// Linear constraints
		// - Total budget
		checkOk(task.AppendCons(1))
		checkOk(task.PutConName(coff_bud, "budget"))
		for j := int32(0); j < n; j++ {
			/* Coefficients in the first row of A */
			checkOk(task.PutAij(coff_bud, voff_x+j, 1))
		}
		U := w
		for i := int32(0); i < n; i++ {
			U += x0[i]
		}
		checkOk(task.PutConBound(coff_bud, gmsk.BK_FX, U, U))

		// - Absolute value
		checkOk(task.AppendCons(2 * n))
		for i := int32(0); i < n; i++ {
			checkOk(task.PutConName(coff_abs1+i, fmt.Sprintf("zabs1[%d]", 1+i)))
			checkOk(task.PutAij(coff_abs1+i, voff_x+i, -1))
			checkOk(task.PutAij(coff_abs1+i, voff_z+i, 1))
			checkOk(task.PutConBound(coff_abs1+i, gmsk.BK_LO, -x0[i], gmsk.INFINITY))
			checkOk(task.PutConName(coff_abs2+i, fmt.Sprintf("zabs2[%d]", 1+i)))
			checkOk(task.PutAij(coff_abs2+i, voff_x+i, 1))
			checkOk(task.PutAij(coff_abs2+i, voff_z+i, 1))
			checkOk(task.PutConBound(coff_abs2+i, gmsk.BK_LO, x0[i], gmsk.INFINITY))
		}

		// - Switch
		checkOk(task.AppendCons(n))
		for i := int32(0); i < n; i++ {
			checkOk(task.PutConName(coff_swi+i, fmt.Sprintf("switch[%d]", i+1)))
			checkOk(task.PutAij(coff_swi+i, voff_z+i, 1))
			checkOk(task.PutAij(coff_swi+i, voff_y+i, -U))
			checkOk(task.PutConBound(coff_swi+i, gmsk.BK_UP, -gmsk.INFINITY, 0))
		}

		// - Cardinality
		checkOk(task.AppendCons(1))
		checkOk(task.PutConName(coff_card, "cardinality"))
		for i := int32(0); i < n; i++ {
			checkOk(task.PutAij(coff_card, voff_y+i, 1))
		}
		checkOk(task.PutConBound(coff_card, gmsk.BK_UP, -gmsk.INFINITY, float64(K)))

		// ACCs
		const aoff_q int64 = 0
		// - (gamma, GTx) in Q(k+1)
		// The part of F and g for variable x:
		//     [0,  0, 0]      [gamma]
		// F = [GT, 0, 0], g = [0    ]
		checkOk(task.AppendAfes(k + 1))
		checkOk(task.PutAfeG(aoff_q, gamma))
		vslice_x := make([]int32, n)
		for i := int32(0); i < n; i++ {
			vslice_x[i] = voff_x + i
		}
		for i := int64(0); i < k; i++ {
			checkOk(task.PutAfeFRow(aoff_q+i+1, n, &vslice_x[0], &GT[i][0]))
		}

		res, qdom := task.AppendQuadraticConeDomain(k + 1)
		checkOk(res)
		checkOk(task.AppendAccSeq(qdom, k+1, aoff_q, nil))
		checkOk(task.PutAccName(aoff_q, "risk"))

		// Objective: maximize expected return mu^T x
		for j := int32(0); j < n; j++ {
			checkOk(task.PutCJ(voff_x+j, mu[j]))
		}
		checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

		/* No log output */
		checkOk(task.PutIntParam(gmsk.IPAR_LOG, 0))

		/* Dump the problem to a human readable PTF file. */
		checkOk(task.WriteDataHandle(os.Stderr, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE))

		res, _ = task.OptimizeTrm()

		/* Display the solution summary for quick inspection of results. */
		checkOk(task.SolutionSummary(gmsk.STREAM_LOG))
		checkOk(res)

		res, _ = task.GetXxSlice(gmsk.SOL_ITG, voff_x, voff_x+n, xx)

		return res
	}

	for K := int32(1); K <= n; K++ {
		checkOk(markowitz_with_card(K))
		var expret float64 = 0
		fmt.Printf("Bound %d:  x = ", K)
		for i := int32(0); i < n; i++ {
			fmt.Printf("%.5f ", xx[i])
			expret += xx[i] * mu[i]
		}
		fmt.Printf("  Return:  %.5f\n", expret)
	}
	// Output:
	// Bound 1:  x = 0.00000 0.00000 1.00000 0.00000 0.00000 0.00000 0.00000 0.00000   Return:  0.17535
	// Bound 2:  x = 0.00000 0.00000 0.35691 0.00000 0.00000 0.64309 0.00000 0.00000   Return:  0.31527
	// Bound 3:  x = 0.00000 0.00000 0.19261 0.00000 0.00000 0.54597 0.26142 0.00000   Return:  0.33240
	// Bound 4:  x = 0.00000 0.00000 0.20391 0.00000 0.06710 0.49181 0.23718 0.00000   Return:  0.33408
	// Bound 5:  x = 0.00000 0.03197 0.17028 0.00000 0.07074 0.49551 0.23150 0.00000   Return:  0.33434
	// Bound 6:  x = 0.00000 0.03196 0.17028 0.00000 0.07073 0.49551 0.23152 0.00000   Return:  0.33434
	// Bound 7:  x = 0.00000 0.02699 0.16706 0.00000 0.07124 0.49559 0.22943 0.00969   Return:  0.33436
	// Bound 8:  x = 0.00000 0.02699 0.16706 0.00000 0.07125 0.49559 0.22943 0.00969   Return:  0.33436
}
