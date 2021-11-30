package main

import (
	"log"; "fmt"
	"os"; "os/signal"; "syscall"
	"github.com/tedski999/tjsj.dev/pkg/webserver"
)

func main() {

	// Ensure correct number of command-line arguments have been pased
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s <sitefile> <statfile> <certfile> <keyfile>\n", os.Args[0])
		os.Exit(1)
	}

	// Create web server
	log.Printf("Starting web server serving site %s...\n", os.Args[1])
	server, err := webserver.Create(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize web server:\n%v", err)
		os.Exit(1)
	}

	// Start web server
	errChan := make(chan error)
	server.Start(errChan, os.Args[3], os.Args[4])
	log.Println("Web server up and listening on :80 and :443")

	// Print exceptions, exit on signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			log.Println("Error: " + err.Error())
		case <-sigChan:
			server.Stop()
			os.Exit(0)
		}
	}
}
