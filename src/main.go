package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Loading HTML templates...")
	LoadTemplates("templates/*.html")
	LoadTitleGoofs("./data/titlegoofs.txt")

	log.Println("Creating routes...")
	r := NewRouter()

	log.Println("Starting HTTPS server...")
	err := http.ListenAndServeTLS(":443", "certs/fullchain.pem", "certs/privkey.pem", r)
	log.Fatal(err)
}

