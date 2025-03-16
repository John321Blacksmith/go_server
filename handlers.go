package main

import (
	"encoding/json"
	"os"
)

// get a list of books
func getBooksList() ([]Book, error) {
	var books []Book
	file_bytes, err := os.ReadFile("./books.json")
	if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(file_bytes, &books)
		if err != nil {
			return nil, err
		} else {
			return books, nil
		}
	}
}

// save changes to
// the json file
func saveChanges(books []Book, file string) error {
	book_bytes, err := json.Marshal(books)
	if err != nil {
		return err
	} else {
		err = os.WriteFile(file, book_bytes, 0644)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

// get a single book
func GetBookObject(id string) (Book, error) {
	books, err := getBooksList()
	if err != nil {
		return Book{}, err
	} else {
		if len(books) != 0 {
			var book Book = Book{}
			for _, b := range books {
				if b.Id == id {
					book = b
				}
			}
			return book, nil
		} else {
			return Book{}, nil
		}
	}
}

// update an existing book
func UpdateBook(id string, body []byte) (Book, error) {
	books, err := getBooksList()
	var book Book
	updated_book := &book
	if err != nil {
		return Book{}, err
	} else {
		if len(books) != 0 {
			err = json.Unmarshal(body, updated_book)
			if err != nil {
				return Book{}, err
			} else {
				for i, b := range books {
					if b.Id == id {
						books[i] = *updated_book
					}
				}
				err = saveChanges(books, "./books.json")
				if err != nil {
					return Book{}, err
				} else {
					return *updated_book, nil
				}
			}
		}
	}
	return *updated_book, nil
}

// add a new book
func AddBook() {

}

// delete an existing book
func DeleteBook(id string) Book {
	return Book{}
}
