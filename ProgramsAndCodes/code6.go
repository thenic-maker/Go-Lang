package main

import "fmt"

func vals() (int, int, int) {
	return 3, 7, 56
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func tmultiple(num int) func() int {
	i := 0
	return func() int {
		i++
		return num * i
	}
}
func main() {

	a, b, d := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	s := make([]int, 3)
	s[0], s[1], s[2] = vals()
	fmt.Println(s)

	nextInt := intSeq()
	//closures as an anonmyous
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	tnum := tmultiple(8)
	for i := 0; i < 10; i++ {
		fmt.Println(tnum())
	}
}
