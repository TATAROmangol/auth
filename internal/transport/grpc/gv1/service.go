package gv1

import (
	ssov1 "auth/pkg/grpc/auth"
	"context"
)

type Service struct{}

func NewService() Service{
	return Service{}
}

func (h *Service) Login(context.Context, *ssov1.LoginRequest) (*ssov1.LoginResponse, error){
	return &ssov1.LoginResponse{}, nil
}
func (h *Service) Register(context.Context, *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error){
	return &ssov1.RegisterResponse{}, nil
}