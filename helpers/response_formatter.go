package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateResponse(rw http.ResponseWriter, responseCode int, responseBody interface{}) error {
	rw.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		log.Println(err)
	}
	return err
}
