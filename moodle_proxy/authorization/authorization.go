package authorization

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"net/http"
	"reflect"
)

/**
 * send authorization request to the policy engine
 */
func AuthorizeUser(req *http.Request, userCredentials jwt.UserCredentialForJWT) (success bool, errMessage string) {

	authorizationPayload := getAuthorizationPayload(req, userCredentials)
	body := request.GetBodyFromStruct(authorizationPayload)
	authorizationRequest, err := http.NewRequest(http.MethodPost, "https://"+env.GetPolicyEngineAddress(), body)

	if err != nil {
		fmt.Println(err)
	}

	authorizationResponse := sendRequestToPolicyEngine(authorizationRequest)

	return authorizationResponse.Success, authorizationResponse.Message
}

/**
 * send request to policy engine and parse body into authorization response struct
 */
func sendRequestToPolicyEngine(authorizationRequest *http.Request) (authorizationResponse response.Authorization) {

	client := &http.Client{}
	res, err := client.Do(authorizationRequest)

	if err != nil {
		fmt.Println(err)
	}

	response.ParseResponseBodyIntoStruct(res, &authorizationResponse)

	return
}

/**
 * add each authorization payload field to Request
 */
func addAuthorizationPayloadToRequest(authorizationPayload request.AuthorizationPayload, req *http.Request) {

	payload := reflect.ValueOf(authorizationPayload)

	for i := 0; i < payload.NumField(); i++ {

		value := payload.Field(i).String()
		key := payload.Type().Field(i).Name

		request.AddQueryParamToRequest(req, key, value)
	}
}

/**
 * get authorization payload by user credentials and request info
 */
func getAuthorizationPayload(req *http.Request, userCredentials jwt.UserCredentialForJWT) request.AuthorizationPayload {

	requestUrl := req.URL

	return request.AuthorizationPayload{
		Username:             userCredentials.Username,
		DeviceId:             userCredentials.DeviceId,
		UserAuthentication:   userCredentials.UserAuthentication,
		DeviceAuthentication: userCredentials.DeviceAuthentication,
		IPAddress:			  userCredentials.IPAdress,
		Scope:                "moodle",
		RequestedPath:        requestUrl.Path,
		GivenParams:          requestUrl.RawQuery}
}
