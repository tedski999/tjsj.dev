package webserver

import (
	"net/http"
	"context"
)

// Start the server on a separate goroutine
func (server *Server) Start(errChan chan<- error) {
	server.errChan = errChan

	server.doneWG.Add(1)
	go func () {
		defer server.doneWG.Done()
		err := server.http.ListenAndServeTLS(server.certFilePath, server.keyFilePath)
		if err != http.ErrServerClosed {
			server.errChan <- err
		}
	}()
}

// Gracefully shutdown the server
func (server *Server) Stop() error {
	err := server.http.Shutdown(context.Background())
	server.doneWG.Wait()
	return err
}
