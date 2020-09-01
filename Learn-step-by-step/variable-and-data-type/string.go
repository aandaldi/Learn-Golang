package main

import "fmt"

func main() {
	var firstName = "aan"
	var lastName string
	lastName = "aldi"

	fmt.Println("Your Name: ", firstName, lastName) // your name:  aan aldi

	// using fmt.Printf() function

    fmt.Printf("Welcome \n") //Welcome
	fmt.Printf("Halo %s %s \n", firstName, lastName) // Halo aan aldi 
	
	// declare variabel with no data type
	var fName = "aan"
	lName := "aldi" 

	fmt.Println(fName +" "+ lName)	//aan aldi
}
