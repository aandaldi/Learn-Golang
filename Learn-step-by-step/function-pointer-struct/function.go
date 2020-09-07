package main

import "fmt"

// declare simple function1
func printName() {
	var name = "Jhon"
	fmt.Println("Name is ", name)
}

// declare simple function2 with param
func printStudent(name string, class int) {
	fmt.Println("================Student====================")
	fmt.Println("Name : ", name)
	fmt.Println("Class : ", class)
}

// declare simple function3 with return value
func returnAddition(val1, val2 int) int { //can create param with same data type like this
	fmt.Println("=================Addition===================")
	result := val1 + val2

	return result
}

// declare simple function4 with multiple return value
func returnAdditionAndSubtraction(val1, val2 int) (int, int) { //can create param with same data type like this
	fmt.Println("=================Addition and Substraction===================")
	add := val1 + val2
	sub := val1 - val2

	return add, sub
}

// declare simple function5 with variadic function
func calculateAddition(numbers ...int) int {
	add := 0
	for _, num := range numbers {
		add += num
	}
	return add
}

func main() {
	// call simple funtion1
	printName()

	// call simple function2 with Param
	printStudent("willy", 23)

	// declare simple function3 with return value
	var addition1 = returnAddition(4, 6)
	fmt.Println(addition1)

	// declare simple function4 with return value
	var resAdd, resSub = returnAdditionAndSubtraction(4, 6)
	fmt.Println("result Add ", resAdd, " and result Sub ", resSub)

	// declare simple function5 with variadic function
	resCalculate := calculateAddition(2, 4, 8, 3, 6)
	fmt.Println("Result add calculate is ", resCalculate)

}
