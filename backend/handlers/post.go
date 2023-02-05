package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaneferrellwv/backend/initializers"
	"github.com/shaneferrellwv/backend/objects"
)

func Post(c *gin.Context) {
	// Get data off request body

	// Create test struct
	postTest := objects.User{
		Name:     "Shane",
		ID:       69,
		Password: "password123",
	}
	result := initializers.DB.Create(&postTest)
	if result.Error != nil {
		log.Fatal("Failed to create entry in database.")
		return
	}

	// Return to sender
	c.JSON(200, gin.H{
		"User": postTest,
	})
}
