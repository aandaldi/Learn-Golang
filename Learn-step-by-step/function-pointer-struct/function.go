package main

import (
	"fmt"
	"strings"
)

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

// declare simple closure with function is return value
func findMax(numbers []int) func() int {

	var result = func() int {
		max := numbers[0]
		for _, num := range numbers {
			if max < num {
				max = num
			}
		}
		return max
	}
	return result
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

	text := `
		hello
		my name
		is jhon
	`

	fmt.Println(text)
	word := strings.Fields(text)
	fmt.Println(word)

	// declare simple closure function1, save as variable
	var getMinMax = func(numbers []int) (int, int) {
		min, max := numbers[0], numbers[0]
		for _, num := range numbers {
			if min > num {
				min = num
			}
			if max < num {
				max = num
			}
		}
		return min, max
	}

	numbers := []int{2, 3, 42, 3, 5, 5, 2, 5, 66, 34}
	fmt.Println(getMinMax(numbers))

	// declare simple closure with IIFE expression
	var getMin = func(numbers []int) int {
		min := numbers[0]
		for _, num := range numbers {
			if min > num {
				min = num
			}
		}
		return min
	}(numbers)

	fmt.Println(getMin)

	// call funtion simple closure with function as return value
	res := findMax(numbers) //res is function
	fmt.Println(res())
}
