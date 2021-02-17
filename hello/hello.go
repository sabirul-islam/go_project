package main

import "fmt"

func main(){
	// var name string
	// var age int
	// var result float32
	// var isFound bool

	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(result)
	// fmt.Println(isFound)

	// var chr rune
	// chr = 'F'
	// fmt.Println(chr)
	// fmt.Printf("%c", chr)

	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(true || true)
	// fmt.Println(true || false)
	// fmt.Println(!true)


	fmt.Println("Please enter your name and age:")
	var name string
	var age int

	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Your name is %s and age is %d", name, age)

}