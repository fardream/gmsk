package gmsk_test

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/fardream/gmsk"
)

// Example of geometric program with exponential cones and log-sum-exp,
// reproduced from gp1.c in MOSEK C api.
//
//	maximize     h*w*d
//	subjecto to  2*(h*w + h*d) <= Awall
//	             w*d <= Afloor
//	             alpha <= h/w <= beta
//	             gamma <= d/w <= delta
//
//	Variable substitutions:  h = exp(x), w = exp(y), d = exp(z).
//
//	maximize     x+y+z
//	subject      log( exp(x+y+log(2/Awall)) + exp(x+z+log(2/Awall)) ) <= 0
//	                             y+z <= log(Afloor)
//	             log( alpha ) <= x-y <= log( beta )
//	             log( gamma ) <= z-y <= log( delta )
func Example_geometricProgram1() {
	checkOk := func(r gmsk.ResCode) {
		if !r.IsOk() {
			_, sym, desc := gmsk.GetCodeDescSimple(r)
			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	const Aw float64 = 200.0
	const Af float64 = 50.0
	const alpha float64 = 2.0
	const beta float64 = 10.0
	const gamma float64 = 2.0
	// const delta float64 = 10.0

	r := gmsk.RES_OK
	hwd := [3]float64{}

	// max_volume_box - begin ------------------------------------------------

	// Basic dimensions of our problem
	const numvar int32 = 3 // Variables in original problem
	const numcon int32 = 3 // Linear constraints in original problem

	// Linear part of the problem involving x, y, z
	cval := []float64{1, 1, 1}
	asubi := []int32{0, 0, 1, 1, 2, 2}
	asubj := []int32{1, 2, 0, 1, 2, 1}
	const alen int32 = 6
	aval := []float64{1, 1, 1, -1, 1, -1}
	bkc := []gmsk.BoundKey{gmsk.BK_UP, gmsk.BK_RA, gmsk.BK_RA}
	blc := []float64{-gmsk.INFINITY, math.Log(alpha), math.Log(gamma)}
	buc := []float64{math.Log(Af), math.Log(beta), math.Log(beta)}

	// Affine conic constraint data of the problem
	var expdomidx, rzerodomidx int64
	const numafe, f_nnz int64 = 6, 8
	afeidx := [8]int64{0, 1, 2, 2, 3, 3, 5, 5}
	varidx := [8]int32{3, 4, 0, 1, 0, 2, 3, 4}
	f_val := [8]float64{1, 1, 1, 1, 1, 1, 1, 1}
	g := [6]float64{0, 0, math.Log(2 / Aw), math.Log(2 / Aw), 1, -1}

	xyz := make([]float64, 3)

	task, err := gmsk.MakeTask(nil, 0, 0)
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	checkOk(task.AppendVars(numvar))

	checkOk(task.AppendCons(numcon))

	checkOk(task.AppendAfes(numafe))

	// Objective is the sum of three first variables
	checkOk(task.PutObjsense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))
	checkOk(task.PutCSlice(0, numvar, &cval[0]))
	checkOk(task.PutVarboundSliceConst(0, numvar, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	// Add the three linear constraints
	checkOk(task.PutAijList(alen, &asubi[0], &asubj[0], &aval[0]))
	checkOk(task.PutConboundSlice(0, numvar, &bkc[0], &blc[0], &buc[0]))

	acc1_afeidx := []int64{0, 4, 2}
	acc2_afeidx := []int64{1, 4, 3}
	acc3_afeidx := []int64{5}

	// Affine expressions appearing in affine conic constraints
	// in this order:
	// u1, u2, x+y+log(2/Awall), x+z+log(2/Awall), 1.0, u1+u2-1.0
	checkOk(task.AppendVars(2))
	checkOk(task.PutVarboundSliceConst(numvar, numvar+2, gmsk.BK_FR, -gmsk.INFINITY, gmsk.INFINITY))

	checkOk(task.PutAfeFEntryList(f_nnz, &afeidx[0], &varidx[0], &f_val[0]))
	checkOk(task.PutAfeGSlice(0, numafe, &g[0]))

	/* Append the primal exponential cone domain */
	r, expdomidx = task.AppendPrimalExpConeDomain()
	checkOk(r)

	/* (u1, 1, x+y+log(2/Awall)) \in EXP */
	checkOk(task.AppendAcc(expdomidx, 3, &acc1_afeidx[0], nil))

	/* (u2, 1, x+z+log(2/Awall)) \in EXP */
	checkOk(task.AppendAcc(expdomidx, 3, &acc2_afeidx[0], nil))

	/* The constraint u1+u2-1 \in \ZERO is added also as an ACC */
	r, rzerodomidx = task.AppendRZeroDomain(1)
	checkOk(r)
	checkOk(task.AppendAcc(rzerodomidx, 1, &acc3_afeidx[0], nil))

	// Solve and map to original h, w, d
	r, trmcode := task.OptimizeTerm()
	checkOk(r)

	r, solsta := task.GetSolSta(gmsk.SOL_ITR)
	checkOk(r)

	if solsta != gmsk.SOL_STA_OPTIMAL {
		fmt.Printf("Solution not optimal, termination code %d.\n", trmcode)
		checkOk(trmcode)
	}

	r, xyz = task.GetXxSlice(gmsk.SOL_ITR, 0, numvar, xyz)
	checkOk(r)
	for i := 0; i < int(numvar); i++ {
		hwd[i] = math.Exp(xyz[i])
	}

	// max_volume_box - end   ------------------------------------------------

	fmt.Printf("Solution h=%.4f w=%.4f d=%.4f\n", hwd[0], hwd[1], hwd[2])
	// Output: Solution h=8.1641 w=4.0821 d=8.1666
}
