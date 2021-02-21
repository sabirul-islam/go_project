package main

import (
	"fmt"
)

func main()  {
	a := func (reverse string) (b string) {
		for i, character := range reverse{
			var reversed = ""
			reversed = string(character) + reversed
			fmt.Println(i,reversed,string(character))
		}
		return
	}("shimul")
	
	fmt.Println(a)
}