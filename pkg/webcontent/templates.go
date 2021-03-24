package webcontent

import (
	"html/template"
	"path"
)

func (content *Content) LoadHTMLTemplates(pattern string) {
	fullPattern := path.Join(pattern, "*.html")
	content.htmlTemplates = template.Must(template.ParseGlob(fullPattern))
}

func (content *Content) GetHTMLTemplate(name string) *template.Template {
	t := content.htmlTemplates.Lookup(name)
	if t == nil {
		panic("HTML template not found: " + name)
	}
	return t
}
