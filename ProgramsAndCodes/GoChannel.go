package main

import "fmt"

func myfunc(chan1 chan int) {
	fmt.Println(230 + <-chan1)
}

func myfunc2(chan2 chan string) {
	for i := 0; i < 4; i++ {
		chan2 <- "hello nitin chauhan"
	}
	close(chan2)
}

func main() {

	var ch1 chan int

	ch2 := make(chan int)

	fmt.Println("Value of the channel: ", ch1)
	fmt.Printf("Type of the channel: %T \n", ch1)

	fmt.Println("Value of the channel: ", ch2)
	fmt.Printf("Type of the channel: %T \n", ch2)

	//var chan1 chan int
	chan1 := make(chan int)
	go myfunc(chan1)
	chan1 <- 30

	chan2 := make(chan string)

	go myfunc2(chan2)

	for {
		res, ok := <-chan2
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
