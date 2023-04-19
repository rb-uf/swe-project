package handlers

import (
	"log"
	"net/http"
	"swe-project/backend/datamgr"
)

/*
 * DeleteSubject: Deletes specified subject by the request packet.
 * Body should contain a subject object
 */

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	// Read object to delete
	var request datamgr.Subject
	ReadRequest(w, r, &request)

	if !CheckCookieAndPermissions(w, r, false, "", false) {
		return
	}

	// Soft delete (just sets deleted_at field and keeps the entry in the db)
	// Check if entry exists, if it doesn't return a bad request
	var p datamgr.Subject
	datamgr.DB.Find(&p, request.ID)
	if p.ID != request.ID {
		log.Println("Error deleting object: ", request.Name)
		w.WriteHeader(400) // Bad request
		return
	}

	datamgr.DB.Delete(&p)

	log.Println("Subject deleted:", request.Name)
	w.WriteHeader(200) // OK
}

/*
 * DeleteReview: Deletes a specified review by subject by the request packet. Verifies the
 * user if built with authentication using the cookie.
 * Body should contain a review object
 */

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	// Read request body
	var request datamgr.Review
	ReadRequest(w, r, &request)

	// Probably a better way to do this
	// Check the requests cookie against cookies stored in cookie jar
	if !CheckCookieAndPermissions(w, r, true, request.Author, false) {
		return
	}

	// Soft delete the entry in the database
	var p datamgr.Review
	datamgr.DB.Find(&p, request.ID)
	if p.ID != request.ID {
		log.Println("Object not found: ", request.ID)
		w.WriteHeader(400)
		return
	}

	datamgr.DB.Delete(&p)

	log.Println("Review deleted")
	w.WriteHeader(http.StatusOK)
}
