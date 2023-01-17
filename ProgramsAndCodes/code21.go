package main

import (
	"fmt"
)

// func main() {
// 	mess := make(chan string)
// 	done := make(chan bool)
// 	go func() { mess <- "Nitin chauhan" }()
// 	go func() { fmt.Println("this is annymous goroutines") }()
// 	go hello(done)
// 	fmt.Println(<-done)
// 	fmt.Println(<-mess)
// }

// func hello(done chan bool) {
// 	fmt.Println("this is hello function")
// 	done <- true
// }


// func main(){
// 	var mywg sync.WaitGroup
// 	mywg.Add(1)
// 	go func (){
// 		fmt.Print("A ")
// 		mywg.Done() 
// 	}()
// 	mywg.Add(1)
// 	go func (){
// 		fmt.Print("B ")
// 		mywg.Done() 
// 	}()
	
// 	mywg.Wait()
// 	fmt.Print("C ")
// }

func sq(num int , sqch chan int){
   sum:=0
	for num != 0 {
     digit:= num % 10
	 sum += digit*digit
	 num /= 10
	}
	sqch <- sum 
}


func cube(num int , cubech chan int){
	sum:=0
	 for num != 0 {
	  digit:= num % 10
	  sum += digit*digit*digit
	  num /= 10
	 }
	 cubech <- sum
 }
func main(){
	number:= 123

	sqch := make(chan int)
	cubech := make(chan int)
    go sq(number,sqch)
    go cube(number, cubech)
	fmt.Println(<-sqch)
	fmt.Println(<-cubech)
}