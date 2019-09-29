package session

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/cookie"
	enumAuth "github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"github.com/ma-zero-trust-prototype/shared_lib/request/params"
	"github.com/ma-zero-trust-prototype/sso_auth/env"
	"github.com/ma-zero-trust-prototype/sso_auth/keys"
	"net/http"
)

/*
 * create cookie and redirect to proxy
 * TODO domain = userCredentials.Domain
 * TODO Redirect to specific PROXY
 */
func CreateJWTokenForSpecificDomainAndRedirectToProxy (res http.ResponseWriter, req *http.Request,
	userCredentials jwt.UserCredentialForJWT) {

	authenticationMethod := userCredentials.UserAuthentication
	newJWT := jwt.CreateAndSignNewJWT(userCredentials, keys.GetSignKey())
	params := fmt.Sprintf("?%s=%s&%s=%s", params.JwtKey, newJWT, params.AuthenticationMethodKey, authenticationMethod)

	http.Redirect(res, req, "https://" + env.GetProxyAddress() + "/setSessionCookie" + params, http.StatusSeeOther)
}

/*
 * create cookie and redirect to proxy
 */
func CreateMasterCookieForAuthServer (res http.ResponseWriter, authenticationMethod enumAuth.Method,
	userCredentials jwt.UserCredentialForJWT) {

	newJWT := jwt.CreateAndSignNewJWT(userCredentials, keys.GetSignKey())
	cookieName := fmt.Sprintf("%v_%v", env.GetSsoAuthSessionCookieName(), authenticationMethod)
	ssoMasterCookie := cookie.CreateNewCookieForDomain(cookieName, newJWT, env.GetListenAddressURL())

	http.SetCookie(res, ssoMasterCookie)
}
