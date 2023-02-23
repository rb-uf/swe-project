package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

// MasterHandler: calls the rest of the handler functions
func MasterHandler(r *mux.Router) {

	// serve frontend
	var frontend_path = "../angular-front-end/dist/angular-front-end/"
	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir(frontend_path))))

	r.HandleFunc("/get-test", GetTest)
	r.HandleFunc("/post-test", PostTest)


	r.HandleFunc("/create-subject", CreateSubject).
	  Methods("POST")
}
