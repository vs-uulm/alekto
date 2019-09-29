package proxy

import (
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/**
 * get new proxy
 */
func GetNewProxy(host string) *httputil.ReverseProxy {

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host:   host,
	})

	return proxy
}

/**
 * redirect
 */
func toSSOAuthenticationService(res http.ResponseWriter, req *http.Request) {

	proxy := GetNewProxy(env.GetSsoAuthAddress())
	proxy.ServeHTTP(res, req)
}
