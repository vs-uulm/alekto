package database

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
)

/**
 * check if given session id exists in moodle database for given username
 */
func CheckIfSessionIsAuthenticatedByUserCredentials(sessionId string, userCredentials jwt.UserCredentialForJWT) bool {

	if CheckConnection() == false || sessionId == "" {
		return false
	}

	sessionExists := getValidSessionBySessionIdAndUsername(sessionId, userCredentials.Username)

	fmt.Printf("sessionExists: %v \n", sessionExists)

	return sessionExists
}
