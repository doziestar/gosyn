package routes

import (
	"github.com/doziestar/gosyn/app/bookstore/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Test()).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks()).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook()).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook()).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook()).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook()).Methods("DELETE")
}
