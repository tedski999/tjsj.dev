package webserver

import (
	"net/http"
	"strings"
	"time"
)

// Middleware to trim any requests prefixed with "www."
func (server *Server) trimWWWRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Redirect instead if the request starts with "www."
		if strings.HasPrefix(r.Host, "www.") {
			u := *r.URL
			u.Host = strings.TrimPrefix(r.Host, "www.")
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Middleware to serve static files if found
func (server *Server) serveStaticFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Serve the static file instead if it exists
		if file, err := server.content.OpenStaticFile(r.URL.Path); err == nil {
			defer file.Close()
			if info, err := file.Stat(); err == nil && !info.IsDir() {
				http.ServeContent(w, r, r.URL.Path, time.Time{}, file)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
