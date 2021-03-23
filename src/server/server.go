package server

import "net/http"

func StartServer() error {
	// TODO: subdomains: www should redirect to @
	mux := http.NewServeMux()
	mux.Handle("/", newRootFileServer(homeHandler, "./public/"))
	mux.HandleFunc("/projects", projectsListHandler)
	mux.HandleFunc("/projects/", projectHandler)
	mux.HandleFunc("/posts", postsListHandler)
	mux.HandleFunc("/posts/", postHandler)
	return http.ListenAndServeTLS(":443", "certs/fullchain.pem", "certs/privkey.pem", mux)
}
