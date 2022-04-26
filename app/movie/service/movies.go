package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/doziestar/gosyn/app/movie/data"
	"github.com/gorilla/mux"
)

var movies []data.Movie
var directors []data.Director

func MovieRoutes(r *mux.Router) *mux.Router {
	directors = append(directors, data.Director{ID: 1, Name: "Steven Spielberg"})
	movies = append(movies, data.Movie{ID: 1, Title: "movie1", Year: 2019, Director: &directors[0]})
	movies = append(movies, data.Movie{ID: 2, Title: "movie2", Year: 2019, Director: &directors[0]})
	r.HandleFunc("/movies", getMovies()).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie()).Methods("GET")
	r.HandleFunc("/movies", createMovie()).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie()).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie()).Methods("DELETE")

	return r
}

func getMovies() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	}
}

func getMovie() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get user id from request
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		// get movie from movies slice
		for _, movie := range movies {
			if movie.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(movie)
				return
			}
		}
	}
}

func createMovie() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// create movies and append to movies slice
		var movie data.Movie
		json.NewDecoder(r.Body).Decode(&movie)
		movie.Director = &directors[0]
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
	}
}

func updateMovie() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		for index, movie := range movies {
			if movie.ID == id {
				json.NewDecoder(r.Body).Decode(&movie)
				movie.Director = &directors[0]
				movies[index] = movie
				json.NewEncoder(w).Encode(movie)
				return
			}
		}
	}
}

func deleteMovie() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		for index, movie := range movies {
			if movie.ID == id {
				movies = append(movies[:index], movies[index+1:]...)
				return
			}
		}
	}
}
