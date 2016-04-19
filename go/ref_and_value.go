package main 

import (
        "fmt"
)

type Book struct {
        name string
}

func main() {
        book := getBook()
        book.name = "bar" 
        fmt.Println(book)  
        fmt.Printf("%p\n", &book) 
}

func getBook() Book {
        book := Book{"foo"}
        fmt.Printf("%p\n", &book)
        return book
}

