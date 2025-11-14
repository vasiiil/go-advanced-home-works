package order

import (
	"api-orders/internal/models"
	"api-orders/pkg/db"
	"fmt"
	"time"
)

type OrderRepository struct {
	Db *db.Db
}

func NewRepository(_db *db.Db) *OrderRepository {
	return &OrderRepository{
		Db: _db,
	}
}

func (repo *OrderRepository) Create(userId uint, productIds []uint) (*models.Order, error) {
	tx := repo.Db.DB.Begin()
	defer tx.Rollback()

	var user models.User
	if err := tx.First(&user, userId).Error; err != nil {
		return nil, fmt.Errorf("пользователь с ID %d не найден: %w", userId, err)
	}

	var products []models.Product
	if err := tx.Where("id IN ?", productIds).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("ошибка при поиске продуктов: %w", err)
	}
	if len(products) != len(productIds) {
		return nil, fmt.Errorf("не все указанные продукты были найдены")
	}

	newOrder := models.Order{
		UserID:    userId,
		OrderDate: time.Now(),
		Products:  products,
	}
	if err := tx.Create(&newOrder).Error; err != nil {
		return nil, fmt.Errorf("ошибка при создании заказа: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("ошибка при коммите транзакции: %w", err)
	}

	return &newOrder, nil
}

func (repo *OrderRepository) GetByUserId(userId uint, page, pageSize uint) ([]models.Order, error) {
	var orders []models.Order
	err := repo.Db.DB.Scopes(db.PaginateScope(page, pageSize)).
		Model(&models.Order{}).
		Order("order_date DESC").
		Where("user_id = ?", userId).
		Preload("Products").
		Find(&orders).
		Error

	if err != nil {
		return nil, fmt.Errorf("ошибка при получении заказов: %w", err)
	}

	return orders, nil
}

func (repo *OrderRepository) GetById(userId, id uint) (*models.Order, error) {
	var order models.Order
	result := repo.Db.DB.
		Preload("Products").
		First(&order, "id = ? AND user_id = ?", id, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}
