package env

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func GetListenAddressHost() string {
	return getEnv("SSO_AUTH_HOST", "localhost")
}

func GetLDAPHost() string {
	return getEnv("LDAP_SERVER", "localhost")
}

func GetLDAPPort() string {
	port := getEnv("LDAP_PORT", "389")
	return port
}

func GetListenAddressPort() string {
	port := getEnv("SSO_AUTH_PORT", "8085")
	return ":" + port
}

func GetListenAddressURL() string {
	return getEnv("SSO_AUTH_URL", "sso.uni-ulm.de")
}

func GetListenAddressPortTLS() string {
	port := getEnv("SSO_AUTH_PORT_TLS", "4435")
	return ":" + port
}

func GetProxyAddress() string {
	return getEnv("PROXY_URL", "moodle.uni-ulm.de:10443")
}

func GetRootCertificate() string {
	return getEnv("ROOT_CRT", "")
}

func GetRootIntermediateCertificate() string {
	return getEnv("ROOT_INTERMEDIATE_CRT", "")
}

func GetCertificate() string {
	return getEnv("SSO_CRT", "keys/sso.crt")
}

func GetPrivateKey() string {
	return getEnv("SSO_KEY", "keys/sso.key")
}

func GetSsoAuthSessionCookieName() string {
	return getEnv("SSO_AUTH_SESSION_COOKIE", "_sso_auth")
}

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/**
 * init env variables and log setup
 */
func LogSetup() {

	err := gotenv.Load()

	if err != nil {
		fmt.Println("Failed to load env variables")
		panic(err)
	}

	log.Printf("SSO Auth Server will run on: %s\n", GetListenAddressPort())
	log.Printf("SSO Auth Server will run secure connection on: %s\n", GetListenAddressPortTLS())
	log.Printf("Proxy Server is running on url: %s\n", GetProxyAddress())
}
