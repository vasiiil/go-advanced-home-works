package product

import (
	"api-orders/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Db *db.Db
}

func NewRepository(_db *db.Db) *ProductRepository {
	return &ProductRepository{
		Db: _db,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Db.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) GetAll(page, pageSize int) ([]Product, error) {
	var products []Product
	result := repo.Db.Scopes(db.PaginateScope(page, pageSize)).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.Db.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Db.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Db.DB.Delete(&Product{}, id)
	return result.Error
}
