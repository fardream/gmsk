package gmsk_test

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/fardream/gmsk"
)

// Example on how to use included BLAS/LAPACK routines in MOSEK,
// reproduced from blas_lapack.c in MOSEK C api.
func Example_blas_lapack() {
	print_matrix := func(x *float64, r, c int32) {
		var i, j int32
		s := unsafe.Slice(x, r*c)
		for i = 0; i < r; i++ {
			for j = 0; j < c; j++ {
				if j >= c-1 {
					fmt.Printf("%f\n", s[j*r+i])
				} else {
					fmt.Printf("%f ", s[j*r+i])
				}
			}
		}
	}

	r := gmsk.RES_OK

	const n, m, k int32 = 3, 2, 3
	const alpha, beta float64 = 2.0, 0.5

	x := []float64{1.0, 1.0, 1.0}
	y := []float64{1.0, 2.0, 3.0}
	z := []float64{1.0, 1.0}
	A := []float64{1.0, 1.0, 2.0, 2.0, 3.0, 3.0}
	B := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
	C := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	D := []float64{1.0, 1.0, 1.0, 1.0}
	Q := []float64{1.0, 0.0, 0.0, 2.0}
	v := []float64{0.0, 0.0, 0.0}

	var xy float64
	/* BLAS routines*/
	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)
	fmt.Printf("n=%d m=%d k=%d\n", m, n, k)
	fmt.Printf("alpha=%f\n", alpha)
	fmt.Printf("beta=%f\n", beta)

	r, xy = env.Dot(n, &x[0], &y[0])
	fmt.Printf("dot results= %f r=%d\n", xy, r)

	print_matrix(&x[0], 1, n)
	print_matrix(&y[0], 1, n)

	r = env.Axpy(n, alpha, &x[0], &y[0])
	fmt.Println("axpy results is")
	print_matrix(&y[0], 1, n)

	r = env.Gemv(gmsk.TRANSPOSE_NO, m, n, alpha, &A[0], &x[0], beta, &z[0])
	fmt.Printf("gemv results is (r=%d)\n", r)
	print_matrix(&z[0], 1, m)

	r = env.Gemm(gmsk.TRANSPOSE_NO, gmsk.TRANSPOSE_NO, m, n, k, alpha, &A[0], &B[0], beta, &C[0])
	fmt.Printf("gemm results is (r=%d)\n", r)
	print_matrix(&C[0], m, n)

	r = env.Syrk(gmsk.UPLO_LO, gmsk.TRANSPOSE_NO, m, k, 1., &A[0], beta, &D[0])
	fmt.Printf("syrk results is (r=%d)\n", r)
	print_matrix(&D[0], m, m)

	/* LAPACK routines*/

	r = env.Potrf(gmsk.UPLO_LO, m, &Q[0])
	fmt.Printf("potrf results is (r=%d)\n", r)
	print_matrix(&Q[0], m, m)

	r = env.Syeig(gmsk.UPLO_LO, m, &Q[0], &v[0])
	fmt.Printf("syeig results is (r=%d)\n", r)
	print_matrix(&v[0], 1, m)

	r = env.Syevd(gmsk.UPLO_LO, m, &Q[0], &v[0])
	fmt.Printf("syevd results is (r=%d)\n", r)
	print_matrix(&v[0], 1, m)
	print_matrix(&Q[0], m, m)
	// Output:
	// n=2 m=3 k=3
	// alpha=2.000000
	// beta=0.500000
	// dot results= 6.000000 r=0
	// 1.000000 1.000000 1.000000
	// 1.000000 2.000000 3.000000
	// axpy results is
	// 3.000000 4.000000 5.000000
	// gemv results is (r=0)
	// 12.500000 12.500000
	// gemm results is (r=0)
	// 12.500000 13.500000 14.500000
	// 13.000000 14.000000 15.000000
	// syrk results is (r=0)
	// 14.500000 1.000000
	// 14.500000 14.500000
	// potrf results is (r=0)
	// 1.000000 0.000000
	// 0.000000 1.414214
	// syeig results is (r=0)
	// 1.000000 1.414214
	// syevd results is (r=0)
	// 1.000000 1.414214
	// 1.000000 0.000000
	// 0.000000 1.000000
}
