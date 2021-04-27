package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"s.ab/gofiles/v1handlers"
)

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
	router.HandleFunc("/createFile", v1handlers.CreateFile).Methods("POST")
	router.HandleFunc("/readFile/{fileName}", v1handlers.GetFile).Methods("GET")
	router.HandleFunc("/modifyFile/{fileName}", v1handlers.ModifyFile).Methods("PUT")
	log.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
