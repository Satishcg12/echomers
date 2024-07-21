package services

import (
	"errors"
	"strings"
	"time"

	"github.com/satishcg12/echomers/internal/repos"
	"github.com/satishcg12/echomers/internal/types"
)

type (
	AuthService interface {
		Login(email, password string) (string, error)
		Register(fullName, email, password string) error
		VerifyEmail(email, token string) error
		ResetPassword(email string) error
	}

	authService struct {
		repo repos.UserRepository
	}
)

func NewAuthService(repo repos.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(fullName, email, password string) error {
	if _, err := s.repo.FindByEmail(email); err != nil {
		return errors.New("email already exists")
	}
	userName := strings.Split(email, "@")[0]
	if _, err := s.repo.FindByUsername(userName); err != nil {
		userName = userName + time.Now().String()
	}

	user := types.Users{
		FullName: fullName,
		Email:    email,
		Password: password,
		Username: userName,
	}

	if _, err := s.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(email, password string) (string, error) {
	return "", nil

}

func (s *authService) VerifyEmail(email, token string) error {
	return nil
}

func (s *authService) ResetPassword(email string) error {
	return nil
}
