package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice constaining
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// User the http.NewServeMux() function to initialize a new servemux, then
	// registerthe home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() function to start a new web server. We pass
	// in two parms: TCP network addr to listen on (:4000)
	// and the servemux we just created. If http.ListenAndServe() returns
	// an error,  we use log.Fatal() function to log the error message and exit
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
