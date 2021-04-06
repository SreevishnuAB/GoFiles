package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"s.ab/go-files/models"
)

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
	err = os.WriteFile(file.Name, fileJson, 0666)
	if err != nil {
		log.Panic(err)
	}
	rw.WriteHeader(201)
	err = json.NewEncoder(rw).Encode(map[string]string{"status": "File created"})
	if err != nil {
		log.Panic(err)
	}
}

func GetFile(rw http.ResponseWriter, r *http.Request) {

}

func Test(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	err := json.NewEncoder(rw).Encode(map[string]string{"status": "Hello world"})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Test).Methods("GET")
	router.HandleFunc("/createFile", CreateFile).Methods("POST")
	router.HandleFunc("/readFile/{fileName}", GetFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
