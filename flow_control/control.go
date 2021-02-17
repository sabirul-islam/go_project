package main

import (
	"fmt"
)

func main()  {
	fmt.Print("Please enter your age:")
	var age int
	fmt.Scanf("%d", &age)

	//// If else
	// if age<3{
	// 	fmt.Println("infant")
	// }else if age<7 {
	// 	fmt.Println("baby")
	// }else if age<18 {
	// 	fmt.Println("teen")
	// }else if age<33 {
	// 	fmt.Println("adult")
	// }else{
	// 	fmt.Println("old")
	// }

	// //Switch Case
	// switch age {
	// case 2:
	// 	fmt.Println("baby")
	// case 9:
	// 	fmt.Println("old")
	// default:
	// fmt.Println("adult")
	// }
	
	// //For loop
	// for i := 1; i <= age; i++ {
	// 	fmt.Println(i)
	// }

	// Range For Loop
	students := []string{"sabirul", "shimul", "shumit"}
	for i, value := range students{
		fmt.Println(i, value)
	}

	// while loop
	i := 1
	for i <= age {
		fmt.Println(i, "shimul")
		i++
	}
}