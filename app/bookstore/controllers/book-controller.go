package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/doziestar/gosyn/app/bookstore/config"
	"github.com/doziestar/gosyn/app/bookstore/models"
	"github.com/gorilla/mux"
)

func Test() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test")
	}
}

func GetBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		book, err := models.GetBookById(id)
		if err != nil {
			log.Fatal(err)
		}
		// if book == nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	json.NewEncoder(w).Encode("Book not found")
		// 	return
		// }
		json.NewEncoder(w).Encode(book)

	}
}

func GetBooks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		books, _ := models.GetBooks()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}

func CreateBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		book := models.Book{}
		json.NewDecoder(r.Body).Decode(&book)
		book.Create()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	}
}

func UpdateBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		book, err := models.GetBookById(id)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		book.Author = vars["author"]
		book.Name = vars["name"]
		book.Publication = vars["publication"]
		fmt.Println(&book)
		book.Update()
		config.GetDB().Save(&book)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(&book)
	}
}

func DeleteBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		book, err := models.GetBookById(id)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		book.Delete()
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode("Book with id " + vars["id"] + "  deleted successfully")
	}
}
