package main

import "fmt"

func main() {
	a := []int{0, 2, 4, 6}
	b := []int{9, 7, 5, 3}

	c := a
	for _, num := range b {
		c = append(c, num)
	}
	fmt.Println(c)

}
