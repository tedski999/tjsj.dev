package sitegen

import (
	"os"; "path"; "path/filepath"
	"fmt"; "strings"
	"net/http"; "html/template"
	"github.com/tdewolff/minify/v2"
)

func (siteTemplate *SiteTemplateFormat) generateFromTemplates(srcDir, dstDir string, m *minify.M) error {

	// Parse JSON data file
	data, err := ParseDataFile(path.Join(srcDir, siteTemplate.Data))
	if err != nil { return err }

	// Handy functions that can be used within HTML templates
	templates := template.New("")
	templates.Funcs(template.FuncMap{
		"split": func(text string) []string { return strings.Split(text, "") },
		"list": func(list ...string) []string { return list },
		"data": func(keys ...interface{}) interface{} { return indexData(data, keys) },
		"dict": func(kv ...interface{}) map[string]interface{} { return makeDict(kv) },
	})

	// General function for generating HTML files from templates
	generateFile := func(srcFile, dstFile string, templateData interface{}) error {
		srcFile = path.Join(srcDir, srcFile)
		dstFile = path.Join(dstDir, dstFile)
		t, err := templates.Clone()
		if err != nil { return err }
		t, err = t.ParseFiles(srcFile)
		if err != nil { return err }
		err = os.MkdirAll(filepath.Dir(dstFile), os.ModePerm)
		if err != nil { return err }
		f, err := os.Create(dstFile)
		if err != nil { return err }
		defer f.Close()
		mw := m.Writer("text/html", f)
		defer mw.Close()
		return t.ExecuteTemplate(mw, filepath.Base(srcFile), templateData)
	}

	// Generate segments for generating following templates
	for _, file := range siteTemplate.Segments {
		templates, err = templates.ParseFiles(path.Join(srcDir, file))
		if err != nil { return err }
	}

	// Generate pages
	for page := range siteTemplate.Pages {
		err = generateFile(siteTemplate.Pages[page], siteTemplate.Pages[page], data)
		if err != nil { return err }
	}

	// Generate error files
	var errorTemplateData struct { Code int; Message string }
	errorTemplateData.Code = http.StatusNotFound
	errorTemplateData.Message = http.StatusText(errorTemplateData.Code)
	err = generateFile(siteTemplate.ErrorTemplate, siteTemplate.Errors.NotFound, errorTemplateData)
	if err != nil { return err }
	errorTemplateData.Code = http.StatusInternalServerError
	errorTemplateData.Message = http.StatusText(errorTemplateData.Code)
	return generateFile(siteTemplate.ErrorTemplate, siteTemplate.Errors.Internal, errorTemplateData)
}

func indexData(data interface{}, keys []interface{}) interface{} {
	if len(keys) == 0 { return data }
	switch k := keys[0].(type) {
	case string:
		d, ok := data.(map[string]interface{})
		if !ok { panic("Unable to index data with string") }
		return indexData(d[k], keys[1:])
	case int:
		d, ok := data.([]interface{})
		if !ok { panic("Unable to index data with int") }
		return indexData(d[k], keys[1:])
	default:
		panic("Invalid key type")
	}
}

func makeDict(kv []interface{}) map[string]interface{} {
	size := len(kv)
	if size % 2 != 0 { panic("Missing value for last key") }
	dict := make(map[string]interface{}, size / 2)
	for i := 0; i < size; i += 2 {
		dict[fmt.Sprintf("%v", kv[i+0])] = kv[i+1]
	}
	return dict
}
