package params

import (
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"net/http"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
)

const (
	AuthenticationMethodKey ="authenticationMethod"
)


func GetAuthenticationMethodFromRequest (req *http.Request) (authenticationMethod authentication.Method) {

	var rawValue = req.URL.Query().Get(AuthenticationMethodKey)

	authenticationMethod = authentication.ParseString(rawValue)

	return
}


func SetAuthenticationMethodToRequest (req *http.Request, authenticationMethod authentication.Method) {

	request.AddQueryParamToRequest(req, AuthenticationMethodKey, authenticationMethod.String())
}