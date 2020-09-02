package main

import (
	"fmt"
)

func main() {
	// create variable with var and data type
	var firstNum float64
	firstNum = 1.1
	var secondNum float32 = 2.2

	fmt.Println(firstNum, secondNum) // 1.1 2.2

	// create variable without data type, = with var or :=
	var thirdthNum = 3.210
	fourthNum := 4.321

	fmt.Println(thirdthNum, fourthNum) // 3.21 4.321

	// create multivarible
	var fifthNum, sixthNum = 5.4321, 6.54321
	seventhNum, eightNum := 7.654321, 8.7654321

	fmt.Println(fifthNum, sixthNum)   // 5.4321 6.54321
	fmt.Println(seventhNum, eightNum) // 7.654321 8.7654321

	// treat the unused variable with _
	_ = 8
	var ninthNum, _ = 9.87654321, 100.000323
	_, tenth := 1000.2342, 10.98765321

	fmt.Println(ninthNum, tenth) // 9.87654321 10.98765321

	// declare with 'new' keyword
	var eleventhNum = new(float32)
	*eleventhNum = 11.000340

	fmt.Println(eleventhNum)  // 0xc000014108
	fmt.Println(*eleventhNum) // 11.000340 dereference

	var twelfthNum = 12.293845234
	fmt.Printf("ini integer to string %f \n", twelfthNum)   //this is float to string 12.293845234
	fmt.Printf("ini integer to string %.2f \n", twelfthNum) //this is float to string 12.29

}
