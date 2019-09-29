package main

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/cookie"
	enumAuth "github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"github.com/ma-zero-trust-prototype/shared_lib/request/params"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication/basicAuth"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication/deviceAuth"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication/twoFactorAuth"
	"github.com/ma-zero-trust-prototype/sso_auth/env"
	"github.com/ma-zero-trust-prototype/sso_auth/keys"
	"github.com/ma-zero-trust-prototype/sso_auth/server"
	"github.com/ma-zero-trust-prototype/sso_auth/session"
	"github.com/ma-zero-trust-prototype/sso_auth/templateHandler"
	"net/http"
)

func main() {

	env.LogSetup()

	ssoServer := server.GetMTLSConfiguredServer()

	http.HandleFunc("/", authenticateUserByExistingSessionCookie)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/loginBasicAuth", handleBasicLoginRequest)
	http.HandleFunc("/loginTwoFactor", handleTwoFactorLoginRequest)
	http.HandleFunc("/authorizationFailed", handleAuthorizationFailure)

	err := ssoServer.ListenAndServeTLS(env.GetCertificate(), env.GetPrivateKey())

	if err != nil {
		fmt.Println(err)
	}
}

/**
 * handle LOGOUT
 * -> delete all session cookies for this user
 * -> TODO Save Logout in Session Manager
 */
func handleLogout(res http.ResponseWriter, req *http.Request) {

	cookie.DeleteAllCookiesInRequest(env.GetListenAddressURL(), res, req)

	templateHandler.ShowLogoutPage(res)
}

/**
 * handle LOGIN redirection
 */
func handleLogin(res http.ResponseWriter, req *http.Request) {

	authenticationMethod := params.GetAuthenticationMethodFromRequest(req)

	templateHandler.ShowLoginPage(res, authenticationMethod, "New Login Required")
}

/**
 * handle AUTHENTICATION
 */
func authenticateUserByExistingSessionCookie(res http.ResponseWriter, req *http.Request) {
	fmt.Println("---- NEW AUTH REQUEST ----")

	authenticationMethod := params.GetAuthenticationMethodFromRequest(req)
	userCredentials := jwt.UserCredentialForJWT{
		DeviceId:           params.GetDeviceIdFromRequest(req),
		UserAuthentication: authenticationMethod.String()}

	authToken, cookieExists := cookie.GetExistingJWTByUserCredentials(req, env.GetSsoAuthSessionCookieName(), userCredentials)

	if !cookieExists {
		templateHandler.ShowLoginPage(res, authenticationMethod, "New Login required.")
		return
	}

	authenticated := jwt.VerifyJWTAndParseClaimsToStruct(authToken, keys.GetVerifyKey(), &userCredentials)

	if authenticated {
		session.CreateJWTokenForSpecificDomainAndRedirectToProxy(res, req, userCredentials)
	} else {
		templateHandler.ShowLoginPage(res, authenticationMethod, "New Login required.")
	}
}

/**
 * handle login request
 */
func handleBasicLoginRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println("---- NEW BASIC LOGIN REQUEST ----")

	userCredentials := jwt.UserCredentialForJWT{Domain: "moodle.uni-ulm.de"}
	loginPayload := basicAuth.GetLoginRequestPayload(req)
	_, _ = deviceAuth.Authenticate(req, &userCredentials) // disabled mandatory device authentication

	authenticated, message := basicAuth.Authenticate(loginPayload, userCredentials)

	if authenticated {
		// create new jwt and set a cookie for specific domain
		basicAuth.ParsePayloadDataIntoUserCredentials(loginPayload, &userCredentials)
		session.CreateMasterCookieForAuthServer(res, enumAuth.BasicAuth, userCredentials)
		session.CreateJWTokenForSpecificDomainAndRedirectToProxy(res, req, userCredentials)

	} else {
		templateHandler.ShowLoginPage(res, enumAuth.BasicAuth, message)
	}
}

/**
 * handle login request
 */
func handleAuthorizationFailure(res http.ResponseWriter, req *http.Request) {
	fmt.Println("---- AUTHORIZATION FAILURE ----")

	message := req.URL.Query().Get("message")

	templateHandler.ShowAuthorizationFailurePage(res, message)
}

/**
 * TODO Shreya handle two factor login request
 * Hint: You can access this page directly through localhost:8085/loginTwoFactor
 */
func handleTwoFactorLoginRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println("---- NEW TWO FACTOR AUTH REQUEST ----")

	authenticated := twoFactorAuth.Authenticate()

	if authenticated {
		// create new jwt and set a cookie for specific domain
		userCredentials := twoFactorAuth.GetJWTCredentials()
		session.CreateMasterCookieForAuthServer(res, enumAuth.TwoFactorAuth, userCredentials)
		session.CreateJWTokenForSpecificDomainAndRedirectToProxy(res, req, userCredentials)

	} else {
		// show login page with error message
		templateHandler.ShowLoginPage(res, enumAuth.TwoFactorAuth, "Login Failed")
	}
}
