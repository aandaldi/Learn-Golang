package main

import "fmt"

func main() {
	// just using if
	if 1 == 1 {
		fmt.Println("it will be print out")
	}

	// another example just use if
	if 1 == 2 {
		fmt.Println("this will never run")
	}

	// using if and else
	name := "admin"

	if name == "admin" {
		fmt.Printf("welcome %s \n", name)
	} else {
		fmt.Println("your not admin")
	}

	// using if, else if, and else
	result := 8
	fmt.Printf("your result = %v \n", result)

	if result > 7 {
		fmt.Println("your did very well")
	} else if result >= 5 && result < 7 {
		fmt.Println("your did well")
	} else {
		fmt.Println("you have to study hard")
	}
}
