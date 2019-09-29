package certs

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"github.com/dgrijalva/jwt-go"
	"github.com/ma-zero-trust-prototype/sso_auth/env"
	"io/ioutil"
	"log"
)

const (
	ssoAuthPubKeyPath = "certs/sso_auth.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

func GetOwnX509KeyPair () (certificate tls.Certificate) {

	certificate, err := tls.LoadX509KeyPair(env.GetCertificate(), env.GetPrivateKey())

	if err != nil {
		log.Fatalf("failed to load proxy cert: %s", err)
	}

	return
}


func GetRootCaPools () (caCertPool *x509.CertPool) {

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


func getVerifyBytes() (verifyBytes []byte) {

	verifyBytes, err := ioutil.ReadFile(ssoAuthPubKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}


func GetVerifyKey() (verifyKey *rsa.PublicKey) {

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(getVerifyBytes())
	if err != nil {
		log.Fatal(err)
	}
	return
}
