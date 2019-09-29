package twoFactorAuth

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
)

/**
 * TODO Shreya
 */
func Authenticate() (success bool) {

	return false
}

/**
 * TODO Shreya
 */
func GetJWTCredentials() jwt.UserCredentialForJWT {

	userCredentials := jwt.UserCredentialForJWT{
		Username:           "John Doe",
		UserAuthentication: authentication.TwoFactorAuth.String(),
	}

	return userCredentials
}
