package handlers

import (
	"log"
	"net/http"
	"swe-project/backend/datamgr"
)

/*
 * A user should be able to edit their post's text. I think the location should not be editable
 * Input should be review ID, and new text string to replace the old text field
 */
func UpdateReview(w http.ResponseWriter, r *http.Request) {
	// Read body
	request := struct {
		ID      uint
		NewText string
	}{}
	ReadRequest(w, r, &request)

	var entry datamgr.Review
	datamgr.DB.First(&entry, "ID = ?", request.ID)

	// Verify requester
	if !CheckCookieAndPermissions(w, r, true, entry.Author, false) {
		return
	}

	if entry.ID != request.ID {
		log.Println("Entry not found: ", request.ID)
		w.WriteHeader(400)
		return
	}

	// Update db entry
	datamgr.DB.Model(&entry).Update("text", request.NewText)

	// Return the new object to indicate success
	WriteResponse(w, entry, 200)
}
