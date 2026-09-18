package service

import (
	"errors"
	"mypos-backend/internal/model"
	"mypos-backend/internal/repository"
	"mypos-backend/pkg/token"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.User, string, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", errors.New("username atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, "", errors.New("username atau password salah")
	}

	jwtToken, err := token.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, jwtToken, nil
}
