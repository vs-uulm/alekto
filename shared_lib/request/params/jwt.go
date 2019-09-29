package params

import (
	"net/http"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
)

const (
	JwtKey = "jwt"
)

func GetJwtFromRequest (req *http.Request) (token string) {

	token = req.URL.Query().Get(JwtKey)

	return
}


func SetJwtToRequest (req *http.Request, token string) {

	request.AddQueryParamToRequest(req, JwtKey, token)
}