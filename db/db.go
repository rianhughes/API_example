// db/db.go
package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := "user:password@tcp(localhost:5432)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
		return nil, err
	}

	fmt.Println("Database connection successfully opened")

	return db, nil
}
