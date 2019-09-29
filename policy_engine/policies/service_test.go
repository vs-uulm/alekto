package policies

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/policy_engine/data"
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"testing"
)

func TestParseYamlIntoPolicy(t *testing.T) {

	initPolicies()

	for _, policy := range policies {

		fmt.Printf("%+v\n", policy)
	}
}

/**
 * test policy evaluation with given authorization payload
 */
func TestCheckPoliciesAgainstRequestAndNetworkAgent(t *testing.T) {

	var authorizationPayload = request.AuthorizationPayload{
		Username:             "bender",
		DeviceId:             "bendersDevice",
		UserAuthentication:   "basicAuth",
		DeviceAuthentication: "",
		Scope:                "moodle",
		RequestedPath:        "/my",
		GivenParams:          "id=123"}
	var agent = structs.CreateAgentFromAuthorizationPayload(authorizationPayload)

	data.GetUserAndDeviceInfoForAgent(&agent)
	agent.User.TrustScore = 0.9
	agent.Device.TrustScore = 0
	agent.TrustScore = 0.9

	fmt.Printf("Resulting Agent: %v \n", agent)

	response := CheckPoliciesAgainstRequestAndNetworkAgent(agent, authorizationPayload)

	fmt.Printf("--- POLICY RESPONSE\n%+v \n", response)
}

/**
 * test policy evaluation with given authorization payload and network agent
 */
func TestCheckPoliciesWithGivenNetworkAgent(t *testing.T) {

	var requestPayload = request.AuthorizationPayload{
		Username:           "alexanderBias",
		DeviceId:           "deviceId123",
		UserAuthentication: "u2f",
		Scope:              "moodle",
		RequestedPath:      "/",
		GivenParams:        ""}

	var agent = structs.NetworkAgent{
		User: structs.User{
			TrustScore:     0.5,
			Id:             requestPayload.Username,
			Role:           "admin",
			Category:       "highprivileged",
			Authentication: requestPayload.UserAuthentication},
		Device: structs.Device{
			Id:             requestPayload.DeviceId,
			TrustScore:     0.9,
			Type:           "mobile",
			Authentication: "ipsec",
			OS:             "Ubuntu",
			Owner:          "uniulm",
			IPAddress:      "134.60.112.69",
			Location:       ip2location.GetLocationByIpAddress("134.60.112.69")},
		TrustScore: 0.8}

	fmt.Printf("Resulting Agent: %+v \n", agent)

	response := CheckPoliciesAgainstRequestAndNetworkAgent(agent, requestPayload)

	fmt.Printf("--- POLICY RESPONSE\nSuccess:%+v\nMessage:%+v\n", response.Success, response.Message)
}
