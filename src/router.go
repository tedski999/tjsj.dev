package main

import (
	"net/http"
	"strings"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/projects/", projectsHandler)
	mux.HandleFunc("/posts/", postsHandler)

	// Static file server
	staticFileSystem := http.Dir("./static/")
	staticFileServer := http.FileServer(staticFileSystem)
	staticHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, err := staticFileSystem.Open(req.URL.Path)
		if err == nil {
			staticFileServer.ServeHTTP(w, req)
		} else {
			errorHandler(w, req, http.StatusNotFound)
		}
	})
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	return mux
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		errorHandler(w, req, http.StatusNotFound)
		return
	}

	// TODO: pass list of recent projects and posts into the template
	ExecuteTemplate(w, "home.html", nil)
}

func projectsHandler(w http.ResponseWriter, req *http.Request) {
	selection := strings.TrimPrefix(req.URL.Path, "/projects/")
	if len(selection) == 0 {
		// TODO: pass the list of projects into the template
		ExecuteTemplate(w, "projects.html", nil)
	} else {
		// TODO: find the project with a name equal to selection
		// TODO: if none found, 404
		errorHandler(w, req, http.StatusNotFound)
	}
}

func postsHandler(w http.ResponseWriter, req *http.Request) {
	selection := strings.TrimPrefix(req.URL.Path, "/posts/")
	if len(selection) == 0 {
		// TODO: pass the list of posts into the template
		ExecuteTemplate(w, "posts.html", nil)
	} else {
		// TODO: find the post with a name equal to selection
		// TODO: if none found, 404
		errorHandler(w, req, http.StatusNotFound)
	}
}

func errorHandler(w http.ResponseWriter, req *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	ExecuteTemplate(w, "error.html", data)
}
