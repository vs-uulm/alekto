package trustScore

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/logger"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

func GetAuthenticationAttempts(subject structs.Subject, lastUpdated int64) (authAttempts []structs.AuthAttempt) {

	body := request.GetBodyFromStruct(subject)
	path := fmt.Sprintf("/getAuthAttempts?since=%v", lastUpdated)
	loggerResponse := logger.SendRequest(body, path)

	if !loggerResponse.Success {
		fmt.Println("Failed to Load Authentication Attempts: " + loggerResponse.Message)
	}

	return loggerResponse.AuthAttempts
}
