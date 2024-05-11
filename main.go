package main

import (
	"github.com/gorilla/mux"
)

type Movie struct {
	id       string    `json: "id"`
	isbn     string    `json: "isbn"`
	title    string    `json: "title"`
	director *Director `json: "director"`
}

type Director struct {
	firstName string `json: "firstname"`
	lastName  string `json: "lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies.Methods("GET"))
	r.HandleFunc("/movies/{ID}", getMovie.Methods("GET"))
	r.HandleFunc("/movies", createMovie.Methods("POST"))
	r.HandleFunc("/movies/{id}", updateMovie.Methods("PUT"))
	r.HandleFunc("/movies/{id}", deleteMovie.Methods("DELETE"))
}
