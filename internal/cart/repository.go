package cart

import (
	"errors"
	"links/internal/cart/models"
	"links/pkg/db"
)

type CartRepository struct {
	Database *db.Db
}

func NewCartRepository(database *db.Db) *CartRepository {
	return &CartRepository{
		Database: database,
	}
}

func (repo *CartRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	tx := repo.Database.DB.First(&product, "product_id=?", product.ProductId)
	if tx.Error == nil {
		return nil, errors.New("product is exists")
	}
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *CartRepository) UpdateProduct(product *models.Product) (*models.Product, error) {
	p := models.Product{}
	tx := repo.Database.DB.Where("product_id=?", product.ProductId).First(&p)
	if tx.Error == nil {
		p.Name = product.Name
		p.Description = product.Description
		p.Price = product.Price
		p.Images = product.Images
		repo.Database.DB.Save(&p)
		return &p, nil
	}
	return nil, errors.New("Error update")
}

func (repo *CartRepository) GetProduct(product *models.Product) (*models.Product, error) {
	tx := repo.Database.DB.Where("product_id=?", product.ProductId).First(&product)
	if tx.Error != nil {
		return nil, errors.New("product not exists")
	}
	return product, nil
}

func (repo *CartRepository) DeleteProduct(product *models.Product) (*models.Product, error) {
	tx := repo.Database.DB.Where("product_id=?", product.ProductId).First(&product)
	if tx.Error != nil {
		return nil, errors.New("product not exists")
	}
	repo.Database.DB.Where("product_id=?", product.ProductId).Delete(&product)
	return product, nil
}

func (repo *CartRepository) CreateOrder(products []ProductOrder, userId int) (*models.Order, error) {

	tx := repo.Database.DB.Begin()

	order := models.Order{
		UserID: uint(userId),
	}

	tx.Save(&order)

	for _, v := range products {
		product_order := models.OrderProduct{
			ProductID: uint(v.ProductId),
			OrderID:   order.ID,
			Price:     v.Price,
			Quantity:  v.Quantity,
		}
		tx.Save(&product_order)
	}

	tx.Commit()

	return &order, nil
}

func (repo *CartRepository) GetOrder(orderID, userId int) (*models.Order, error) {

	var order models.Order

	tx := repo.Database.DB.Where("ID=? and user_id=?", orderID, userId).First(&order)
	if tx.Error != nil {
		return nil, errors.New("order not exists")
	}

	var order_products []models.OrderProduct

	repo.Database.Table("order_products").Where("order_id=?", order.ID).Scan(&order_products)
	order.Products = order_products

	return &order, nil
}

func (repo *CartRepository) OrderList(userId int) (*[]models.Order, error) {

	var orders []models.Order

	repo.Database.Table("orders").
		Where("user_id=?", userId).
		Limit(100).
		Order("id desc").
		Scan(&orders)

	return &orders, nil
}
