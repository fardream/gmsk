// Automatically generated by github.com/fardream/gen-gmsk
// function deinitions

package gmsk

// #include <stdlib.h> // for C.free
// #include <mosek.h>
import "C"

import (
	"unsafe"

	"github.com/fardream/gmsk/res"
)

// GetAccName is wrapping [MSK_getaccname],
// Obtains the name of an affine conic constraint.
//
// Arguments:
//
//   - `accidx` Index of an affine conic constraint.
//
// Returns:
//
//   - `name` Returns the required name.
//
// [MSK_getaccname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getaccname
func (task *Task) GetAccName(
	accidx int64,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getaccname(
			task.task,
			C.MSKint64t(accidx),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetAccNameLen is wrapping [MSK_getaccnamelen],
// Obtains the length of the name of an affine conic constraint.
//
// Arguments:
//
//   - `accidx` Index of an affine conic constraint.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// [MSK_getaccnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getaccnamelen
func (task *Task) GetAccNameLen(
	accidx int64,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getaccnamelen(
			task.task,
			C.MSKint64t(accidx),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetBarvarName is wrapping [MSK_getbarvarname],
// Obtains the name of a semidefinite variable.
//
// Arguments:
//
//   - `i` Index of the variable.
//
// Returns:
//
//   - `name` The requested name is copied to this buffer.
//
// [MSK_getbarvarname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getbarvarname
func (task *Task) GetBarvarName(
	i int32,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getbarvarname(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetBarvarNameLen is wrapping [MSK_getbarvarnamelen],
// Obtains the length of the name of a semidefinite variable.
//
// Arguments:
//
//   - `i` Index of the variable.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// [MSK_getbarvarnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getbarvarnamelen
func (task *Task) GetBarvarNameLen(
	i int32,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getbarvarnamelen(
			task.task,
			C.MSKint32t(i),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetConeName is wrapping [MSK_getconename],
// Obtains the name of a cone.
//
// Arguments:
//
//   - `i` Index of the cone.
//
// Returns:
//
//   - `name` The required name.
//
// Deprecated: [MSK_getconename]/GetConeName is deprecated by mosek and will be removed in a future release.
//
// [MSK_getconename]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getconename
func (task *Task) GetConeName(
	i int32,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getconename(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetConeNameLen is wrapping [MSK_getconenamelen],
// Obtains the length of the name of a cone.
//
// Arguments:
//
//   - `i` Index of the cone.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// Deprecated: [MSK_getconenamelen]/GetConeNameLen is deprecated by mosek and will be removed in a future release.
//
// [MSK_getconenamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getconenamelen
func (task *Task) GetConeNameLen(
	i int32,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getconenamelen(
			task.task,
			C.MSKint32t(i),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetConName is wrapping [MSK_getconname],
// Obtains the name of a constraint.
//
// Arguments:
//
//   - `i` Index of the constraint.
//
// Returns:
//
//   - `name` The required name.
//
// [MSK_getconname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getconname
func (task *Task) GetConName(
	i int32,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getconname(
			task.task,
			C.MSKint32t(i),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetConNameLen is wrapping [MSK_getconnamelen],
// Obtains the length of the name of a constraint.
//
// Arguments:
//
//   - `i` Index of the constraint.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// [MSK_getconnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getconnamelen
func (task *Task) GetConNameLen(
	i int32,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getconnamelen(
			task.task,
			C.MSKint32t(i),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetDjcName is wrapping [MSK_getdjcname],
// Obtains the name of a disjunctive constraint.
//
// Arguments:
//
//   - `djcidx` Index of a disjunctive constraint.
//
// Returns:
//
//   - `name` Returns the required name.
//
// [MSK_getdjcname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdjcname
func (task *Task) GetDjcName(
	djcidx int64,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getdjcname(
			task.task,
			C.MSKint64t(djcidx),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetDjcNameLen is wrapping [MSK_getdjcnamelen],
// Obtains the length of the name of a disjunctive constraint.
//
// Arguments:
//
//   - `djcidx` Index of a disjunctive constraint.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// [MSK_getdjcnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdjcnamelen
func (task *Task) GetDjcNameLen(
	djcidx int64,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getdjcnamelen(
			task.task,
			C.MSKint64t(djcidx),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetDomainName is wrapping [MSK_getdomainname],
// Obtains the name of a domain.
//
// Arguments:
//
//   - `domidx` Index of a domain.
//
// Returns:
//
//   - `name` Returns the required name.
//
// [MSK_getdomainname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdomainname
func (task *Task) GetDomainName(
	domidx int64,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getdomainname(
			task.task,
			C.MSKint64t(domidx),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetDomainNameLen is wrapping [MSK_getdomainnamelen],
// Obtains the length of the name of a domain.
//
// Arguments:
//
//   - `domidx` Index of a domain.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// [MSK_getdomainnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getdomainnamelen
func (task *Task) GetDomainNameLen(
	domidx int64,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getdomainnamelen(
			task.task,
			C.MSKint64t(domidx),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetInfName is wrapping [MSK_getinfname],
// Obtains the name of an information item.
//
// Arguments:
//
//   - `inftype` Type of the information item.
//   - `whichinf` An information item.
//
// Returns:
//
//   - `infname` Name of the information item.
//
// [MSK_getinfname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getinfname
func (task *Task) GetInfName(
	inftype InfType,
	whichinf int32,
) (r res.Code, infname string) {
	c_infname := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_infname))

	r = res.Code(
		C.MSK_getinfname(
			task.task,
			C.MSKinftypee(inftype),
			C.MSKint32t(whichinf),
			c_infname,
		),
	)

	if r.IsOk() {
		infname = C.GoString(c_infname)
	}

	return
}

// GetMaxNameLen is wrapping [MSK_getmaxnamelen],
// Obtains the maximum length (not including terminating zero character) of any objective, constraint, variable, domain or cone name.
//
// Arguments:
//
//   - `maxlen` The maximum length of any name.
//
// [MSK_getmaxnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getmaxnamelen
func (task *Task) GetMaxNameLen() (r res.Code, maxlen int32) {
	r = res.Code(
		C.MSK_getmaxnamelen(
			task.task,
			(*C.MSKint32t)(&maxlen),
		),
	)

	return
}

// GetObjName is wrapping [MSK_getobjname],
// Obtains the name assigned to the objective function.
//
// Returns:
//
//   - `objname` Assigned the objective name.
//
// [MSK_getobjname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getobjname
func (task *Task) GetObjName(
	sizeobjname int32,
) (r res.Code, objname string) {
	c_objname := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_objname))

	r = res.Code(
		C.MSK_getobjname(
			task.task,
			C.MSKint32t(sizeobjname),
			c_objname,
		),
	)

	if r.IsOk() {
		objname = C.GoString(c_objname)
	}

	return
}

// GetObjNameLen is wrapping [MSK_getobjnamelen],
// Obtains the length of the name assigned to the objective function.
//
// Returns:
//
//   - `len` Assigned the length of the objective name.
//
// [MSK_getobjnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getobjnamelen
func (task *Task) GetObjNameLen() (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getobjnamelen(
			task.task,
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetParamName is wrapping [MSK_getparamname],
// Obtains the name of a parameter.
//
// Arguments:
//
//   - `partype` Parameter type.
//   - `param` Which parameter.
//
// Returns:
//
//   - `parname` Parameter name.
//
// [MSK_getparamname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getparamname
func (task *Task) GetParamName(
	partype ParameterType,
	param int32,
) (r res.Code, parname string) {
	c_parname := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_parname))

	r = res.Code(
		C.MSK_getparamname(
			task.task,
			C.MSKparametertypee(partype),
			C.MSKint32t(param),
			c_parname,
		),
	)

	if r.IsOk() {
		parname = C.GoString(c_parname)
	}

	return
}

// GetTaskName is wrapping [MSK_gettaskname],
// Obtains the task name.
//
// Returns:
//
//   - `taskname` Returns the task name.
//
// [MSK_gettaskname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.gettaskname
func (task *Task) GetTaskName(
	sizetaskname int32,
) (r res.Code, taskname string) {
	c_taskname := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_taskname))

	r = res.Code(
		C.MSK_gettaskname(
			task.task,
			C.MSKint32t(sizetaskname),
			c_taskname,
		),
	)

	if r.IsOk() {
		taskname = C.GoString(c_taskname)
	}

	return
}

// GetTaskNameLen is wrapping [MSK_gettasknamelen],
// Obtains the length the task name.
//
// Returns:
//
//   - `len` Returns the length of the task name.
//
// [MSK_gettasknamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.gettasknamelen
func (task *Task) GetTaskNameLen() (r res.Code, len int32) {
	r = res.Code(
		C.MSK_gettasknamelen(
			task.task,
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// GetVarName is wrapping [MSK_getvarname],
// Obtains the name of a variable.
//
// Arguments:
//
//   - `j` Index of a variable.
//
// Returns:
//
//   - `name` Returns the required name.
//
// [MSK_getvarname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getvarname
func (task *Task) GetVarName(
	j int32,
	sizename int32,
) (r res.Code, name string) {
	c_name := (*C.char)(C.calloc(MAX_STR_LEN+1, 1))
	defer C.free(unsafe.Pointer(c_name))

	r = res.Code(
		C.MSK_getvarname(
			task.task,
			C.MSKint32t(j),
			C.MSKint32t(sizename),
			c_name,
		),
	)

	if r.IsOk() {
		name = C.GoString(c_name)
	}

	return
}

// GetVarNameLen is wrapping [MSK_getvarnamelen],
// Obtains the length of the name of a variable.
//
// Arguments:
//
//   - `i` Index of a variable.
//
// Returns:
//
//   - `len` Returns the length of the indicated name.
//
// [MSK_getvarnamelen]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.getvarnamelen
func (task *Task) GetVarNameLen(
	i int32,
) (r res.Code, len int32) {
	r = res.Code(
		C.MSK_getvarnamelen(
			task.task,
			C.MSKint32t(i),
			(*C.MSKint32t)(&len),
		),
	)

	return
}

// IsDouParName is wrapping [MSK_isdouparname],
// Checks a double parameter name.
//
// Arguments:
//
//   - `parname` Parameter name.
//   - `param` Returns the parameter corresponding to the name, if one exists.
//
// [MSK_isdouparname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.isdouparname
func (task *Task) IsDouParName(
	parname string,
	param *DParam,
) res.Code {
	c_parname := C.CString(parname)
	defer C.free(unsafe.Pointer(c_parname))

	return res.Code(
		C.MSK_isdouparname(
			task.task,
			c_parname,
			(*C.MSKdparame)(param),
		),
	)
}

// IsIntParName is wrapping [MSK_isintparname],
// Checks an integer parameter name.
//
// Arguments:
//
//   - `parname` Parameter name.
//   - `param` Returns the parameter corresponding to the name, if one exists.
//
// [MSK_isintparname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.isintparname
func (task *Task) IsIntParName(
	parname string,
	param *IParam,
) res.Code {
	c_parname := C.CString(parname)
	defer C.free(unsafe.Pointer(c_parname))

	return res.Code(
		C.MSK_isintparname(
			task.task,
			c_parname,
			(*C.MSKiparame)(param),
		),
	)
}

// IsStrParName is wrapping [MSK_isstrparname],
// Checks a string parameter name.
//
// Arguments:
//
//   - `parname` Parameter name.
//   - `param` Returns the parameter corresponding to the name, if one exists.
//
// [MSK_isstrparname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.isstrparname
func (task *Task) IsStrParName(
	parname string,
	param *SParam,
) res.Code {
	c_parname := C.CString(parname)
	defer C.free(unsafe.Pointer(c_parname))

	return res.Code(
		C.MSK_isstrparname(
			task.task,
			c_parname,
			(*C.MSKsparame)(param),
		),
	)
}

// PutAccName is wrapping [MSK_putaccname],
// Sets the name of an affine conic constraint.
//
// Arguments:
//
//   - `accidx` Index of the affine conic constraint.
//   - `name` The name of the affine conic constraint.
//
// [MSK_putaccname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putaccname
func (task *Task) PutAccName(
	accidx int64,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putaccname(
			task.task,
			C.MSKint64t(accidx),
			c_name,
		),
	)
}

// PutBarvarName is wrapping [MSK_putbarvarname],
// Sets the name of a semidefinite variable.
//
// Arguments:
//
//   - `j` Index of the variable.
//   - `name` The variable name.
//
// [MSK_putbarvarname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putbarvarname
func (task *Task) PutBarvarName(
	j int32,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putbarvarname(
			task.task,
			C.MSKint32t(j),
			c_name,
		),
	)
}

// PutConeName is wrapping [MSK_putconename],
// Sets the name of a cone.
//
// Arguments:
//
//   - `j` Index of the cone.
//   - `name` The name of the cone.
//
// Deprecated: [MSK_putconename]/PutConeName is deprecated by mosek and will be removed in a future release.
//
// [MSK_putconename]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putconename
func (task *Task) PutConeName(
	j int32,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putconename(
			task.task,
			C.MSKint32t(j),
			c_name,
		),
	)
}

// PutConName is wrapping [MSK_putconname],
// Sets the name of a constraint.
//
// Arguments:
//
//   - `i` Index of the constraint.
//   - `name` The name of the constraint.
//
// [MSK_putconname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putconname
func (task *Task) PutConName(
	i int32,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putconname(
			task.task,
			C.MSKint32t(i),
			c_name,
		),
	)
}

// PutDjcName is wrapping [MSK_putdjcname],
// Sets the name of a disjunctive constraint.
//
// Arguments:
//
//   - `djcidx` Index of the disjunctive constraint.
//   - `name` The name of the disjunctive constraint.
//
// [MSK_putdjcname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putdjcname
func (task *Task) PutDjcName(
	djcidx int64,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putdjcname(
			task.task,
			C.MSKint64t(djcidx),
			c_name,
		),
	)
}

// PutDomainName is wrapping [MSK_putdomainname],
// Sets the name of a domain.
//
// Arguments:
//
//   - `domidx` Index of the domain.
//   - `name` The name of the domain.
//
// [MSK_putdomainname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putdomainname
func (task *Task) PutDomainName(
	domidx int64,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putdomainname(
			task.task,
			C.MSKint64t(domidx),
			c_name,
		),
	)
}

// PutObjName is wrapping [MSK_putobjname],
// Assigns a new name to the objective.
//
// Arguments:
//
//   - `objname` Name of the objective.
//
// [MSK_putobjname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putobjname
func (task *Task) PutObjName(
	objname string,
) res.Code {
	c_objname := C.CString(objname)
	defer C.free(unsafe.Pointer(c_objname))

	return res.Code(
		C.MSK_putobjname(
			task.task,
			c_objname,
		),
	)
}

// PutTaskName is wrapping [MSK_puttaskname],
// Assigns a new name to the task.
//
// Arguments:
//
//   - `taskname` Name assigned to the task.
//
// [MSK_puttaskname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.puttaskname
func (task *Task) PutTaskName(
	taskname string,
) res.Code {
	c_taskname := C.CString(taskname)
	defer C.free(unsafe.Pointer(c_taskname))

	return res.Code(
		C.MSK_puttaskname(
			task.task,
			c_taskname,
		),
	)
}

// PutVarName is wrapping [MSK_putvarname],
// Sets the name of a variable.
//
// Arguments:
//
//   - `j` Index of the variable.
//   - `name` The variable name.
//
// [MSK_putvarname]: https://docs.mosek.com/latest/capi/alphabetic-functionalities.html#mosek.task.putvarname
func (task *Task) PutVarName(
	j int32,
	name string,
) res.Code {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return res.Code(
		C.MSK_putvarname(
			task.task,
			C.MSKint32t(j),
			c_name,
		),
	)
}
