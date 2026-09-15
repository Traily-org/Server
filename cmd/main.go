package main

import (
	"github.com/charmbracelet/log"

	"github.com/traily-org/server/internal/adapters/http"
	"github.com/traily-org/server/internal/config"
	"github.com/traily-org/server/internal/domain/greeting"
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
