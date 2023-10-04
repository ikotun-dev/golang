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
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header.
}

func main() { 
	r := mux.NewRouter()

	r.Handle
}
