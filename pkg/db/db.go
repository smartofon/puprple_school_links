package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

type DbConfig struct {
	Dsn string
}

// объявляем глобальную переменную пакета,
// которая будет содержать коннектор  для базы данных и будет дооступна из других пакетов
// но ее нужно будет инициализировать в пакете main

var DbConnector *Db

func (dbconnector *Db) Init(conf DbConfig) {
	database, err := gorm.Open(postgres.Open(conf.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DbConnector = &Db{database}
}
