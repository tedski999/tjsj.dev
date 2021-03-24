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

	return <- errChan
}

// Gracefully shutdown the server
func (srv *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	return srv.httpServer.Shutdown(ctx);
}
