package main

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/moodle_proxy/authentication"
	"github.com/ma-zero-trust-prototype/moodle_proxy/authorization"
	"github.com/ma-zero-trust-prototype/moodle_proxy/certs"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"github.com/ma-zero-trust-prototype/moodle_proxy/proxy"
	"github.com/ma-zero-trust-prototype/moodle_proxy/server"
	"github.com/ma-zero-trust-prototype/shared_lib/cookie"
	enumAuth "github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"github.com/ma-zero-trust-prototype/shared_lib/request/params"
	"net/http"
)

func main() {

	env.LogSetup()

	proxyServer := server.GetMTLSConfiguredServer()

	http.HandleFunc("/", handleRequestAuthenticationAndAuthorization)
	http.HandleFunc("/setSessionCookie", handleSetSessionCookie)
	http.HandleFunc(env.GetLogoutUrl(), handleLogout)

	err := proxyServer.ListenAndServeTLS(env.GetCertificate(), env.GetPrivateKey())

	if err != nil {
		fmt.Println(err)
	}
}

/**
 * handle request AUTHENTICATION and AUTHORIZATION
 */
func handleRequestAuthenticationAndAuthorization(res http.ResponseWriter, req *http.Request) {
	fmt.Println("---- NEW REQUEST ----")

	userCredentials := jwt.UserCredentialForJWT{
		UserAuthentication: enumAuth.BasicAuth.String(),
		Domain:             env.GetProxyDomain()}

	_, _ = authentication.AuthenticateDevice(req, &userCredentials) // disabled mandatory mtls authentication

	authenticatedUser := authentication.AuthenticateUser(res, req, &userCredentials)

	if !authenticatedUser {
		redirectToSSOLoginPage(res, req, userCredentials)

	} else {
		authorized, errMessage := authorization.AuthorizeUser(req, userCredentials)

		if authorized || env.IgnoreFailedAuthorization() {
			proxy.ToService(res, req)

		} else {
			proxy.ToSSOAuthorizationFailedPage(res, req, errMessage)
		}
	}
}

/**
 * handle LOGOUT
 */
func handleLogout(res http.ResponseWriter, req *http.Request) {

	fmt.Println("---- LOGOUT ----")

	cookie.DeleteAllCookiesInRequest(env.GetProxyDomain(), res, req)

	redirectToSSOLogoutPage(res, req)
}

/**
 * set session cookie and redirect to default url
 */
func handleSetSessionCookie(res http.ResponseWriter, req *http.Request) {

	userCredentials := jwt.UserCredentialForJWT{}
	userToken := params.GetJwtFromRequest(req)

	jwt.ParseClaimsIntoStruct(userToken, certs.GetVerifyKey(), &userCredentials)

	cookieName := env.GetSsoAuthMoodleSessionCookieName() + "_" + userCredentials.UserAuthentication
	jwtCookie := cookie.CreateNewCookieForDomain(cookieName, userToken, env.GetProxyDomain())

	http.SetCookie(res, jwtCookie)
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}

/**
 * redirect to sso service login page
 */
func redirectToSSOLoginPage(res http.ResponseWriter, req *http.Request, userCredentials jwt.UserCredentialForJWT) {

	param := fmt.Sprintf("?%s=%s", params.AuthenticationMethodKey, userCredentials.UserAuthentication)

	http.Redirect(res, req, "https://"+env.GetSsoAuthAddress()+"/login"+param, http.StatusTemporaryRedirect)
}

/**
 * redirect to sso service logout page
 */
func redirectToSSOLogoutPage(res http.ResponseWriter, req *http.Request) {

	http.Redirect(res, req, "https://"+env.GetSsoAuthAddress()+"/logout", http.StatusTemporaryRedirect)
}
