package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewCustomer(name, code string) *Customer {
	return &Customer{
		Name: name,
		Code: code,
	}
}
