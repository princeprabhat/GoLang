package database

import (
	models "github.com/princeprabhat/Authentication/Models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	db.AutoMigrate(&models.User{})

	return db
}
