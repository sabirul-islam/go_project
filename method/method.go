package main

import "fmt"

type Doctor struct{
	name string
	education string
	age int
	salary float32
}

func (d Doctor) Speak()  {
	fmt.Println(d.name, "can speak")
}

func (d Doctor) Salary() float32 {
	return d.salary
}

func main()  {

	// var d = Doctor {"tarek", "mbbs", 35, 65.98}

	/*
	var d = Doctor {
		name : "tarek",
		education : "mbbs",
		age : 36,
		salary: 89.88, 
	}
	*/

	var d Doctor
	d.name = "tarek"
	d.age = 35
	d.education ="mbbs"
	d.salary =5065.89

	fmt.Print(d.name, d.education, d.age)
	d.Speak()
	fmt.Println(d.Salary())
}