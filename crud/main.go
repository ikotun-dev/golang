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
	Lastname  string `json:"lastname"`
}

var movies []Movie

// function to get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// function to delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

// function to get a single movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}
}

// function to create a movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// var movie Movie
	// err := json.NewDecoder(r.Body).Decode(&movie)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	// 	return
	// }
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

	// movie.ID = strconv.Itoa(rand.Intn(100000))

	// movies = append(movies, movie)
	// json.NewEncoder(w).Encode(movie)

}

// function
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//setting the header type
	w.Header().Set("Content-type", "application/json")
	//get params
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// main function
func main() {
	r := mux.NewRouter()

	movie1 := Movie{
		ID:    "1",
		Isbn:  "12345",
		Title: "Spiderman",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	}

	movie2 := Movie{
		ID:    "2",
		Isbn:  "12323",
		Title: "Superman",
		Director: &Director{
			Firstname: "Black",
			Lastname:  "Boss",
		},
	}

	movie3 := Movie{
		ID:    "3",
		Isbn:  "18003",
		Title: "Batman",
		Director: &Director{
			Firstname: "Mike",
			Lastname:  "Don",
		},
	}

	movies = append(movies, movie1)
	movies = append(movies, movie2)
	movies = append(movies, movie3)

	/* r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", updateMovie).Methods("PUT")
	r.HandleFunc("/movies", deleteMovie).Methods("DELETE") */

	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", updateMovie).Methods("PUT")
	r.HandleFunc("/movies", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
