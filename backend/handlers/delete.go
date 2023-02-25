package handlers

import (
	"fmt"
	"net/http"
	"swe-project/backend/datamgr"
)

/*
 * DeleteSubject: Deletes specified subject by the request packet. For now this has no authentication
 * for who can delete what but when users and account levels are implemented later this will be added to
 */

func DeleteSubjecet(w http.ResponseWriter, r *http.Request) {
	// Read object to delete
	var request datamgr.Subject
	ReadRequest(w, r, &request)

	// TODO: Ultimatley we should check the permissions of the requester (probably just admins)
	// note for furture 401 is unauthorized return code

	// Soft delete (just sets deleted_at field and keeps the entry in the db)
	// Check if entry exists, if it doesn't return a bad request
	var p datamgr.Subject
	datamgr.DB.Find(&p, request.ID)

	if p.ID != request.ID {
		fmt.Println("Error deleting object: ", request.Name)
		w.WriteHeader(400) // Bad request
		return
	}

	datamgr.DB.Delete(&p)

	w.WriteHeader(200) // OK
}
