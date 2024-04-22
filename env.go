// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"unsafe"
)

// Axpy is wrapping [MSK_axpy],
// performs y = a*x + y where x/y are vectors.
//
// [MSK_axpy]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.axpy
func (env *Env) Axpy(
	n int32,
	alpha float64,
	x []float64,
	y []float64,
) error {
	return ResCode(
		C.MSK_axpy(
			env.getEnv(),
			C.MSKint32t(n),
			C.MSKrealt(alpha),
			(*C.MSKrealt)(getPtrToFirst(x)),
			(*C.MSKrealt)(getPtrToFirst(y)),
		),
	).ToError()
}

// CheckInAll is wrapping [MSK_checkinall],
// Check in all unused license features to the license token server.
//
// [MSK_checkinall]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.checkinall
func (env *Env) CheckInAll() error {
	return ResCode(
		C.MSK_checkinall(
			env.getEnv(),
		),
	).ToError()
}

// CheckInLicense is wrapping [MSK_checkinlicense],
// Check in a license feature back to the license server ahead of time.
//
// Arguments:
//
//   - `feature` Feature to check in to the license system.
//
// [MSK_checkinlicense]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.checkinlicense
func (env *Env) CheckInLicense(
	feature Feature,
) error {
	return ResCode(
		C.MSK_checkinlicense(
			env.getEnv(),
			C.MSKfeaturee(feature),
		),
	).ToError()
}

// CheckMemenv is wrapping [MSK_checkmemenv]
//
// [MSK_checkmemenv]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.checkmemenv
func (env *Env) CheckMemenv(
	file string,
	line int32,
) error {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	return ResCode(
		C.MSK_checkmemenv(
			env.getEnv(),
			c_file,
			C.MSKint32t(line),
		),
	).ToError()
}

// CheckOutLicense is wrapping [MSK_checkoutlicense],
// Check out a license feature from the license server ahead of time.
//
// Arguments:
//
//   - `feature` Feature to check out from the license system.
//
// [MSK_checkoutlicense]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.checkoutlicense
func (env *Env) CheckOutLicense(
	feature Feature,
) error {
	return ResCode(
		C.MSK_checkoutlicense(
			env.getEnv(),
			C.MSKfeaturee(feature),
		),
	).ToError()
}

// CheckVersion is wrapping [MSK_checkversion],
// Compares a version of the MOSEK DLL with a specified version.
//
// Arguments:
//
//   - `major` Major version number.
//   - `minor` Minor version number.
//   - `revision` Revision number.
//
// [MSK_checkversion]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.checkversion
func (env *Env) CheckVersion(
	major int32,
	minor int32,
	revision int32,
) error {
	return ResCode(
		C.MSK_checkversion(
			env.getEnv(),
			C.MSKint32t(major),
			C.MSKint32t(minor),
			C.MSKint32t(revision),
		),
	).ToError()
}

// Dot is wrapping [MSK_dot],
// Computes the inner product of two vectors.
//
// Arguments:
//
//   - `n` Length of the vectors.
//   - `x` The x vector.
//   - `y` The y vector.
//   - `xty` The result of the inner product.
//
// [MSK_dot]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.dot
func (env *Env) Dot(
	n int32,
	x []float64,
	y []float64,
) (xty float64, r error) {
	r = ResCode(
		C.MSK_dot(
			env.getEnv(),
			C.MSKint32t(n),
			(*C.MSKrealt)(getPtrToFirst(x)),
			(*C.MSKrealt)(getPtrToFirst(y)),
			(*C.MSKrealt)(&xty),
		),
	).ToError()

	return
}

// EchoIntro is wrapping [MSK_echointro],
// Prints an intro to message stream.
//
// Arguments:
//
//   - `longver` If non-zero, then the intro is slightly longer.
//
// [MSK_echointro]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.echointro
func (env *Env) EchoIntro(
	longver int32,
) error {
	return ResCode(
		C.MSK_echointro(
			env.getEnv(),
			C.MSKint32t(longver),
		),
	).ToError()
}

// Expirylicenses is wrapping [MSK_expirylicenses],
// Reports when the first license feature expires.
//
// Arguments:
//
//   - `expiry` If nonnegative, then it is the minimum number days to expiry of any feature that has been checked out.
//
// [MSK_expirylicenses]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.expirylicenses
func (env *Env) Expirylicenses() (expiry int64, r error) {
	r = ResCode(
		C.MSK_expirylicenses(
			env.getEnv(),
			(*C.MSKint64t)(&expiry),
		),
	).ToError()

	return
}

// Gemm is wrapping [MSK_gemm],
// Performs a dense matrix multiplication.
//
// Arguments:
//
//   - `transa` Indicates whether the matrix A must be transposed.
//   - `transb` Indicates whether the matrix B must be transposed.
//   - `m` Indicates the number of rows of matrix C.
//   - `n` Indicates the number of columns of matrix C.
//   - `k` Specifies the common dimension along which op(A) and op(B) are multiplied.
//   - `alpha` A scalar value multiplying the result of the matrix multiplication.
//   - `a` The pointer to the array storing matrix A in a column-major format.
//   - `b` The pointer to the array storing matrix B in a column-major format.
//   - `beta` A scalar value that multiplies C.
//   - `c` The pointer to the array storing matrix C in a column-major format.
//
// [MSK_gemm]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.gemm
func (env *Env) Gemm(
	transa Transpose,
	transb Transpose,
	m int32,
	n int32,
	k int32,
	alpha float64,
	a []float64,
	b []float64,
	beta float64,
	c []float64,
) error {
	return ResCode(
		C.MSK_gemm(
			env.getEnv(),
			C.MSKtransposee(transa),
			C.MSKtransposee(transb),
			C.MSKint32t(m),
			C.MSKint32t(n),
			C.MSKint32t(k),
			C.MSKrealt(alpha),
			(*C.MSKrealt)(getPtrToFirst(a)),
			(*C.MSKrealt)(getPtrToFirst(b)),
			C.MSKrealt(beta),
			(*C.MSKrealt)(getPtrToFirst(c)),
		),
	).ToError()
}

// Gemv is wrapping [MSK_gemv],
// calculates y = aAx + by, where A is matrix, x,y is vector, and a b are scalars.
//
// [MSK_gemv]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.gemv
func (env *Env) Gemv(
	transa Transpose,
	m int32,
	n int32,
	alpha float64,
	a []float64,
	x []float64,
	beta float64,
	y []float64,
) error {
	return ResCode(
		C.MSK_gemv(
			env.getEnv(),
			C.MSKtransposee(transa),
			C.MSKint32t(m),
			C.MSKint32t(n),
			C.MSKrealt(alpha),
			(*C.MSKrealt)(getPtrToFirst(a)),
			(*C.MSKrealt)(getPtrToFirst(x)),
			C.MSKrealt(beta),
			(*C.MSKrealt)(getPtrToFirst(y)),
		),
	).ToError()
}

// GetSymbcondim is wrapping [MSK_getsymbcondim]
//
// [MSK_getsymbcondim]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.getsymbcondim
func (env *Env) GetSymbcondim(
	num []int32,
	maxlen []uint64,
) error {
	return ResCode(
		C.MSK_getsymbcondim(
			env.getEnv(),
			(*C.MSKint32t)(getPtrToFirst(num)),
			(*C.size_t)(getPtrToFirst(maxlen)),
		),
	).ToError()
}

// Iparvaltosymnam is wrapping [MSK_iparvaltosymnam]
//
// [MSK_iparvaltosymnam]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.iparvaltosymnam
func (env *Env) Iparvaltosymnam(
	whichparam IParam,
	whichvalue int32,
	symbolicname *byte,
) error {
	return ResCode(
		C.MSK_iparvaltosymnam(
			env.getEnv(),
			C.MSKiparame(whichparam),
			C.MSKint32t(whichvalue),
			(*C.char)(unsafe.Pointer(symbolicname)),
		),
	).ToError()
}

// LinkFiletoenvstream is wrapping [MSK_linkfiletoenvstream]
//
// [MSK_linkfiletoenvstream]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.linkfiletoenvstream
func (env *Env) LinkFiletoenvstream(
	whichstream StreamType,
	filename string,
	append int32,
) error {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	return ResCode(
		C.MSK_linkfiletoenvstream(
			env.getEnv(),
			C.MSKstreamtypee(whichstream),
			c_filename,
			C.MSKint32t(append),
		),
	).ToError()
}

// Potrf is wrapping [MSK_potrf],
// Computes a Cholesky factorization of a dense matrix.
//
// Arguments:
//
//   - `uplo` Indicates whether the upper or lower triangular part of the matrix is stored.
//   - `n` Dimension of the symmetric matrix.
//   - `a` A symmetric matrix stored in column-major order.
//
// [MSK_potrf]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.potrf
func (env *Env) Potrf(
	uplo UpLo,
	n int32,
	a []float64,
) error {
	return ResCode(
		C.MSK_potrf(
			env.getEnv(),
			C.MSKuploe(uplo),
			C.MSKint32t(n),
			(*C.MSKrealt)(getPtrToFirst(a)),
		),
	).ToError()
}

// PutLicenseCode is wrapping [MSK_putlicensecode],
// Input a runtime license code.
//
// Arguments:
//
//   - `code` A license key string.
//
// [MSK_putlicensecode]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.putlicensecode
func (env *Env) PutLicenseCode(
	code []int32,
) error {
	return ResCode(
		C.MSK_putlicensecode(
			env.getEnv(),
			(*C.MSKint32t)(getPtrToFirst(code)),
		),
	).ToError()
}

// PutLicenseDebug is wrapping [MSK_putlicensedebug],
// Enables debug information for the license system.
//
// Arguments:
//
//   - `licdebug` Enable output of license check-out debug information.
//
// [MSK_putlicensedebug]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.putlicensedebug
func (env *Env) PutLicenseDebug(
	licdebug int32,
) error {
	return ResCode(
		C.MSK_putlicensedebug(
			env.getEnv(),
			C.MSKint32t(licdebug),
		),
	).ToError()
}

// PutLicensePath is wrapping [MSK_putlicensepath],
// Set the path to the license file.
//
// Arguments:
//
//   - `licensepath` A path specifying where to search for the license.
//
// [MSK_putlicensepath]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.putlicensepath
func (env *Env) PutLicensePath(
	licensepath string,
) error {
	c_licensepath := C.CString(licensepath)
	defer C.free(unsafe.Pointer(c_licensepath))

	return ResCode(
		C.MSK_putlicensepath(
			env.getEnv(),
			c_licensepath,
		),
	).ToError()
}

// PutLicenseWait is wrapping [MSK_putlicensewait],
// Control whether mosek should wait for an available license if no license is available.
//
// Arguments:
//
//   - `licwait` Enable waiting for a license.
//
// [MSK_putlicensewait]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.putlicensewait
func (env *Env) PutLicenseWait(
	licwait int32,
) error {
	return ResCode(
		C.MSK_putlicensewait(
			env.getEnv(),
			C.MSKint32t(licwait),
		),
	).ToError()
}

// ResetExpiryLicenses is wrapping [MSK_resetexpirylicenses],
// Reset the license expiry reporting startpoint.
//
// [MSK_resetexpirylicenses]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.resetexpirylicenses
func (env *Env) ResetExpiryLicenses() error {
	return ResCode(
		C.MSK_resetexpirylicenses(
			env.getEnv(),
		),
	).ToError()
}

// SparseTriangularSolveDense is wrapping [MSK_sparsetriangularsolvedense],
// Solves a sparse triangular system of linear equations.
//
// Arguments:
//
//   - `transposed` Controls whether the solve is with L or the transposed L.
//   - `lnzc` lnzc\[j\] is the number of nonzeros in column j.
//   - `lptrc` lptrc\[j\] is a pointer to the first row index and value in column j.
//   - `lsubc` Row indexes for each column stored sequentially.
//   - `lvalc` The value corresponding to row indexed stored lsubc.
//   - `b` The right-hand side of linear equation system to be solved as a dense vector.
//
// [MSK_sparsetriangularsolvedense]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.sparsetriangularsolvedense
func (env *Env) SparseTriangularSolveDense(
	transposed Transpose,
	n int32,
	lnzc []int32,
	lptrc []int64,
	lensubnval int64,
	lsubc []int32,
	lvalc []float64,
	b []float64,
) error {
	return ResCode(
		C.MSK_sparsetriangularsolvedense(
			env.getEnv(),
			C.MSKtransposee(transposed),
			C.MSKint32t(n),
			(*C.MSKint32t)(getPtrToFirst(lnzc)),
			(*C.MSKint64t)(getPtrToFirst(lptrc)),
			C.MSKint64t(lensubnval),
			(*C.MSKint32t)(getPtrToFirst(lsubc)),
			(*C.MSKrealt)(getPtrToFirst(lvalc)),
			(*C.MSKrealt)(getPtrToFirst(b)),
		),
	).ToError()
}

// Syeig is wrapping [MSK_syeig],
// Computes all eigenvalues of a symmetric dense matrix.
//
// Arguments:
//
//   - `uplo` Indicates whether the upper or lower triangular part is used.
//   - `n` Dimension of the symmetric input matrix.
//   - `a` Input matrix A.
//   - `w` Array of length at least n containing the eigenvalues of A.
//
// [MSK_syeig]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.syeig
func (env *Env) Syeig(
	uplo UpLo,
	n int32,
	a []float64,
	w []float64,
) error {
	return ResCode(
		C.MSK_syeig(
			env.getEnv(),
			C.MSKuploe(uplo),
			C.MSKint32t(n),
			(*C.MSKrealt)(getPtrToFirst(a)),
			(*C.MSKrealt)(getPtrToFirst(w)),
		),
	).ToError()
}

// Syevd is wrapping [MSK_syevd],
// Computes all the eigenvalues and eigenvectors of a symmetric dense matrix, and thus its eigenvalue decomposition.
//
// Arguments:
//
//   - `uplo` Indicates whether the upper or lower triangular part is used.
//   - `n` Dimension of the symmetric input matrix.
//   - `a` Input matrix A.
//   - `w` Array of length at least n containing the eigenvalues of A.
//
// [MSK_syevd]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.syevd
func (env *Env) Syevd(
	uplo UpLo,
	n int32,
	a []float64,
	w []float64,
) error {
	return ResCode(
		C.MSK_syevd(
			env.getEnv(),
			C.MSKuploe(uplo),
			C.MSKint32t(n),
			(*C.MSKrealt)(getPtrToFirst(a)),
			(*C.MSKrealt)(getPtrToFirst(w)),
		),
	).ToError()
}

// Syrk is wrapping [MSK_syrk],
// Performs a rank-k update of a symmetric matrix.
//
// Arguments:
//
//   - `uplo` Indicates whether the upper or lower triangular part of C is used.
//   - `trans` Indicates whether the matrix A must be transposed.
//   - `n` Specifies the order of C.
//   - `k` Indicates the number of rows or columns of A, and its rank.
//   - `alpha` A scalar value multiplying the result of the matrix multiplication.
//   - `a` The pointer to the array storing matrix A in a column-major format.
//   - `beta` A scalar value that multiplies C.
//   - `c` The pointer to the array storing matrix C in a column-major format.
//
// [MSK_syrk]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.syrk
func (env *Env) Syrk(
	uplo UpLo,
	trans Transpose,
	n int32,
	k int32,
	alpha float64,
	a []float64,
	beta float64,
	c []float64,
) error {
	return ResCode(
		C.MSK_syrk(
			env.getEnv(),
			C.MSKuploe(uplo),
			C.MSKtransposee(trans),
			C.MSKint32t(n),
			C.MSKint32t(k),
			C.MSKrealt(alpha),
			(*C.MSKrealt)(getPtrToFirst(a)),
			C.MSKrealt(beta),
			(*C.MSKrealt)(getPtrToFirst(c)),
		),
	).ToError()
}

// UnlinkFuncfromenvstream is wrapping [MSK_unlinkfuncfromenvstream]
//
// [MSK_unlinkfuncfromenvstream]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.env.unlinkfuncfromenvstream
func (env *Env) UnlinkFuncfromenvstream(
	whichstream StreamType,
) error {
	return ResCode(
		C.MSK_unlinkfuncfromenvstream(
			env.getEnv(),
			C.MSKstreamtypee(whichstream),
		),
	).ToError()
}
