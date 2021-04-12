package v1handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"s.ab/gofiles/constants"
	"s.ab/gofiles/helpers"
)

func ModifyFile(rw http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["fileName"]
	decoder := json.NewDecoder(r.Body)
	var reqBody interface{}
	err := decoder.Decode(&reqBody)
	newContent := reqBody.(map[string]interface{})
	if err != nil {
		log.Println(err)
		// TODO add proper error response
		return
	}
	log.Println("Queried file = ", fileName)
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		// TODO proper error response
		return
	}
	fileName = fileName + ".json"
	filePath := filepath.Join(wd, constants.FileDir, fileName)
	log.Println("Queried filepath = ", filePath)
	newContentJson, err := json.Marshal(newContent["content"].([]interface{}))
	if err != nil {
		log.Println(err)
		// TODO proper error response
		return
	}
	log.Println("New content", newContentJson)

	if _, err = os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		log.Println("File", fileName, "does not exist")
		// TODO proper error response
		return
	}
	err = os.WriteFile(filePath, newContentJson, 0666)
	if err != nil {
		log.Println("File", fileName, "could not be modified")
		// TODO proper error response
		return
	}

	err = helpers.CreateResponse(rw, http.StatusOK, map[string]string{"status": "File modified"})

	if err != nil {
		log.Println("Error creating response body", err)
		helpers.CreateResponse(rw, http.StatusInternalServerError, map[string]string{"error": "Something went wrong"})
	}

}
