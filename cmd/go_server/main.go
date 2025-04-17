package main

import (
	"fmt"
	"log"
	"net/http"

	"books_api/internal/pres"
)

// define a listening port
const PORT string = ":8000"

// starting the server
// and handling the incoming
// requests
func main() {
	// define handling
	http.HandleFunc("/", pres.HandleGetBooksList)

	http.HandleFunc("/book", pres.HandleGetBookObject)

	http.HandleFunc("/update", pres.HandleUpdateBook)

	http.HandleFunc("/add", pres.HandleAddBook)

	http.HandleFunc("/delete", pres.HandleDeleteBook)

	fmt.Printf("Server is listening on port %v\n", PORT)

	// launching the server
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal(err)
	}
}
