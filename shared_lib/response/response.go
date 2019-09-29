package response

import (
	"encoding/json"
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"net/http"
)

type Authorization struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Logger struct {
	Success      bool                  `json:"success"`
	Message      string                `json:"message"`
	AuthAttempts []structs.AuthAttempt `json:"authAttempts"`
}

/**
 * send any interface as json response
 */
func SendStructAsJson(res http.ResponseWriter, responseData interface{}) {

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	err := json.NewEncoder(res).Encode(responseData)

	if err != nil {
		fmt.Println(err)
	}
}

/**
 * get an authorization object from response body
 */
func ParseResponseBodyIntoStruct(response *http.Response, responseData interface{}) {

	defer closeBodyOfResponse(response)

	if response.Body == nil {
		fmt.Println("ParseResponseBodyIntoStruct: empty response body")
		return
	}

	err := json.NewDecoder(response.Body).Decode(&responseData)

	if err != nil {
		fmt.Println(err)
	}
}

/**
 * close body of response
 */
func closeBodyOfResponse(response *http.Response) {

	if response != nil {
		if response.Body != nil {

			err := response.Body.Close()

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
