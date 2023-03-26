package datamgr

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDB(file string) {
	var err error

	DB, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database")
	}

	DB.AutoMigrate(&User{}, &Subject{}, &Review{})
}
