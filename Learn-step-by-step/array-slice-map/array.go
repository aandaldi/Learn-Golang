package main

import (
	"fmt"
)

func main() {
	// declare simple array 1
	var arr1 [2]string
	arr1[0] = "hello"
	arr1[1] = "world"

	fmt.Println(arr1) // [hello world]

	// declare simple array 2
	var arr2 = [4]int{4, 6, 45, 3}
	fmt.Println(arr2) //[4, 6, 45, 3]

	// declare simple array 3
	var arr3 = [3]string{
		"welcome",
		"to",
		"my world",
	}
	fmt.Println(arr3) // [welcome to my world]

	// declare array without limit
	arr4 := [...]float32{4.234, 65.2342, 3.23, 2.0234}
	fmt.Println(arr4) // [4.234 65.2342 3.23 2.0234]

	// declare an array multi dimensions
	arr5 := [2][3]int{
		{3, 4},
		{1, 2, 3},
	}
	fmt.Println(arr5) //	[[3 4 0] [1 2 3]] if any value is not input, will be replace with defaul value

	arr6 := [...][3]int{
		[3]int{2, 3},
		[3]int{4, 2},
	}
	fmt.Println(arr6) //	[[2 3 0] [4 2 0]]

	// using simple for loop in array 1

	arr7 := [...]int{3, 4, 5, 3}
	for i := 0; i < len(arr7); i++ {
		fmt.Print(arr7[i]) //	3453
	}
	fmt.Println("")

	// using simple for loop in array 2 with range
	arr8 := [4]string{"banana", "apple", "grape", "mellon"}
	for i, fruit := range arr8 {
		fmt.Printf("fruit index %v is %v \n", i, fruit)
	}

	// using simple for loop in array 2 with range and keyword (_)
	for _, fruit := range arr8 {
		fmt.Printf("fruit is %v without index number \n", fruit)
	}

	// declare element array with keword make
	arr9 := make([]int, 2)
	arr9[0] = 3
	arr9[1] = 9

	fmt.Println(arr9) // [3 9]
}
