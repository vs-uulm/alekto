package twoFactorAuth

import (
	"html/template"
	"log"
	"net/http"
)

/**
 * TODO Shreya
 */
func ShowLoginPage (res http.ResponseWriter) {

	indexTemplate, err := template.ParseFiles("static/twoFactorLogin.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = indexTemplate.Execute(res, nil)
}