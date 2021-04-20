package webcontent

import (
	"html/template"
	"path/filepath"
	"io/ioutil"
	"errors"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

func (content *Content) GetHTMLTemplate(name string) *template.Template {
	return content.htmlTemplates.Lookup(name)
}

func (content *Content) loadHTMLTemplates() error {

	// Find files within templateDirPath
	filenames, err := filepath.Glob(filepath.Join(content.templateDirPath, "*.html"))
	if err != nil {
		return err
	}
	if len(filenames) == 0 {
		return errors.New("No HTML template files found in '" + content.templateDirPath + "'!")
	}

	// Setup HTML minifier
	minifier := minify.New()
	minifier.AddFunc("text/html", html.Minify)

	// Minify every file before adding it to the templates
	content.htmlTemplates = template.New("")
	for _, filename := range filenames {

		// Minify HTML template gile
		fileBytes, err := ioutil.ReadFile(filename)
		if err != nil { return err }
		minifiedBytes, err := minifier.Bytes("text/html", fileBytes)
		if err != nil { return err }

		// Parse minified HTML and add to templates
		content.htmlTemplates = content.htmlTemplates.New(filepath.Base(filename))
		_, err = content.htmlTemplates.Parse(string(minifiedBytes))
		if err != nil { return err }
	}

	return nil
}
