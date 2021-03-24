package webserver

import (
	"net/http"
	"time"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
)

type Server struct {
	httpServer *http.Server
	content *webcontent.Content
}

func Create() *Server {

	// Configure server
	httpServer := &http.Server {
		Addr: ":https",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Load web content for server
	content := webcontent.Create()
	content.LoadHTMLTemplates(templateFilePaths)
	content.LoadSplashes(splashesFilePath)

	// Setup server
	srv := &Server {
		httpServer: httpServer,
		content: content,
	}
	srv.registerHandlers()

	return srv
}
