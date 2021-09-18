package models

import (
	// "gorm.io/gorm"
)

type Item struct {
	OrderId int `json:"order_id"`
	ItemCode string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}