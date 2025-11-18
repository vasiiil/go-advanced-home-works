package auth_test

import (
	"api-project/internal/auth"
	"api-project/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error)     {
	return u, nil
}
func (repo *MockUserRepository) GetByEmail(email string) (*user.User, error) {
	return nil, nil
}
func (repo *MockUserRepository) GetById(id uint) (*user.User, error)         {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initEmail = "a@a.ru"
	authService := auth.NewService(&MockUserRepository{})
	email, err := authService.Register(initEmail, "123", "Test V")
	if err != nil {
		t.Fatal(err)
	}
	if email != initEmail {
		t.Fatalf("Email %s does not match init email %s", email, initEmail)
	}
}
