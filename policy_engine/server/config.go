package server

import (
	"crypto/tls"
	"fmt"
	"github.com/ma-zero-trust-prototype/policy_engine/certs"
	"github.com/ma-zero-trust-prototype/policy_engine/env"
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
