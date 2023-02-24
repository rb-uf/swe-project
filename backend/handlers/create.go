package handlers

import (
	"fmt"
	"net/http"

	"swe-project/backend/datamgr"
)

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	// Get the json body off of the request and store in a struct
	var subject datamgr.Subject
	ReadRequest(w, r, &subject)

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

	WriteResponse(w, subject, 201)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review datamgr.Review
	ReadRequest(w, r, &review)

	// If subject specified does not exist, log it and return an error
	// TODO: maybe just have this create the subject as well
	result := datamgr.DB.First(&datamgr.Subject{}, "name = ?", review.Subject)
	if result.Error != nil {
		fmt.Println("Subject specified DNE")
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

	WriteResponse(w, review, 201)
}
