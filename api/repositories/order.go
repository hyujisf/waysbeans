package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	CreateOrder(Order models.Order) (models.Order, error)
	UpdatesOrder(Order []models.Order) ([]models.Order, error)
	UpdateOrder(Order models.Order) (models.Order, error)
	DeleteOrder(Order models.Order) (models.Order, error)
	CreateTransactionID(transaction models.Transaction) (models.Transaction, error)
	FindProductID(ProductID []int) ([]models.Product, error)
	FindOrdersTransaction(TrxID int) ([]models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

// mengambil semua order
func (r *repository) FindOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Product").Find(&orders).Error

	return orders, err
}

// mengambil 1 order berdasarkan id

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Product").First(&order, ID).Error

	return order, err
}

// menambahkan order baru
func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Preload("Product").Create(&order).Error

	return order, err
}

// mengupdate order
func (r *repository) UpdatesOrder(order []models.Order) ([]models.Order, error) {
	err := r.db.Save(&order).Error

	return order, err
}

func (r *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Save(&order).Error

	return order, err
}

// menghapus order

func (r *repository) DeleteOrder(order models.Order) (models.Order, error) {
	err := r.db.Delete(&order).Error

	return order, err
}

func (r *repository) CreateTransactionID(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindProductID([]int) ([]models.Product, error) {
	var product []models.Product
	err := r.db.Find(&product).Error

	return product, err
}

func (r *repository) FindOrdersTransaction(TrxID int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Product").Find(&orders, "user_id = ? AND status = ?", TrxID, "on").Error

	return orders, err
}
