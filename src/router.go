package main

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return mux
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		errorHandler(w, req, http.StatusNotFound)
		return
	}

	ExecuteTemplate(w, "home.html", nil)
}

func errorHandler(w http.ResponseWriter, req *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	ExecuteTemplate(w, "error.html", data)
}
