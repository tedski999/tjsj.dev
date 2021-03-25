package webserver

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/radovskyb/watcher"
	"github.com/tedski999/tjsj.dev/pkg/fileserver"
	"github.com/tedski999/tjsj.dev/pkg/htmlpages"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
	"github.com/tedski999/tjsj.dev/pkg/splashes"
)

type Server struct {
	httpServer *http.Server
	pages *htmlpages.Pages
	content *webcontent.Content
	splashes *splashes.Splashes
	fileWatcher *watcher.Watcher
}

func Create() *Server {

	// Configure server
	router := mux.NewRouter()
	httpServer := &http.Server {
		Addr: ":https",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Setup file watching
	fileWatcher := watcher.New()
	if err := fileWatcher.AddRecursive(pagesDir); err != nil {
		panic(err.Error())
	}
	if err := fileWatcher.AddRecursive(contentDir); err != nil {
		panic(err.Error())
	}
	if err := fileWatcher.Add(splashesFilePath); err != nil {
		panic(err.Error())
	}

	// Setup server
	srv := &Server {
		httpServer: httpServer,
		pages: htmlpages.Load(pagesDir),
		content: webcontent.Load(contentDir),
		splashes: splashes.Load(splashesFilePath),
		fileWatcher: fileWatcher,
	}

	// Setup HTTP route multiplexing
	// TODO: subdomain handling
	router.StrictSlash(true)
	router.HandleFunc("/", srv.homeHandler)
	router.HandleFunc("/posts/", srv.postsHandler)
	router.HandleFunc("/posts/{id}", srv.postHandler)
	router.PathPrefix("/").Handler(fileserver.Create(staticFilesDir, srv.errorHandler))

	return srv
}
