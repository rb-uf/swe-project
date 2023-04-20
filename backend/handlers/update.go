package handlers

import (
	"log"
	"net/http"
	"swe-project/backend/datamgr"
)

/*
 * A user should be able to edit their post's text. I think the location should not be editable
 * Input should be review ID, and new text string to replace the old text field
 * Request format:
 * {
 * 		ID	uint
 *		Text string
 * }
 * and the text of the review with ID == ID is replaced with the text in the packet
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

/*
 * A user should be able to agree or disagree with someone elses review, so they can increment or
 * decrement a public "Ups" counter. The body of the request should be as follows:
 * {
 * 		"ReviewID": <ID number>
 *		"Up:"		<negative number for down, positive or 0 for up>
 * }
 * Returns the updated review object
 */

func UpdateReviewUps(w http.ResponseWriter, r *http.Request) {
	request := struct {
		ReviewID int
		Up       int
	}{}

	ReadRequest(w, r, &request)

	var review datamgr.Review
	datamgr.DB.First(&review, "ID = ?", request.ReviewID)

	if review.ID != uint(request.ReviewID) {
		log.Println("Entry not found:", request.ReviewID)
		w.WriteHeader(400)
		return
	}

	var diff int
	if request.Up >= 0 {
		diff = 1
	} else {
		diff = -1
	}

	review.Ups += diff

	datamgr.DB.Model(&review).Update("Ups", review.Ups)

	WriteResponse(w, review, 200)

}
