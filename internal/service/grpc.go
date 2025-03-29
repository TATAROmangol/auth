package service

import (
	ssov1 "auth/pkg/grpc/auth"
	"context"
	"database/sql"
)

type Service struct{
	ctx context.Context
	db *sql.DB
}

func NewService(ctx context.Context, db *sql.DB) *Service{
	return &Service{ctx, db}
}

func (h *Service) AddUser(log, pas string) (int, error){
	return 0, nil
}
func (h *Service) Register(context.Context, *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error){
	return &ssov1.RegisterResponse{}, nil
}