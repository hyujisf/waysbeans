package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"waysbeans/models"
	"waysbeans/pkg/bcrypt"
	"waysbeans/pkg/sql"

	"gorm.io/gorm"
)

func RunSeeder() {
	// ==================================
	// CREATE SUPER ADMIN ON MIGRATION
	// ==================================

	// cek is user table exist
	if sql.DB.Migrator().HasTable(&models.User{}) {
		// check is user table has minimum 1 user as admin
		err := sql.DB.First(&models.User{}, "role = ?", "admin").Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create 1 admin
			newUser := models.User{
				Name:  "Admin",
				Role:  "admin",
				Email: os.Getenv("ADMIN_EMAIL"),
				Image: "https://api.dicebear.com/5.x/thumbs/svg?seed=" + strconv.Itoa(int(time.Now().Unix())),
			}

			hashPassword, err := bcrypt.HashingPassword(os.Getenv("ADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Hash password failed")
			}

			newUser.Password = hashPassword

			// insert admin to database
			errAddUser := sql.DB.Select("Name", "Role", "Email", "Password", "Image").Create(&newUser).Error
			if errAddUser != nil {
				fmt.Println(errAddUser.Error())
				log.Fatal("Seeding failed")
			}
		}
	}

	fmt.Println("Seeding completed successfully")
}
