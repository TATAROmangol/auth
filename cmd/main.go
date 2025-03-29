package main

import (
	"auth/internal/config"
	"auth/internal/service"
	"auth/internal/transport/grpc/gv1"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"context"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	l := logger.New()

	db, err := postgres.NewConnect(cfg.PG)
	if err != nil{
		l.ErrorContext(ctx, "failed to connect postgres", "err", err)
	}

	service := service.NewService(ctx, db)
	server := gv1.New(ctx, cfg.GRPC, l, service)
	server.Run()
}
