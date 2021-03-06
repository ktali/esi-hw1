package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Status       string `json:"status"`
	ErrorContent string `json:"error"`
}

type Do struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed"`
}

type Dos []Do

var todos Dos

func listDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getRandomDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(todos) > 0 {
		randId := rand.Intn(len(todos))
		json.NewEncoder(w).Encode(todos[randId])
	}
}

func addDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	content := r.FormValue("content")
	if len(todos) == 0 {
		todos = append(todos, Do{Id: 0, Content: content, IsCompleted: false})
	} else {
		newId := todos[len(todos)-1].Id + 1
		todos = append(todos, Do{Id: newId, Content: content, IsCompleted: false})
	}
	json.NewEncoder(w).Encode(Response{Status: "ok"})
}

func removeDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err == nil {
		for idx, item := range todos {
			if item.Id == id {
				removeFromDos(idx)
				json.NewEncoder(w).Encode(Response{Status: "ok"})
				return
			}
		}
		json.NewEncoder(w).Encode(Response{Status: "error", ErrorContent: "no such entry"})
	} else {
		json.NewEncoder(w).Encode(Response{Status: "error", ErrorContent: "failed to parse id"})
	}

}

func removeFromDos(i int) {
	copy(todos[i:], todos[i+1:]) // Shift a[i+1:] left one index.
	todos[len(todos)-1] = Do{}   // Erase last element (write zero value).
	todos = todos[:len(todos)-1] // Truncate slice.
}

func completeDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err == nil {
		for idx, item := range todos {
			if item.Id == id {
				todos[idx].IsCompleted = true
				json.NewEncoder(w).Encode(Response{Status: "ok"})
				return
			}
		}
		json.NewEncoder(w).Encode(Response{Status: "error", ErrorContent: "no such entry"})
	}
}

func handleRequests() {

	//Mock data
	todos = append(todos, Do{Id: 0, Content: "Fake todo #0", IsCompleted: true})
	todos = append(todos, Do{Id: 1, Content: "Fake todo #1", IsCompleted: false})
	todos = append(todos, Do{Id: 2, Content: "Fake todo #2", IsCompleted: false})

	router := mux.NewRouter()
	router.HandleFunc("/", listDos).Methods("GET")
	router.HandleFunc("/random", getRandomDo).Methods("GET")
	router.HandleFunc("/add", addDos).Methods("POST")
	router.HandleFunc("/remove/{id:[0-9]+}", removeDos).Methods("DELETE")
	router.HandleFunc("/complete/{id:[0-9]+}", completeDos).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
