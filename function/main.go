package main
import "fmt"

/*
// example: 1
func add(x int, y int) int {
	r := x + y
	return r
}

// example:2
func add(x , y int) int {
	r := x + y
	return r
}

// example:3
func add(x , y int) (r int){
	r = x + y
	return r
}

// example:4
func add(x , y int) (r int){
	r = x + y
	return
}
*/

func main()  {
	a := add(10, 15)
	fmt.Println(a)
}