package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type chairRating struct {
	ID            string `gorm:"primary_key" json:"id"`
	ClassroomName string `json:"classroomName"`
	Rating        int    `json:"rating"`
}

type App struct {
	db *gorm.DB
	r  *mux.Router
}

func (a *App) start() {
	a.db.AutoMigrate(&chairRating{})
	a.r.HandleFunc("/chairRatings", a.getAllChairRatings).Methods("GET")
	a.r.HandleFunc("/chairRatings", a.addChairRating).Methods("POST")
	a.r.HandleFunc("/chairRatings/{id}", a.updateChairRating).Methods("PUT")
	a.r.HandleFunc("/chairRatings/{id}", a.deleteChairRating).Methods("DELETE")
	a.r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/dist/webapp/")))
	log.Fatal(http.ListenAndServe(":8080", a.r))
}

func (a *App) getAllChairRatings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var all []chairRating
	err := a.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) addChairRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s chairRating
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = uuid.New().String()
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) updateChairRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s chairRating
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = mux.Vars(r)["id"]
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) deleteChairRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := a.db.Unscoped().Delete(chairRating{ID: mux.Vars(r)["id"]}).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
