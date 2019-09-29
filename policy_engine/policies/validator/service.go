package validator

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
)

/**
 * validate stringArray struct against given value
 * -> check if value is set in array and if not flag is set to true/false
 */
func ValidateStringArrayBySingleValue(name fields.Kind, array *policy.StringArray, value string) (valid bool, message string) {

	if array == nil {
		return true, "no requirements"
	}

	valid = validateStringArray(*array, []string{value})

	if valid {
		message = fmt.Sprintf("%v matches required values", name)
	} else if array.Not {
		message = fmt.Sprintf("%v [%v] does not match required values: NOT %v", name, value, array.Values)
	} else if !array.Not {
		message = fmt.Sprintf("%v [%v] does not match required values: %v", name, value, array.Values)
	}

	return
}

/**
 * check if the given trust score ist greater or equal than the required score
 */
func ValidateTrustScore(name fields.Kind, trustScore *structs.TrustScore, value structs.TrustScore) (
	valid bool, message string) {

	if trustScore == nil {
		return true, "no requirements"
	}

	valid = validateTrustScore(*trustScore, value)

	if valid {
		message = fmt.Sprintf("%v matches requirements", name)
	} else {
		message = fmt.Sprintf("%v [%v] is not high enough for required score [%v]", name, value, *trustScore)
	}

	return
}

/**
 * check if the given location matches the requirements
 */
func ValidateLocation(name fields.Kind, location *ip2location.Location, userLocation ip2location.Location) (
	valid bool, message string) {

	if location == nil {
		return true, "no requirements"
	}

	valid = validateLocation(*location, userLocation)

	if valid {
		message = fmt.Sprintf("%v matches requirements", name)
	} else {
		message = fmt.Sprintf("%v [%+v] is not in the required range [%+v]",
			name, userLocation, *location)
	}

	return
}
