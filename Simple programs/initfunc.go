package main

import "fmt"

func init() {
	fmt.Println("Hello first init function")
}

func init() {
	fmt.Println("Hello second init function ")
}

func main() {

	fmt.Println("hello MAIN function ")
}
