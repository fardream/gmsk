package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Portfolio optimization example with 3/2 impact. Reproduced from portfolio_3_impact.c in MOSEK C api.
func Example_portfolio_3_impact() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
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
	const gamma float64 = 0.36
	var totalBudget float64

	m := make([]float64, n)
	for i := int32(0); i < n; i++ {
		m[i] = 0.01
	}

	// Offset of variables into the API variable.
	const numvar int32 = 3 * n
	const voff_x int32 = 0
	const voff_c int32 = n
	const voff_z int32 = 2 * n

	// Offset of constraints.
	// const numcon int32 = 2*n + 1
	const coff_bud int32 = 0
	const coff_abs1 int32 = 1
	const coff_abs2 int32 = 1 + n

	var expret float64

	var res error

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

	checkOk(task.AppendVars(numvar))
	for j := int32(0); j < n; j++ {
		/* Optionally we can give the variables names */
		checkOk(task.PutVarName(voff_x+j, fmt.Sprintf("x[%d]", j+1)))
		checkOk(task.PutVarName(voff_c+j, fmt.Sprintf("c[%d]", j+1)))
		checkOk(task.PutVarName(voff_z+j, fmt.Sprintf("z[%d]", j+1)))
		/* Apply variable bounds (x >= 0, c and z free) */
		checkOk(task.PutVarBound(voff_x+j, gmsk.BK_LO, 0, gmsk.INFINITY))
		checkOk(task.PutVarBound(voff_c+j, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))
		checkOk(task.PutVarBound(voff_z+j, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))
	}

	// Linear constraints
	// - Total budget
	checkOk(task.AppendCons(1))
	checkOk(task.PutConName(coff_bud, "budget"))
	for j := int32(0); j < n; j++ {
		/* Coefficients in the first row of A */
		checkOk(task.PutAij(coff_bud, voff_x+j, 1))
		checkOk(task.PutAij(coff_bud, voff_c+j, m[j]))
	}
	totalBudget = w
	for i := int32(0); i < n; i++ {
		totalBudget += x0[i]
	}
	checkOk(task.PutConBound(coff_bud, gmsk.BK_FX, totalBudget, totalBudget))

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

	// ACCs
	const aoff_q int64 = 0
	const aoff_pow int64 = k + 1
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

	qdom, res := task.AppendQuadraticConeDomain(k + 1)
	checkOk(res)
	checkOk(task.AppendAccSeq(qdom, k+1, aoff_q, nil))
	checkOk(task.PutAccName(aoff_q, "risk"))

	// - (c_j, 1, z_j) in P3(2/3, 1/3)
	// The part of F and g for variables [c, z]:
	//     [0, I, 0]      [0]
	// F = [0, 0, I], g = [0]
	//     [0, 0, 0]      [1]
	checkOk(task.AppendAfes(2*int64(n) + 1))
	for i := int32(0); i < n; i++ {
		checkOk(task.PutAfeFEntry(aoff_pow+int64(i), voff_c+i, 1.0))
		checkOk(task.PutAfeFEntry(aoff_pow+int64(n+i), voff_z+i, 1.0))
	}
	checkOk(task.PutAfeG(aoff_pow+2*(int64(n)), 1.0))
	// We use one row from F and g for both c_j and z_j, and the last row of F and g for the constant 1.
	// NOTE: Here we reuse the last AFE and the power cone n times, but we store them only once.
	exponents := []float64{2, 1}
	powdom, res := task.AppendPrimalPowerConeDomain(3, 2, &exponents[0])
	checkOk(res)
	flat_afe_list := make([]int64, 3*n)
	dom_list := make([]int64, n)
	for i := int64(0); i < int64(n); i++ {
		flat_afe_list[3*i+0] = aoff_pow + i
		flat_afe_list[3*i+1] = aoff_pow + 2*int64(n)
		flat_afe_list[3*i+2] = aoff_pow + int64(n) + i
		dom_list[i] = powdom
	}
	checkOk(task.AppendAccs(int64(n), &dom_list[0], 3*int64(n), &flat_afe_list[0], nil))
	for i := int64(0); i < int64(n); i++ {
		checkOk(task.PutAccName(i+1, fmt.Sprintf("market_impact[%d]", i)))
	}

	// Objective: maximize expected return mu^T x
	for j := int32(0); j < n; j++ {
		checkOk(task.PutCJ(voff_x+j, mu[j]))
	}
	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	/* No log output */
	checkOk(task.PutIntParam(gmsk.IPAR_LOG, 0))

	/* Dump the problem to a human readable PTF file. */
	checkOk(task.WriteDataHandle(os.Stderr, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE))

	_, res = task.OptimizeTrm()

	/* Display the solution summary for quick inspection of results. */
	checkOk(task.SolutionSummary(gmsk.STREAM_LOG))
	checkOk(res)

	for j := int32(0); j < n; j++ {
		xx, res := task.GetXxSlice(gmsk.SOL_ITR, voff_x+j, voff_x+j+1, nil)
		checkOk(res)
		xj := xx[0]
		expret += mu[j] * xj
	}

	fmt.Printf("\nExpected return %e for gamma %e\n", expret, gamma)
	// Output: Expected return 4.165712e-01 for gamma 3.600000e-01
}
