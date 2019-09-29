package policies

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/policy_engine/policies/deviceRequirements"
	"github.com/ma-zero-trust-prototype/policy_engine/policies/networkAgentRequirements"
	"github.com/ma-zero-trust-prototype/policy_engine/policies/userRequirements"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/device"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/scope"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/user"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/userCategory"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
	"regexp"
)

var policies []policy.ClientPolicy

/**
 * check policies against given request and network agent
 */
func CheckPoliciesAgainstRequestAndNetworkAgent(agent structs.NetworkAgent,
	payload request.AuthorizationPayload) (response response.Authorization) {

	relevantPolicies := getRelevantPoliciesForTheAgentsRequest(agent, payload)
	var policyEvaluations []policy.PolicyEvaluation

	for _, relevantPolicy := range relevantPolicies {

		evaluation := evaluateUserDeviceAndNetworkRequirements(relevantPolicy, agent)

		policyEvaluations = append(policyEvaluations, evaluation)
	}

	return getAuthorizationResponse(policyEvaluations)
}

/**
 * evaluate user, device and network agent if requirements are given
 */
func evaluateUserDeviceAndNetworkRequirements(clientPolicy policy.ClientPolicy,
	agent structs.NetworkAgent) policy.PolicyEvaluation {

	evaluation := getNewEvaluation(clientPolicy)

	if clientPolicy.Requires.User != nil {
		success, message := userRequirements.Validate(*clientPolicy.Requires.User, agent.User)
		updateEvaluation(success, message, &evaluation)
	}

	if clientPolicy.Requires.Device != nil {
		success, message := deviceRequirements.Validate(*clientPolicy.Requires.Device, agent.Device)
		updateEvaluation(success, message, &evaluation)
	}

	if clientPolicy.Requires.NetworkAgent != nil {
		success, message := networkAgentRequirements.Validate(*clientPolicy.Requires.NetworkAgent, agent)
		updateEvaluation(success, message, &evaluation)
	}

	return evaluation
}

/**
 * return the relevant policies for given network agent
 */
func getRelevantPoliciesForTheAgentsRequest(agent structs.NetworkAgent,
	payload request.AuthorizationPayload) (relevantPolicies []policy.ClientPolicy) {

	initPolicies()

	for _, clientPolicy := range policies {

		if !(policyMatchesScope(clientPolicy.Metadata.Scope, payload.Scope) &&
			policyMatchesPath(clientPolicy.Metadata.Path, payload.RequestedPath)) {
			continue
		}

		if policySubjectMatchesNetworkAgent(clientPolicy, agent) {
			relevantPolicies = append(relevantPolicies, clientPolicy)
		}
	}

	return
}

/**
 * check if policy subject matches network agent
 */
func policySubjectMatchesNetworkAgent(policy policy.ClientPolicy, agent structs.NetworkAgent) bool {

	for _, subject := range policy.Subjects {

		switch subject.Kind {

		case fields.UserRole:
			if policyMatchesUserRole(subject.Name, agent.User.Role) {
				return true
			}
		case fields.UserCategory:
			if policyMatchesUserCategory(subject.Name, agent.User.Category) {
				return true
			}
		case fields.DeviceType:
			if policyMatchesDeviceType(subject.Name, agent.Device.Type) {
				return true
			}
		}
	}

	return false
}

/**
 * get authorization response
 */
func getAuthorizationResponse(policyEvaluations []policy.PolicyEvaluation) response.Authorization {

	evaluationMap := make(map[string]bool)
	success, message := true, ""

	for _, evaluation := range policyEvaluations {
		evaluationMap[evaluation.Name] = evaluation.Success
	}

	for _, evaluation := range policyEvaluations {

		if evaluation.Success {
			continue
		}

		if !exchangeablePolicyExists(evaluation.Exchangeable, evaluationMap) {
			success, message = false, evaluation.Message
			break
		}
	}

	return response.Authorization{
		Success: success,
		Message: message}
}

/**
 * check if an exchangeable policy exists and if it is valid
 */
func exchangeablePolicyExists(exchangeablePolicies structs.Exchangeable, evaluationMap map[string]bool) bool {

	for _, exchangeable := range exchangeablePolicies {

		if evaluationMap[exchangeable] {
			return true
		}
	}

	return false
}

/**
 * check if policy is relevant for given network agent
 */
func policyMatchesScope(policyScope scope.Service, userScope string) bool {
	return policyScope == scope.Default || policyScope == userScope
}
func policyMatchesPath(policyPath string, requestedPath string) bool {
	path := regexp.MustCompile(fmt.Sprintf("^.*%v.*$", policyPath))
	return policyPath == "" || path.MatchString(requestedPath)
}
func policyMatchesUserRole(role string, userRole string) bool {
	return role == string(user.Default) || role == userRole
}
func policyMatchesUserCategory(category string, value string) bool {
	return category == userCategory.Default || category == value
}
func policyMatchesDeviceType(deviceType string, userDeviceType string) bool {
	return deviceType == string(device.Default) || deviceType == userDeviceType
}
