package server

import "net/http"

// Creates a new HTTP handler which serves 'rootHandler' to requests for /
// Otherwise, it serves the static files rooted at 'path'
func newRootFileServer(rootHandler http.HandlerFunc, path string) http.Handler {
	staticFileSystem := http.Dir(path)
	staticFileServer := http.FileServer(staticFileSystem)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Handle root requests separately
		if r.URL.Path == "/" {
			rootHandler(w, r)
			return
		}

		// Otherwise we could be serving a public static file
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
