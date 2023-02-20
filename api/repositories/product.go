package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	AddProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product, ID int) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("status = ?", "active").Find(&products).Error
	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Where("status = ?", "active").First(&product, ID).Error
	return product, err
}

func (r *repository) AddProduct(newProduct models.Product) (models.Product, error) {
	err := r.db.Create(&newProduct).Error
	return newProduct, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) DeleteProduct(product models.Product, ID int) (models.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
