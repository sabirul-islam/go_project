package main
import "fmt"
import "reflect"

func main(){
	// slice method
	var fruits []string
	fruits = append(fruits, "apple", "banana", "orange")
	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fmt.Println(len(fruits))

	// checking default value
	x := make([]string, 0)
	fmt.Println(x)

	// slice method with value
	y := make([]string, 0)
	y = append(y, "dal", "vat", "dim_vaja")
	fmt.Println(y)

	// data type
	fmt.Printf("%T", fruits)

	// dtat type with reflect tool
	a := reflect.TypeOf(fruits).Kind().String()
	fmt.Println(a)
}