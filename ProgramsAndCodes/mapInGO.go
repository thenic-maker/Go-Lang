package main

import (
	"fmt"
)

func main() {

	var map_1 map[int]string
	fmt.Println(map_1)

	var map_2 = map[int]string{
		1: "Nitin chauhan",
		2: "Nagesh kumar",
	}

	map_2[3] = "Rudra Pratap"
	fmt.Println(map_2)

	//using make function
	m := make(map[int]string)

	m[1] = "apple"
	m[2] = "mango"

	fmt.Println(" Fruits : ", m)

	ch, ok := m[2]

	fmt.Println(" check m[2] is present or not ", ch, ok)

	delete(m, 2)

	ch1, ok1 := m[2]
	fmt.Println(" check m[2] is present or not ", ch1, ok1)

	newmap := m

	newmap[2] = "grapes"
	newmap[3] = "gavava"

	fmt.Println(" new map  Fruits : ", newmap)

	fmt.Println("old map Fruits : ", m)

}
