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

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var signals = []string{"test"}
// var wg sync.WaitGroup
// var mutex sync.Mutex

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go nitin("Nitin Chauhan")
// 		wg.Add(1)
// 	}
// 	wg.Wait()
// 	fmt.Println(signals)
// }

// func nitin(name string) {
// 	mutex.Lock()
// 	signals = append(signals, name)
// 	mutex.Unlock()
// 	fmt.Println(name)
// 	defer wg.Done()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup
// var mutex sync.Mutex
 
// func main() {
// 	go rudra("Rudra")
// 	go rudra("Pratap")
// 	go rudra("singh")
// 	wg.Add(2)
// 	wg.Wait()
// }

// func rudra(s string) {
// 	mutex.Lock()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// 	mutex.Unlock()

// 	defer wg.Done()
// }
