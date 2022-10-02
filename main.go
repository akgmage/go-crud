package main

import "github.com/gorilla/mux"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	r.HandleFunc()
}