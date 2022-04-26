package movie

import (
	"log"
	"net/http"

	"github.com/doziestar/gosyn/app/movie/service"
	"github.com/gorilla/mux"
)

func Movie() {
	r := mux.NewRouter()

	service.MovieRoutes(r)

	// fmt.Println(r.ServeHTTP(nil, nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}
