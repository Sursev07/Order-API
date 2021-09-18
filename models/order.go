package models

import (
	// "gorm.io/gorm"
	"time"
)

type Order struct {
	ID int `gorm:"default:uuid_generate_v3()"`
	CustomerName string `json:"customer_name" binding:"required"`
	OrderedAt time.Time `json:"ordered_at" binding:"required"`
}