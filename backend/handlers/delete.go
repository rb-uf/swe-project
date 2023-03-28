package handlers

import (
	"fmt"
	"net/http"
	"swe-project/backend/datamgr"
)

/*
 * DeleteSubject: Deletes specified subject by the request packet. For now this has no authentication
 * for who can delete what but when users and account levels are implemented later this will be added to
 */

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	// Read object to delete
	var request datamgr.Subject
	ReadRequest(w, r, &request)

	// Probably a better way to do this
	// Check the requests cookie against cookies stored in cookie jar
	c, _ := r.Cookie("rater-gator user cookie")
	present, user := VerifyCookie(*c)

	// If cookie was not issued return unauthorized
	if !present {
		fmt.Println("Error, request's cookie not found in cookiejar")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get user info from db
	var temp datamgr.User
	datamgr.DB.Find(&temp, "Name = ?", user)

	// If user is not an admin do not let them delete the subject
	if !temp.Admin {
		fmt.Println("Error, requester does not have permission to delete a subject")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Soft delete (just sets deleted_at field and keeps the entry in the db)
	// Check if entry exists, if it doesn't return a bad request
	var p datamgr.Subject
	datamgr.DB.Find(&p, request.ID)
	if p.ID != request.ID {
		fmt.Println("Error deleting object: ", request.Name)
		w.WriteHeader(400) // Bad request
		return
	}

	datamgr.DB.Delete(&p)

	fmt.Println("Subject deleted:", request.Name)
	w.WriteHeader(200) // OK
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	// Read request body
	var request datamgr.Review
	ReadRequest(w, r, &request)

	// Probably a better way to do this
	// Check the requests cookie against cookies stored in cookie jar
	c, _ := r.Cookie("rater-gator user cookie")
	present, user := VerifyCookie(*c)

	// If cookie was not issued return unauthorized
	if !present {
		fmt.Println("Error, request's cookie not found in cookiejar")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get user info from db
	var temp datamgr.User
	datamgr.DB.Find(&temp, "Name = ?", user)

	// If user is not an admin nor the author do not let them delete the subject
	if !temp.Admin || user != request.Author {
		fmt.Println("Error, requester does not have permission to delete a subject")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Soft delete the entry in the database
	var p datamgr.Review
	datamgr.DB.Find(&p, request.ID)
	if p.ID != request.ID {
		fmt.Println("Object not found: ", request.ID)
		w.WriteHeader(400)
		return
	}

	datamgr.DB.Delete(&p)

	fmt.Println("Review deleted")
	w.WriteHeader(http.StatusOK)
}
