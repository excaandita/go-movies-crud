package main

import (
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.newRouter()

	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "1-1",
		Title:    "Pengabdi Setan 1",
		Director: &Director{Firstname: "Joko", Lastname: "Anwar"},
	})
	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "1-2",
		Title:    "Pengabdi Setan 2",
		Director: &Director{Firstname: "Joko", Lastname: "Anwar"},
	})
	movies = append(movies, Movie{
		ID:       "3",
		Isbn:     "2-1",
		Title:    "Manusia Setengah Salmon",
		Director: &Director{Firstname: "Raditya", Lastname: "Dika"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
