package webserver

import "net/http"

// Respond with the HTML template "home.html"
func (srv *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of recent posts metadata
	data := struct { SplashText string } { srv.splashes.GetRandom() }
	srv.pages.Get("home.html").Execute(w, data)
}

// Respond with a list of posts in the HTML template "posts.html"
func (srv *Server) postsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of posts metadata
	data := struct { SplashText string } { srv.splashes.GetRandom() }
	srv.pages.Get("posts.html").Execute(w, data)
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
	srv.pages.Get("error.html").Execute(w, data)
}
