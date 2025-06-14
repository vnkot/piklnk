package auth

import (
	"errors"

	"github.com/vnkot/piklnk/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Login(email, password string) (*uint, error) {
	user, _ := service.UserRepository.FindByEmail(email)
	if user == nil {
		return nil, errors.New(ErrWrongCredetials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New(ErrWrongCredetials)
	}

	return &user.ID, nil
}

func (service *AuthService) Register(email, password, name string) (*uint, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return nil, errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}
