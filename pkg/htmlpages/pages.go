package htmlpages

import (
	"html/template"
	"path"
)

type Pages struct {
	templates *template.Template
}

func Load(dirpath string) *Pages {
	// TODO: HTML minifying
	pattern := path.Join(dirpath, "*.html")
	templates := template.Must(template.ParseGlob(pattern))
	return &Pages {
		templates: templates,
	}
}

func (pages *Pages) Get(name string) *template.Template {
	template := pages.templates.Lookup(name)
	if template == nil {
		panic("Could not get HTML page " + name)
	}
	return template
}
