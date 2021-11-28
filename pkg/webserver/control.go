package webserver

import (
	"log"
	"time"
	"context"
	"net/http"
)

func (server *Server) Start(errChan chan<- error, certfile, keyfile string) {
	startChan := make(chan bool)

	go func() {
		log.Println("Starting HTTP server...")
		defer server.httpWG.Done()
		server.httpWG.Add(1)
		startChan <- true
		err := server.http.ListenAndServe()
		if err != nil && err != http.ErrServerClosed  { errChan <- err }
	}()
	go func() {
		log.Println("Starting HTTPS server...")
		server.httpsWG.Add(1)
		defer server.httpsWG.Done()
		startChan <- true
		err := server.https.ListenAndServeTLS(certfile, keyfile)
		if err != nil && err != http.ErrServerClosed { errChan <- err }
	}()

	<-startChan
	<-startChan
}

func (server *Server) Stop() error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	log.Println("Stopping HTTP server...")
	if err := server.http.Shutdown(ctx); err != nil { return err }
	log.Println("Stopping HTTPS server...")
	if err := server.https.Shutdown(ctx); err != nil { return err }
	server.httpWG.Wait()
	server.httpsWG.Wait()
	return nil
}
