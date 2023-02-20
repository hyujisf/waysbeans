package repositories

import "waysbeans/models"

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	FindTransactionsByUser(UserId int) ([]models.Transaction, error)
	GetTransaction(ID string) (models.Transaction, error)
	CreateTransaction(newTransaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, trxId string) (models.Transaction, error)
	UpdateTokenTransaction(token string, trxId string) (models.Transaction, error)
}

// mengambil semua transaksi
func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Order").Preload("Order.Product").Order("order_date desc").Find(&transaction).Error

	return transaction, err
}

// mengambil semua transaksi berdasarkan user tertentu
func (r *repository) FindTransactionsByUser(UserId int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Order").Preload("Order.Product").Where("user_id = ?", UserId).Order("order_date desc").Find(&transaction).Error

	return transaction, err
}

// mengambil 1 transaksi berdasarkan id
func (r *repository) GetTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Order").Preload("Order.Product").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

// menambahkan transaksi baru
func (r *repository) CreateTransaction(newTransaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&newTransaction).Error

	return newTransaction, err
}

// mengupdate status transaksi berdasarkan id
func (r *repository) UpdateTransaction(status string, trxId string) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Order").Preload("Order.Product").First(&transaction, "id = ?", trxId)

	// If is different & Status is "success" decrement available quota on data trip
	if status != transaction.Status && status == "success" {
		for _, ordr := range transaction.Order {
			var product models.Product
			r.db.First(&product, ordr.Product.ID)
			product.Stock = product.Stock - ordr.OrderQty
			r.db.Model(&product).Updates(product)
		}
	}

	// If is different & Status is "reject" decrement available quota on data trip
	if status != transaction.Status && status == "rejected" {
		for _, ordr := range transaction.Order {
			var product models.Product
			r.db.First(&product, ordr.Product.ID)
			product.Stock = product.Stock + ordr.OrderQty
			r.db.Model(&product).Updates(product)
		}
	}

	// change transaction status
	transaction.Status = status

	// fmt.Println(status)
	// fmt.Println(transaction.Status)
	// fmt.Println(transaction.ID)

	err := r.db.Model(&transaction).Updates(transaction).Error

	return transaction, err
}

// mengupdate token midtrans pada transaksi tertentu berdasarkan id
func (r *repository) UpdateTokenTransaction(token string, trxId string) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Order").Preload("Order.Product").First(&transaction, "id = ?", trxId)

	// change transaction token
	transaction.MidtransID = token
	err := r.db.Model(&transaction).Updates(transaction).Error

	return transaction, err
}

// menghapus transaksi
func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
