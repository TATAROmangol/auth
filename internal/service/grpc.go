package service

import (
	ssov1 "auth/pkg/grpc/auth"
	"context"
)

type Repo interface{
	CheckLogin(string) (bool, error)
	CreateUser(string, string) (int, error)
	CheckUser(string, string) (int, error)
}

type Service struct{
	ctx context.Context
	repo Repo
}

func NewService(ctx context.Context, repo Repo) *Service{
	return &Service{ctx, repo}
}

func (s *Service) CreateUser(log, pas string) (int, error){
	return s.repo.CreateUser(log, pas)
}
func (h *Service) Register(context.Context, *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error){
	return &ssov1.RegisterResponse{}, nil
}