package service

import (
	"fmt"
)

//go:generate go run github.com/vektra/mockery/v2@v2.53.3

type Repo interface{
	CheckLogin(string) (bool, error)
	CreateUser(string, string) (int, error)
	CheckPassword(string, string) (int, error)
}

type JWT interface{
	GenerateToken(int) (string, error)
}

type Service struct{
	repo Repo
	jwt JWT
}

func NewService(repo Repo, jwt JWT) *Service{
	return &Service{repo, jwt}
}

func (s *Service) Register(log, pas string) (string, error){
	exist, err := s.repo.CheckLogin(log)
	if err != nil{
		return "", fmt.Errorf("failed in db: %v", err)
	}
	if exist{
		return "", fmt.Errorf("login already taken")
	}

	id, err := s.repo.CreateUser(log, pas)
	if err != nil{
		return "", fmt.Errorf("failed in db: %v", err)
	}

	token, err := s.jwt.GenerateToken(id)
	if err != nil{
		return "", fmt.Errorf("failed in jwt: %v", err)
	}

	return token, nil
}

func (s *Service) Login(log, pas string) (string, error){
	exist, err := s.repo.CheckLogin(log)
	if err != nil{
		return "", fmt.Errorf("failed in db: %v", err)
	}
	if exist{
		return "", fmt.Errorf("login already taken")
	}

	id, err := s.repo.CheckPassword(log, pas)
	if err != nil{
		return "", fmt.Errorf("failed in db: %v", err)
	}

	token, err := s.jwt.GenerateToken(id)
	if err != nil{
		return "", fmt.Errorf("failed in jwt: %v", err)
	}

	return token, nil
}
