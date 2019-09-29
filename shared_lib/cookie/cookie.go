package cookie

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"net/http"
	"time"
)

/**
 * set cookie for request and response
 */
func SetCookieForRequestAndResponse(res http.ResponseWriter, req *http.Request, cookie *http.Cookie) {

	http.SetCookie(res, cookie)
	req.AddCookie(cookie)
}

/**
 * create a new session cookie for domain
 */
func CreateNewCookieForDomain(cookieName string, jwt string, domain string) *http.Cookie {

	expireTomorrow := time.Now().AddDate(0, 0, 1)

	newCookie := http.Cookie{
		Name:       cookieName,
		Value:      jwt,
		Path:       "/",
		Domain:     domain, // TODO userCredentials.Domain
		Expires:    expireTomorrow,
		RawExpires: expireTomorrow.Format(time.UnixDate),
		MaxAge:     86400, // a negative value means that the cookie is not stored persistently and will be deleted when the Web browser exits
		Secure:     false, // should be true for HTTPS only
		HttpOnly:   true,  // can not be accessed from JS
		Raw:        cookieName + "=" + jwt,
		Unparsed:   []string{cookieName + "=" + jwt},
	}

	return &newCookie
}

/**
 * delete all cookies in request
 */
func DeleteAllCookiesInRequest(domain string, res http.ResponseWriter, req *http.Request) {

	cookies := req.Cookies()

	for _, cookie := range cookies {

		DeleteCookieByNameAndDomain(cookie.Name, domain, res, req)
	}
}

/**
 * delete cookie by given name in request and response
 */
func DeleteCookieByNameAndDomain(cookieName, domain string, res http.ResponseWriter, req *http.Request) {

	if !Exists(cookieName, req) {
		return
	}

	expireOneWeekAgo := time.Now().Add(-7 * 24 * time.Hour)

	cookie, _ := req.Cookie(cookieName)
	cookie.Expires = expireOneWeekAgo
	cookie.MaxAge = -1
	cookie.Domain = domain
	cookie.Path = "/"

	fmt.Printf("deleted cookie: %v \n", cookie.Name)
	SetCookieForRequestAndResponse(res, req, cookie)
}

/**
 * get existing json web token for given authentication method in request param (default to basic Auth)
 * TODO check if cookie exists for given deviceID
 */
func GetExistingJWTByUserCredentials(req *http.Request, domain string,
	userCredentials jwt.UserCredentialForJWT) (jwt string, cookieExists bool) {

	cookieName := fmt.Sprintf("%v_%v", domain, userCredentials.UserAuthentication)

	cookieExists = Exists(cookieName, req)

	if cookieExists {

		jwt = GetValueByName(cookieName, req)
	}

	return
}

/**
 * try to get a cookies value by given name
 */
func GetValueByName(cookieName string, req *http.Request) (value string) {

	if !Exists(cookieName, req) {
		return
	}

	cookie, _ := req.Cookie(cookieName)

	return cookie.Value
}

/**
 * check if a cookie exists with given name
 */
func Exists(name string, req *http.Request) (exists bool) {

	cookie, err := req.Cookie(name)

	if err != nil || cookie == nil {

		fmt.Printf("Cookie with the name %v doesnt exist. - err: %v \n", name, err)
		return false
	}

	return true
}
