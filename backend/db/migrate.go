package main

import (
	"swe-project/backend/initializers"
	"swe-project/backend/objects"
)

func main() {
	initializers.LoadEnv()
	initializers.Connect_db()
	initializers.DB.AutoMigrate(&objects.User{})
}
