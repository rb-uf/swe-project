package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MasterHandler: calls the rest of the handler functions
func MasterHandler(r *mux.Router, frontendPath string) {
	// Routes for user requests
	r.HandleFunc("/sign-up", CreateUser).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("DELETE")

	// Handle "subject" requests
	r.HandleFunc("/create-subject", CreateSubject).Methods("POST")
	r.HandleFunc("/get-subject/{name}", GetSubject).Methods("GET")
	r.HandleFunc("/get-subjects", GetSubjects).Methods("GET")
	r.HandleFunc("/delete-subject", DeleteSubject).Methods("DELETE")

	// Handle "review" requests
	r.HandleFunc("/create-review", CreateReview).Methods("POST")
	r.HandleFunc("/get-subject-reviews", GetSubjectReviews).Methods("GET")
	r.HandleFunc("/get-reviews-by-subjects", GetReviewsBySubjects).Methods("GET")
	r.HandleFunc("/get-reviews-by-author", GetReviewsByAuthor).Methods("GET")
	r.HandleFunc("/delete-review", DeleteReview).Methods("DELETE")
	r.HandleFunc("/update-review", UpdateReview).Methods("PUT")

	// Serve frontend files
	// (Essentially if nothing else matches, try matching frontend files)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(frontendPath)))
}
