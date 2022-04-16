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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMoviesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	paramas := mux.Vars(r)

	for _, item := range movies {
		if item.Id == paramas["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMoviesByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMoviesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = strconv.Itoa(rand.Intn(1000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "45667", Title: "Surya", Director: &Director{FirstName: "RajMAuli", LastName: "Gond"}})
	movies = append(movies, Movie{Id: "2", Isbn: "45668", Title: "Surya1", Director: &Director{FirstName: "RajMAuli1", LastName: "Gond1"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviesByID).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMoviesByID).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMoviesByID).Methods("DELETE")

	fmt.Printf("Starting Server at PORT :8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
