package webserver

import (
	"net/http"
	"html/template"
	"errors"
)

type homeResponseData struct {
	SplashText string
	ProjectsList []template.HTML
	RecentPostsList []template.HTML
}

type statsResponseData struct {
	TotalHits int
	Uptime string
}

// General method for executing HTML templates by name
func (server *Server) executeHTMLTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	template := server.content.GetHTMLTemplate(templateName)
	if template == nil {
		server.errChan <- errors.New("HTML template '" + templateName + "' not found!")
		return
	}

	template.Execute(w, data)
}

// Respond with the HTML template "home.html"
func (server *Server) homeResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of recent posts metadata
	data := homeResponseData {
		server.content.GetRandomSplashText(),
		nil,
		nil,
	}

	server.executeHTMLTemplate(w, "home.html", data)
}

// Respond with a list of posts in the HTML template "posts.html"
func (server *Server) postsResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of posts metadata
	// server.executeHTMLTemplate(w, "posts.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the post page of the id given in the URL
func (server *Server) postResponse(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//postID := vars["id"]
	// TODO: find the post
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the statistics page
func (server *Server) statsResponse(w http.ResponseWriter, r *http.Request) {
	data := statsResponseData {
		server.stats.GetTotalHits(),
		server.stats.GetUptime(),
	}

	server.executeHTMLTemplate(w, "stats.html", data)
}

// Respond with the error page with an appropriate message
func (server *Server) errorResponse(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	server.executeHTMLTemplate(w, "error.html", data)
}

