package sitegen

import (
	"log";
	"os"; "os/exec"
	"path"; "path/filepath"
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
	dstDir, err = filepath.Abs(dstDir)
	if err != nil { return err }

	log.Println("Running pre-generation hooks...")
	for _, args := range siteTemplate.Hooks.Pregen {
		args[0], err = exec.LookPath(args[0])
		if err != nil { return err }
		cmd := exec.Cmd {
			Path: args[0],
			Args: args,
			Dir: srcDir,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			Env: append(os.Environ(), "SITEGEN_DST=" + dstDir),
		}
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if err = os.RemoveAll(dstDir); err != nil { return err }
	log.Printf("Generating minified static directory %s from %s...\n", path.Join(dstDir, siteTemplate.Static.Directory), path.Join(srcDir, siteTemplate.Static.Directory))
	if err = siteTemplate.generateStaticDir(srcDir, dstDir, m); err != nil { return err }
	log.Printf("Generating %d HTML pages from %d templates...\n", len(siteTemplate.Pages), len(siteTemplate.Pages) + len(siteTemplate.Segments))
	if err = siteTemplate.generateFromTemplates(srcDir, dstDir, m); err != nil { return err }
	log.Printf("Generating site file %s from template file %s...\n", path.Join(dstDir, path.Base(templateFile)), templateFile)
	if err = siteTemplate.generateSiteFile(templateFile, dstDir, m); err != nil { return err }

	log.Println("Running post-generation hooks...")
	for _, args := range siteTemplate.Hooks.Postgen {
		args[0], err = exec.LookPath(args[0])
		if err != nil { return err }
		cmd := exec.Cmd {
			Path: args[0],
			Args: args,
			Dir: srcDir,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			Env: append(os.Environ(), "SITEGEN_DST=" + dstDir),
		}
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

