package main

import (
	"crypto/x509"
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/trust_engine/certs"
	"github.com/ma-zero-trust-prototype/trust_engine/env"
	"github.com/ma-zero-trust-prototype/trust_engine/server"
	"github.com/ma-zero-trust-prototype/trust_engine/trustScore"
	"net/http"
)

func main() {

	env.LogSetup()

	trustServer := server.GetMTLSConfiguredServer()

	http.HandleFunc("/", handleRequestAndInspectPayload)

	err := trustServer.ListenAndServeTLS(env.GetCertificate(), env.GetKey())

	if err != nil {
		fmt.Println(err)
	}
}

func handleRequestAndInspectPayload(res http.ResponseWriter, req *http.Request) {

	if success, message := authorizedConnectingEntity(req); !success {

		authorizationResponse := response.Authorization{
			Success: false,
			Message: message}

		response.SendStructAsJson(res, authorizationResponse)
	}

	var agent structs.NetworkAgent
	request.ParseRequestBodyIntoStruct(req, &agent)

	trustScores := trustScore.CalcTrustScoresForAgent(agent)

	response.SendStructAsJson(res, trustScores)
}

/**
 * Only Valid Certificate of Policy engine will be accepted
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

	if certificate.Subject.CommonName != "policy.uni-ulm.de" {
		return false, "Invalid Certificate. Only Connections of Policy Engine will be accepted."
	}

	return true, "Valid Certificate"
}
