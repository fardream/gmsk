package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Portfolio optimization example, reproduced from portfolio_1_basic.c in MOSEK C api example.
func Example_portfolio_1_basic() {
	checkOk := func(err error) {
		if err != nil {
			log.Fatalf("failed: %s", err.Error())
		}
	}

	var r error

	const (
		n     int32   = 8
		gamma float64 = 36
	)
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
	x0 := []float64{8.0, 5.0, 3.0, 5.0, 2.0, 9.0, 3.0, 6.0}
	const w float64 = 59

	// Offset of variables into the API variable.
	var numvar int32 = n
	var voff_x int32 = 0

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
	// No other auxiliary variables are needed in this formulation
	checkOk(task.AppendVars(numvar))
	// Setting up variable x
	for j := int32(0); j < n; j++ {
		/* Optionally we can give the variables names */
		checkOk(task.PutVarName(voff_x+j, fmt.Sprintf("x[%d]", 1+j)))
		/* No short-selling - x^l = 0, x^u = inf */
		checkOk(task.PutVarBound(voff_x+j, gmsk.BK_LO, 0, gmsk.INFINITY))
	}
	// One linear constraint: total budget
	checkOk(task.AppendCons(numcon))
	checkOk(task.PutConName(coff_bud, "budget"))
	for j := int32(0); j < n; j++ {
		/* Coefficients in the first row of A */
		checkOk(task.PutAij(coff_bud, voff_x+j, 1))
	}
	totalBudget := w
	for _, ax0 := range x0 {
		totalBudget += ax0
	}
	checkOk(task.PutConBound(coff_bud, gmsk.BK_FX, totalBudget, totalBudget))

	// Input (gamma, GTx) in the AFE (affine expression) storage
	// We need k+1 rows
	checkOk(task.AppendAfes(k + 1))
	// The first affine expression = gamma
	checkOk(task.PutAfeG(0, gamma))
	// The remaining k expressions comprise GT*x, we add them row by row
	// In more realisic scenarios it would be better to extract nonzeros and input in sparse form
	vslice_x := make([]int32, n)
	for i := int32(0); i < n; i++ {
		vslice_x[i] = voff_x + i
	}
	for i := int64(0); i < k; i++ {
		checkOk(task.PutAfeFRow(i+1, n, vslice_x, GT[i][:]))
	}

	// Input the affine conic constraint (gamma, GT*x) \in QCone
	// Add the quadratic domain of dimension k+1
	qdom, r := task.AppendQuadraticConeDomain(k + 1)
	checkOk(r)
	// Add the constraint
	checkOk(task.AppendAccSeq(qdom, k+1, 0, nil))
	checkOk(task.PutAccName(0, "risk"))

	// Objective: maximize expected return mu^T x
	for j := int32(0); j < n; j++ {
		checkOk(task.PutCJ(voff_x+j, mu[j]))
	}
	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	/* No log output */
	checkOk(task.PutIntParam(gmsk.IPAR_LOG, 0))

	/* Dump the problem to a human readable PTF file. */
	checkOk(task.WriteDataHandle(os.Stderr, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE)) // dump to stderr instead.

	_, r = task.OptimizeTrm()
	checkOk(r)

	var expret float64

	for j := int32(0); j < n; j++ {
		xx, r := task.GetXxSlice(gmsk.SOL_ITR, voff_x+j, voff_x+j+1, nil)
		checkOk(r)
		xj := xx[0]
		expret += mu[j] * xj
	}

	/* Read the value of s. This should be gamma. */
	fmt.Printf("\nExpected return %e for gamma %e\n", expret, gamma)
	// Output: Expected return 4.192247e+01 for gamma 3.600000e+01
}
