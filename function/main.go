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

// Multiple Value
func rectangle(l int, b int) (parameter int, area int){
	parameter = 2 * (l + b)
	area = l * b
	return
}

// Pointer
func update(a *int, t *string)  {
	*a = *a + 5
	*t = *t + "Islam"
	return
}
*/

func main()  {
	/*
	// for function
	a := add(10, 15)
	fmt.Println(a)
	
	// for multiple value
	a, p := rectangle(10, 10)
	fmt.Println(a, p)
	*/

	/*
	// for pointer
	number := 10
	name := "Sabirul "
	update(&number, &name)
	fmt.Println(number, name)
	*/

	/*
	// Anonumous Function example:1
	a := func (x, y int) (r int) {
		r = x * y
		return
	}

	fmt.Println(a(10,10))
	*/

	// Anonumous Function example:2
	a := func (x, y int) (r int) {
		r = x * y
		return
	}(10,10)

	fmt.Println(a)
}