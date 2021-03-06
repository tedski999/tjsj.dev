package webcontent

import (
	"html/template"
	"path/filepath"
	"io/ioutil"
    "strings"
	"fmt"
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

	// Add util functions for templates
	content.htmlTemplates = template.New("")
	content.htmlTemplates.Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"split": func(text string) []string {
			return strings.Split(text, "")
		},
		"dict": func(keyvalues ...interface{}) map[string]interface{} {
			dict := make(map[string]interface{}, len(keyvalues) / 2)
			for i := 0; i < len(keyvalues); i += 2 {
				key := keyvalues[i].(string)
				dict[key] = keyvalues[i + 1]
			}
			return dict
		},
		"list": func(list ...interface{}) []interface{} {
			return list
		},
		"toHTML": func(text template.HTML) template.HTML {
			return text
		},
		"join": func(slice []interface{}, separator string) string {
			strs := make([]string, len(slice))
			for i, v := range slice {
				strs[i] = fmt.Sprintf("%v", v)
			}
			return strings.Join(strs, separator)
		},
		"joinstr": func(slice []string, separator string) string {
			return strings.Join(slice, separator)
		},
	})

	// Minify every file before adding it to the templates
	minifier := minify.New()
	minifier.AddFunc("text/html", html.Minify)
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
