package main

import (
	"links/configs"
	"links/pkg/db"
)

func main() {

	// инициализация конфигурации - пакет config
	configs.Config.LoadConfig()

	// инициализация коннеткора GORM
	db.DbConnector.Init(configs.Config.DB)

}
