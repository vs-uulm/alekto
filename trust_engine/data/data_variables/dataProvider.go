package data_variables

import (
	"flag"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
)


func getVariablesBySubject (givenSubject structs.Subject) (vars []structs.Variable) {

	for _, variable := range getAllVariables() {

		if variable.Subject.Equals(givenSubject) {
			vars = append(vars, variable)
		}
	}

	return
}


func getBaselineVariablesByKind (kind fields.Kind) (baselineVariables []structs.Variable) {

	for _, variable := range getAllBaselineVariables() {

		if variable.Subject.Kind == kind {
			baselineVariables = append(baselineVariables, variable)
		}
	}

	return
}


func getAllVariables () (variables []structs.Variable) {

	yaml.LoadStructFromFile(getFilename(), &variables)
	return
}


func getAllBaselineVariables () (variables []structs.Variable) {

	yaml.LoadStructFromFile(getDefaultFilename(), &variables)
	return
}


func getFilename () string {

	filename := "./data/data_variables/variables.yaml"

	if flag.Lookup("test.v") == nil {
		return filename
	} else {
		return "." + filename
	}
}


func getDefaultFilename () string {

	filename := "./data/data_variables/default.yaml"

	if flag.Lookup("test.v") == nil {
		return filename
	} else {
		return "." + filename
	}
}