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
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Println("Starting HTTPS server...")
	err := http.ListenAndServeTLS(":443", "certs/fullchain.pem", "certs/privkey.pem", nil)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		errorHandler(w, req, http.StatusNotFound)
		return
	}

	templates.ExecuteTemplate(w, "home.html", nil)
}

func errorHandler(w http.ResponseWriter, req *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	templates.ExecuteTemplate(w, "error.html", data)
}
