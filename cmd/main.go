package main

import (
	"fmt"
	"links/configs"
	"links/internal/auth"
	"links/internal/cart"
	"links/internal/user"
	"links/pkg/db"
	"links/pkg/jwt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func main() {

	// инициализация конфигурации - пакет config
	configs.Config.LoadConfig()

	// инициализация коннеткора GORM
	db.DbConnector.Init(configs.Config.DB)

	repo := cart.NewCartRepository(db.DbConnector)
	u := user.NewUserRepository(db.DbConnector)
	j := jwt.NewJWT(configs.Config.Secret)
	s := auth.NewAuthService(u, j)

	handler := &cart.CartHandler{
		CartRepository: repo,
		UserRepository: u,
		AuthHService:   s,
	}

	// Путь к директории для логов (можно заменить на нужный)
	logDir := "logs"

	// Создаем директорию, если она не существует
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Не удалось создать директорию для логов: %v", err))
	}

	// Путь к файлу логов
	logFilePath := filepath.Join(logDir, "app.log")

	// Открываем файл для записи (добавление в конец)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Не удалось открыть файл логов: %v", err))
	}
	defer file.Close()

	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{ // Используем текстовый формат
		FullTimestamp: true,
	})

	router := http.NewServeMux()
	cart.NewCartHandler(router, handler)

	server := http.Server{
		Addr:    ":8081",
		Handler: cart.LoggingMiddleware(router),
	}

	server.ListenAndServe()

}
