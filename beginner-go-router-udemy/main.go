package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	//Set header with json
	w.Header().Set("Content-Type", "application/json")
	//Render dart object to json
	json.NewEncoder(w).Encode(rolls)
}

// Show single sushi
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Get and parsing params
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Add single sushi
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Set var for new roll
	var newRoll Roll
	//Decode the body to newRoll
	json.NewDecoder(r.Body).Decode(&newRoll)

	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

// Update single sushi
func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			//remove the found id from slice by using append
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

// Delete single sushi
func DeleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			json.NewEncoder(w).Encode(item)
		}
	}
}
func main() {
	//Generate Mock Data

	rolls = append(rolls,
		Roll{
			ID:          "1",
			Name:        "Salmon Roll",
			Description: "crab stick, tamago sushi",
			Ingredient:  "Salmon, Soy Sauce",
		},
		Roll{
			ID:          "2",
			Name:        "California Roll",
			Description: "california stick, tamago sushi",
			Ingredient:  "California fish, Soy Sauce",
		})
	//init Router
	router := mux.NewRouter()

	//Handle Endpoints/Routing

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("PUT")
	router.HandleFunc("/sushi/{id}", DeleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
