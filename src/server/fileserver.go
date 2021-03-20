package server

import "net/http"

func newFileServer(path string) http.Handler {
	staticFileSystem := http.Dir(path)
	staticFileServer := http.FileServer(staticFileSystem)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		file, err := staticFileSystem.Open(req.URL.Path)
		if err == nil {
			file.Close()
			staticFileServer.ServeHTTP(w, req)
		} else {
			errorHandler(w, req, http.StatusNotFound)
		}
	})
}
