package validator

import (
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

/**
 * check if the given trust score ist greater or equal than the required score
 */
func validateTrustScore(trustScore structs.TrustScore, value structs.TrustScore) (valid bool) {

	if trustScore <= 0 {
		return true
	}

	valid = value >= trustScore

	return
}
