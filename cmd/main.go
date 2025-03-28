package main

import (
	"auth/internal/config"
	"auth/pkg/postgres"
	"context"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	db, err := postgres.NewConnect(cfg.PG)

}
