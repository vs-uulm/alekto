package params

import (
	"net/http"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
)

const (
	messageKey = "message"
)

func GetMessageFromRequest (req *http.Request) (message string) {

	message = req.URL.Query().Get(messageKey)

	return
}

func SetMessageToRequest (req *http.Request, message string) {

	request.AddQueryParamToRequest(req, messageKey, message)
}