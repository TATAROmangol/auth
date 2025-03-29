package main

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/transport/grpc/gv1"
	"auth/pkg/logger"
	"auth/pkg/migrator"
	"auth/pkg/postgres"
	"context"
	"os"
)

const (
	migrationPath = "file://internal/repository/migrations"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	l := logger.New()

	db, err := postgres.NewConnect(cfg.PG)
	if err != nil{
		l.ErrorContext(ctx, "failed to connect postgres", "err", err)
		os.Exit(1)
	}

	m, err := migrator.New(migrationPath, cfg.PG)
	if err != nil{
		l.ErrorContext(ctx, "failed in create migrator", "error", err.Error())
		os.Exit(1)
	}
	l.InfoContext(ctx, "migrator loaded")

	if err := m.Up(); err != nil{
		l.ErrorContext(ctx, "failed in up migrate", "error", err.Error())
		os.Exit(1)
	}
	l.InfoContext(ctx, "migrations complete")

	repo := repository.New(db)

	service := service.NewService(ctx, repo)
	server := gv1.New(ctx, cfg.GRPC, l, service)
	server.Run()
}
