package gv1

import (
	ssov1 "auth/pkg/grpc/auth"
	"context"

	"google.golang.org/grpc"
)

//go:generate mockery --all --output=./mocks 

type Auth interface {
	Register(string, string) (string, error)
	Login(string, string) (string, error)
}

type Api struct {
	ssov1.UnimplementedAuthServer
	service Auth
}

func Register(gRPCServer *grpc.Server, service Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &Api{service: service})
}

func (s *Api) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.TokenResponse, error) {
	return &ssov1.TokenResponse{}, nil
}

func (s *Api) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.TokenResponse, error) {
	return &ssov1.TokenResponse{}, nil
}
