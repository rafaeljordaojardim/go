package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	// HTTP E UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE SERVIDOR

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"pedro", "email"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
