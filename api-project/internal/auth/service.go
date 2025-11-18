package auth

import (
	"api-project/internal/user"
	"api-project/pkg/di"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Login(email, password string) (string, error) {
	existsUser, _ := service.UserRepository.GetByEmail(email)
	if existsUser == nil {
		return "", errors.New(ErrWrongCreadentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCreadentials)
	}
	return existsUser.Email, nil
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existsUser, _ := service.UserRepository.GetByEmail(email)
	if existsUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	newUser := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}
	_, err = service.UserRepository.Create(newUser)
	if err != nil {
		return "", err
	}

	return newUser.Email, nil
}
