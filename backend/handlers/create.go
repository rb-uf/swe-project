package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"

	"swe-project/backend/datamgr"
)

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	// Get body of the request
	// (it should contain a JSON description of the "subject" to create).
	// Convert JSON description into a Subject struct
	var newSubject datamgr.Subject
	json.NewDecoder(r.Body).Decode(&newSubject)
	fmt.Printf("New subject received. \"Name\": %s", newSubject.Name)
	r.Body.Close()

	// Store new subject in database.
	result := datamgr.DB.Create(&newSubject)
	if result.Error != nil {
		log.Fatal("Failed to create entry in database.")
		return
	}

	// Respond to sender.
	w.Write([]byte("Success.\n"))
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	// Get data off request body

	// Create test struct
	postTest := datamgr.User{
		Name:     "Shane",
		ID:       69,
		Password: "password123",
	}
	result := datamgr.DB.Create(&postTest)
	if result.Error != nil {
		log.Fatal("Failed to create entry in database.")
		return
	}

	// Return to sender
	w.Write([]byte("PostTest\n"))
}
