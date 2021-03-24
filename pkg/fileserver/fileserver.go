package fileserver

import "net/http"

type ErrorHandler func(http.ResponseWriter, *http.Request, int)

// Creates a HTTP handler which serves static files rooted at 'path'
func Create(path string, errorHandler ErrorHandler) http.Handler {
	staticFileSystem := http.Dir(path)
	staticFileServer := http.FileServer(staticFileSystem)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check if the requested file exists
		file, err := staticFileSystem.Open(r.URL.Path)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		defer file.Close()

		// Serve the file if it does
		staticFileServer.ServeHTTP(w, r)
	})
}
