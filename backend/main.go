package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"swe-project/backend/handlers"
	"swe-project/backend/initializers"
	"swe-project/backend/objects"
)

func main() {
	fmt.Println("Starting swe-project/backend server.")

	initializers.LoadEnv()
	initializers.Connect_db()
	initializers.DB.AutoMigrate(&objects.User{}, &objects.Subject{})

	r := mux.NewRouter()
	handlers.MasterHandler(r)

	log.Fatal(http.ListenAndServe(":3000", r))

}
