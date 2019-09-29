package keys

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

const (
	privateKeyPath = "keys/sso_auth.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath     = "keys/sso_auth.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// Get Keys for Signing and Verifying
func getSignBytes() (signBytes []byte) {

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getVerifyBytes() (verifyBytes []byte) {

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetSignKey() (signKey *rsa.PrivateKey) {

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(getSignBytes())
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
