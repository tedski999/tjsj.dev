package sitegen

import (
	"os"; "path"
	"strings"
	 "io/ioutil"
	"github.com/tdewolff/minify/v2"
	. "github.com/otiai10/copy"
)

func (siteTemplate *SiteTemplateFormat) generateStaticDir(srcDir, dstDir string, m *minify.M) error {

	// Copy static directory into destination, minify selected files
	Copy(
		path.Join(srcDir, siteTemplate.Static.Directory),
		path.Join(dstDir, siteTemplate.Static.Directory))
	for file, mediatype := range siteTemplate.Static.Minify {
		file = path.Join(dstDir,  siteTemplate.Static.Directory, file)
		b, err := ioutil.ReadFile(file)
		if err != nil { return err }
		f, err := os.Create(file)
		if err != nil { return err }
		err = m.Minify(mediatype, f, strings.NewReader(string(b)))
		f.Close()
		if err != nil { return err }
	}

	// Generate sitemap.xml
	f, err := os.Create(path.Join(dstDir, siteTemplate.Static.Directory, siteTemplate.Static.Sitemap))
	if err != nil { return err }
	f.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
	f.Write([]byte(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">`))
	for route, file := range siteTemplate.Pages {
		info, err := os.Stat(path.Join(srcDir, file))
		if err != nil { return err }
		f.Write([]byte("<url>"))
		f.Write([]byte("<loc>https://" + siteTemplate.Name + route + "</loc>"))
		f.Write([]byte("<lastmod>" + info.ModTime().Format("2006-1-2T15:04:05+00:00") + "</lastmod>"))
		f.Write([]byte("</url>"))
	}
	f.Write([]byte("</urlset>"))
	f.Close()

	return nil
}
