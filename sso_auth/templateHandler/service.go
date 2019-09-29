package templateHandler

import (
	"fmt"
	enumAuth "github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication/basicAuth"
	"github.com/ma-zero-trust-prototype/sso_auth/authentication/twoFactorAuth"
	"net/http"
)


func ShowLoginPage (res http.ResponseWriter, authenticationMethod enumAuth.Method, message string) {

	fmt.Println("Show Login Page for " +  authenticationMethod.String())

	switch authenticationMethod {

		case enumAuth.BasicAuth:
			basicAuth.ShowLoginPage(res, message)

		case enumAuth.TwoFactorAuth:
			twoFactorAuth.ShowLoginPage(res)

		default:
			basicAuth.ShowLoginPage(res, message)
	}
}