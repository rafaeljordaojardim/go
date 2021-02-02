package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}
func main() {
	// HTTP E UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE SERVIDOR
	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
