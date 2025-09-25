package database

import (
	// "fmt"
	// "errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var (
// 	ErrCantFindProduct = errors.New("No product found")
// )

func NewConnection()(*gorm.DB, error){
	
	dsn  := os.Getenv("DATABASE_URL")

	db,err :=gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err!=nil {
		return db,err
	}
    fmt.Println("Connected to the database")
	return db,nil
}