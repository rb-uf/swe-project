package datamgr

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_FILE")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database")
	}

	DB.AutoMigrate(&User{}, &Subject{}, &Review{})
}
