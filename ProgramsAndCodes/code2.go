package main

import "fmt"

func main() {

	var arr [5]int
	fmt.Println("emp:", arr)

	arr[2] = 25
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])
	fmt.Println("len:", len(arr))

	arr2 := [5]int{12, 23, 34, 45, 56}
	fmt.Println("dcl:", arr2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
