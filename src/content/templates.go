package content

import (
	"net/http"
	"html/template"
)

var htmlTemplates *template.Template

func LoadTemplates(pattern string) {
	htmlTemplates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) {
	htmlTemplates.ExecuteTemplate(w, name, data)
}
