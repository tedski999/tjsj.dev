package main

import (
	"os"
	"fmt"
	"log"
	"github.com/tedski999/tjsj.dev/pkg/sitegen"
)

func main() {

	// Ensure correct number of command-line arguments have been pased
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <template file> <destination>\n", os.Args[0])
		os.Exit(1)
	}

	// Generate site
	log.Printf("Generating site in %s from template file %s...\n", os.Args[2], os.Args[1])
	if err := sitegen.Generate(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}

	log.Println("Done!")
}
