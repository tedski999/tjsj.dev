package webserver

import (
	"time"
	"log"
	"github.com/tedski999/tjsj.dev/pkg/htmlpages"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
	"github.com/tedski999/tjsj.dev/pkg/splashes"
)

// File watcher thread
func (srv *Server) watchFiles(errChan chan error) {
	err := srv.fileWatcher.Start(1 * time.Minute)
	if err != nil {
		errChan <-err
	}
}

// File event handler thread
func (srv *Server) handleFileEvents(errChan chan error) {
	for {
		select {
			case event := <-srv.fileWatcher.Event:
				// TODO: only reload a top-level directory which changed
				log.Println(event)
				srv.pages = htmlpages.Load(pagesDir)
				srv.content = webcontent.Load(contentDir)
				srv.splashes = splashes.Load(splashesFilePath)
			case err := <-srv.fileWatcher.Error:
				errChan <-err
			case <-srv.fileWatcher.Closed:
				return
		}
	}
}
