package pres

import (
	"encoding/json"
	"log"
	"net/http"

	"books_api/internal/infra"
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

// Handle a request for
// a books list
func HandleGetBooksList(w http.ResponseWriter, r *http.Request) {
	books, err := infra.GetBooksList()
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(500)
		w.Write(encodeData("Internal Server Error"))
	} else {
		w.WriteHeader(200)
		w.Write(encodeData(books))
	}
}

// Handle a request for
// a book object
func HandleGetBookObject(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	book, _, err := infra.GetBookObject(id)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(500)
		w.Write(encodeData("Internal Server Error"))
	} else {
		w.WriteHeader(200)
		w.Write(encodeData(book))
	}
}

// Handle a request for
// book update
func HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if checkMethod(r, "POST") {
		id := query.Get("id")
		book, err := infra.UpdateBook(id, encodeData(r.Body))
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

// Handle a request
// for book creation
func HandleAddBook(w http.ResponseWriter, r *http.Request) {
	if checkMethod(r, "POST") {
		book, err := infra.AddBook(encodeData(r.Body))
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

// Handle a reqeust for
// book deletion
func HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if checkMethod(r, "DELETE") {
		id := query.Get("id")
		result, err := infra.DeleteBook(id)
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
