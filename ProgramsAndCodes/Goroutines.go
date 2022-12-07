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

func portal1(channel1 chan string) {
	time.Sleep(1 * time.Second)
	channel1 <- "welcome to channel1"

}

func portal2(channel1 chan string) {
	time.Sleep(1 * time.Second)
	channel1 <- "welcome to channel2"
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
	// select in go lang
	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal1(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)

	case op2 := <-R2:
		fmt.Println(op2)

		//default:
		//	fmt.Println("not found")
	}

}
