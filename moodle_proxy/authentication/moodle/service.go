package moodle

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"github.com/ma-zero-trust-prototype/shared_lib/jwt"
	"net/http"
	"net/url"
	"strings"
)

/**
 * Send new login request to moodle server with username and password (jwt)
 * return moodle session cookie "MoodleSession"
 */
func SendNewLoginRequestAndReturnSessionCookie(userJWT string, userCredentials jwt.UserCredentialForJWT) (
	sessionCookie *http.Cookie) {

	// stop moodle from redirecting
	httpClient := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}

	form := getRequiredLoginForm(userCredentials.Username, userJWT)
	req := getNewLoginRequest(form)
	response, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	sessionCookie = getSessionCookieOfResponse(response)

	return
}

func getRequiredLoginForm(username string, userJWT string) url.Values {

	form := url.Values{}

	form.Add("username", username)
	form.Add("password", userJWT)

	return form
}

func getNewLoginRequest(form url.Values) (req *http.Request) {

	req, err := http.NewRequest(http.MethodPost, env.GetMoodleLoginUrl(), strings.NewReader(form.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return
}

func getSessionCookieOfResponse(response *http.Response) (sessionCookie *http.Cookie) {

	if response.Cookies() == nil {
		return
	}

	for _, sessCookie := range response.Cookies() {

		if sessCookie.Name == env.GetMoodleSessionCookieName() {
			sessionCookie = sessCookie
		}
	}

	return
}
