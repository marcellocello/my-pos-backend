package service

import (
	"errors"
	"mypos-backend/internal/model"
	"mypos-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(req *model.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
		Role:         req.Role,
	}

	return s.userRepo.Create(user)
}

func (s *UserService) GetByUsername(username string) (*model.User, error) {
	return s.userRepo.GetByUsername(username)
}

func (s *UserService) ChangePassword(req *model.ChangePasswordRequest) error {
	user, err := s.userRepo.GetByID(req.ID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword))
	if err != nil {
		return errors.New("password lama salah")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(req.ID, string(newHashedPassword))
}
