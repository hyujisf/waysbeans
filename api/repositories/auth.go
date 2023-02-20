package repositories

import (
	"fmt"
	"waysbeans/models"
)

type AuthRepository interface {
	Register(newUser models.User) (models.User, error)
	Login(email string) (models.User, error)
	IsUserRegistered(email string) bool
}

// menambahkan user baru
func (r *repository) Register(newUser models.User) (models.User, error) {
	err := r.db.Create(&newUser).Error

	return newUser, err
}

// mengambil salah satu data user dari database berdasarkan email
func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

// memeriksa apakah user sudah terdaftar
func (r *repository) IsUserRegistered(email string) bool {
	// validasi data user, jika email sudah terdaftar maka kirimkan true
	var user models.User
	errCekUser := r.db.First(&user, "email=?", email).Error

	fmt.Println(errCekUser)
	return errCekUser == nil
}
