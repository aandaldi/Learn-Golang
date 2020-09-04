package main

import (
	"fmt"
	"strings"
)

func main() {
	// declare simple map
	var user1 = map[string]string{}
	user1["name"] = "muhammad"
	user1["class"] = "12"

	fmt.Println(user1) // map[class:12 name:muhammad]

	// declare with keyword make
	var user2 = make(map[string]string)
	user2["name"] = "ali"
	user2["class"] = "11"

	fmt.Println(user2)

	// loop the map with for loop
	var user3 = map[string]string{}
	user3["name"] = "usman"
	user3["class"] = "12"
	user3["gender"] = "Male"

	for key, value := range user3 {
		fmt.Println(key, " is ", value)
	}

	// remove an item in map
	delete(user1, "name")
	fmt.Println(user1) //	map[class:12]

	// combine slice with map
	var slice1 = []map[string]string{
		{"name": "ahmad", "class": "3"},
		{"name": "Aisyah"},
		{"name": "mustofa", "class": "6", "gender": "male"},
	}
	fmt.Println(slice1) // [map[class:3 name:ahmad] map[name:Aisyah] map[class:6 gender:male name:mustofa]]

	for _, el := range slice1 {
		for key, value := range el {
			if key == "name" {
				fmt.Printf("=============== %v ===============\n", strings.ToUpper(value))
			}
			fmt.Printf("%v is %v \n", strings.Title(key), strings.Title(value))
		}
		fmt.Println()
	}
}
