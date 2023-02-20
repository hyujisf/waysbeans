package repositories

import "waysbeans/models"

type OrderRepository interface {
	FindOrders(UserID int) ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	GetOrderByProduct(ProductID int, UserID int) (models.Order, error)
	CreateOrder(newOrder models.Order) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(order models.Order) (models.Order, error)
}

// mengambil semua order
func (r *repository) FindOrders(UserID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Where("user_id = ?", UserID).Where("transaction_id IS NULL").Find(&order).Error
	return order, err
}

// mengambil 1 order berdasarkan id
func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Product").Where("transaction_id IS NULL").First(&order, "id = ?", ID).Error
	return order, err
}

// mengambil 1 order berdasarkan id product
func (r *repository) GetOrderByProduct(ProductID int, UserID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Product").Where("user_id = ?", UserID).Where("transaction_id IS NULL").First(&order, "product_id = ?", ProductID).Error
	return order, err
}

// menambahkan order baru
func (r *repository) CreateOrder(newOrder models.Order) (models.Order, error) {
	err := r.db.Select("ProductID", "OrderQty", "UserID").Create(&newOrder).Error
	return newOrder, err
}

// mengupdate order tertentu berdasarkan id
func (r *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Model(&order).Updates(order).Error
	return order, err
}

// menghapus order
func (r *repository) DeleteOrder(order models.Order) (models.Order, error) {
	err := r.db.Delete(&order).Error
	return order, err
}
