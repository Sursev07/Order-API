package models

import (
	// "gorm.io/gorm"
)

type ItemOrder struct {
	CustomerName string `json:"customer_name"`
	// OrderId string `json:"order_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	OrderedAt string `json:"ordered_at"`
}