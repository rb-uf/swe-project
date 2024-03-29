package handlers

import (
	"log"
	"net/http"

	"swe-project/backend/datamgr"
)

/*
 * CreateSubject: Creates a new subject object in the database, returns the new structure on success
 * Incoming packet body should look like:
 * {
 *	"Name": <string>
 * }
 */

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	// Get the json body off of the request and store in a struct
	var subject datamgr.Subject
	ReadRequest(w, r, &subject)

	if !CheckCookieAndPermissions(w, r, false, "", true) {
		return
	}

	// Create the new entry in the db
	result := datamgr.DB.Create(&subject)
	if result.Error != nil {
		log.Println("Failed to create entry in database")
		w.WriteHeader(400) // Return error code to client (just 400 for now)
		return
	}

	log.Println("Subject created:", subject.Name)
	WriteResponse(w, subject, 201)
}

/*
 * CreateReview: Creates a new subject object in the database, returns the new structure on success
 * Incoming packet body should look like:
 * {
    "Subject": "CSE",
    "Rating": 5,
    "Text": "The dungeon has good A/C",
    "Author": "Emmett",
    "AuthorID": 420
 * }
*/

func CreateReview(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Subject  string
		Rating   uint
		Text     string
		Author   string
		AuthorID int
	}{}

	ReadRequest(w, r, &request)

	var review datamgr.Review

	review.Subject = request.Subject
	review.Rating = request.Rating
	review.Text = request.Text
	review.Author = request.Author
	review.AuthorID = uint(request.AuthorID)
	review.Ups = 0

	if !CheckCookieAndPermissions(w, r, false, "", true) {
		return
	}

	// If subject specified does not exist, log it and return an error
	// TODO: maybe just have this create the subject as well
	result := datamgr.DB.First(&datamgr.Subject{}, "name = ?", review.Subject)
	if result.Error != nil {
		log.Println("Subject specified DNE")
		w.WriteHeader(400)
		return
	}

	// Otherwise the subject does exist and we can create the review and return the created object
	result = datamgr.DB.Create(&review)
	if result.Error != nil {
		log.Println("Failed to create DB entry")
		w.WriteHeader(400)
		return
	}

	log.Println("Review created for", review.Subject, "by AuthorID", review.AuthorID)
	WriteResponse(w, review, 201)
}
