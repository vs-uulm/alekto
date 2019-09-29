package networkAgentRequirements

import (
	"github.com/ma-zero-trust-prototype/policy_engine/policies/validator"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
)

func Validate(agentPolicy policy.NetworkAgent, agent structs.NetworkAgent) (success bool, message string) {

	success, message = validator.ValidateTrustScore(fields.NetworkAgentTrustScore, agentPolicy.TrustScore, agent.TrustScore)

	return
}
