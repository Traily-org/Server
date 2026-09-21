package main

import (
	"context"

	"github.com/charmbracelet/log"

	"github.com/traily-org/server/internal/config"
	"github.com/traily-org/server/internal/infrastructure/postgres"
	"github.com/traily-org/server/migrations"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	pool, err := postgres.Connect(ctx, cfg.DB.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	if err := migrations.Run(ctx, pool); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Info("migrations applied")
}
