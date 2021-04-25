package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice constaining
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Creating a new snippet..."))
}

func main() {
	// User the http.NewServeMux() function to initialize a new servemux, then
	// registerthe home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass
	// in two parms: TCP network addr to listen on (:4000)
	// and the servemux we just created. If http.ListenAndServe() returns
	// an error,  we use log.Fatal() function to log the error message and exit
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
