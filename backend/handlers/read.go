package handlers

import (
	"fmt"
	"log"
	"net/http"
	"swe-project/backend/datamgr"

	"github.com/gorilla/mux"
)

func GetSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Received: ", vars["name"])
}

/*
 * GetSubjects: Returns a list of all subjects in the database as an array of JSON representations.
 * The body does not matter as it is not used.
 */

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	var subjects []datamgr.Subject
	datamgr.DB.Find(&subjects)
	log.Println("Request for list of subjects received.")
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

	// Refactored to use the limit feature of gorm instead of the jank if statement I made
	var reviews []datamgr.Review
	datamgr.DB.Limit(request.MaxReviews).Find(&reviews, "Subject = ?", request.Name)

	log.Println("Request for reviews of", request.Name, "received.")
	//log.Println(reviews)
	WriteResponse(w, reviews, 200)
}

/*
 * Get Reviews by list of subjects:
 * A requested function begrudgingly being implemented so that the front end
 * can request the reviews for multiple subjects at a time (they should really just make mmultiple
 * calls to the GetSubjectReview function. There is absolutley no reason for this to ever be used
 * and will likely be removed in later versions)
 * An example of a body is the following:
 * {
 *		SubjectNames[]: {name1, name2, name3}
 * }
 */

func GetReviewsBySubjects(w http.ResponseWriter, r *http.Request) {
	// Decode body into a workable object
	request := struct {
		Subjects []string
	}{}

	ReadRequest(w, r, &request)

	// Get list of reviews from DB
	var reviews []datamgr.Review
	datamgr.DB.Where("Subject IN ?", request.Subjects).Find(&reviews)

	// Write response and log request
	log.Println("Request for reviews of ", request.Subjects, " received")
	WriteResponse(w, reviews, 200)
}

/*
 * Get Reviews by Author
 * Gets a list of reviews by author so that when users look at an author's page they can see
 * their posts. This request expects a json body of the following:
 * {
 *		Author:	<string>
 * }
 */

func GetReviewsByAuthor(w http.ResponseWriter, r *http.Request) {
	// Decode body into a workable object
	request := struct {
		Author string
	}{}

	ReadRequest(w, r, &request)

	// Get list of reviews from DB
	var reviews []datamgr.Review
	datamgr.DB.Find(&reviews, "Author = ?", request.Author)

	log.Println("Request for reviews by", request.Author, "received.")
	WriteResponse(w, reviews, 200)
}
