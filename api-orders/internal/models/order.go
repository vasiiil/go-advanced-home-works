package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderDate time.Time `json:"orderDate"`
	UserID    uint      `json:"userId"`
	Products  []Product `json:"products" gorm:"many2many:order_products;"`
}
