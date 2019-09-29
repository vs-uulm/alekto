package ldap

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/sso_auth/env"
	"gopkg.in/ldap.v3"
	"log"
)

const (
	baseDn = "dc=planetexpress,dc=com"
)

/**
 * authenticate given user credentials
 */
func Authenticate(loginData request.BasicLoginPayload) bool {

	fmt.Printf("Trying to authenticate user: '%v' with password: '%v' \n", loginData.Username, loginData.Password)

	connection := getLdapConnection()
	defer connection.Close()

	searchRequest := getNewLdapSearchRequest(loginData.Username)

	// search ldap entry for given username
	searchResult, err := connection.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// no entry found
	if len(searchResult.Entries) != 1 {
		fmt.Println("User does not exist or too many entries returned")
		return false
	}

	userDN := searchResult.Entries[0].DN
	fmt.Printf("User exists with DN=(%s). \nVerifying password...\n", userDN)

	// Bind as the user to verify their password
	err = connection.Bind(userDN, loginData.Password)
	success := err == nil

	if success {
		fmt.Println("User " + loginData.Username + " successfully verified")
	} else {
		fmt.Println("User " + loginData.Username + " could not be verified")
	}

	return success
}

/**
 * get new ldap connection
 */
func getLdapConnection() *ldap.Conn {

	connection, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", env.GetLDAPHost(), env.GetLDAPPort()))

	if err != nil {
		log.Fatal(err)
	}

	return connection
}

/**
 * get new ldap search request
 */
func getNewLdapSearchRequest(username string) *ldap.SearchRequest {

	searchRequest := ldap.NewSearchRequest(
		baseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn"},
		nil,
	)

	return searchRequest
}
