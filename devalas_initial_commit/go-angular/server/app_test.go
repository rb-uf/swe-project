package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestGetAllChairRatings(t *testing.T) {
	app := initApp()
	st := chairRating{ID: "id-1", Rating: 20, ClassroomName: "John Doe"}
	app.db.Save(st)
	req, _ := http.NewRequest("GET", "/chairRatings", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getAllChairRatings)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, st, t)
}

func TestAddChairRating(t *testing.T) {
	app := initApp()
	var rqBody = toReader(`{"classroomName":"John Doe", "rating":20}`)
	req, _ := http.NewRequest("POST", "/chairRatings", rqBody)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.addChairRating)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusCreated, t)
	checkContentType(r, t)
	checkProperties(firstChairRating(app), t)
}

func TestUpdateChairRating(t *testing.T) {
	app := initApp()
	app.db.Save(chairRating{ID: "id-1", Rating: 25, ClassroomName: "Peter Doe"})
	var rqBody = toReader(`{"classroomName":"John Doe", "rating":20}`)
	req, _ := http.NewRequest("PUT", "/chairRatings/id", rqBody)
	req = mux.SetURLVars(req, map[string]string{"id": "id-1"})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.updateChairRating)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkProperties(firstChairRating(app), t)
}

func TestDeleteChairRating(t *testing.T) {
	app := initApp()
	app.db.Save(chairRating{ID: "id-1", Rating: 20, ClassroomName: "John Doe"})
	req, _ := http.NewRequest("DELETE", "/chairRatings/id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "id-1"})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.deleteChairRating)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkDbIsEmpty(app.db, t)
}

func initApp() App {
	db, _ := gorm.Open("sqlite3", ":memory:")
	db.AutoMigrate(&chairRating{})
	return App{db: db}
}

func firstChairRating(app App) chairRating {
	var all []chairRating
	app.db.Find(&all)
	return all[0]
}

func toReader(content string) io.Reader {
	return bytes.NewBuffer([]byte(content))
}

func checkStatusCode(code int, want int, t *testing.T) {
	if code != want {
		t.Errorf("Wrong status code: got %v want %v", code, want)
	}
}

func checkContentType(r *httptest.ResponseRecorder, t *testing.T) {
	ct := r.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Wrong Content Type: got %v want application/json", ct)
	}
}

func checkProperties(st chairRating, t *testing.T) {
	if st.ClassroomName != "John Doe" {
		t.Errorf("Classroom Name should match: got %v want %v", st.ClassroomName, "Peter Doe")
	}
	if st.Rating != 20 {
		t.Errorf("Rating should match: got %v want %v", st.Rating, 20)
	}
}

func checkBody(body *bytes.Buffer, st chairRating, t *testing.T) {
	var chairRatings []chairRating
	_ = json.Unmarshal(body.Bytes(), &chairRatings)
	if len(chairRatings) != 1 {
		t.Errorf("Wrong lenght: got %v want 1", len(chairRatings))
	}
	if chairRatings[0] != st {
		t.Errorf("Wrong body: got %v want %v", chairRatings[0], st)
	}
}

func checkDbIsEmpty(db *gorm.DB, t *testing.T) {
	var chairRatings []chairRating
	db.Find(&chairRatings)
	if len(chairRatings) != 0 {
		t.Errorf("ChairRating has not been deleted")
	}
}
