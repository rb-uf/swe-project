package datamgr

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

// Adding a cookie tracking data structure for now to keep track of active sessions

var CookieJar []http.Cookie
