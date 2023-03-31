package main

import (
	"fmt"
)

func main() {
	row := 4

	a, b := 0, 1
	for i := 0; i < row; i++ {
		for j := 1; j < row-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print(a, " ")
			a, b = b, a+b
		}

		fmt.Println()
	}
}
