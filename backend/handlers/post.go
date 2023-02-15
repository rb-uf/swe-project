package handlers

import (
	"log"
	"net/http"

	"swe-project/backend/initializers"
	"swe-project/backend/objects"
)

func PostTest(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("PostTest\n"))
}
