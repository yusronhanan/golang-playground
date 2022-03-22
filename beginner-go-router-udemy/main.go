package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Roll is model for Sushi API
type Roll struct {
	ID          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Ingredient  string `json: "ingredient"`
}

//Init rolls var as a slice
var rolls []Roll

// Index
func getRolls(w http.ResponseWriter, r *http.Request) {

}

// Show single sushi
func getRoll(w http.ResponseWriter, r *http.Request) {

}

// Add single sushi
func createRoll(w http.ResponseWriter, r *http.Request) {

}

// Update single sushi
func updateRoll(w http.ResponseWriter, r *http.Request) {

}

// Delete single sushi
func DeleteRoll(w http.ResponseWriter, r *http.Request) {

}
func main() {
	//init Router
	router := mux.NewRouter()

	//Handle Endpoints/Routing

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", DeleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
