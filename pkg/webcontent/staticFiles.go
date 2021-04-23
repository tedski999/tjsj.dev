package webcontent

import "net/http"

func (content *Content) OpenStaticFile(path string) (http.File, error) {
	return content.staticFilesDir.Open(path)
}

func (content *Content) createStaticFileDir() error {
	content.staticFilesDir = http.Dir(content.staticFilesDirPath)
	return nil
}
