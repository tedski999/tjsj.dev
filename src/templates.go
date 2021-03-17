package main

import (
	"net/http"
	"html/template"
)

var templates *template.Template

func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) {
	templates.ExecuteTemplate(w, name, data)
}
