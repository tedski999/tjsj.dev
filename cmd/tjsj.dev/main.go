package main

import (
	"log"
	"os"
	"os/signal"
	"github.com/tedski999/tjsj.dev/pkg/webserver"
	"github.com/tedski999/tjsj.dev/pkg/webstats"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
)

func main() {

	// Start/Resume web server statistics
	stats, err := webstats.Create("./data/stats.bin")

	// Load all web content
	log.Println("Loading web content...")
	content, err := webcontent.Create("./web/static/", "./web/templates/", "./web/posts/", "./web/splashes.txt")
	if err != nil {
		log.Println("An error occurred while loading web content:\n" + err.Error())
		return
	}

	// Create a new web server
	log.Println("Creating web server...")
	server, err := webserver.Create(content, stats, "./web/certs/fullchain.pem", "./web/certs/privkey.pem")
	if err != nil {
		log.Println("An error occurred while creating the web server:\n" + err.Error())
		return
	}

	// Setup channels for signals and goroutine errors
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	exitChan := make(chan bool, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	// Start the web server
	log.Println("Starting web server...")
	stats.Start(errChan)
	server.Start(errChan)
	log.Println("Web server listening on :443")

	// Exception handling
	go func() {
		for {
			// Wait for either a signal or an error
			select {
			case sig := <-sigChan:
				log.Println("Received " + sig.String())
			case err := <-errChan:
				log.Println("An error occurred while running the web server:\n" + err.Error())
			}

			// Let main start to exit
			exitChan <- true
		}
	}()

	// Attempt to gracefully exit if an exception occurs
	<-exitChan
	log.Println("Stopping web server...")
	server.Stop()
	stats.Stop()
	log.Println("Goodbye!")
}
