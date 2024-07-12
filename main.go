package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.NotFound(w, r)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	_, err := fmt.Fprintf(w, "Post Request Successful\nEmail: %s\nPassword: %s\n\n", email, password)
	if err != nil {
		return
	}

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello")
	if err != nil {
		return
	}

}
func main() {
	fileSerer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileSerer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Starting Server at Port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
