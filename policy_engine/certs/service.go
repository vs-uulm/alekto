package certs

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/ma-zero-trust-prototype/policy_engine/env"
	"io/ioutil"
	"log"
)

func GetOwnX509KeyPair() (certificate tls.Certificate) {

	certificate, err := tls.LoadX509KeyPair(env.GetCertificate(), env.GetKey())

	if err != nil {
		log.Fatalf("failed to load OwnX509KeyPair: %s", err)
	}

	return
}

func GetRootCaPools() (caCertPool *x509.CertPool) {

	caCert, err := ioutil.ReadFile(env.GetRootCertificate())
	if err != nil {
		log.Fatalf("failed to load root cert: %s", err)
	}

	caIntermediateCert, err := ioutil.ReadFile(env.GetRootIntermediateCertificate())
	if err != nil {
		log.Fatalf("failed to load intermediate cert: %s", err)
	}

	caCertPool = x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	caCertPool.AppendCertsFromPEM(caIntermediateCert)

	return
}