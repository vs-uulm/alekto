package server

import (
	"crypto/tls"
	"fmt"
	"github.com/ma-zero-trust-prototype/moodle_proxy/certs"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"net/http"
)

/**
 * get mutual TLS authentication server
 */
func GetMTLSConfiguredServer() *http.Server {

	fmt.Println("---- NEW REQUEST ----")

	caCertPool := certs.GetRootCaPools()
	proxyCertificate := certs.GetProxyX509KeyPair()

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{proxyCertificate},
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
		ClientCAs:          caCertPool}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		Certificates:       []tls.Certificate{proxyCertificate},
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
		ClientCAs:          caCertPool}

	return &http.Server{
		Addr:      env.GetListenAddressPortTLS(),
		TLSConfig: tlsConfig}
}
