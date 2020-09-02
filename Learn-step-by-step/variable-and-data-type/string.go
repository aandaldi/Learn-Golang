package main

import "fmt"

func main() {
	// declare with var and data type:
	var str1 string
	str1 = "one"
	var str2 string = "two"
	fmt.Println(str1, str2)

	// declare without data type
	var str3 = "three"
	str4 := "four"

	fmt.Println(str3, str4)

	// declare multivariable
	var str5, str6 = "five", "six"
	str7, str8 := "seven", "eight"

	fmt.Println(str5, str6, str7, str8)

	// using variable underscore (_) for unused variable
	var str100 = "100"
	var _ = str100

	// declare with keyword new
	var str9 = new(string)
	*str9 = "nine"
	fmt.Println(*str9, str9)

	// using fmt.Prinf() function

	fmt.Printf("the value %s in memory address %v \n", *str9, str9) //the value nine in memory address 0xc000010280
	// %v for print value from every data type and %T for print the value data type
}
