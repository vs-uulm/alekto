package deviceAuth

import (
	"crypto/x509"
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"github.com/ma-zero-trust-prototype/sso_auth/certs"
	"net/http"
	"strings"
)

/**
 * Authenticate the users device with mutual TLS
 */
func Authenticate(req *http.Request, userCredentials *jwt.UserCredentialForJWT) (success bool, message string) {

	userCredentials.IPAdress = getIPv4Address(req.RemoteAddr)
	userCredentials.Domain = req.Host

	if len(req.TLS.PeerCertificates) <= 0 {
		return unAuthenticatedDevice("no client certificate found", userCredentials)
	}

	caCertPool := certs.GetRootCaPools()
	clientCertificate := req.TLS.PeerCertificates[0]
	opts := x509.VerifyOptions{Roots: caCertPool}

	if _, err := clientCertificate.Verify(opts); err != nil {
		message := fmt.Sprintf("failed to verify client certificate of %s: %s",
			clientCertificate.Subject.CommonName, err.Error())
		return unAuthenticatedDevice(message, userCredentials)
	}

	userCredentials.DeviceId = clientCertificate.Subject.CommonName
	userCredentials.DeviceAuthentication = "mtls"

	return true, ""
}

/**
 * values for an unauthenticated device
 */
func unAuthenticatedDevice(message string, userCredentials *jwt.UserCredentialForJWT) (bool, string) {

	userCredentials.DeviceId = ""
	userCredentials.DeviceAuthentication = ""

	fmt.Println("Unauthenticated Device: " + message)

	return false, message
}

func getIPv4Address(remoteAddr string) string {

	parts := strings.Split(remoteAddr, ":")
	return parts[0]
}
