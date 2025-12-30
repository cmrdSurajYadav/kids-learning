package main

import (
	"log"

	"github.com/cmrdSurajYadav/auth-service/internal/config"
	"github.com/cmrdSurajYadav/auth-service/internal/server"
)

func main() {
	cfg := config.LoadConfig()   // load config from .env or defaults
	srv := server.NewServer(cfg) // create server instance
	log.Println("Starting server on", cfg.AppPort)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
