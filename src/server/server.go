package server

import "net/http"

func StartServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/projects", projectsListHandler)
	mux.HandleFunc("/projects/", projectHandler)
	mux.HandleFunc("/posts", postsListHandler)
	mux.HandleFunc("/posts/", postHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", newFileServer("./static/")))
	return http.ListenAndServeTLS(":443", "certs/fullchain.pem", "certs/privkey.pem", mux)
}
