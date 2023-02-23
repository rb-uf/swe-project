package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"swe-project/backend/handlers"
	"swe-project/backend/datamgr"
)

func main() {
	fmt.Println("Starting swe-project/backend server.")

	// Load environment variables, used for filenames and port numbers.
	// Access an env variable with os.Getenv("ENV_VAR").
	err := godotenv.Load(); if err != nil {
		log.Fatal("Error loading .env file")
	}


	datamgr.ConnectDB()

	r := mux.NewRouter()
	handlers.MasterHandler(r)

	log.Fatal(http.ListenAndServe(":3000", r))
}
