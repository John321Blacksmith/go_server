package main

import (
	"fmt"
	"log"
	"net/http"
)

// define a listening port
const PORT string = ":8000"

// starting the server
// and handling the incoming
// requests
func main() {
	// define handling
	http.HandleFunc("/", handleGetBooksList)

	http.HandleFunc("/book", handleGetBookObject)

	http.HandleFunc("/update", handleUpdateBook)

	http.HandleFunc("/add", handleAddBook)

	http.HandleFunc("/delete", handleDeleteBook)

	fmt.Printf("Server is listening on port %v\n", PORT)

	// launching the server
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal(err)
	}
}

// handle a request for
// a books list
func handleGetBooksList(w http.ResponseWriter, r *http.Request) {
}

// handle a request for
// a book object
func handleGetBookObject(w http.ResponseWriter, r *http.Request) {

}

// handle a request for
// book update
func handleUpdateBook(w http.ResponseWriter, r *http.Request) {

}

// handle a request
// for book creation
func handleAddBook(w http.ResponseWriter, r *http.Request) {

}

// handle a reqeust for
// book deletion
func handleDeleteBook(w http.ResponseWriter, r *http.Request) {

}
