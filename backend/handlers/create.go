package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"swe-project/backend/datamgr"
)

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	// Get the json body off of the request and store in a struct
	w.Header().Set("Content-Type", "application/json")
	var subject datamgr.Subject
	json.NewDecoder(r.Body).Decode(&subject)

	// Log the request in the console
	fmt.Println("CreateSubject called: created ", subject.Name)

	// Create the new entry in the db
	result := datamgr.DB.Create(&subject)

	// If it fails to create, return an error code
	if result.Error != nil {
		log.Fatal("Failed to create entry in database")

		// Return error code to whomever made the request, for now I'm just going to do 400
		w.WriteHeader(400)
		return

	} 

	// Create JSON version of subject object
	response, err := json.Marshal(subject)
	if err != nil {
		log.Fatal("Failed to convert subject into JSON")

		w.WriteHeader(406)	// Not acceptable
	}

	// Return to the requester a JSON of the created object on our end
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)	// Created status code
	w.Write(response)


}