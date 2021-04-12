package v1handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"s.ab/gofiles/constants"
	"s.ab/gofiles/helpers"
	"s.ab/gofiles/models"
)

func CreateFile(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var file models.File
	err := decoder.Decode(&file)
	log.Println(file)
	if err != nil {
		log.Println(err)
		// TODO add proper error response
		return
	}
	fileJson, err := json.Marshal(file.Content)
	log.Println(fileJson)
	if err != nil {
		log.Println(err)
		// TODO add proper error response
		return
	}
	log.Println(file.Name, fileJson)

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		// TODO add proper error response
		return
	}
	log.Println("working directory = ", wd)
	dirPath := filepath.Join(wd, constants.FileDir)
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		log.Println("Directory does not exist. Creating directory")
		err := os.Mkdir(dirPath, 0766)
		if err != nil {
			log.Println(err)
			// TODO add proper error response
			return
		}
		log.Println("Directory created")
	}
	fileName := file.Name + ".json"
	filePath := filepath.Join(dirPath, fileName)
	log.Println("File path = ", filePath)
	if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		log.Println("File with name", file.Name, "already exists")
		// TODO add proper error response
		return
	}

	err = os.WriteFile(filePath, fileJson, 0666)
	if err != nil {
		log.Println(err)
		// TODO add proper error response
		return
	}
	err = helpers.CreateResponse(rw, http.StatusCreated, map[string]string{"status": "File created"})

	if err != nil {
		log.Println("Error creating response body", err)
		helpers.CreateResponse(rw, http.StatusInternalServerError, map[string]string{"error": "Something went wrong"})
	}
}
