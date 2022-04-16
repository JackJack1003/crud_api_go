package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lasstname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mymovie Movie
	_ = json.NewDecoder(r.Body).Decode(&Mymovie)
	Mymovie.ID = strconv.Itoa((rand.Intn(1000000)))
	movies = append(movies, Mymovie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mymovie Movie
	params := mux.Vars(r)
	for index, items := range movies {
		if params["id"] == items.ID {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&Mymovie)
			Mymovie.ID = params["id"]
			movies = append(movies, Mymovie)
			json.NewEncoder(w).Encode(movies)
			return
		}

	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(movies[index])
			break
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "Movie 1", Director: &Director{Firstname: "Shaan", Lastname: "kort"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting Jack's server\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
