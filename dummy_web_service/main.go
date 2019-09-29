package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dummy Web Sever content")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServeTLS(":4438", "dummy-web.crt", "dummy-web.key", nil)

	if err != nil {
		fmt.Println(err)
	}
}
