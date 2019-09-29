package basicAuth

import (
	authMethod "github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/time"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication/ldap"
	"net/http"
)

func Authenticate(loginPayload request.BasicLoginPayload, userCredentials jwt.UserCredentialForJWT) (
	success bool, message string) {

	success = ldap.Authenticate(loginPayload)

	if !success {
		message = "ldap authentication failed"
	}

	authentication.LogAuthenticationAttempt(getAuthenticationAttempt(success, loginPayload, userCredentials))

	return
}

func getAuthenticationAttempt(success bool, loginPayload request.BasicLoginPayload,
	userCredentials jwt.UserCredentialForJWT) structs.AuthAttempt {

	return structs.AuthAttempt{
		Subject: structs.Subject{
			Kind: fields.User,
			Name: loginPayload.Username},
		UserAuth:   authMethod.BasicAuth,
		DeviceAuth: userCredentials.DeviceAuthentication,
		IPAddress:  userCredentials.IPAdress,
		DeviceId:   userCredentials.DeviceId,
		Timestamp:  time.NowTimestamp(),
		Failed:     !success}
}

/**
 * write login payload into user credentials
 */
func ParsePayloadDataIntoUserCredentials(loginPayload request.BasicLoginPayload,
	userCredentials *jwt.UserCredentialForJWT) {

	userCredentials.Username = loginPayload.Username
	userCredentials.UserAuthentication = authMethod.BasicAuth.String()
}

/**
 * get login request payload
 */
func GetLoginRequestPayload(req *http.Request) (loginRequestPayload request.BasicLoginPayload) {

	request.ParseRequestFormIntoStruct(req, &loginRequestPayload)

	return loginRequestPayload
}
