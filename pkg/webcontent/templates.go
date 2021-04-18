package webcontent

import (
	"html/template"
	"path"
)

func (content *Content) GetHTMLTemplate(name string) *template.Template {
	return content.htmlTemplates.Lookup(name)
}

func (content *Content) loadHTMLTemplates() error {
	// TODO: HTML minifying
	var err error
	pattern := path.Join(content.templateDirPath, "*.html")
	content.htmlTemplates, err = template.ParseGlob(pattern)
	return err
}
