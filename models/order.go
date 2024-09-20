package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	UserId    int     `json:"userId"`
	Total     float64 `json:"total"`
}

func NewOrder(productID, quantity int, total float64, userId int) *Order {
	return &Order{
		ProductID: productID,
		Quantity:  quantity,
		Total:     total,
		UserId:    userId,
	}
}
