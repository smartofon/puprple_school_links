package configs

import (
	"links/pkg/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB db.DbConfig
}

// глобальный журнал натсроек приложения
var Config AppConfig

// функция загрузчик конфигурации приложения
func (conf *AppConfig) LoadConfig() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error: Не удалось загрузить файл конфигурации: %v", err)
		os.Setenv("config_init", "0")
	} else {
		os.Setenv("config_init", "1")
	}

	conf.DB.Dsn = os.Getenv("DSN")
}
