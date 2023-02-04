package database

import (
	"github.com/princeprabhat/GormApi/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	db.AutoMigrate(&models.Product{})

	return db
}
