package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Author struct {
	Name string
}

type Book struct {
	Name   string
	Author *Author
}

func main() {
	author := &Author{"alice"}
	book := Book{
		Name:   "a book",
		Author: author,
	}
	bytes, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal
	unm := Book{}
	print(unm)
	err = json.Unmarshal(bytes, &unm)
	if err != nil {
		log.Fatal(err)
	}
	print(unm)
}

func print(book Book) {
	fmt.Println(book.Name)
	fmt.Println(book.Author)
}
