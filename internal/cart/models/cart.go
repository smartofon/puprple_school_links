package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductId   int            `json:"uid" gorm:"iniqueIndex"`
	Name        string         `json:"name" gorm:"required"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price" gorm:"not null"`
}