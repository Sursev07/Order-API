package models

import (
	// "time"
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