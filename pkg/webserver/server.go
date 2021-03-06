package webserver

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/NYTimes/gziphandler"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"

	"github.com/tedski999/tjsj.dev/pkg/webcontent"
	"github.com/tedski999/tjsj.dev/pkg/webstats"
)

type Server struct {
	http *http.Server
	certFilePath, keyFilePath string
	content *webcontent.Content
	stats *webstats.Statistics
	doneWG sync.WaitGroup
	errChan chan<- error
}

func Create(content *webcontent.Content, stats *webstats.Statistics, certFilePath, keyFilePath string) (*Server, error) {

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
		certFilePath: certFilePath,
		keyFilePath:keyFilePath,
		content: content,
		stats: stats,
	}

	// Setup the CSS minifier middleware
	minifier := minify.New()
	minifier.AddFunc("text/css", css.Minify)

	// Setup HTTP route multiplexing
	router.StrictSlash(true)
	router.HandleFunc("/", server.homeResponse)
	router.HandleFunc("/projects/", server.projectsResponse)
	router.HandleFunc("/projects/{id}", server.projectResponse)
	router.HandleFunc("/posts/", server.postsResponse)
	router.HandleFunc("/posts/{id}", server.postResponse)
	router.HandleFunc("/stats", server.statsResponse)
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.errorResponse(w, r, http.StatusNotFound)
	})

	router.Use(server.recordCompressedData)
	router.Use(gziphandler.GzipHandler)
	router.Use(minifier.Middleware)
	router.Use(server.recordData)
	router.Use(server.trimWWWRequests)
	router.Use(server.serveStaticFiles)

	return server, nil
}
