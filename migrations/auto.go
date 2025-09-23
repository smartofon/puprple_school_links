package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"links/configs"
	"links/internal/cart/models"
	"links/pkg/db"
	"os"
)

type JsonProducts []struct {
	ID          int     `json:id`
	Price       float64 `json:price`
	Name        string  `json:name`
	Description string  `json:description`
}

func main() {
	// инициализация конфигурации - пакет config
	configs.Config.LoadConfig()

	// инициализация коннеткора GORM
	db.DbConnector.Init(configs.Config.DB)

	// автомиграция
	db.DbConnector.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderProduct{})
	Seeder(db.DbConnector)
}

func Seeder(db_connector *db.Db) {

	jsonFile, err := os.Open("products.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var products JsonProducts

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &products)

	fmt.Println(products)

	for _, v := range products {
		product := models.Product{
			ProductId:   v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		}
		db_connector.Save(&product)
	}

}
