package templateHandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ShowLogoutPage (res http.ResponseWriter) {

	fmt.Println("---- LOGOUT ----")

	indexTemplate, err := template.ParseFiles("static/logout.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = indexTemplate.Execute(res, nil)
}
