package webserver

import (
	"log"
	"net/http"
	"os"; "path"
	"time"; "sync"
	"github.com/gorilla/mux"
	"github.com/NYTimes/gziphandler"
	"github.com/tedski999/tjsj.dev/pkg/sitegen"
)

type Server struct {
	static http.Dir
	http, https http.Server
	wg sync.WaitGroup
	stats statistics
}

func Create(siteFile, statFile string) (*Server, error) {

	log.Println("Parsing site file " + siteFile + "...")
	site, err := sitegen.ParseSiteFile(siteFile)
	if err != nil { return nil, err }
	root := path.Dir(siteFile)

	var server Server
	server.static = http.Dir(path.Join(root, site.StaticDir))
	server.http = http.Server {
		Addr: ":http",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.https = http.Server {
		Addr: ":https",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.stats.file = statFile
	for k := range site.Pages {
		server.stats.pages = append(server.stats.pages, k)
	}

	log.Println("Registering HTTP route handlers...")
	server.http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://" + r.Host + r.RequestURI, http.StatusMovedPermanently)
	})

	log.Println("Registering HTTPS route handlers...")
	httpsHandler := mux.NewRouter()
	httpsHandler.Use(server.recordCompression)
	httpsHandler.Use(gziphandler.GzipHandler)
	httpsHandler.Use(server.recordRequest)
	httpsHandler.Use(server.trimWWWRequests)
	httpsHandler.Use(server.serveStaticFiles)
	for route := range site.Pages {
		file := path.Join(root, site.Pages[route])
		httpsHandler.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			if _, err := os.Stat(file); err != nil {
				http.ServeFile(w, r, path.Join(root, site.Errors.Internal))
			} else {
				http.ServeFile(w, r, file)
			}
		})
	}
	httpsHandler.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, path.Join(root, site.Errors.NotFound))
	})
	server.https.Handler = httpsHandler

	return &server, nil
}
