package templateHandler

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func ShowAuthorizationFailurePage (res http.ResponseWriter, message string) {

	indexTemplate, err := template.ParseFiles("static/authorizationFailed.html")
	safe := template.HTMLEscapeString(message)
	safe = strings.Replace(safe, "\n", "<br>", -1)

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = indexTemplate.Execute(res, template.HTML(safe))
}
