package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// define a listening port
const PORT string = ":8000"

// translate incoming
// data to bytes
func encodeData(data any) []byte {
	encoded_data, _ := json.Marshal(data)
	return encoded_data
}

// check for the
// right request method
func checkMethod(r *http.Request, meth string) bool {
	return r.Method != meth
}

// handle a request for
// a books list
func handleGetBooksList(w http.ResponseWriter, r *http.Request) {
	books, err := GetBooksList()
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(500)
		w.Write(encodeData("Internal Server Error"))
	} else {
		w.WriteHeader(200)
		w.Write(encodeData(books))
	}
}

// handle a request for
// a book object
func handleGetBookObject(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	book, _, err := GetBookObject(id)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(500)
		w.Write(encodeData("Internal Server Error"))
	} else {
		w.WriteHeader(200)
		w.Write(encodeData(book))
	}
}

// handle a request for
// book update
func handleUpdateBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if checkMethod(r, "POST") {
		id := query.Get("id")
		book, err := UpdateBook(id, encodeData(r.Body))
		if err != nil {
			log.Printf("Error: %v", err)
			w.WriteHeader(500)
			w.Write(encodeData("Internal Server Error"))
		} else {
			log.Printf("The result of update is: %v", book)
			w.WriteHeader(201)
			w.Write(encodeData(book))
		}
	} else {
		log.Printf("Error: invalid method %v", r.Method)
		w.WriteHeader(405)
		w.Write(encodeData("Method Not Allowed"))
	}
}

// handle a request
// for book creation
func handleAddBook(w http.ResponseWriter, r *http.Request) {
	if checkMethod(r, "POST") {
		book, err := AddBook(encodeData(r.Body))
		if err != nil {
			log.Printf("Error: %v", err)
			w.WriteHeader(500)
			w.Write(encodeData("Internal Server Error"))
		} else {
			log.Printf("Result: %v", book)
			w.WriteHeader(201)
			w.Write(encodeData(book))
		}
	} else {
		log.Printf("Error: invalid method %v", r.Method)
		w.WriteHeader(405)
		w.Write(encodeData("Method Not Allowed"))
	}
}

// handle a reqeust for
// book deletion
func handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if checkMethod(r, "DELETE") {
		id := query.Get("id")
		result, err := DeleteBook(id)
		if err != nil {
			log.Printf("Error: %v", err)
			w.WriteHeader(500)
			w.Write(encodeData("Internal Server Error"))
		} else {
			log.Printf("Deletion the book %v\n", id)
			w.WriteHeader(204)
			w.Write(encodeData(result))
		}
	}
}

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
