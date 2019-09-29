package main

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/logger/data"
	"github.com/ma-zero-trust-prototype/logger/env"
	"github.com/ma-zero-trust-prototype/logger/server"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"github.com/ma-zero-trust-prototype/shared_lib/stringUtil"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"net/http"
)

func main() {

	env.LogSetup()

	loggerServer := server.GetMTLSConfiguredServer()

	http.HandleFunc("/", handleRequestAndInspectPayload)
	http.HandleFunc("/addAuthAttempt", addAuthAttempt)
	http.HandleFunc("/getAuthAttempts", getAuthAttempts)

	err := loggerServer.ListenAndServeTLS(env.GetCertificate(), env.GetKey())

	if err != nil {
		fmt.Println(err)
	}
}

func getAuthAttempts(res http.ResponseWriter, req *http.Request) {

	var subject structs.Subject
	var loggerResponse = authenticateDevice(req)

	if !loggerResponse.Success {
		response.SendStructAsJson(res, loggerResponse)
	}

	request.ParseRequestBodyIntoStruct(req, &subject)
	since := stringUtil.ToInt64(req.URL.Query().Get("since"))

	if authAttempts, err := data.GetAuthAttemptsOfSubjectSince(subject, since); err != nil {
		loggerResponse.Message = "Getting the Entity's Auth Attempts Failed"
	} else {
		loggerResponse.Success = true
		loggerResponse.AuthAttempts = authAttempts
	}

	fmt.Printf("Loaded (success: %v) (Amount: %+v)\n", loggerResponse.Success, len(loggerResponse.AuthAttempts))

	response.SendStructAsJson(res, loggerResponse)
}

func addAuthAttempt(res http.ResponseWriter, req *http.Request) {

	var authAttempt structs.AuthAttempt
	var loggerResponse = authenticateDevice(req)

	if !loggerResponse.Success {
		response.SendStructAsJson(res, loggerResponse)
	}

	request.ParseRequestBodyIntoStruct(req, &authAttempt)

	if err := data.LogAuthAttempt(authAttempt); err != nil {
		loggerResponse.Message = "Logging Auth Attempt Failed"
	} else {
		loggerResponse.Success = true
	}

	fmt.Printf("Logged (success: %v): %+v\n", loggerResponse.Success, authAttempt)

	response.SendStructAsJson(res, loggerResponse)
}

func handleRequestAndInspectPayload(res http.ResponseWriter, req *http.Request) {

	var loggerResponse = authenticateDevice(req)

	if !loggerResponse.Success {
		response.SendStructAsJson(res, loggerResponse)
	}
}

/**
 * authenticate device
 */
func authenticateDevice(req *http.Request) response.Logger {

	var loggerResponse response.Logger

	if success, message := server.AuthorizedConnectingEntity(req); !success {
		loggerResponse.Success = false
		loggerResponse.Message = message

	} else {
		loggerResponse.Success = true
	}

	return loggerResponse
}
