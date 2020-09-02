package main

import "fmt"

func main() {
	// create variable with var and data type
	var exist bool
	exist = true
	var valid bool = false

	fmt.Println(exist, valid)

	// create multi variable with := and keyword new

	active, is_false := true, new(bool)
	*is_false = false

	fmt.Println(active, is_false, *is_false)

	// using fmt.Printf() function
	fmt.Printf("the value is %t \n", valid)
}
