package sitegen

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

func ParseSiteTemplateFile(templateFile string) (SiteTemplateFormat, error) {
	var siteTemplate SiteTemplateFormat
	file, err := os.Open(templateFile)
	if err != nil { return siteTemplate, err }
	bytes, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil { return siteTemplate, err }
	err = json.Unmarshal(bytes, &siteTemplate)
	if err != nil { return siteTemplate, err }
	return siteTemplate, nil
}

func ParseSiteFile(siteFile string) (SiteFormat, error) {
	var site SiteFormat
	file, err := os.Open(siteFile)
	if err != nil { return site, err }
	bytes, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil { return site, err }
	err = json.Unmarshal(bytes, &site)
	if err != nil { return site, err }
	return site, nil
}

func ParseDataFile(dataFile string) (map[string]interface{}, error) {
	file, err := os.Open(dataFile)
	if err != nil { return nil, err }
	bytes, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil { return nil, err }
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil { return nil, err }
	return data, nil
}
