package main

import "fmt"

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.title = "Go语言2"
}

func changeBook2(book *Book) {
	book.title = "Go语言3"
}

func main() {
	var book1 Book
	book1.title = "Go语言"
	book1.auth = "Google"

	fmt.Printf("book1.title = %s, book1.auth = %s\n", book1.title, book1.auth)

	changeBook(book1)

	fmt.Printf("book1.title = %s, book1.auth = %s\n", book1.title, book1.auth)

	changeBook2(&book1)

	fmt.Printf("book1.title = %s, book1.auth = %s\n", book1.title, book1.auth)
}
