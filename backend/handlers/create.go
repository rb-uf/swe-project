package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Failed to create entry in database")

		// Return error code to whomever made the request, for now I'm just going to do 400
		w.WriteHeader(400)
		return

	} 

	// Create JSON version of subject object
	response, err := json.Marshal(subject)
	if err != nil {
		fmt.Println("Failed to convert subject into JSON")

		w.WriteHeader(406)	// Not acceptable
	}

	// Return to the requester a JSON of the created object on our end
	w.WriteHeader(201)	// Created status code
	w.Write(response)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var review datamgr.Review
	json.NewDecoder(r.Body).Decode(&review)

	// If subject specified does not exist, log it and return an error
	result := datamgr.DB.First(&datamgr.Subject{}, "name = ?", review.Subject)
	if result.Error != nil {
		fmt.Println("Subjcet specified DNE")
		w.WriteHeader(400)
		return
	}

	// Otherwise the subject does exist and we can create the review and return the created object
	result = datamgr.DB.Create(&review)

	if result.Error != nil {
		fmt.Println("Failed to create DB entry")
		w.WriteHeader(400)
		return
	}

	response, err := json.Marshal(review)
	if err != nil {
		fmt.Println("Failed to convert subject into JSON")
		w.WriteHeader(406)
	}

	w.WriteHeader(200)
	w.Write(response)
}