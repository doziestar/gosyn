package bookstore

import (
	"fmt"
	"net/http"

	"github.com/doziestar/gosyn/app/bookstore/routes"
	"github.com/gorilla/mux"
)

func Start() {
	// fmt.Println("My bookstore!")
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Println("Serving at http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", r))
}
