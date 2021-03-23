package server

import (
	"net/http"
	"strings"
	"github.com/tedski999/tjsj.dev/src/content"
)

// Respond with the HTML template "home.html"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: pass list of recent projects and posts into the template
	data := struct { SplashText string } { content.GetRandomSplash() }
	content.ExecuteTemplate(w, "home.html", data)
}

// Respond with a list of projects in the HTML template "projects.html"
func projectsListHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of projects metadata
	data := struct { SplashText string } { content.GetRandomSplash() }
	content.ExecuteTemplate(w, "projects.html", data)
}

// Respond with a list of posts in the HTML template "posts.html"
func postsListHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of posts metadata
	data := struct { SplashText string } { content.GetRandomSplash() }
	content.ExecuteTemplate(w, "posts.html", data)
}

// Respond with the project page of the id given in the URL
func projectHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id of the project requested from the URL
	// Respond with the posts list page instead if the id is empty
	id := strings.TrimPrefix(r.URL.Path, "/projects/")
	if len(id) == 0 {
		http.Redirect(w, r, "/projects", http.StatusMultipleChoices)
		return
	}

	// TODO: find the project with the id, if none found, 404
	errorHandler(w, r, http.StatusNotFound)
}

// Respond with the post page of the id given in the URL
func postHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id of the post requested from the URL
	// Respond with the posts list page instead if the id is empty
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if len(id) == 0 {
		http.Redirect(w, r, "/posts", http.StatusMultipleChoices)
		return
	}

	// TODO: find the post with the id, if none found, 404
	errorHandler(w, r, http.StatusNotFound)
}

// Respond with the error page with appropriate code and message
func errorHandler(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	content.ExecuteTemplate(w, "error.html", data)
}
