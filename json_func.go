package main

import (
	"encoding/json"
	"os"
)

func GetBooks() ([]Book, error) {
	var books []Book = []Book{}
	// Get a list of all the books
	file_bytes, err := os.ReadFile("./books.json")
	if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(file_bytes, books)
		if err != nil {
			return nil, err
		}
	}
	return books, err
}

// get a book objects
// via its id
func GetBook(id string) (Book, int, error) {
	books, err := GetBooks()

	var requested_book Book
	var requested_book_id int

	if err != nil {
		return Book{}, 0, err
	} else {
		for i, b := range books {
			if b.Id == id {
				requested_book = b
				requested_book_id = i
			}
		}
	}
	return requested_book,
		requested_book_id,
		nil
}

// receive a list of Books
// and serialize them to\
// the json file
func SaveBooks(books []Book) error {
	books_bytes, err := json.Marshal(books)
	if err != nil {
		return err
	} else {
		os.WriteFile("./books.json", books_bytes, 0664)
	}
	return err
}
