package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Author struct {
	Name string
	Book *Book
}

type Book struct {
	Name string
	Author Author
}

func main() {
	// basic types
	i := 1
	fmt.Println(reflect.TypeOf(i))

	// customized
	book := Book{
		Name: "TAOCP",
	}
	fmt.Println(reflect.TypeOf(book))

	//
	author := Author{
		Name: "Alice",
		Book: &book,
	}
	book.Author = author
	json, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(json)
}