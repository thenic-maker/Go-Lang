package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{12, 6, 76, 98, 5}

	sort.Ints(arr)

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i], len(arr))
	}

	// multi-dimensionla array
	newarr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}

	// array Using for loop
	fmt.Println("Elements of Array 1")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {

			fmt.Println(newarr[x][y])
		}
	}

	// multi-dimensional array using index
	var arr1 [2][2]int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	// Accessing the values of the array
	fmt.Println("Elements of array 2")
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr1[p][q])
		}
	}

	// ellipse
	myarr := [...]int{26, 456, 386, 6, 368}
	sort.Ints(myarr[:])
	fmt.Println("ellipse len and and its element: ", len(myarr), myarr)

	newArr1 := myarr
	newArr2 := &myarr

	fmt.Println("ARRAY BY VALUE ", newArr1)
	fmt.Println("ARRAY BY REFERENCE ", *newArr2)

	myarr[0] = 5

	fmt.Println("ARRAY BY VALUE ", newArr1)
	fmt.Println("ARRAY BY REFERENCE ", *newArr2)

}
