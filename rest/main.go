package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Do struct {
	Id          int64  `json:"id"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed"`
}

type Dos []Do

func listDos(w http.ResponseWriter, r *http.Request) {
	dos := Dos{
		Do{Id: 1, Content: "Test"},
	}
	json.NewEncoder(w).Encode(dos)
}

func addDos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func removeDos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func completeDos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", listDos).Methods("GET")
	router.HandleFunc("/add", addDos).Methods("POST")
	router.HandleFunc("/remove{id:[0-9]+}", removeDos).Methods("DELETE")
	router.HandleFunc("/complete", completeDos).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
