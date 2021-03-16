package main

import (
	"net/http"
	"html/template"
	"log"
)

var templates *template.Template

func main() {
	log.Println("Loading HTML templates...")
	templates = template.Must(template.ParseGlob("templates/*.html"))

	log.Println("Setting up routes...")
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Println("Starting HTTPS server...")
	err := http.ListenAndServeTLS(":443", "certs/fullchain.pem", "certs/privkey.pem", nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// DEBUG: reloading templates at runtime
	templates = template.Must(template.ParseGlob("templates/*.html"))
	templates.ExecuteTemplate(w, "index.html", nil)
}
