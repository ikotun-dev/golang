package main
import (
	"fmt",
	"log",
	"encoding/json",
	"math/rand",
	"net/http",
	"strconv",
	"github.com/gorilla/mux"
)
 
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
	
}

type Director struct { 
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

//function to get all movies 
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

//function to delete a movie 
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"]{
			movies = append(movies[:index], movies[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
	
}
//function to get a single movie 
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies { 
		if item.ID == params["ID"]{
			json.NewEncoder(w).Encode(item)

			return
		}
	}
}


//main function 
func main() { 
	r := mux.NewRouter()

	movie1 := Movie{
		ID : "1",
		Isbn : "12345",
		Title : "Spiderman",
		Director:  : &Director{
			Firstname : "John",
			Lastname : "Doe"
		},
	}

	movie2 := Movie{
		ID : "2",
		Isbn : "12323",
		Title : "Superman",
		Director:  : &Director{
			Firstname : "John",
			Lastname : "Doe"
		},
	}

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", updateMovie).Methods("PUT")	
	r.HandleFunc("/movies", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
}


