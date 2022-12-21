package main

import "fmt"

func main() {
	// complex and imaginary
	a := complex(2, 5)

	fmt.Printf(" %v  %v  %v ", a, real(a), imag(a))
}
