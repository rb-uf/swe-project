package handlers

import (
	"encoding/json"
	"log"
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

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	// Get all subjects
	var subjects []datamgr.Subject
	datamgr.DB.Find(&subjects)

	// Create JSON version of subjects to write back to requester
	response, err := json.Marshal(subjects)

	if err != nil {
		log.Fatal("Subjects failed to be converted to JSON")
		w.WriteHeader(424)	// Failed dependency
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)

}
