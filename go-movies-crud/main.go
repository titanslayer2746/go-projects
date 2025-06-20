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

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func main() {

	r := mux.NewRouter() 
	
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454545", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Coppola"}})
	movies = append(movies, Movie{ID: "3", Isbn: "454546", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})
	movies = append(movies, Movie{ID: "4", Isbn: "454547", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "5", Isbn: "454548", Title: "Fight Club", Director: &Director{Firstname: "David", Lastname: "Fincher"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Server starting at port 8000\n")
	log.Fatal(http.ListenAndServe(
		":8000",r,
	))
}