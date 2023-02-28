package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"swe-project/backend/datamgr"
	"swe-project/backend/handlers"
)

func main() {
	fmt.Println("Starting swe-project/backend server.")

	loadEnv()
	datamgr.ConnectDB(os.Getenv("DB_FILE"))

	r := mux.NewRouter()

	handlers.MasterHandler(r)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}

// loadEnv: load environment variables
// Env vars are being used for filenames and port numbers.
// Access an env variable with os.Getenv("ENV_VAR").
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
