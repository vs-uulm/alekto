package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/ma-zero-trust-prototype/logger/certs"
	"github.com/ma-zero-trust-prototype/logger/env"
	"net/http"
)

/**
 * get mutual TLS authentication server
 */
func GetMTLSConfiguredServer() *http.Server {

	fmt.Println("---- NEW REQUEST ----")

	caCertPool := certs.GetRootCaPools()
	ownCertificate := certs.GetOwnX509KeyPair()

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{ownCertificate},
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
		ClientCAs:          caCertPool}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		Certificates:       []tls.Certificate{ownCertificate},
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
		ClientCAs:          caCertPool}

	return &http.Server{
		Addr:      env.GetListenAddressPort(),
		TLSConfig: tlsConfig}
}

/**
 * Only Valid Certificate of SSO and Trust engine will be accepted
 */
func AuthorizedConnectingEntity (req *http.Request) (bool, string) {

	if len(req.TLS.PeerCertificates) <= 0 {
		return false, fmt.Sprintf("Missing Certificate. You tried to connect with an unauthorized device.")
	}

	certificate := req.TLS.PeerCertificates[0]
	opts := x509.VerifyOptions{Roots: certs.GetRootCaPools()}

	if _, err := certificate.Verify(opts); err != nil {
		return false, fmt.Sprintf("Invalid Certificate. Failed to verify certificate of %s: %s",
			certificate.Subject.CommonName, err.Error())
	}

	if certificate.Subject.CommonName != "trust.uni-ulm.de" &&
		certificate.Subject.CommonName != "sso.uni-ulm.de" {

		return false, "Invalid Certificate. Only Connections of SSO and Trust Engine will be accepted."
	}

	return true, "Valid Certificate"
}
