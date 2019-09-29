package env

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

// Get the port to listen on
func GetListenAddressPort() string {
	port := getEnv("LISTEN_ADDRESS_PORT", "8091")
	return ":" + port
}

// Get address of policy engine
func GetPolicyEngineAddress() string {
	serverAddress := getEnv("POLICY_ENGINE_ADDRESS", "localhost:8090")
	return serverAddress
}

func GetCertificate() string {
	return getEnv("OWN_CRT", "")
}

func GetKey() string {
	return getEnv("OWN_KEY", "")
}

func GetRootCertificate() string {
	return getEnv("ROOT_CRT", "")
}

func GetRootIntermediateCertificate() string {
	return getEnv("ROOT_INTERMEDIATE_CRT", "")
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

	log.Printf("Trust Engine will run on: %s\n", GetListenAddressPort())
	log.Printf("Policy Engine will run on: %s\n", GetPolicyEngineAddress())
}
