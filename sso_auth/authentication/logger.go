package authentication

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/logger"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

func LogAuthenticationAttempt(authAttempt structs.AuthAttempt) {

	body := request.GetBodyFromStruct(authAttempt)
	loggerResponse := logger.SendRequest(body, "/addAuthAttempt")

	if !loggerResponse.Success {
		fmt.Println("Logging Failure: " + loggerResponse.Message)
	}
}
