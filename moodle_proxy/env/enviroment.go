package env

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func IgnoreFailedAuthorization() bool {
	debug := getEnv("IGNORE_FAILED_AUTHORIZATION", "0")
	return debug == "1"
}

// Get the proxy port to listen on
func GetListenAddress() string {
	port := getEnv("PROXY_PORT", "8080")
	return ":" + port
}

// Get the proxy port to listen on
func GetListenAddressPortTLS() string {
	port := getEnv("PROXY_PORT_TLS", "10443")
	return ":" + port
}

func GetPolicyEngineAddress() string {
	return getEnv("POLICY_ENGINE_ADDRESS", "localhost:8090")
}

func MoodleEnabled() bool {
	disabled := getEnv("MOODLE_DISABLED", "0")
	return disabled == "0"
}

// Get the server address to listen on
func GetServerAddress() string {
	serverAddress := getEnv("MOODLE_SERVER_URL", "localhost:8443")
	return serverAddress
}

func GetDummyWebAddress() string {
	serverAddress := getEnv("DUMMY_WEB_ADDRESS", "dummy-web.uni-ulm.de:4438")
	return serverAddress
}

func GetMariaDbInfo() string {
	user := getEnv("MOODLE_DB_USER", "bn_moodle")
	password := getEnv("MOODLE_DB_PASSWORD", "")
	host := getEnv("MOODLE_DB_HOST", "mariadb")
	port := getEnv("MOODLE_DB_PORT", "3306")
	dbName := getEnv("MOODLE_DB_NAME", "bitnami_moodle")

	mariaDbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbName)

	return mariaDbInfo
}

// Get moodle login url
func GetMoodleLoginUrl() string {
	loginUrl := "https://" + GetServerAddress() + "/login/index.php"
	return loginUrl
}

// Get the server address to listen on
func GetSsoAuthAddress() string {
	serverAddress := getEnv("SSO_AUTH_URL", "sso.uni-ulm.de:4435")
	return serverAddress
}

// Get the server address to listen on
func GetMoodleSessionCookieName() string {
	cookie := getEnv("MOODLE_SESSION_COOKIE", "MoodleSession")
	return cookie
}

// Get the server address to listen on
func GetSsoAuthMoodleSessionCookieName() string {
	cookie := getEnv("SSO_AUTH_MOODLE_SESSION_COOKIE", "_sso_auth_moodle")
	return cookie
}

// Get the logout url for the server
func GetLogoutUrl() string {
	logoutUrl := getEnv("LOGOUT_URL", "/login/logout.php")
	return logoutUrl
}

func GetCertificate() string {
	return getEnv("PROXY_CRT", "")
}

func GetRootCertificate() string {
	return getEnv("ROOT_CRT", "")
}

func GetRootIntermediateCertificate() string {
	return getEnv("ROOT_INTERMEDIATE_CRT", "")
}

func GetPrivateKey() string {
	return getEnv("PROXY_KEY", "keys/moodle_proxy.key")
}

// Get the logout url for the server
func GetProxyDomain() string {
	return getEnv("PROXY_DOMAIN", "moodle.uni-ulm.de")
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

	log.Printf("Proxy Server will run on: %s\n", GetListenAddressPortTLS())
	log.Printf("Moodle Server is running on url: %s\n", GetServerAddress())
	log.Printf("SSO Auth Server is running on url: %s\n", GetSsoAuthAddress())
	log.Printf("Policy Engine is running on url: %s\n", GetPolicyEngineAddress())
	log.Printf("Ignore failed authorization: %v\n", IgnoreFailedAuthorization())
}
