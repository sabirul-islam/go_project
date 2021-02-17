package main
import "fmt"

	// Package Scope
	type name string
	var name1 name

func main () {
	// // structure data type
	// type Book struct {
	// 	Title string
	// 	Author string
	// 	ISBN string
	// 	Price float32
	// 	Pages int
	// }
	// var b1 Book
	// b1.Title = "An Introduction To Programming In Go\n"
	// b1.Author = "CALEB DOXY\n"
	// b1.ISBN = "978-1478355823\n"
	// b1.Price = 10.5
	// b1.Pages = 165
	// fmt.Println(b1)

	// Shorthand Method (Anonymous/Literal struct)
	b1 := struct{
		Title string
		Author string
		ISBN string
		Price float32
		Pages int
	}{
		Title: "An Introduction To Programming In Go\n",
		Author: "CALEB DOXY\n",
		ISBN: "978-1478355823\n",
		Price: 10.5,
		Pages: 165,
	}
	fmt.Println(b1, b1.Title)

	// Custom Data Type (Primitive)
	// type weight float32
	var w1 weight
	w1 = 1.5
	fmt.Println(w1)

	// From package scope
	name1 = "shimul"
	fmt.Println(name1)

	


}
