package proxy

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"github.com/ma-zero-trust-prototype/shared_lib/request/params"
	"net/http"
	"net/http/httputil"
)

func ToService(res http.ResponseWriter, req *http.Request) {

	var proxy *httputil.ReverseProxy

	if env.MoodleEnabled() {
		proxy = GetNewProxy(env.GetServerAddress())
	} else {
		proxy = GetNewProxy(env.GetDummyWebAddress())
	}

	proxy.ServeHTTP(res, req)
}

func ToSSOAuthorizationFailedPage(res http.ResponseWriter, req *http.Request, errMessage string) {

	fmt.Println("Authorization Failed: " + errMessage)

	params.SetMessageToRequest(req, errMessage)

	req.URL.Path = "authorizationFailed"
	req.URL.RawPath = "authorizationFailed"

	toSSOAuthenticationService(res, req)
}
