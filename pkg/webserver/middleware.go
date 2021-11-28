package webserver

import (
	"time"
	"strings"
	"net/http"
)

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
