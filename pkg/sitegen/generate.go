package sitegen

import (
	"log"; "os"; "path"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/xml"
)

func Generate(templateFile, dstDir string) error {

	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)
	m.AddFunc("application/xml", xml.Minify)

	log.Println("Parsing template file " + templateFile + "...")
	siteTemplate, err := ParseSiteTemplateFile(templateFile)
	if err != nil { return err }
	srcDir := path.Dir(templateFile)

	if err = os.RemoveAll(dstDir); err != nil { return err }
	log.Printf("Generating minified static directory %s from %s...\n", path.Join(dstDir, siteTemplate.Static.Directory), path.Join(srcDir, siteTemplate.Static.Directory))
	if err = siteTemplate.generateStaticDir(srcDir, dstDir, m); err != nil { return err }
	log.Printf("Generating %d HTML pages from %d templates...\n", len(siteTemplate.Pages), len(siteTemplate.Pages) + len(siteTemplate.Segments))
	if err = siteTemplate.generateFromTemplates(srcDir, dstDir, m); err != nil { return err }
	log.Printf("Generating site file %s from template file %s...\n", path.Join(dstDir, path.Base(templateFile)), templateFile)
	if err = siteTemplate.generateSiteFile(templateFile, dstDir, m); err != nil { return err }

	return nil
}

