package main

import (
	"github.com/charmbracelet/log"

	"traily/backend/internal/adapters/http"
	"traily/backend/internal/config"
	"traily/backend/internal/domain/greeting"
)

func main() {
	cfg := config.Load()

	greetingService := greeting.NewService()
	greetingHandler := http.NewGreetingHandler(greetingService)
	server := http.NewServer(greetingHandler)

	if err := server.Run(cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
