package sitegen

import (
	"os"; "path"
	 "io/ioutil"; "encoding/json"
	"github.com/tdewolff/minify/v2"
)

func (siteTemplate *SiteTemplateFormat) generateSiteFile(templateFile, dstDir string, m *minify.M) error {
	site := SiteFormat {
		siteTemplate.Static.Directory,
		siteTemplate.Pages,
		siteTemplate.Errors,
	}
	siteData, err := json.Marshal(site)
	if err != nil { return err }
	info, err := os.Stat(templateFile)
	if err != nil { return err }
	return ioutil.WriteFile(
		path.Join(dstDir, path.Base(templateFile)),
		siteData, info.Mode())
}
