package main

import "fmt"

func main() {
	// create with var and data type
	var firstNum int
	var secondNum int = 11
	firstNum = 10
	fmt.Println(firstNum, secondNum) // 10 11

	// create variable without data type, = or :=
	var thirdNum = 13
	fourthNum := 14

	fmt.Println(thirdNum, fourthNum) // 13 14

	// create multivariable
	var fifthNum, sixsthNum = 5, 6
	seventhNum, eighthNum := 7, 8

	fmt.Println(fifthNum, sixsthNum)   // 5 6
	fmt.Println(seventhNum, eighthNum) // 7 8

	// treat the unused variable with _
	_ = 8
	var ninthNum, _ = 9, 8
	_, tenthNum := 11, 10

	fmt.Println(ninthNum, tenthNum) // 9 10

	// declare with new keyword
	var eleventhNum = new(int)
	*eleventhNum = 11         // assign value to pointer
	fmt.Println(eleventhNum)  // 0xc000190008
	fmt.Println(*eleventhNum) // 11 dereference

	// using fmt.Printf() function

	var twelfthNum = 12
	fmt.Printf("ini integer to string %d \n", twelfthNum) //ini integer to string 12

}
