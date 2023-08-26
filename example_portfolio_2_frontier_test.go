package gmsk_test

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/fardream/gmsk"
)

// Portfolio frontier optimization, reproduced from portfolio_2_frontier.c in MOSEK C api.
func Example_portfolio_2_frontier() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	const n int32 = 8
	mu := []float64{0.07197349, 0.15518171, 0.17535435, 0.0898094, 0.42895777, 0.39291844, 0.32170722, 0.18378628}
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
	alphas := []float64{0.0, 0.01, 0.1, 0.25, 0.30, 0.35, 0.4, 0.45, 0.5, 0.75, 1.0, 1.5, 2.0, 3.0, 10.0}
	const numalpha int32 = 15
	var totalBudget float64

	// Offset of variables into the API variable.
	const numvar int32 = n + 1
	const voff_x int32 = 0
	const voff_s int32 = n

	// Constraints offsets
	var numcon int32 = 1
	var coff_bud int32 = 0

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

	// Holding variable x of length n
	checkOk(task.AppendVars(numvar))
	for j := int32(0); j < n; j++ {
		/* Optionally we can give the variables names */
		checkOk(task.PutVarName(voff_x+j, fmt.Sprintf("x[%d]", 1+j)))
		/* No short-selling - x^l = 0, x^u = inf */
		checkOk(task.PutVarBound(voff_x+j, gmsk.BK_LO, 0, gmsk.INFINITY))
	}

	checkOk(task.PutVarName(voff_s, "s"))
	checkOk(task.PutVarBound(voff_s, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	// One linear constraint: total budget
	checkOk(task.AppendCons(numcon))
	checkOk(task.PutConName(coff_bud, "budget"))
	for j := int32(0); j < n; j++ {
		/* Coefficients in the first row of A */
		checkOk(task.PutAij(coff_bud, voff_x+j, 1))
	}

	totalBudget = w
	for _, x0i := range x0 {
		totalBudget += x0i
	}
	checkOk(task.PutConBound(coff_bud, gmsk.BK_FX, totalBudget, totalBudget))

	// Input (gamma, GTx) in the AFE (affine expression) storage
	// We build the following F and g for variables [x, s]:
	//     [0, 1]      [0  ]
	// F = [0, 0], g = [0.5]
	//     [GT,0]      [0  ]
	// We need k+2 rows
	checkOk(task.AppendAfes(k + 2))
	// The first affine expression is variable s (last variable, index n)
	checkOk(task.PutAfeFEntry(0, n, 1))
	// The second affine expression is constant 0.5
	checkOk(task.PutAfeG(1, 0.5))
	// The remaining k expressions comprise GT*x, we add them row by row
	// In more realisic scenarios it would be better to extract nonzeros and input in sparse form
	vslice_x := make([]int32, n)
	for i := int32(0); i < n; i++ {
		vslice_x[i] = voff_x + i
	}
	for i := int64(0); i < k; i++ {
		checkOk(task.PutAfeFRow(i+2, n, &vslice_x[0], &GT[i][0]))
	}

	// Input the affine conic constraint (gamma, GT*x) \in QCone
	// Add the quadratic domain of dimension k+1
	rqdom, res := task.AppendRQuadraticConeDomain(k + 2)
	checkOk(res)
	// Add the constraint
	checkOk(task.AppendAccSeq(rqdom, k+2, 0, nil))
	checkOk(task.PutAccName(rqdom, "risk"))

	// Objective: maximize expected return mu^T x
	for j := int32(0); j < n; j++ {
		checkOk(task.PutCJ(voff_x+j, mu[j]))
	}
	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	/* Set the log level */
	checkOk(task.PutIntParam(gmsk.IPAR_LOG, 0)) // #define LOGLEVEL 0

	fmt.Println(strings.TrimRight(fmt.Sprintf("%-12s  %-12s  %-12s", "alpha", "exp ret", "std. dev"), " ")) // 3rd column will be width 12, so trim right to match the below output

	for i := int32(0); i < numalpha; i++ {
		alpha := alphas[i]

		/* Sets the objective function coefficient for s. */
		checkOk(task.PutCJ(voff_s+0, -alpha))

		_, res := task.OptimizeTrm()
		checkOk(res)

		solsta, res := task.GetSolSta(gmsk.SOL_ITR)
		checkOk(res)

		if solsta != gmsk.SOL_STA_OPTIMAL {
			fmt.Printf("An error occurred when solving for alpha=%e\n", alpha)
		}

		var expret, stddev float64

		for j := int32(0); j < n; j++ {
			xx, res := task.GetXxSlice(gmsk.SOL_ITR, voff_x+j, voff_x+j+1, nil)
			checkOk(res)
			xj := xx[0]
			expret += mu[j] * xj
		}

		stddevd, res := task.GetXxSlice(gmsk.SOL_ITR, voff_s, voff_s+1, nil)
		stddev = stddevd[0]
		fmt.Println(strings.TrimRight(fmt.Sprintf("%-12.3e  %-12.3e  %-12.3e", alpha, expret, math.Sqrt(float64(stddev))), " ")) // the last column will be width 12, so we need to trim to match the below output
	}
	// Output:
	// alpha         exp ret       std. dev
	// 0.000e+00     4.290e-01     1.000e+00
	// 1.000e-02     4.290e-01     4.152e-01
	// 1.000e-01     4.290e-01     4.152e-01
	// 2.500e-01     4.216e-01     3.725e-01
	// 3.000e-01     4.175e-01     3.518e-01
	// 3.500e-01     4.146e-01     3.386e-01
	// 4.000e-01     4.124e-01     3.298e-01
	// 4.500e-01     4.107e-01     3.236e-01
	// 5.000e-01     4.093e-01     3.191e-01
	// 7.500e-01     4.052e-01     3.083e-01
	// 1.000e+00     4.032e-01     3.044e-01
	// 1.500e+00     3.932e-01     2.915e-01
	// 2.000e+00     3.847e-01     2.828e-01
	// 3.000e+00     3.567e-01     2.632e-01
	// 1.000e+01     2.379e-01     2.123e-01
}
