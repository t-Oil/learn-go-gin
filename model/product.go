package model

import (
	database "learn/go-gin/config"
	"time"
)

type Product struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	IsActive  bool      `gorm:"DEFAULT:1" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ProductRequest struct {
	Title string  `json:"title" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func GetProducts() ([]*Product, error) {
	var products []*Product

	if err := database.Database.Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func FindProduct(id int) (Product, error) {
	var product Product

	if err := database.Database.Find(&product, id).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

func StoreProduct(request *ProductRequest) (err error) {
	product := new(Product)
	product.Title = request.Title
	product.Price = request.Price
	if err := database.Database.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id int, request *ProductRequest) (product Product, err error) {
	if err := database.Database.Where("id = ?", id).Take(&product).UpdateColumns(map[string]interface{}{
		"title":      request.Title,
		"price":      request.Price,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return product, err
	}

	return product, nil
}
