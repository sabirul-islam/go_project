package main
import "fmt"

func main(){
	// // lonhand way
	// var students[3]string
	// students[0] = "shimul"
	// students[1] = "sabirul"
	// students[2] = "shumit"
	// fmt.Println(students)
	// fmt.Println(len(students))
	// fmt.Println(students[2])

	// shorthand way
	// students := [3]string{"shimul", "sabirul", "shumit"}
	students := [...]string{"shimul", "sabirul", "shumit"}
	fmt.Println(students)
	fmt.Println(len(students))
	fmt.Println(students[0])
}