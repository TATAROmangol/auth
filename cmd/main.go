package main

import (
	"auth/internal/config"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"context"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()
	ctx = logger.ImportInContext(ctx)

	db, err := postgres.NewConnect(cfg.PG)
	if err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, "failed to connect postgres", "err", err)
	}

	
}
