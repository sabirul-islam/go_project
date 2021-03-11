package main

import (
	"fmt"
	"os"
	"text/template"
)


func main() {
	tmpl, err := template.ParseFiles("one.gohtml", "two.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, 100.50)
}