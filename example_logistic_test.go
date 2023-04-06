package gmsk_test

import (
	"fmt"
	"log"
	"math"

	"github.com/fardream/gmsk"
)

func MSKCALL(r gmsk.ResCode) {
	if r.IsOk() {
		return
	}

	b, sym, desc := gmsk.GetCodeDescSimple(r)
	if b.NotOk() {
		log.Panicf("cannot get the description of error %d - the error code is %d for getting the description", r, b)
	} else {
		log.Panicf("failed: %d %s %s", r, sym, desc)
	}
}

// Adds ACCs for t_i >= log ( 1 + exp((1-2*y[i]) * theta' * X[i]) )
// Adds auxiliary variables, AFE rows and constraints
func softplus(task *gmsk.Task, d int32, n int32, theta int32, t int32, X []float64, y []int32) gmsk.ResCode {
	var thetaafe, tafe, z1afe, z2afe, oneafe, expdomain int64
	var z1, z2, zcon int32
	subi := make([]int32, 2*n)
	subj := make([]int32, 3*n)
	aval := make([]float64, 2*n)
	afeidx := make([]int64, d*n+4*n)
	varidx := make([]int32, d*n+4*n)
	fval := make([]float64, d*n+4*n)
	idx := make([]int64, 3)
	var k, i, j int32
	res := gmsk.RES_OK

	res, nvar := task.GetNumVar()
	if res.NotOk() {
		return res
	}
	res, ncon := task.GetNumCon()
	if res.NotOk() {
		return res
	}
	res, nafe := task.GetNumAfe()
	if res.NotOk() {
		return res
	}

	MSKCALL(task.AppendVars(2 * n))        // z1, z2
	MSKCALL(task.AppendCons(n))            // z1 + z2 = 1
	MSKCALL(task.AppendAfes(4 * int64(n))) // theta * X[i] - t[i], -t[i], z1[i], z2[i]

	z1 = nvar
	z2 = nvar + n
	zcon = ncon
	thetaafe = nafe
	tafe = nafe + int64(n)
	z1afe = nafe + int64(2*n)
	z2afe = nafe + int64(3*n)

	// Linear constraints
	k = 0
	for i = 0; i < n; i++ {
		// z1 + z2 = 1
		subi[k] = zcon + i
		subj[k] = z1 + i
		aval[k] = 1
		k++
		subi[k] = zcon + i
		subj[k] = z2 + i
		aval[k] = 1
		k++
	}

	MSKCALL(task.PutAijList(2*n, &subi[0], &subj[0], &aval[0]))
	MSKCALL(task.PutConboundSliceConst(zcon, zcon+n, gmsk.BK_FX, 1, 1))
	MSKCALL(task.PutVarboundSliceConst(nvar, nvar+2*n, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	// Affine conic expressions
	k = 0

	// Thetas
	for i = 0; i < n; i++ {
		for j = 0; j < d; j++ {
			afeidx[k] = thetaafe + int64(i)
			varidx[k] = theta + j
			if y[i] != 0 {
				fval[k] = -1 * X[i*d+j]
			} else {
				fval[k] = X[i*d+j]
			}

			k++
		}
	}

	// -t[i]
	for i = 0; i < n; i++ {
		afeidx[k] = thetaafe + int64(i)
		varidx[k] = t + i
		fval[k] = -1
		k++
		afeidx[k] = tafe + int64(i)
		varidx[k] = t + i
		fval[k] = -1
		k++
	}

	// z1, z2
	for i = 0; i < n; i++ {
		afeidx[k] = z1afe + int64(i)
		varidx[k] = z1 + i
		fval[k] = 1
		k++
		afeidx[k] = z2afe + int64(i)
		varidx[k] = z2 + i
		fval[k] = 1
		k++
	}

	// Add the expressions
	MSKCALL(task.PutAfeFEntryList(int64(d*n+4*n), &afeidx[0], &varidx[0], &fval[0]))

	// Add a single row with the constant expression "1.0"
	res, oneafe = task.GetNumAfe()
	MSKCALL(res)
	MSKCALL(task.AppendAfes(1))
	MSKCALL(task.PutAfeG(oneafe, 1))

	// Add an exponential cone domain
	res, expdomain = task.AppendPrimalExpConeDomain()

	// Conic constraints
	for i = 0; i < n; i++ {
		idx[0] = z1afe + int64(i)
		idx[1] = oneafe
		idx[2] = thetaafe + int64(i)
		MSKCALL(task.AppendAcc(expdomain, 3, &idx[0], nil))
		idx[0] = z2afe + int64(i)
		idx[1] = oneafe
		idx[2] = tafe + int64(i)
		MSKCALL(task.AppendAcc(expdomain, 3, &idx[0], nil))
	}

	return res
}

// Model logistic regression (regularized with full 2-norm of theta)
// X - n x d matrix of data points
// y - length n vector classifying training points
// lamb - regularization parameter
func logisticRegression(env *gmsk.Env,
	n int32, // num samples
	d int32, // dimension
	X []float64,
	y []int32,
	lamb float64,
	thetaVal []float64, // result
) gmsk.ResCode {
	res := gmsk.RES_OK
	var nvar int32 = 1 + d + n
	var r, theta, t int32 = 0, 1, 1 + d
	var numafe, quadDom int64
	var i int32

	task, err := env.MakeTask(0, 0)
	if err != nil {
		log.Panic(err)
	}

	// Variables [r; theta; t]
	MSKCALL(task.AppendVars(nvar))
	MSKCALL(task.PutVarboundSliceConst(0, nvar, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	// Objective lambda*r + sum(t)
	MSKCALL(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MINIMIZE))
	MSKCALL(task.PutCj(r, lamb))
	for i = 0; i < n && res.IsOk(); i++ {
		MSKCALL(task.PutCj(t+i, 1))
	}

	// Softplus function constraints
	MSKCALL(softplus(task, d, n, theta, t, X, y))

	// Regularization
	// Append a sequence of linear expressions (r, theta) to F
	res, numafe = task.GetNumAfe()
	MSKCALL(res)
	MSKCALL(task.AppendAfes(1 + int64(d)))
	MSKCALL(task.PutAfeFEntry(numafe, r, 1.0))
	for i = 0; i < d; i++ {
		MSKCALL(task.PutAfeFEntry(numafe+int64(i)+1, theta+i, 1.0))
	}

	// Add the constraint
	res, quadDom = task.AppendQuadraticConeDomain(1 + int64(d))
	MSKCALL(res)
	MSKCALL(task.AppendAccSeq(quadDom, int64(d)+1, numafe, nil))

	res, _ = task.OptimizeTerm()
	MSKCALL(res)
	MSKCALL(task.SolutionSummary(gmsk.STREAM_LOG))

	res, _ = task.GetXxSlice(gmsk.SOL_ITR, theta, theta+d, thetaVal)
	return res
}

// Logistic regression example with MOSEK, reproduced from logistic.c in MOSEK C api.
func Example_logistic() {
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	const n int32 = 30
	X := make([]float64, 6*n*n)
	Y := make([]int32, n*n)
	var i, j int32
	theta := make([]float64, 6)

	// Test: detect and approximate a circle using degree 2 polynomials
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			k := i*n + j
			var x float64 = -1.0 + 2.0*float64(i)/float64(n-1)
			var y float64 = -1.0 + 2.0*float64(j)/float64(n-1)
			X[6*k+0] = 1.0
			X[6*k+1] = x
			X[6*k+2] = y
			X[6*k+3] = x * y
			X[6*k+4] = x * x
			X[6*k+5] = y * y
			if x*x+y*y >= 0.69 {
				Y[k] = 1
			} else {
				Y[k] = 0
			}
		}
	}

	MSKCALL(logisticRegression(env, n*n, 6, X, Y, 0.1, theta))

	for i = 0; i < 6; i++ {
		if math.Abs(theta[i]) <= 1e-6 {
			theta[i] = 0
		}
		fmt.Printf("%.2e\n", theta[i])
	}
	// Output:
	// -5.37e+01
	// 0.00e+00
	// 0.00e+00
	// 0.00e+00
	// 7.72e+01
	// 7.72e+01
}
