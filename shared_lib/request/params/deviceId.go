package params

import (
	"net/http"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
)

const (
	deviceIdKey = "deviceId"
)

func GetDeviceIdFromRequest (req *http.Request) (deviceId string) {

	deviceId = req.URL.Query().Get(deviceIdKey)

	return
}


func SetDeviceIdToRequest (req *http.Request, deviceId string) {

	request.AddQueryParamToRequest(req, deviceIdKey, deviceId)
}