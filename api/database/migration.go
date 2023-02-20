package database

import (
	"fmt"
	"waysbeans/models"
	"waysbeans/pkg/sql"
)

func RunMigration() {
	err := sql.DB.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
