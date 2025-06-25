package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE-APP: ", log.LstdFlags|log.Lshortfile)
	srv := server.New(logger)
	err := srv.Start()
	if err != nil {
		logger.Fatalf("FATAL: Server failed to start: %v", err)
	}
}
