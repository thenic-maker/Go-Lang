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

	// blocking and send
	chan3 := make(chan string)

	// Anonymous goroutine
	go func() {
		chan3 <- "hello"
		chan3 <- "nitin"
		chan3 <- "chauhan"
		chan3 <- "EE"
		close(chan3)
	}()

	// Using for loop
	for res := range chan3 {
		fmt.Println(res)
	}
	fmt.Println(len(chan3))

	mychnl := make(chan string, 5)
	mychnl <- "hello"
	mychnl <- "nitin"
	mychnl <- "chauhan"
	mychnl <- "EE"

	// Finding the length of the channel
	// Using len() function
	fmt.Println("Length of the channel is: ", len(mychnl))
	fmt.Println("Capacity of the channel is: ", cap(mychnl))

	//unidirectional channel
	c1 := make(<-chan string)
	c2 := make(chan<- string)

	fmt.Printf("%T\n", c1)
	fmt.Printf("%T", c2)

}
