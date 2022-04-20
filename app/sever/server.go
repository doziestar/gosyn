package sever

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Serve() {
	fileServer := http.FileServer(http.Dir("../../static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler2)

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello, world!</h1>")
	// w.Write([]byte("Hello, world!"))
}

func formHandler2(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/form" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	// if r.Method != "POST" {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() err", http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
	fmt.Println(r.Form)
	fmt.Println(r.Body.Read([]byte{}))
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, Form!")
}
