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
	Director *Director `json:"director"` // points to director struct - of type Director
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// encode the movies array to json then send --
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) // send back json response
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // getting request parameters
	for index, item := range movies {
		if item.ID == params["id"] {
			// slice out movie from movies by indexing
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // send back json response
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // send back json response
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie                            // new movie object of type struct Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // decode request body into new movie object
	movie.ID = strconv.Itoa(rand.Intn(9999))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // send back json response
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		// search delete movie first
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
		// then recreate movie
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = params["id"]
		json.NewEncoder(w).Encode(movie)
		return
	}
}

func main() {
	r := mux.NewRouter()

	// movies won't be empty on first call
	movies = append(movies, Movie{ID: "1", Isbn: "1242", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Jonah"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("startint server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
