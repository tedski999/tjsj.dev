package sitegen

import (
	"io/ioutil"
	"encoding/json"
)

func ParseSiteTemplateFile(file string) (SiteTemplateFormat, error) {
	var siteTemplate SiteTemplateFormat
	bytes, err := ioutil.ReadFile(file)
	if err != nil { return siteTemplate, err }
	err = json.Unmarshal(bytes, &siteTemplate)
	if err != nil { return siteTemplate, err }
	return siteTemplate, nil
}

func ParseSiteFile(file string) (SiteFormat, error) {
	var site SiteFormat
	bytes, err := ioutil.ReadFile(file)
	if err != nil { return site, err }
	err = json.Unmarshal(bytes, &site)
	if err != nil { return site, err }
	return site, nil
}

func ParseDataFile(file string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil { return nil, err }
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil { return nil, err }
	return data, nil
}
