package authentication

import (
	"crypto/x509"
	"fmt"
	"github.com/ma-zero-trust-prototype/moodle_proxy/authentication/moodle"
	"github.com/ma-zero-trust-prototype/moodle_proxy/authentication/moodle/database"
	"github.com/ma-zero-trust-prototype/moodle_proxy/certs"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"github.com/ma-zero-trust-prototype/shared_lib/cookie"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"net/http"
	"strings"
)

/**
 * Authenticate the users device with mutual TLS
 */
func AuthenticateDevice(req *http.Request, userCredentials *jwt.UserCredentialForJWT) (success bool, message string) {

	userCredentials.IPAdress = getIPv4Address(req.RemoteAddr)

	if len(req.TLS.PeerCertificates) <= 0 {
		return unAuthenticatedDevice("no client certificate found", userCredentials)
	}

	caCertPool := certs.GetRootCaPools()
	clientCertificate := req.TLS.PeerCertificates[0]
	opts := x509.VerifyOptions{Roots: caCertPool}

	if _, err := clientCertificate.Verify(opts); err != nil {
		return unAuthenticatedDevice(fmt.Sprintf("failed to verify client certificate of %s: %s",
			clientCertificate.Subject.CommonName, err.Error()), userCredentials)
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

/**
 * try to authenticate user
 */
func AuthenticateUser(res http.ResponseWriter, req *http.Request, userCredentials *jwt.UserCredentialForJWT) (success bool) {

	userJWT, cookieExists := cookie.GetExistingJWTByUserCredentials(req,
		env.GetSsoAuthMoodleSessionCookieName(), *userCredentials)

	if cookieExists {
		success = jwt.VerifyJWTAndParseClaimsToStruct(userJWT, certs.GetVerifyKey(), userCredentials)
	}

	if env.MoodleEnabled() {
		success = moodleUserAuthentication(res, req, success, userJWT, userCredentials)
	}

	return
}

func moodleUserAuthentication(res http.ResponseWriter, req *http.Request, success bool, userJWT string,
	userCredentials *jwt.UserCredentialForJWT) bool {

	if success {
		fmt.Printf("User successfully verified: %+v \n", *userCredentials)
		success = createValidMoodleSessionIfNonAlreadyExists(res, req, userJWT, *userCredentials) // TODO Bug Fix if moodle account is invalid

	} else {
		fmt.Printf("User could not be verified successfully by PROXY: %+v \n", *userCredentials)
		cookie.DeleteCookieByNameAndDomain(env.GetMoodleSessionCookieName(), env.GetProxyDomain(), res, req)
	}

	return success
}

/**
 * check if a valid moodle session exists for given user
 * -> if not create a new session and set the MoodleSession Cookie
 */
func createValidMoodleSessionIfNonAlreadyExists(res http.ResponseWriter, req *http.Request,
	userJWT string, userCredentials jwt.UserCredentialForJWT) bool {

	sessionId := cookie.GetValueByName(env.GetMoodleSessionCookieName(), req)
	validSessionExists := database.CheckIfSessionIsAuthenticatedByUserCredentials(sessionId, userCredentials)

	if !validSessionExists {

		fmt.Printf("No valid Session exists.. Sending new Login Request to Moodle.. ")

		sessionCookie := moodle.SendNewLoginRequestAndReturnSessionCookie(userJWT, userCredentials)
		if sessionCookie != nil {
			validSessionExists = database.CheckIfSessionIsAuthenticatedByUserCredentials(sessionCookie.Value, userCredentials)

			if validSessionExists {
				cookie.SetCookieForRequestAndResponse(res, req, sessionCookie)
			}

		} else {
			validSessionExists = true
		}

		fmt.Printf("-> Success: %v \n", validSessionExists)
	}

	return validSessionExists
}
