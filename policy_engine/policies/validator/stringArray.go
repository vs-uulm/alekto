package validator

import (
	"github.com/ma-zero-trust-prototype/shared_lib/stringUtil"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
)

/**
 * validate stringArray struct against given value
 * -> check if value is set in array and if not flag is set to true/false
 */
func validateStringArray(requirement policy.StringArray, values []string) (valid bool) {

	if len(requirement.Values) <= 0 {
		return true
	}

	listed, allListed := checkIfRequiredValuesAreListed(requirement.Values, values)
	requiredListed := false

	if andOperator(requirement.Operator) {
		requiredListed = allListed
	} else if orOperator(requirement.Operator) {
		requiredListed = listed
	}

	return requirement.Not && !requiredListed || !requirement.Not && requiredListed
}

func andOperator(op string) bool {
	return op == "and"
}

func orOperator(op string) bool {
	return op == "" || op == "or"
}

func checkIfRequiredValuesAreListed(requiredValues []string, values []string) (valueListed bool, allValuesListed bool) {

	for _, value := range values {

		valueListed = valueListed || stringUtil.InSlice(value, requiredValues)

		if !valueListed {
			allValuesListed = false
		}
	}

	return
}
