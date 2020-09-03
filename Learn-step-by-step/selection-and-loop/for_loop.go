package main

import (
	"fmt"
)

func main() {
	// simple for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// for loop like while concept
	i := 1

	for i <= 5 {
		fmt.Println("no", i)
		i++
	}

	// for loop like do while
	id := 20
	for {
		fmt.Println("id ", id)

		id++
		if id == 25 {
			break
		}
	}

	// using keyword break and continue

	ganjil := 1
	fmt.Print("Bilangan Ganjil ", ganjil)
	for {
		ganjil++
		if ganjil%2 == 0 {
			continue
		} else {
			fmt.Print(", ", ganjil, " ")
		}

		if ganjil == 9 {
			break
		}

	} 
	fmt.Println()

	// using label for
	val := 0

	loopfor:
	for {
		val++
		fmt.Println("val ", val)

		if val == 5 {
			break loopfor
		}
	}

}
