package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {

	book := Book{
		Title:  "The go programming language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
	}
	fmt.Printf("Title : %s \nAuthor : %s \nPages : %d \n", book.Title, book.Author, book.Pages)

}
