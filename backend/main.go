package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"swe-project/backend/datamgr"
	"swe-project/backend/handlers"
)

var portFlag = flag.String("port", "3000", "port number on which to listen for requests")
var dbFlag = flag.String("db", "datamgr/database.db", "database file path")
var frontendFlag = flag.String("frontend", "../frontend/rater-gator/dist/rater-gator/", "frontend dist path")

func main() {
	flag.Parse()

	log.Println("Starting swe-project/backend server.")

	datamgr.ConnectDB(*dbFlag)
	r := mux.NewRouter()
	handlers.MasterHandler(r, *frontendFlag)

	log.Fatal(http.ListenAndServe(":" + *portFlag, r))
}
