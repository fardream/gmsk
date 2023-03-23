package gmsk_test

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/fardream/gmsk"
)

// Portfolio optimization example with factor model, reproduced from
// portfolio_6_factor.c in MOSEK C api.
func Example_portfolio6Factor() {
	res := gmsk.RES_OK
	checkOk := func(r uint32) {
		res = r
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodeDescSimple(r)

			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	get_nr_nc := func(m [][]gmsk.Realt) (nr int, nc int) {
		nr = len(m)
		if nr > 0 {
			nc = len(m[0])
		}

		return
	}

	// array_print := func(a []gmsk.Realt) {
	// 	fmt.Print("[")
	// 	for _, aj := range a {
	// 		fmt.Printf("%f, ", aj)
	// 	}
	// 	fmt.Print("\b\b]\n")
	// }

	// matrix_print := func(m [][]gmsk.Realt) {
	// 	var i, j int
	// 	nr, nc := get_nr_nc(m)
	// 	for i = 0; i < nr; i++ {
	// 		array_print(m[i])
	// 	}
	// }

	matrix_alloc := func(dim1, dim2 int) [][]gmsk.Realt {
		result := make([][]gmsk.Realt, dim1)
		for i := 0; i < dim1; i++ {
			result[i] = make([]gmsk.Realt, dim2)
		}

		return result
	}

	vector_alloc := func(dim int) []gmsk.Realt {
		return make([]gmsk.Realt, dim)
	}

	// sum := func(x []gmsk.Realt) gmsk.Realt {
	// 	var r gmsk.Realt
	// 	for _, ax := range x {
	// 		r += ax
	// 	}
	// 	return r
	// }

	// Vectorize matrix (column-major order)
	mat_to_vec_c := func(m [][]gmsk.Realt) []gmsk.Realt {
		ni, nj := get_nr_nc(m)
		c := make([]gmsk.Realt, ni*nj)
		for j := 0; j < nj; j++ {
			for i := 0; i < ni; i++ {
				c[j*ni+i] = m[i][j]
			}
		}

		return c
	}

	// Reshape vector to matrix (column-major order)
	vec_to_mat_c := func(c []gmsk.Realt, ni, nj int) [][]gmsk.Realt {
		m := matrix_alloc(ni, nj)
		for j := 0; j < nj; j++ {
			for i := 0; i < ni; i++ {
				m[i][j] = c[j*ni+i]
			}
		}

		return m
	}
	// Reshape vector to matrix (row-major order)
	vec_to_mat_r := func(r []gmsk.Realt, ni, nj int) [][]gmsk.Realt {
		m := matrix_alloc(ni, nj)
		for i := 0; i < ni; i++ {
			for j := 0; j < nj; j++ {
				m[i][j] = r[i*nj+j]
			}
		}

		return m
	}

	cholesky := func(env *gmsk.Env, m [][]gmsk.Realt) [][]gmsk.Realt {
		nr, _ := get_nr_nc(m)
		n := nr
		vecs := mat_to_vec_c(m)
		checkOk(gmsk.POTRF(env, gmsk.UPLO_LO, int32(n), &vecs[0]))
		s := vec_to_mat_c(vecs, n, n)
		// Zero out upper triangular part (MSK_potrf does not use it, original matrix values remain there)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				s[i][j] = 0
			}
		}

		return s
	}

	// Matrix multiplication
	matrix_mul := func(env *gmsk.Env, a [][]gmsk.Realt, b [][]gmsk.Realt) [][]gmsk.Realt {
		anr, _ := get_nr_nc(a)
		bnr, bnc := get_nr_nc(b)

		na := anr
		nb := bnc
		k := bnr

		vecm := vector_alloc(na * nb)
		veca := mat_to_vec_c(a)
		vecb := mat_to_vec_c(b)

		checkOk(gmsk.GEMM(env, gmsk.TRANSPOSE_NO, gmsk.TRANSPOSE_NO, int32(na), int32(nb), int32(k), 1, &veca[0], &vecb[0], 1, &vecm[0]))

		ab := vec_to_mat_c(vecm, na, nb)

		return ab
	}

	var expret gmsk.Realt

	const n int32 = 8
	var w gmsk.Realt = 1.0
	mu := []gmsk.Realt{0.07197, 0.15518, 0.17535, 0.08981, 0.42896, 0.39292, 0.32171, 0.18379}
	x0 := []gmsk.Realt{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

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

	// NOTE: Here we specify matrices as vectors (row major order) to avoid having
	// to initialize them as double(*)[] type, which is incompatible with double**.

	// Factor exposure matrix
	vecB := []gmsk.Realt{
		0.4256, 0.1869,
		0.2413, 0.3877,
		0.2235, 0.3697,
		0.1503, 0.4612,
		1.5325, -0.2633,
		1.2741, -0.2613,
		0.6939, 0.2372,
		0.5425, 0.2116,
	}

	B := vec_to_mat_r(vecB, int(n), 2)
	// Factor covariance matrix
	vecS_F := []gmsk.Realt{
		0.0620, 0.0577,
		0.0577, 0.0908,
	}
	S_F := vec_to_mat_r(vecS_F, 2, 2)
	// Specific risk components
	theta := []gmsk.Realt{0.0720, 0.0508, 0.0377, 0.0394, 0.0663, 0.0224, 0.0417, 0.0459}

	P := cholesky(env, S_F)
	G_factor := matrix_mul(env, B, P)

	_, _k := get_nr_nc(G_factor)
	k := gmsk.Int64t(_k)

	gammas := []gmsk.Realt{0.24, 0.28, 0.32, 0.36, 0.4, 0.44, 0.48}
	num_gammas := int32(len(gammas))
	var totalBudget gmsk.Realt

	// Offset of variables into the API variable.
	const numvar, voff_x int32 = 8, 0

	// Constraint offset
	const coff_bud int32 = 0

	// Holding variable x of length n
	// No other auxiliary variables are needed in this formulation
	checkOk(task.AppendVars(numvar))
	// Setting up variable x
	for j := int32(0); j < n; j++ {
		/* Optionally we can give the variables names */
		checkOk(task.PutVarName(voff_x+j, fmt.Sprintf("x[%d]", 1+j)))
		/* No short-selling - x^l = 0, x^u = inf */
		checkOk(task.PutVarbound(voff_x+j, gmsk.BK_LO, 0, gmsk.INFINITY))
	}

	// One linear constraint: total budget
	checkOk(task.AppendCons(1))
	checkOk(task.PutConName(0, "budget"))
	for j := int32(0); j < n; j++ {
		/* Coefficients in the first row of A */
		checkOk(task.PutAij(coff_bud, voff_x+j, 1))
	}
	totalBudget = w
	for i := int32(0); i < n; i++ {
		totalBudget += x0[i]
	}
	checkOk(task.PutConBound(coff_bud, gmsk.BK_FX, totalBudget, totalBudget))

	// Input (gamma, G_factor_T x, diag(sqrt(theta))*x) in the AFE (affine expression) storage
	// We need k+n+1 rows and we fill them in in three parts
	task.AppendAfes(k + gmsk.Int64t(n) + 1)
	// 1. The first affine expression = gamma, will be specified later
	// 2. The next k expressions comprise G_factor_T*x, we add them column by column since
	//    G_factor is stored row-wise and we transpose on the fly
	afeidx := make([]gmsk.Int64t, k)
	for i := gmsk.ZeroInt64t; i < k; i++ {
		afeidx[i] = i + 1
	}
	for i := int32(0); i < n; i++ {
		checkOk(task.PutAfeFCol(i, k, &afeidx[0], &G_factor[i][0])) // i-th row of G_factor goes in i-th column of F
	}
	// 3. The remaining n rows contain sqrt(theta) on the diagonal
	for i := int32(0); i < n; i++ {
		checkOk(task.PutAfeFEntry(k+1+gmsk.Int64t(i), voff_x+i, gmsk.Realt(math.Sqrt(float64(theta[i])))))
	}

	// Input the affine conic constraint (gamma, G_factor_T x, diag(sqrt(theta))*x) \in QCone
	// Add the quadratic domain of dimension k+n+1
	res, qdom := task.AppendQuadraticConeDomain(k + 1 + gmsk.Int64t(n))
	checkOk(res)
	// Add the constraint
	checkOk(task.AppendAccSeq(qdom, k+1+gmsk.Int64t(n), 0, nil))
	checkOk(task.PutAccName(0, "risk"))

	// Objective: maximize expected return mu^T x
	for j := int32(0); j < n; j++ {
		checkOk(task.PutCj(voff_x+j, mu[j]))
	}
	checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	/* No log output */
	checkOk(task.PutIntParam(gmsk.IPAR_LOG, 0))

	for i := int32(0); i < num_gammas; i++ {
		gamma := gammas[i]

		// Specify gamma in ACC
		checkOk(task.PutAfeG(0, gamma))

		/* Dump the problem to a human readable PTF file. */
		checkOk(task.WriteDataHandle(os.Stderr, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE))

		res, _ = task.OptimizeTerm()
		checkOk(res)

		/* Display the solution summary for quick inspection of results. */
		task.SolutionSummary(gmsk.STREAM_LOG) // not using MSG becasue MSG is going to Stdout right now

		expret = 0

		/* Read the x variables one by one and compute expected return. */
		/* Can also be obtained as value of the objective. */
		for j := int32(0); j < n; j++ {
			res, xx := task.GetXxSlice(gmsk.SOL_ITR, voff_x+j, voff_x+j+1, nil)
			checkOk(res)
			xj := xx[0]
			expret += mu[j] * xj
		}

		fmt.Printf("\nExpected return %e for gamma %e\n", expret, gamma)
	}
	// Output: Expected return 3.162054e-01 for gamma 2.400000e-01
	//
	// Expected return 3.776816e-01 for gamma 2.800000e-01
	//
	// Expected return 4.081833e-01 for gamma 3.200000e-01
	//
	// Expected return 4.186580e-01 for gamma 3.600000e-01
	//
	// Expected return 4.264498e-01 for gamma 4.000000e-01
	//
	// Expected return 4.289600e-01 for gamma 4.400000e-01
	//
	// Expected return 4.289600e-01 for gamma 4.800000e-01
}
