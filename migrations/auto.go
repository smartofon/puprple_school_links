package main

import (
	"links/configs"
	"links/internal/models"
	"links/pkg/db"
)

func main() {
	// инициализация конфигурации - пакет config
	configs.Config.LoadConfig()

	// инициализация коннеткора GORM
	db.DbConnector.Init(configs.Config.DB)

	// автомиграция
	db.DbConnector.AutoMigrate(&models.Product{})
}
