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

// Middleware to record data about every OK request made to the server
type statusWriter struct {
	http.ResponseWriter
	status, length int
}
func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 { w.status = 200 }
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}
func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
func (server *Server) recordRequestData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter { ResponseWriter: w }
		next.ServeHTTP(&sw, r)
		if sw.status >= 200 && sw.status < 300 {
			server.stats.RecordRequest(r)
		}
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
