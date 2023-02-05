/*
 * Written by Emmett Kogan
 * Last modified 2/1/23@11:19am
 * Basic JSON CRUD API Example to be built on
 * Followed:
 * 1. https://www.youtube.com/watch?v=lf_kiH_NPvM (general go project with gin/gorm)
 * 2. https://www.youtube.com/watch?v=wXEZZ2JT3-k (sqlite tutorial)
 */

package main

import (
	"fmt"

	"github.com/Emmett-Kogan/go-crud/initializers"
	"github.com/Emmett-Kogan/go-crud/models"
	"github.com/gin-gonic/gin"
)

/*
 * Iinitializes some constants like port
 */

func init() {
	initializers.LoadEnv()
	initializers.Connect_db()
}

func main() {

	fmt.Println("Hello")

	r := gin.Default()

	// Default port is 8080, port is currently set in .env to 3000

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/test", CreateTest)

	r.Run()
}

func CreateTest(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body string
	}

	c.Bind(&body)

	// Create test struct
	newTest := models.Test{Text: body.Body}
	result := initializers.DB.Create(&newTest)

	// I don't understand why yet but the test is not being pushed
	// to the database

	if result.Error != nil {
		fmt.Println(result.Error)
		fmt.Println("Error doing result?")
		c.Status(400)
		return
	}

	newTest.Text = "Server's response: " + newTest.Text

	// return to sender
	c.JSON(200, gin.H{
		"test": newTest,
	})

}
