package main

import "fmt"

func main () {
	// structure data type
	type Book struct {
		Title string
		Author string
		ISBN string
		Price float32
		Pages int
	}
	var b1 Book
	b1.Title = "An Introduction To Programming In Go\n"
	b1.Author = "CALEB DOXY\n"
	b1.ISBN = "978-1478355823\n"
	b1.Price = 10.5
	b1.Pages = 165
	fmt.Println(b1)
}
