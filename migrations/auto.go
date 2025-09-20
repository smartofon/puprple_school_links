package main

import (
	"links/configs"
	"links/internal/cart/models"
	"links/internal/user"
	"links/pkg/db"
)

func main() {
	// инициализация конфигурации - пакет config
	configs.Config.LoadConfig()

	// инициализация коннеткора GORM
	db.DbConnector.Init(configs.Config.DB)

	// автомиграция
	db.DbConnector.AutoMigrate(&models.Product{})
	db.DbConnector.AutoMigrate(&user.User{})
}
