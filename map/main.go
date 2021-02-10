package main
import "fmt"

func main(){
	// default value
	var y map[string]string
	fmt.Println(y)

	// map
	x := make((map[string]string))
	x["name"] = "shimul"
	x["height"] = "5.8'"
	x["age"] = "33"
	x["address"] = "jamalpur"
	fmt.Println(x)
	fmt.Println(x["address"])

	// delete from map
	delete(x, "height")
	fmt.Println(x)
}