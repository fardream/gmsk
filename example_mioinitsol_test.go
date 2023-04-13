package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

// Mixed integer programming with initial solution, reproduced from mioinitsol.c in MOSEK C api.
func Example_mixedIntegerProgrammingWithInitialSolution_mioinitsol() {
	checkOk := func(r gmsk.ResCode) {
		if !r.IsOk() {
			_, sym, desc := gmsk.GetCodedesc(r)
			log.Panicf("failed: %s %s", sym, desc)
		}
	}

	const numvar int32 = 4
	const numcon int32 = 1
	const numintvar int32 = 3

	r := gmsk.RES_OK

	c := []float64{7.0, 10.0, 1.0, 5.0}

	bkc := []gmsk.BoundKey{gmsk.BK_UP}
	blc := []float64{-gmsk.INFINITY}
	buc := []float64{2.5}

	bkx := []gmsk.BoundKey{gmsk.BK_LO, gmsk.BK_LO, gmsk.BK_LO, gmsk.BK_LO}
	blx := []float64{0.0, 0.0, 0.0, 0.0}
	bux := []float64{gmsk.INFINITY, gmsk.INFINITY, gmsk.INFINITY, gmsk.INFINITY}

	ptrb := []int32{0, 1, 2, 3}
	ptre := []int32{1, 2, 3, 4}
	asub := []int32{0, 0, 0, 0}

	aval := []float64{1.0, 1.0, 1.0, 1.0}
	intsub := []int32{0, 1, 2}
	var j int32

	xx := make([]float64, 4)

	env, err := gmsk.MakeEnv()
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteEnv(env)

	task, err := env.MakeTask(0, 0)
	if err != nil {
		log.Panic(err)
	}
	defer gmsk.DeleteTask(task)

	checkOk(task.LinkFuncToTaskStream(gmsk.STREAM_LOG, os.Stderr))

	checkOk(task.InputData(
		numcon, numvar,
		numcon, numvar,
		&c[0],
		0,
		&ptrb[0],
		&ptre[0],
		&asub[0],
		&aval[0],
		&bkc[0],
		&blc[0],
		&buc[0],
		&bkx[0],
		&blx[0],
		&bux[0]))

	checkOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MAXIMIZE))

	for j = 0; j < numintvar && r.IsOk(); j++ {
		r = task.PutVarType(intsub[j], gmsk.VAR_TYPE_INT)
	}
	checkOk(r)
	/* Assign values to integer variables
	   (we only set a slice of xx) */
	xxInit := []float64{1, 1, 0}
	checkOk(task.PutXxSlice(gmsk.SOL_ITG, 0, 3, &xxInit[0]))

	/* Request constructing the solution from integer variable values */
	checkOk(task.PutIntParam(gmsk.IPAR_MIO_CONSTRUCT_SOL, gmsk.ON))

	/* solve */
	r, _ = task.OptimizeTrm()
	task.SolutionSummary(gmsk.STREAM_LOG)
	checkOk(r)

	/* Read back the solution */
	r, xx = task.GetXx(gmsk.SOL_ITG, xx)
	checkOk(r)

	fmt.Printf("Solution:\n")
	for j = 0; j < numvar; j++ {
		if j == numvar-1 {
			fmt.Printf("%f\n", xx[j])
		} else {
			fmt.Printf("%f ", xx[j])
		}
	}

	r, constr := task.GetIntInf(gmsk.IINF_MIO_CONSTRUCT_SOLUTION)
	checkOk(r)
	r, constr_obj := task.GetDouInf(gmsk.DINF_MIO_CONSTRUCT_SOLUTION_OBJ)
	checkOk(r)
	fmt.Printf("Construct solution utilization: %d\nConstruct solution objective: %.3f\n", constr, constr_obj)
	// Output:
	// Solution:
	// 0.000000 2.000000 0.000000 0.500000
	// Construct solution utilization: 1
	// Construct solution objective: 19.500
}
