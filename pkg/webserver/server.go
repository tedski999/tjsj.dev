package webserver

import (
	"net/http"
	"sync"
	"time"
	"strings"
	"github.com/gorilla/mux"
	"github.com/tedski999/tjsj.dev/pkg/fileserver"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
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
	router.StrictSlash(true)
	router.HandleFunc("/", server.homeResponse)
	router.HandleFunc("/posts/", server.postsResponse)
	router.HandleFunc("/posts/{id}", server.postResponse)

	// Serve static files, redirect anything else to the error response
	staticFileServer := fileserver.Create("./web/static/", server.errorResponse)
	router.PathPrefix("/").Handler(staticFileServer)

	// Middleware to trim any requests prefixed with "www."
	router.Use(func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.Host, "www.") {
				u := *r.URL
				u.Host = strings.TrimPrefix(r.Host, "www.")
				http.Redirect(w, r, u.String(), http.StatusFound)
			} else {
				next.ServeHTTP(w, r)
			}

		})
	})

	// Setup CSS minifier
	minifier := minify.New()
	minifier.AddFunc("text/css", css.Minify)
	router.Use(minifier.Middleware)

	return server, nil
}
