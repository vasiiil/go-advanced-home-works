package link

import (
	"api-project/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Db *db.Db
}

func NewRepository(_db *db.Db) *LinkRepository {
	return &LinkRepository{
		Db: _db,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Db.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Db.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Db.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Db.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Db.DB.Delete(&Link{}, id)
	return result.Error
}
