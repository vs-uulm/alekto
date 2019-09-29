package main

import (
	"crypto/x509"
	"fmt"
	"github.com/ma-zero-trust-prototype/policy_engine/certs"
	"github.com/ma-zero-trust-prototype/policy_engine/data"
	"github.com/ma-zero-trust-prototype/policy_engine/env"
	"github.com/ma-zero-trust-prototype/policy_engine/policies"
	"github.com/ma-zero-trust-prototype/policy_engine/server"
	"github.com/ma-zero-trust-prototype/policy_engine/trustEngine"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"net/http"
)

func main() {

	env.LogSetup()

	policyServer := server.GetMTLSConfiguredServer()

	http.HandleFunc("/", handleAuthorization)

	err := policyServer.ListenAndServeTLS(env.GetCertificate(), env.GetKey())

	if err != nil {
		fmt.Println(err)
	}
}

/**
 * handle authorization
 */
func handleAuthorization(res http.ResponseWriter, req *http.Request) {

	if success, message := authorizedConnectingEntity(req); !success {

		authorizationResponse := response.Authorization{
			Success: false,
			Message: message}

		response.SendStructAsJson(res, authorizationResponse)
	}

	var authorizationPayload = getAuthorizationPayload(req)
	var agent = structs.CreateAgentFromAuthorizationPayload(authorizationPayload)

	data.GetUserAndDeviceInfoForAgent(&agent)
	trustEngine.CalcTrustScoresForAgent(&agent)

	fmt.Printf("Resulting Agent: %v \n", agent)

	var authorizationResponse = policies.CheckPoliciesAgainstRequestAndNetworkAgent(agent, authorizationPayload)

	response.SendStructAsJson(res, authorizationResponse)
}

/**
 * Only Valid Certificates of listed Access Proxies will be accepted
 */
func authorizedConnectingEntity(req *http.Request) (bool, string) {

	if len(req.TLS.PeerCertificates) <= 0 {
		return false, fmt.Sprintf("Missing Certificate. You tried to connect with an unauthorized device.")
	}

	certificate := req.TLS.PeerCertificates[0]
	opts := x509.VerifyOptions{Roots: certs.GetRootCaPools()}

	if _, err := certificate.Verify(opts); err != nil {
		return false, fmt.Sprintf("Invalid Certificate. Failed to verify certificate of %s: %s",
			certificate.Subject.CommonName, err.Error())
	}

	if certificate.Subject.CommonName != "moodle.uni-ulm.de" {
		return false, "Invalid Certificate. Only certificates of listed Access Proxies will be accepted."
	}

	return true, "Valid Certificate"
}

/**
 * get authorization payload
 */
func getAuthorizationPayload(req *http.Request) (authorizationPayload request.AuthorizationPayload) {

	request.ParseRequestBodyIntoStruct(req, &authorizationPayload)
	return
}
