package webserver

import (
	"net/http"
	"html/template"
)

func (srv *Server) executeHTMLTemplate(w http.ResponseWriter, template *template.Template, data interface{}) {
	err := template.Execute(w, data)
	if err != nil {
		panic(err.Error())
	}
}
