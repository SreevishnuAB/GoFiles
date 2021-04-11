package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"s.ab/gofiles/models"
)

const FileDir = "files"

func CreateFile(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var file models.File
	err := decoder.Decode(&file)
	log.Println(file)
	if err != nil {
		log.Panic(err)
	}

	fileJson, err := json.Marshal(file)
	if err != nil {
		log.Panic(err)
	}
	log.Println(file.Name, fileJson)

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	log.Println("working directory = ", wd)
	dirPath := filepath.Join(wd, FileDir)
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		log.Println("Directory does not exist. Creating directory")
		err := os.Mkdir(dirPath, 0766)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Println("Directory exists")
	}

	filePath := filepath.Join(dirPath, file.Name)
	log.Println(filePath)
	err = os.WriteFile(filePath, fileJson, 0666)
	if err != nil {
		log.Panic(err)
	}
	rw.WriteHeader(201)
	err = json.NewEncoder(rw).Encode(map[string]string{"status": "File created"})
	if err != nil {
		log.Panic(err)
	}
}
