package user

import (
	"api-project/pkg/db"
)

type UserRepository struct {
	Db *db.Db
}

func NewRepository(_db *db.Db) *UserRepository {
	return &UserRepository{
		Db: _db,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Db.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := repo.Db.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) GetById(id uint) (*User, error) {
	var user User
	result := repo.Db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
