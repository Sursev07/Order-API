package models

import (
	"time"
)

type Item struct {
	OrderId int `json:"order_id"`
	ItemCode string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type ItemOrder struct {
	CustomerName string `json:"customer_name"`
	Items []Item
}


type Order struct {
	ID int `gorm:"default:uuid_generate_v3()"`
	CustomerName string `json:"customer_name" binding:"required"`
	OrderedAt time.Time `json:"ordered_at" binding:"required"`
	Items []Item `gorm:"foreignKey:OrderId"`
}