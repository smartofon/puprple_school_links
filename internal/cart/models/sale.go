package models

import (
	"links/internal/user"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	user.User
	Orders []Order
}

type OrderProduct struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Price     float64
	Quantity  int
}

type Order struct {
	gorm.Model
	UserID   uint           `gorm:"not null"`                                            // Внешний ключ к пользователю
	Products []OrderProduct `gorm:"foreignkey:OrderID;association_foreignkey:ProductID"` // Заказ связан с продуктом
}

type Product struct {
	gorm.Model
	ProductId   int            `json:"uid" gorm:"required"`
	Name        string         `json:"name" gorm:"size:255"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Description string         `json:"description"`
	Orders      []OrderProduct `gorm:"foreignkey:ProductID;association_foreignkey:OrderID"` // Заказ связан с продуктом
}
