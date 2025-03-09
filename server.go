package main

import (
	"flag"
	"fmt"
	"os"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "List all the books")
	getOne := getCmd.String("id", "", "Get one book by ID")

	if len(os.Args[0]) < 2 {
		fmt.Println("Please specify a flag")
		os.Exit(1)
	}

	switch os.Args[0] {
	case "get":
		handleGetAll(getCmd, getAll, getOne)
	default:
		fmt.Println("No valid commands were provided")
	}

}
