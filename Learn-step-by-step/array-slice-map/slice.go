package main

import (
	"fmt"
)

func main() {
	// declare simple slice
	slice1 := []string{"hai", "world"}
	fmt.Printf("%v, %T \n", slice1, slice1)

	// function len() and cap
	fmt.Println(len(slice1), cap(slice1))

	// test update value of slice
	fmt.Println(slice1[0]) //hai
	slice1[0] = "hello"
	fmt.Println(slice1[0]) // hello and the next, every value with value referce is slice 1, will be update

	// function copy
	slice2 := []string{}
	copy(slice2, slice1)
	fmt.Println(slice2) // return [] because len of slice2 is nill when initialize

	slice3 := []string{"hai"}
	copy(slice3, slice1)
	fmt.Println(slice3) // return [hello], just copy element with limit of slice3

	slice4 := []string{"hai", "this is mike", "welome to my world"}
	copy(slice4, slice1)
	fmt.Println(slice4) //return [hello world welcome to my world]

	// funtion appen
	appendSlice := append(slice1, "wow great")
	fmt.Println(appendSlice, slice1)
}
