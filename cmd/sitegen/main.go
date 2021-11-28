package main

import (
	"os"
	"errors"
	"log"; "fmt"
	"github.com/tedski999/tjsj.dev/pkg/sitegen"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
	log.Println("Done!")
}

func run() error {
	if len(os.Args) != 3 { return errors.New("Usage: " + os.Args[0] + " <template file> <destination>") }
	log.Printf("Generating site in %s from template file %s...\n", os.Args[2], os.Args[1])
	return sitegen.Generate(os.Args[1], os.Args[2])
}
