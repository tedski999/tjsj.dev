package main

import (
	"log"; "fmt"; "errors"
	"os"; "os/signal"; "syscall"
	"github.com/tedski999/tjsj.dev/pkg/webserver"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
	log.Println("Goodbye!")
}

func run() error {
	if len(os.Args) != 5 { return errors.New("Usage: " + os.Args[0] + " <sitefile> <statfile> <certfile> <keyfile>") }

	errChan := make(chan error)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Starting web server serving site " + os.Args[1] + "...")
	server, err := webserver.Create(os.Args[1], os.Args[2])
	if err != nil { return err }
	server.Start(errChan, os.Args[3], os.Args[4])
	log.Println("Web server up and listening on :80 and :443")

	for {
		select {
		case err := <-errChan: log.Println("Error: " + err.Error())
		case <-sigChan: return server.Stop()
		}
	}
}
