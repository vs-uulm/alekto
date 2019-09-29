package data_variables

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
)

func saveNewVariable(newVariable structs.Variable) {

	var variables = getAllVariables()
	var variablesExist bool
	var err error

	for index, variable := range variables {

		if variable.IsSame(newVariable) {

			fmt.Printf("update var: %+v for: %+v\n", newVariable.Metadata.Name, newVariable.Subject.Name)

			variables[index].Opinion = newVariable.Opinion
			variablesExist = true
			break
		}
	}

	if !variablesExist {
		err = yaml.AppendStructToFile(getFilename(), []structs.Variable{newVariable})
	} else {
		err = yaml.SaveStructToFile(getFilename(), variables)
	}

	if err != nil {
		fmt.Printf("error while saving variables to file: %s", err.Error())
	}
}
