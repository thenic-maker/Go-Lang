package main

import (
	"fmt"
	"time"
)

func Disp(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func main() {
	go Disp("hello")

	Disp("Nitin Chauhan")

	// anonymous go routines
	go func() {

		fmt.Println("this is anonymous Goroutine")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}
