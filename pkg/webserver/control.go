package webserver

import (
	"net/http"
	"context"
	"errors"
	"time"
)

// Start the server
func (srv *Server) Start() error {
	errChan := make(chan error)

	// Main server thread
	go func() {
		err := srv.httpServer.ListenAndServeTLS(certFilePath, keyFilePath)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	// TODO: email server, etc...

	// File watcher threads
	go srv.watchFiles(errChan)
	go srv.handleFileEvents(errChan)

	// TODO: main.go error handling just bails, it should instead
	// at least attempt to shutdown the server gracefully
	return <-errChan
}

// Gracefully shutdown the server
func (srv *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	srv.fileWatcher.Close() // TODO: this blocks until watchFiles wakes up
	// TODO: display details on why we cant shutdown yet
	return srv.httpServer.Shutdown(ctx);
}
