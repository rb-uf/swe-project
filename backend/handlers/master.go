package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// MasterHandler: calls the rest of the handler functions
func MasterHandler(r *mux.Router) {
	// r.HandleFunc("/get-test", GetTest)
	// r.HandleFunc("/post-test", PostTest)

	r.HandleFunc("/create-subject", CreateSubject).
		Methods("POST")

	r.HandleFunc("/get-subject/{name}", GetSubject).Methods("GET")
	r.HandleFunc("/get-subjects", GetSubjects).Methods("GET")

	// If nothing else matches, try matching frontend files.
	r.PathPrefix("/").
		Handler(http.FileServer(http.Dir(os.Getenv("FRONTEND_PATH"))))
}
