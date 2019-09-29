package jwt

import (
	"testing"
)

func TestCreateAndVerifyToken(t *testing.T) {

	user := UserCredentialForJWT{"john", "deviceId", "moodle", "basicAuth"}
	token := CreateAndSignNewJWT(user)

	userCredentials, verified := VerifyJWTAndParseClaimsToStruct(token)

	if !verified {
		t.Errorf("Error while creating and verifying Json Web Token \n Token: %v \n User: %v", token, userCredentials)
	}
}
