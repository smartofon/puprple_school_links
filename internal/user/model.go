package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone       string `json:"phone" gorm:"required"`
	SessionId   string `json:"session" gorm:"required"`
	ConfirmCode string `json:"confirm_code"`
}
