package main

import (
	"linkly/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/shorten", handlers.ShortenHandler)
	r.HandleFunc("/r/{id}", handlers.RedirectHandler)
	r.PathPrefix("/assets/").Handler(fs)

	log.Fatal(http.ListenAndServe(":3000", r))
}
