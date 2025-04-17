package infra

// book object model
type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

// define an interface of the
// response object
type Message struct {
	Msg string
}
