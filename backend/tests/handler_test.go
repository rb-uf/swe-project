package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"swe-project/backend/datamgr"
	"swe-project/backend/handlers"
	"testing"
)

/*
 * TestGetSubjects: tests GetSubjects handler function. Compares the http packet body
 * and code to data retrieved from the database. If there is a mismatch or the return code
 * is not 200, it fails. Otherwise passes.
 */
func TestGetSubjects(t *testing.T) {
	datamgr.ConnectDB("../datamgr/database.db")

	req, err := http.NewRequest("GET", "/get-subjects", nil)
	if err != nil {
		fmt.Println("err not nil")
		t.Fatal(err)
		return
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetSubjects)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			recorder.Code, http.StatusOK)
	}

	// Compare body directly to status of subjects table in db
	var subjects []datamgr.Subject
	datamgr.DB.Find(&subjects)

	raw, err := json.Marshal(subjects)
	if err != nil {
		t.Errorf("Converting the correct output failed?")
	}

	want := string(raw)
	got := recorder.Body.String()

	if want != got {
		t.Error("HTTP packet and database objects do not match")
	}
}
