package data_variables

import (
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/trust_engine/util"
)


func GetVariablesBySubject (givenSubject structs.Subject) (vars []structs.Variable) {

	if util.IsUnIdentifiedDevice(givenSubject) {
		return getBaselineVariablesByKind(givenSubject.Kind)
	}

	vars = getVariablesBySubject(givenSubject)

	if len(vars) <= 0 {
		vars = getBaselineVariablesByKind(givenSubject.Kind)
	}

	return
}


func SaveNewVariablesBySubject (subject structs.Subject, newVariable structs.Variable) {

	newVariable.Subject = subject

	saveNewVariable(newVariable)
}