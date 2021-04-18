package main

import (
	"log"
	"os"
	"os/signal"
	"github.com/tedski999/tjsj.dev/pkg/webserver"
	"github.com/tedski999/tjsj.dev/pkg/webcontent"
)

func main() {
	var err error

	// Create a new content manager
	log.Println("Creating content manager...")
	var content *webcontent.Content
	content, err = webcontent.Create("./web/templates/", "./web/posts/", "./web/splashes.txt")
	if err != nil {
		log.Println("An error occurred while creating the content manager:\n" + err.Error())
		return
	}

	// Create a new web server
	log.Println("Creating web server...")
	var server *webserver.Server
	server, err = webserver.Create(content)
	if err != nil {
		log.Println("An error occurred while creating the web server:\n" + err.Error())
		return
	}

	// Setup channels for signals and goroutine errors
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	exitChan := make(chan bool, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	// Start the content manager and web server
	log.Println("Starting content manager...")
	content.Start(errChan)
	log.Println("Starting web server...")
	server.Start(errChan)
	log.Println("Web server listening on :443")

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
	log.Println("Stopping content manager...")
	content.Stop()
	log.Println("Stopping web server...")
	server.Stop()
	log.Println("Goodbye!")
}
