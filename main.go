package main

import (
	"log"
	"os"

	"github.com/ryszhio/goauth/run"
)

func main() {
	if err := run.InitializeApp(); err != nil {
		log.Fatalln("Failed to start service.")
		log.Fatalf("ERROR: %s", err.Error())
		os.Exit(1)
	}
}
