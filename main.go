package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"s.ab/gofiles/handlers"
)

func GetFile(rw http.ResponseWriter, r *http.Request) {

}

func Test(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	err := json.NewEncoder(rw).Encode(map[string]string{"status": "Hello world"})
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Test).Methods("GET")
	router.HandleFunc("/createFile", handlers.CreateFile).Methods("POST")
	router.HandleFunc("/readFile/{fileName}", GetFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
