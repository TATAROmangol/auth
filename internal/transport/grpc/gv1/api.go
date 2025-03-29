package gv1

import (
	ssov1 "auth/pkg/grpc/auth"
	"context"

	"google.golang.org/grpc"
)

type Api struct {
	ssov1.UnimplementedAuthServer
	service Auther
}

func Register(gRPCServer *grpc.Server, service Auther) {
	ssov1.RegisterAuthServer(gRPCServer, &Api{service: service})
}

func (s *Api) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	return &ssov1.LoginResponse{}, nil
}

func (s *Api) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	return &ssov1.RegisterResponse{}, nil
}
