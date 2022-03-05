package main

import (
	"log"
	"net/http"
	"yofio/intermedio/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/credit-assignment", handler.AsignaMonto{}).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", r))
}


