package basicAuth

import (
	"html/template"
	"log"
	"net/http"
)

func ShowLoginPage (res http.ResponseWriter, message string) {

	indexTemplate, err := template.ParseFiles("static/login.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = indexTemplate.Execute(res, message)
}