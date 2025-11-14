package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone     string `json:"phone" gorm:"index"`
	SessionId string `json:"sessionId" gorm:"index"`
	Orders    []Order
}
