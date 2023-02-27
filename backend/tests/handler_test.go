package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"swe-project/backend/datamgr"
	"swe-project/backend/handlers"
	"testing"
)

/*
 * To run tests properly and not use cached output, use `go test ./ -v -count=1`
 * Where ./ means at this directory, -v for verbose output, and -count=1 to prevent cache use
 */

/*==================== Create Tests ====================*/

/*
 * TestCreateSubject: Tests the CreateSubject handler function by calling it and checking
 * the returned subject and new database entry against expected output. (Also calls DeleteSubject
 * to return the database to its original state)
 */

func TestCreateSubject(t *testing.T) {
	// Open temporary test databse
	datamgr.ConnectDB("temp.db")

	// Create http request to test with
	body := datamgr.Subject{
		Name: "Test Subject",
	}

	raw_test, err := json.Marshal(body)
	if err != nil {
		t.Error("Test body failed JSON conversion")
		return
	}

	req, err := http.NewRequest("POST", "/create-subject", bytes.NewBuffer(raw_test))
	if err != nil {
		t.Error("Issue creating http request")
		return
	}

	// Declare a recorder to get http response from the handler function
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateSubject)

	handler.ServeHTTP(recorder, req)

	// If return code is not a success, we fail
	if recorder.Code != 201 {
		t.Errorf("Handler returned %v, wanted %v", recorder.Code, 201)
	}

	// Get created object from database and check returned object against it
	var subject datamgr.Subject
	datamgr.DB.Find(&subject) // There will only ever be one in this test

	// TODO: maybe clean this up later
	want, _ := json.Marshal(subject)
	got := recorder.Body.String()

	if string(want) != got {
		t.Error("DB entry and packet don't match, something weird occured")
	}

	var output datamgr.Subject
	json.NewDecoder(recorder.Body).Decode(&output)

	if output.ID != 1 || output.Name != "Test Subject" {
		t.Error("Returned object does not match expected output")
	}

	// Close and delete temporary database
	temp, _ := datamgr.DB.DB()
	temp.Close()

	error := os.Remove("temp.db")
	if error != nil {
		t.Error("Failed to remove temporary db")
	}
}

/*
 * TestCreateReview: Tests CreateReview handler function by creating a review for the first subject
 * in the database. Checks the returned review object and new database entry against expected output.
 */

func TestCreateReview(t *testing.T) {
	// Create temporary db file
	datamgr.ConnectDB("temp.db")

	NewSubject := datamgr.Subject{
		Name: "Carelton",
	}

	NewReview := datamgr.Review{
		Subject:  "Carelton",
		Rating:   5,
		Text:     "This is a text string",
		Author:   "Dobra Rocks!",
		AuthorID: 1234,
	}

	datamgr.DB.Create(&NewSubject)

	body, _ := json.Marshal(NewReview)

	req, _ := http.NewRequest("POST", "/create-review", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateReview)

	handler.ServeHTTP(recorder, req)

	var review datamgr.Review
	datamgr.DB.First(&review)
	want, _ := json.Marshal(review)

	if string(want) != recorder.Body.String() {
		t.Error("Database entry and returned object do not match")
	}

	var got datamgr.Review
	json.NewDecoder(recorder.Body).Decode(&got)

	// Just checking text field but if this isn't malformed there's no reason the
	// rest of the non gorm.Model fields should differ
	if got.Text != NewReview.Text {
		t.Error("Bodies of expected output and received do not match")
	}

	// Close and delete temporary database
	temp, _ := datamgr.DB.DB()
	temp.Close()

	error := os.Remove("temp.db")
	if error != nil {
		t.Error("Failed to remove temporary db")
	}
}

/*===================== Read Tests =====================*/

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

/*
 * TestGetSubjectReviews: Tests GetSubjectReviews handler function by comparing the output against
 * the state of the database.
 */

func TestGetSubjectReviews(t *testing.T) {

}

/*==================== Update Tests ====================*/

/*
 * TestUpdateReview: Tests creating and editing a subject by calling CreateReview() and UpdateReview()
 * and checking if the objecet returned from UpdateReview() is the same as the expected output. Cleans up
 * by calling DeleteSubject()
 */

func TestUpdateReview(t *testing.T) {

}

/*==================== Delete Tests ====================*/

/*
 * TestDeleteSubject: Tests creating and deleting a subject by calling CreateSubject() followed by
 * DeleteSubject() and checking if the state of the database matches the initial state
 */

func TestDeleteSubject(t *testing.T) {

}

/*================== Functional Tests ==================*/
