package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string `gorm:"index" json:"name"`
	Price      uint   `json:"price"`
	Picture    string `json:"picture"`
	CategoryID uint   `json:"categoryId"`
}
