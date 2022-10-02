package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title     string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	movies = append(movies, Movie{
		ID: "1",
		Isbn: "1234",
		Director: &Director{
			FirstName: "Abhisek",
			LastName: "Gupta",
		},
	})
	movies = append(movies, Movie{
		ID: "2",
		Isbn: "1235",
		Director: &Director{
			FirstName: "Yam",
			LastName: "Joram",
		},
	})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Strted server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}