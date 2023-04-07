package gmsk

import "github.com/fardream/gmsk/res"

// ResCode is return code from mosek.
// This is a reexport to prevent polluting the namespace of gmsk
type ResCode = res.Code

const (
	RES_OK        res.Code = res.OK // RES_OK indicates success return code.
	RES_ERR_SPACE res.Code = res.ERR_SPACE
)
