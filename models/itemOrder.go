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
	// OrderId string `json:"order_id"`
	// ItemCode string `json:"item_code"`
	// Description string `json:"description"`
	// Quantity int `json:"quantity"`
	// OrderedAt string `json:"ordered_at"`
}