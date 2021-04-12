package v1handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"s.ab/gofiles/constants"
)

func GetFile(rw http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["fileName"] // mux.Vars() returns map of path params
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

	data, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("Requested file", fileName, " not found")
			// TODO proper error response
		} else {
			log.Println(err)
		}
		return
	}

	log.Println(data)

}
