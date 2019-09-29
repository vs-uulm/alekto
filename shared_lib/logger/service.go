package logger

import (
	"bytes"
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/env"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"net/http"
)

func SendRequest (body *bytes.Buffer, path string) (loggerResponse response.Logger) {

	client := &http.Client{}
	loggerRequest, err := http.NewRequest(http.MethodPost, "https://" + env.GetLoggerAddress() + path, body)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(loggerRequest)
	response.ParseResponseBodyIntoStruct(res, &loggerResponse)

	return loggerResponse
}