package main

import "fmt"

func main() {

	fmt.Println("this is the program of defer keywords")

	defer fmt.Println("this is first defer")
	defer fmt.Println("this is second defer")
	defer fmt.Println("this is third defer")
	defer add(10, 23)
	defer add(485, 65)
}

func add(n1 int, n2 int) int {

	fmt.Println("sum of number", n1+n2)
	return 0
}
