package webserver

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tedski999/tjsj.dev/pkg/fileserver"
)

// Setup HTTP route multiplexing
func (srv *Server) registerHandlers() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// TODO: subdomain handling
	r.HandleFunc("/", srv.homeHandler)
	r.HandleFunc("/posts/", srv.postsHandler)
	r.HandleFunc("/posts/{id}", srv.postHandler)
	r.PathPrefix("/").Handler(fileserver.Create(staticFilesDir, srv.errorHandler))

	srv.httpServer.Handler = r
}

// Respond with the HTML template "home.html"
func (srv *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of recent posts metadata
	data := struct { SplashText string } { srv.content.GetRandomSplash() }
	template := srv.content.GetHTMLTemplate("home.html")
	srv.executeHTMLTemplate(w, template, data)
}

// Respond with a list of posts in the HTML template "posts.html"
func (srv *Server) postsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of posts metadata
	data := struct { SplashText string } { srv.content.GetRandomSplash() }
	template := srv.content.GetHTMLTemplate("posts.html")
	srv.executeHTMLTemplate(w, template, data)
}

// Respond with the post page of the id given in the URL
func (srv *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//postID := vars["id"]

	// TODO: find the post

	srv.errorHandler(w, r, http.StatusNotFound)
}

// Respond with the error page with an appropriate message
func (srv *Server) errorHandler(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	template := srv.content.GetHTMLTemplate("error.html")
	srv.executeHTMLTemplate(w, template, data)
}
