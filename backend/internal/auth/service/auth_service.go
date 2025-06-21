package service

import (
	"errors"

	"github.com/vnkot/piklnk/internal/auth/domain"
)

const (
	ErrUserExists      = "user exists"
	ErrWrongCredetials = "wrong email or password"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type AuthService struct {
	UserRepository UserRepository
}

func NewAuthService(userRepository UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Login(email, password string) (uint, error) {
	user, _ := service.UserRepository.FindByEmail(email)
	if user == nil {
		return 0, errors.New(ErrWrongCredetials)
	}

	if !user.IsValidPassword(password) {
		return 0, errors.New(ErrWrongCredetials)
	}

	return user.ID, nil
}

func (service *AuthService) Register(email, password, name string) (uint, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return 0, errors.New(ErrUserExists)
	}

	user := &domain.User{
		Email: email,
		Name:  name,
	}

	err := user.SetPassword(password)
	if err != nil {
		return 0, err
	}

	user, err = service.UserRepository.Create(user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
