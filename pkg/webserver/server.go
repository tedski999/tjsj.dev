package webserver

import (
	"net/http"
	"sync"
	"time"
	"github.com/gorilla/mux"
	"github.com/tedski999/tjsj.dev/pkg/fileserver"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
)

type Server struct {
	http *http.Server
	certFilePath, keyFilePath string
	content *webcontent.Content
	doneWG sync.WaitGroup
	errChan chan<- error
}

func Create(content *webcontent.Content) (*Server, error) {

	// Setup server
	router := mux.NewRouter()
	server := &Server {
		http: &http.Server {
			Addr: ":https",
			Handler: router,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		certFilePath: "./web/certs/fullchain.pem",
		keyFilePath: "./web/certs/privkey.pem",
		content: content,
	}

	// Setup HTTP route multiplexing
	// TODO: subdomain handling
	router.StrictSlash(true)
	router.HandleFunc("/", server.homeResponse)
	router.HandleFunc("/posts/", server.postsResponse)
	router.HandleFunc("/posts/{id}", server.postResponse)

	// Serve static files, redirect anything else to the error response
	staticFileServer := fileserver.Create("./web/static/", server.errorResponse)
	router.PathPrefix("/").Handler(staticFileServer)

	return server, nil
}
