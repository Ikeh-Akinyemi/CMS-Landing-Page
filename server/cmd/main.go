package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	port := ":8080"
	log.Printf("starting server on port %s\n", port)

	mux := pat.New()

	mux.Get("/landing-page", http.HandlerFunc(landingPage))

	// Create a file server which serves files out of the "./ui/static" directory.
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	http.ListenAndServe(port, mux)
}
