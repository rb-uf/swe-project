package handlers

import (
	"fmt"
	"net/http"
	"swe-project/backend/datamgr"

	"github.com/gorilla/mux"
)

func GetTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Received: ", vars["name"])
}

/*
 * GetSubjects: Returns a list of all subjects in the database as an array of JSON representations.
 * The body does not matter as it is not used.
 */

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	// Get all subjects
	var subjects []datamgr.Subject
	datamgr.DB.Find(&subjects)

	WriteResponse(w, subjects, 200)
}

/*
 * GetSubjectReviews: This function returns reviews related to the specified subject. Returning an array
 * of reviews of max length as specified by the http packet. The body should look like:
 * {
 * 		SubjectName: <string>
 * 		MaxReviews:	 <int>
 * }
 */

func GetSubjectReviews(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Name       string
		MaxReviews int
	}{}
	ReadRequest(w, r, &request)

	var reviews []datamgr.Review
	datamgr.DB.Find(&reviews, "Subject = ?", request.Name)

	// TODO: make this more sophisticated so that it takes like, the 10 most liked reviews or the 10 most recent/least recent
	if len(reviews) > request.MaxReviews {
		reviews = reviews[0:request.MaxReviews]
	}

	WriteResponse(w, reviews, 200)
}
