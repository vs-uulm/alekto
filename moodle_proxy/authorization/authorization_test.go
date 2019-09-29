package authorization

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestAddAuthorizationPayloadToRequest(t *testing.T) {

	req, err := http.NewRequest("GET", "http://example.com", nil)

	if err != nil {
		t.Errorf("Error while creating new req: %v", err)
	}

	authorizationPayload := getExampleAuthorizationPayload()

	addAuthorizationPayloadToRequest(authorizationPayload, req)

	payload := reflect.ValueOf(authorizationPayload)

	for i := 0; i < payload.NumField(); i++ {

		value := payload.Field(i).String()
		key := payload.Type().Field(i).Name

		if req.URL.Query().Get(key) != value {
			t.Errorf("Error while parsing values into req - key: %v, value: %v", key, value)
		}
	}
}

func getExampleAuthorizationPayload() (authorizationPayload request.AuthorizationPayload) {

	values := url.Values{}
	values.Add("test", "a")

	return request.AuthorizationPayload{
		Username:           "testUser",
		UserAuthentication: enum.BasicAuth.String(),
		RequestedPath:      "admin/test",
		GivenParams:        values.Encode()}
}
