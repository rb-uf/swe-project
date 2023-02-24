package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"swe-project/backend/datamgr"
	//"github.com/gorilla/mux"
)

func DeleteSubjecet(w http.ResponseWriter, r *http.Request) {
	// Temporary: using url param to specify ID to delete
	//params := mux.Vars(r)
	// ID := params["ID"]


	// Read object to delete
	w.Header().Set("Content-Type", "application/json")
	var request datamgr.Subject

	json.NewDecoder(r.Body).Decode(&request)

	// TODO: Ultimatley we should check the permissions of the requester (probably just admins)
	// note for furture 401 is unauthorized return code

	// Soft delete (just sets deleted_at field and keeps the entry in the db)
	// Check if entry exists, if it doesn't return a bad request
	var p datamgr.Subject
	datamgr.DB.Find(&p, request.ID)

	if (p.ID != request.ID){
		fmt.Println("Error deleting object: ", request.Name)
		w.WriteHeader(400)	// Bad request	
		return
	}

	datamgr.DB.Delete(&p)

	w.WriteHeader(200)	// OK
}