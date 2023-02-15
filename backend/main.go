package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"swe-project/backend/handlers"
	"swe-project/backend/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect_db()
}

func main() {
	fmt.Println("Starting swe-project/backend server.")

	r := mux.NewRouter()
	handlers.MasterHandler(r)

	log.Fatal(http.ListenAndServe(":3000", r))

}
