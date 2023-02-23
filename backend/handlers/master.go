package handlers

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
)

// MasterHandler: calls the rest of the handler functions
func MasterHandler(r *mux.Router) {
	r.HandleFunc("/get-test", GetTest)
	r.HandleFunc("/post-test", PostTest)

	r.HandleFunc("/create-subject", CreateSubject).
	  Methods("POST")

	r.PathPrefix("/").
	  Handler(http.StripPrefix("/", http.FileServer(http.Dir(os.Getenv("FRONTEND_PATH")))))
}
