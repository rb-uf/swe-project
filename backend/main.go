package main

import (
	"fmt"
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"swe-project/backend/handlers"
	"swe-project/backend/datamgr"
)

func main() {
	fmt.Println("Starting swe-project/backend server.")

	loadEnv()
	datamgr.ConnectDB()

	r := mux.NewRouter()
	handlers.MasterHandler(r)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}

// loadEnv: load environment variables
// Env vars are being used for filenames and port numbers.
// Access an env variable with os.Getenv("ENV_VAR").
func loadEnv() {
	err := godotenv.Load(); if err != nil {
		log.Fatal("Error loading .env file")
	}
}
