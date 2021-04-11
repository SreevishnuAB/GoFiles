package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFile(rw http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)
	log.Println("Queried file = ", fileName)

}
