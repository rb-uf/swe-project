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

func TestMain(m *testing.M) {
	// Init
	datamgr.ConnectDB("temp.db")

	// Create some subjects
	var subjects []datamgr.Subject
	subjects = append(subjects, datamgr.Subject{Name: "1"})
	subjects = append(subjects, datamgr.Subject{Name: "2"})
	subjects = append(subjects, datamgr.Subject{Name: "3"})
	subjects = append(subjects, datamgr.Subject{Name: "4"})
	subjects = append(subjects, datamgr.Subject{Name: "5"})
	datamgr.DB.Create(&subjects)

	// Create some reviews
	var reviews []datamgr.Review
	reviews = append(reviews, datamgr.Review{Subject: "1", Rating: 5, Text: "Test text1", Author: "Emmett", AuthorID: 420})
	reviews = append(reviews, datamgr.Review{Subject: "1", Rating: 5, Text: "Test text2", Author: "Emmett", AuthorID: 420})
	reviews = append(reviews, datamgr.Review{Subject: "1", Rating: 5, Text: "Test text3", Author: "Emmett", AuthorID: 420})
	reviews = append(reviews, datamgr.Review{Subject: "1", Rating: 5, Text: "Test text4", Author: "Emmett", AuthorID: 420})
	reviews = append(reviews, datamgr.Review{Subject: "1", Rating: 5, Text: "Test text5", Author: "Emmett", AuthorID: 420})
	datamgr.DB.Create(&reviews)
	// Run tests
	m.Run()

	// Cleanup
	temp, _ := datamgr.DB.DB()
	temp.Close()
	error := os.Remove("temp.db")
	if error != nil {
		fmt.Println("Failed to remove temporary db: ", error)
	}
}

// TODO: make this handle nil correctly
func ExecuteRequest(body interface{}, packet_type string, route string, handler_func func(w http.ResponseWriter, r *http.Request), code int, t *testing.T) *bytes.Buffer {
	// Build http request
	raw, err := json.Marshal(body)
	if err != nil {
		t.Error("Failed to convert body to JSON")
		return nil
	}

	req, err := http.NewRequest(packet_type, route, bytes.NewBuffer(raw))
	if err != nil {
		t.Error("Failed to create http request")
		return nil
	}

	// Set up a recorder to read the response and serve the packet
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handler_func)
	handler.ServeHTTP(recorder, req)

	// If code doesn't match then we fail the test
	if recorder.Code != code {
		t.Errorf("Received response code %v, expected %v", recorder.Code, code)
		return nil
	}

	// Returns body of recorder to be validated by caller
	return recorder.Body
}

/*==================== Create Tests ====================*/

/*
 * TestCreateSubject: Tests the CreateSubject handler function by calling it and checking
 * the returned subject and new database entry against expected output.
 */

func TestCreateSubject(t *testing.T) {
	subject := datamgr.Subject{Name: "Test"}

	body := ExecuteRequest(subject, "POST", "/create-subject", handlers.CreateSubject, 201, t)

	// Make sure returned object has the correct Name
	var output datamgr.Subject
	json.NewDecoder(body).Decode(&output)

	if output.Name != "Test" {
		t.Error("Returned object does not match expected output")
	}
}

/*
 * TestCreateReview: Tests CreateReview handler function by creating a review for the first subject
 * in the database. Checks the returned review object and new database entry against expected output.
 */

func TestCreateReview(t *testing.T) {
	NewReview := datamgr.Review{
		Subject:  "1",
		Rating:   5,
		Text:     "This is a text string",
		Author:   "Dobra Rocks!",
		AuthorID: 1234,
	}

	body := ExecuteRequest(NewReview, "POST", "/create-review", handlers.CreateReview, 201, t)

	// Verify the object returned matches the db entry created and both match expected output

	var got datamgr.Review
	json.NewDecoder(body).Decode(&got)

	// Just checking text field but if this isn't malformed there's no reason the
	// rest of the non gorm.Model fields should differ
	if got.Text != NewReview.Text {
		t.Error("Bodies of expected output and received do not match")
	}
}

/*===================== Read Tests =====================*/

/*
 * TestGetSubjects: tests GetSubjects handler function. Compares the http packet body
 * and code to data retrieved from the database. If there is a mismatch or the return code
 * is not 200, it fails. Otherwise passes.
 */

func TestGetSubjects(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-subjects", nil)
	if err != nil {
		fmt.Println("err not nil")
		t.Fatal(err)
		return
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetSubjects)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("Received response code %v, expected %v", recorder.Code, 200)
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
	datamgr.ConnectDB("temp.db")

	// Add test obhects to temp db

	req_body := struct {
		Name       string
		MaxReviews int
	}{
		Name:       "1",
		MaxReviews: 5,
	}

	body := ExecuteRequest(req_body, "GET", "get-subject-reviews", handlers.GetSubjectReviews, 200, t)

	// Verify output
	var reviews []datamgr.Review
	json.NewDecoder(body).Decode(&reviews)

	if reviews[0].Text != "Test text1" {
		t.Error("Output does not match expected output")
	}
}

/*==================== Update Tests ====================*/

/*
 * TestUpdateReview: Tests creating and editing a subject by calling CreateReview() and UpdateReview()
 * and checking if the objecet returned from UpdateReview() is the same as the expected output. Cleans up
 * by calling DeleteSubject()
 */

func TestUpdateReview(t *testing.T) {
	req_body := struct {
		ID      uint
		NewText string
	}{
		ID:      6,
		NewText: "Emmett rocks",
	}

	body := ExecuteRequest(req_body, "PUT", "/update-review", handlers.UpdateReview, 200, t)

	// Verify the updated entry in the packet's body has the correct string
	var review datamgr.Review
	json.NewDecoder(body).Decode(&review)

	if review.Text != "Emmett rocks" {
		t.Error("Output does not match expected output")
	}
}

/*==================== Delete Tests ====================*/

/*
 * TestDeleteSubject: Tests creating and deleting a subject by calling CreateSubject() followed by
 * DeleteSubject() and checking if the state of the database matches the initial state
 */

func TestDeleteSubject(t *testing.T) {
	var subject datamgr.Subject
	datamgr.DB.Find(&subject, 5)

	ExecuteRequest(subject, "DELETE", "/delete-subject", handlers.DeleteSubjecet, 200, t)

	// Verfiy that the deleted subject is not returned when querying the db
	var subjects []datamgr.Subject
	datamgr.DB.Find(&subjects)
	for i := 0; i < len(subjects); i++ {
		if subjects[i].ID == 5 {
			t.Error("ID found, failed to delete object")
		}
	}

}

/*================== Functional Tests ==================*/
