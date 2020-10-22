package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"pokemon-api/database"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/addPokemon",addPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", myRouter))
}

func addPokemon(w http.ResponseWriter, r *http.Request){
	var newPokemon database.Pokemon
	reqBody, _:= ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newPokemon)
	//if _, ok := database.PokemonDb[nuevoPokemon.ID]; ok{
	//	w.WriteHeader(http.StatusNotModified)
	//	return
	//}
	database.PokemonDb[3]=newPokemon
	w.WriteHeader(http.StatusOK)

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
