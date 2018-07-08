package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Make a get function for the mux router to make it testable.
func newRouter() *mux.Router {
	r := mux.NewRouter()
	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))

	// define routes
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/workorders", getWorkordersHandler).Methods("GET")
	r.HandleFunc("/workorders", createWorkorderHandler).Methods("POST")
	r.PathPrefix("/assets/").Handler(fileServer).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
