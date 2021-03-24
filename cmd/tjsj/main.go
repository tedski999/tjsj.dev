package main

import (
	"log"
	"os"
	"os/signal"
	"github.com/tedski999/tjsj.dev/pkg/webserver"
)

func main() {

	// Register signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	// TODO: tui

	// Create a new server
	log.Println("Creating web server...")
	srv := webserver.Create()

	// Start the server on a new thread
	log.Println("Starting web server...")
	go func() {
		err := srv.Start()
		if err != nil {
			log.Fatal("An error occurred while running the web server:\n" + err.Error())
		}
	}()

	// Gracefully exit if a signal is received
	log.Println("Web server listening on port 443")
	<- sigs
	log.Println("Stopping web server...")
	srv.Stop()

	// TODO: A signal to call something like srv.Reload() instead
	// This would reload certs, content and templates without closing connections and such
	log.Println("Goodbye!")
}
