package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// get a list of books
func getBooksList() ([]Book, error) {
	var books []Book
	file_bytes, err := os.ReadFile("../books_data.json")
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

// get a fixed list of
// books
func GetBooksList() ([]Book, error) {
	queryset, err := getBooksList()
	var books []Book
	if err != nil {
		return nil, err
	} else {
		if len(queryset) != 0 {
			for i := range len(queryset) {
				books = append(books, queryset[i])
			}
		}
	}
	return books, nil
}

// get a single book
func GetBookObject(id string) (Book, int, error) {
	books, err := getBooksList()
	if err != nil {
		return Book{}, -1, err
	} else {
		if len(books) != 0 {
			var book Book = Book{}
			var index int
			for i, b := range books {
				if b.Id == id {
					book = b
					index = i
				}
			}
			return book, index, nil
		} else {
			return Book{}, -1, nil
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
func AddBook(book_bytes []byte) (Book, error) {
	books, err := getBooksList()
	var book_data *Book
	json.Unmarshal(book_bytes, &book_data)
	if err != nil {
		return *book_data, err
	} else {
		books = append(books, *book_data)
		err = saveChanges(books, "./books.json")
		if err != nil {
			return *book_data, err
		} else {
			return *book_data, nil
		}
	}
}

// delete an existing book
func DeleteBook(id string) (string, error) {
	books, err := getBooksList()
	if err != nil {
		return "", err
	} else {
		if len(books) != 0 {
			book, index, err := GetBookObject(id)
			if err != nil {
				return "", err
			} else {
				if index != -1 {
					log.Printf("Deletion of the book with title: %v and id: %v\n", book.Title, book.Id)
					books = append(books[:index], books[index:+1]...)
					err := saveChanges(books, "./books.json")
					if err != nil {
						return "", err
					} else {
						return fmt.Sprintf("Book #%v was deeleted successfully!", id), nil
					}
				} else {
					return fmt.Sprintf("No such book found with id: %v", id), nil
				}
			}
		} else {
			return fmt.Sprintf("There are no books in the DB yet"), nil
		}
	}
}
