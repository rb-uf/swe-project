package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"swe-project/backend/handlers"
	"swe-project/backend/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect_db()
}

func main() {
	fmt.Println("Starting swe-project/backend server.")

	r := gin.Default()

	handlers.MasterHandler(r)

	r.Run()

}
