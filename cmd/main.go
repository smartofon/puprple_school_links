package main

import (
	"links/configs"
	"links/internal/cart"
	"links/pkg/db"
	"net/http"
)

func main() {

	// инициализация конфигурации - пакет config
	configs.Config.LoadConfig()

	// инициализация коннеткора GORM
	db.DbConnector.Init(configs.Config.DB)

	repo := cart.NewCartRepository(db.DbConnector)

	handler := &cart.CartHandler{
		CartRepository: repo,
	}

	router := http.NewServeMux()
	cart.NewCartHandler(router, handler)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()

}
