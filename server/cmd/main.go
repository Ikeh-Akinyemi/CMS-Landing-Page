package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()

	mux.Get("/landing-page", http.HandlerFunc(landingPage))

	// Create a file server which serves files out of the "./ui/static" directory.
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	http.ListenAndServe(":8080", mux)
}
