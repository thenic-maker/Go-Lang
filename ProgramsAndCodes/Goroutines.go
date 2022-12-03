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

	Disp("Nitin chauhan")
}
