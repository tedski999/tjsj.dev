package sitegen

type SiteTemplateFormat struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Static struct {
		Directory string `json:"directory"`
		Sitemap string `json:"sitemap"`
		Minify map[string]string `json:"minify"`
	} `json:"static"`
	Segments []string `json:"segments"`
	Pages map[string]string `json:"pages"`
	ErrorTemplate string `json:"errorTemplate"`
	Errors struct {
		NotFound string `json:"notfound"`
		Internal string `json:"internal"`
	} `json:"errors"`
}

type SiteFormat struct {
	StaticDir string `json:"staticdir"`
	Pages map[string]string `json:"pages"`
	Errors struct {
		NotFound string `json:"notfound"`
		Internal string `json:"internal"`
	} `json:"errors"`
}
