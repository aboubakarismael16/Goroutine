package main

import "fmt"

type Book struct {
	id int
}

func main() {
	bookSelf := make(chan Book, 3)
	for i := 0; i < cap(bookSelf) * 2; i++ {
		select {
		case bookSelf <- Book{id: i} :
			fmt.Println("Succeed put a book", i)
		default:
			fmt.Println("Failed to put a new book")
		}
	}

	for i := 0; i < cap(bookSelf) * 2; i ++ {
		select {
		case book := <-bookSelf:
			fmt.Println("Succeed to get book",book.id )
		default:
			fmt.Println("Failed to get a book")
		}
	}
}
