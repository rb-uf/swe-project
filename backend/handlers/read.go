package handlers

import (
	"net/http"
)


func GetTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}
