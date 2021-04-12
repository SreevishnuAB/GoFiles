package v1handlers

import (
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

	filePath := filepath.Join(wd, constants.FileDir, fileName)
	log.Println("Queried filepath = ", filePath)
	// file, err := os.Stat(filePath)
	// if err != nil {
	// 	if errors.Is(err, os.ErrNotExist){
	// 		log.Println("Requested file not found")
	// 		// TODO proper error response
	// 	} else {
	// 		log.Println(err)
	// 	}
	// 	return
	// }

	// log.Println("File details = ", file)
	// TODO check if ReadFile can replace Stat
	data, err := os.ReadFile(filePath)
	log.Println(data)
	if err != nil {
		log.Println(err)
		return
	}

}
