package gmsk_test

import (
	"fmt"
	"log"
	"os"

	"github.com/fardream/gmsk"
)

func ExampleTask_WriteDataHandle() {
	CheckOk := func(r gmsk.ResCode) {
		if r != gmsk.RES_OK {
			_, sym, desc := gmsk.GetCodedesc(r)
			log.Fatalf("Failed: %s %s", sym, desc)
		}
	}

	task, err := gmsk.MakeTask(nil, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer gmsk.DeleteTask(task)

	CheckOk(task.AppendVars(1))
	CheckOk(task.PutCJ(0, 1.0))
	CheckOk(task.PutVarBound(0, gmsk.BK_RA, 2.0, 3.0))
	CheckOk(task.PutObjSense(gmsk.OBJECTIVE_SENSE_MINIMIZE))

	CheckOk(task.WriteDataHandle(os.Stderr, gmsk.DATA_FORMAT_PTF, gmsk.COMPRESS_NONE))
	fmt.Println("Done")
	// Output: Done
}
