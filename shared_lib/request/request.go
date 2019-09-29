package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"io/ioutil"
	"log"
	"net/http"
)

type BasicLoginPayload struct {
	Username             string `schema:"username"`
	Password             string `schema:"password"`
	DeviceId             string `schema:"deviceId"`
	AuthenticationMethod string `schema:"authenticationMethod"`
}

type TwoFactorLoginPayload struct {
	Username             string `schema:"username"`
	Password             string `schema:"password"`
	DeviceId             string `schema:"deviceId"`
	AuthenticationMethod string `schema:"authenticationMethod"`
}

type AuthenticationPayload struct {
	DeviceId           string `schema:"deviceId"`
	UserAuthentication string `schema:"userAuthentication"`
}

type AuthorizationPayload struct {
	Username             string
	DeviceId             string
	UserAuthentication   string
	DeviceAuthentication string
	IPAddress            string
	Scope                string
	RequestedPath        string
	GivenParams          string
}

func AddQueryParamToRequest(req *http.Request, key string, value string) {

	query := req.URL.Query()

	query.Add(key, value)

	req.URL.RawQuery = query.Encode()
}

func ParseRequestBodyIntoStruct(req *http.Request, requestData interface{}) {

	jsn, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal("Error reading the body", err)
	}

	err = json.Unmarshal(jsn, &requestData)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}
}

func ParseRequestFormIntoStruct(req *http.Request, requestFormData interface{}) {

	decoder := schema.NewDecoder()
	err := req.ParseForm()

	if err != nil {
		log.Fatal("Error parsing the request form", err)
	}

	err = decoder.Decode(requestFormData, req.Form)

	if err != nil {
		log.Fatal("Error decoding the request form", err)
	}
}

func GetBodyFromStruct(data interface{}) *bytes.Buffer {

	json2, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	body := bytes.NewBuffer(json2)

	return body
}
