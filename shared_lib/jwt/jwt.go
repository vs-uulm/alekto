package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

/*
 * TODO add needed credentials
 */
type UserCredentialForJWT struct {
	Username             string `json:"Username"`
	DeviceId             string `json:"DeviceId"`
	IPAdress             string `schema:"ipAddress"`
	Domain               string `json:"Domain"`
	UserAuthentication   string `json:"UserAuthentication"`
	DeviceAuthentication string `json:"DeviceAuthentication"`
}

/**
 * create new Json Web Token with the user credentials and
 * sign it with private key
 */
func CreateAndSignNewJWT(userCredentials UserCredentialForJWT, privateKey *rsa.PrivateKey) string {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, getClaimsByUserCredentials(userCredentials))
	getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"])

	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		fmt.Printf("Err while signing jwt: %v \n", err)
	}

	return tokenString
}

/**
 * Verify Token with public key and validate it's claims
 */
func VerifyJWTAndParseClaimsToStruct(userToken string, publicKey *rsa.PublicKey,
	userCredentials *UserCredentialForJWT) (valid bool) {

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err == nil && token != nil {
		valid = validateToken(token, userCredentials)
	} else {
		fmt.Println("error while verifying user jwt", err.Error())
	}

	return valid
}

func ParseClaimsIntoStruct(userToken string, publicKey *rsa.PublicKey, userCredentials *UserCredentialForJWT) {

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("error while verifying user jwt", err.Error())
		return
	}

	if token == nil {
		fmt.Println("error while verifying user jwt: empty token")
		return
	}

	parseClaimsIntoStruct(token.Claims, &userCredentials)
}

/**
 * validate Token
 * -> claims vs. userCredentials
 * -> remaining Time
 */
func validateToken(token *jwt.Token, userCredentials *UserCredentialForJWT) (valid bool) {

	tokenCredentials := UserCredentialForJWT{}
	remainingTime := getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"])

	parseClaimsIntoStruct(token.Claims, &tokenCredentials)

	userCredentials.Username = tokenCredentials.Username

	valid = token.Valid && getTokenCredentialValidity(tokenCredentials, *userCredentials)

	fmt.Printf("Token Validiation: valid=%v ; user=%v ; time remaining=%vs \n",
		valid, userCredentials, remainingTime)

	return
}
