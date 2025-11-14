package product

import (
	"api-orders/internal/models"
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

func (repo *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	result := repo.Db.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) GetAll(page, pageSize uint) []models.Product {
	var products []models.Product
	result := repo.Db.Scopes(db.PaginateScope(page, pageSize)).Find(&products)
	if result.Error != nil {
		return []models.Product{}
	}

	return products
}

func (repo *ProductRepository) GetById(id uint) (*models.Product, error) {
	var product models.Product
	result := repo.Db.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) Update(product *models.Product) (*models.Product, error) {
	result := repo.Db.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Db.DB.Delete(&models.Product{}, id)
	return result.Error
}
