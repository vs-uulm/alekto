package userRequirements

import (
	"github.com/ma-zero-trust-prototype/policy_engine/policies/validator"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
)

func Validate(userPolicy policy.User, agent structs.User) (success bool, message string) {

	success, message = validator.ValidateStringArrayBySingleValue(fields.UserRole, userPolicy.Role, agent.Role)
	if !success {
		return
	}

	success, message = validator.ValidateStringArrayBySingleValue(fields.UserAuthentication, userPolicy.Authentication, agent.Authentication)
	if !success {
		return
	}

	success, message = validator.ValidateTrustScore(fields.UserTrustScore, userPolicy.TrustScore, agent.TrustScore)

	return
}
