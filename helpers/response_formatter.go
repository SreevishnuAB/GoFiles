package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateResponse(rw http.ResponseWriter, responseCode int, responseBody interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(responseCode)
	err := json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		log.Println(err)
	}
	return err
}
