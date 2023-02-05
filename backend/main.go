package main

import (
	"fmt"

	"github.com/Emmett-Kogan/swe-project/backend/handlers"
	"github.com/Emmett-Kogan/swe-project/backend/initializers"
	"github.com/gin-gonic/gin"
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
