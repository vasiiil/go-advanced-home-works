package di

import "api-project/internal/user"

type IStatRepository interface {
	AddClick(linkId uint)
}

type IUserRepository interface {
	Create(u *user.User) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	GetById(id uint) (*user.User, error)
}
