package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"
)

/*
 * Takes http request body and binds it to the input structure
 */

func ReadRequest(w http.ResponseWriter, r *http.Request, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&obj)
}

/*
 * WriteResponse: Writes JSON to the response and sends the proper http return code as specified
 */

func WriteResponse(w http.ResponseWriter, response interface{}, returnCode int) {
	if response == nil {
		w.WriteHeader(returnCode)
		return
	}

	raw, err := json.Marshal(response)

	if err != nil {
		fmt.Fprint(w, "Failed to convert object into JSON")
		w.WriteHeader(400)
		return
	}

	// NOTE: If I do Write() before WriteHeader() I think it defaults to 200
	w.WriteHeader(returnCode)
	w.Write(raw)
}
