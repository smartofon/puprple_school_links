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
		p.Price = product.Price
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
