package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Do struct {
	Id          int    `json:"id"`
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", listDos).Methods("GET")
	router.HandleFunc("/add", addDos).Methods("POST")
	router.HandleFunc("/remove", removeDos).Methods("DELETE")
	router.HandleFunc("/complete", completeDos).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
