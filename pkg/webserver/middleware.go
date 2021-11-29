package webserver

import (
	"time"
	"strings"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status, length int
	header bool
}

func (w *responseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	if !w.header { w.WriteHeader(http.StatusOK) }
	w.length += n
	return n, err
}

func (w *responseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	if !w.header {
		w.header = true
		w.status = status
	}
}

func (server *Server) trimWWWRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Host, "www.") {
			u := *r.URL
			u.Host = strings.TrimPrefix(r.Host, "www.")
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (server *Server) serveStaticFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if file, err := server.static.Open(r.URL.Path); err == nil {
			defer file.Close()
			if info, err := file.Stat(); err == nil && !info.IsDir() {
				http.ServeContent(w, r, r.URL.Path, time.Time{}, file)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (server *Server) recordRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := responseWriter { ResponseWriter: w }
		next.ServeHTTP(&sw, r)
		server.stats.stats.Data.NoCompression += int64(sw.length)
		server.stats.record(sw.status, r)
	})
}

func (server *Server) recordCompression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := responseWriter { ResponseWriter: w }
		next.ServeHTTP(&sw, r)
		server.stats.stats.Data.Total += int64(sw.length)
	})
}
