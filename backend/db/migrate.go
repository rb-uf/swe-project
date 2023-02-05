package main

import (
	"github.com/shaneferrellwv/backend/initializers"
	"github.com/shaneferrellwv/backend/objects"
)

func main() {
	initializers.LoadEnv()
	initializers.Connect_db()
	initializers.DB.AutoMigrate(&objects.User{})
}
