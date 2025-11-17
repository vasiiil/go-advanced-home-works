package user

import (
	"api-orders/internal/models"
	"api-orders/pkg/db"
)

type UserRepository struct {
	Db *db.Db
}

func NewRepository(_db *db.Db) *UserRepository {
	return &UserRepository{
		Db: _db,
	}
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	result := repo.Db.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	result := repo.Db.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) GetBySessionId(sessionId string) (*models.User, error) {
	var user models.User
	result := repo.Db.DB.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	result := repo.Db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
