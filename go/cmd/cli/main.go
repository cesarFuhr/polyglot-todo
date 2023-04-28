package main

import (
	"flag"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	flag.Parse()

	return nil
}
