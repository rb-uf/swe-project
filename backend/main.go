package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaneferrellwv/backend/handlers"
	"github.com/shaneferrellwv/backend/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect_db()
}

func main() {
	fmt.Println("tet")

	r := gin.Default()

	handlers.MasterHandler(r)

	r.Run()

}
