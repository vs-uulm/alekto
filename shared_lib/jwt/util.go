package jwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func getClaimsByUserCredentials(userCredentials UserCredentialForJWT) jwt.MapClaims {

	expires := expiresIn(24 * time.Hour)

	claims := jwt.MapClaims{
		"username":             userCredentials.Username,
		"deviceId":             userCredentials.DeviceId,
		"userAuthentication":   userCredentials.UserAuthentication,
		"deviceAuthentication": userCredentials.DeviceAuthentication,
		"domain":               userCredentials.Domain,
		"exp":                  expires,
	}

	return claims
}

/**
 * validate if token is expired and return the remaining time
 * -> -1 if token is expired
 */
func getTokenRemainingValidity(timestamp interface{}) int {

	validity, ok := timestamp.(float64)

	if ok {

		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())

		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}

	return -1
}

/**
 * validate if authentication method in token is the expected method
 */
func getTokenCredentialValidity(tokenCredentials UserCredentialForJWT, userCredentials UserCredentialForJWT) bool {

	valid := tokenCredentials.UserAuthentication == userCredentials.UserAuthentication &&
		tokenCredentials.DeviceId == userCredentials.DeviceId &&
		tokenCredentials.Username == tokenCredentials.Username

	return valid
}

/**
 * parse jwt claims into an user cred struct
 */
func parseClaimsIntoStruct(claims jwt.Claims, claimStruct interface{}) {

	jsn, err := json.Marshal(claims)

	if err != nil {
		log.Fatal("Error reading the jwt claims", err)
	}

	err = json.Unmarshal(jsn, &claimStruct)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	return
}

func expiresIn(d time.Duration) int64 {

	return time.Now().Add(d).UnixNano() / int64(time.Second)
}
