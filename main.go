package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func handleRequests() {
	port := os.Getenv("PORT")
	if port == ""{
		port="80"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/addPokemon",addPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func addPokemon(w http.ResponseWriter, r *http.Request){
	var newPuchamon database.Pokemon
	reqBody, _:= ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newPuchamon)

	database.PokemonDb=append(database.PokemonDb,newPuchamon)
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
