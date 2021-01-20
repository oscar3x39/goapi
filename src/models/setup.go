package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDataBase return true
func ConnectDataBase() {

	dsn := "homestead:secret@tcp(127.0.0.1:4306)/homestead?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}