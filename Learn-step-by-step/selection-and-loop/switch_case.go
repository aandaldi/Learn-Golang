package main

import (
	"fmt"
)

func main() {
	// simple switch case when return in case 3
	num1 := 3

	switch num1 {
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2")
	case 3:
		fmt.Println("value is 3")
	}

	// switch in another way will run default case and using keywor fallthrough

	num := 3

	switch num {
	case 1, 2, 3, 4:
		fmt.Println("value is under 5")
		fallthrough
	case 5, 6, 7, 8, 9:
		fmt.Println("value is under 10")
	default:
		fmt.Println("value is above 9")
	}

	// switch with if-else style
	result := 7

	fmt.Println("result =", result)
	switch {
	case result > 8:
		fmt.Println("your doing very well")
	case (result >= 5) && (result <= 8):
		fmt.Println("your doing well")
	case (result >= 3) && (result < 5):
		fmt.Println("you have to study hard")
	default:
		{
			fmt.Print("Error")
			fmt.Print("you have to following this class again")
		}
	}

	// nested switch and if

	name := "admin"
	id := 7

	if name == "admin" {
		switch id {
		case 5:
			fmt.Println("welcome ", name, " wiht id 5")
		case 6:
			fmt.Println("welcome ", name, " wiht id 6")
		case 7:
			fmt.Println("welcome ", name, " wiht id 7")
		}
	} else {
		fmt.Println("your not admin")
	}

}
