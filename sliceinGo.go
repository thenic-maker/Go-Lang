package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	//array
	myarr := [5]string{"Hello ,", "This", "is", " Nitin", "Chauhan"}

	fmt.Println("MyArr", myarr)

	// slice
	myslice := myarr[1:5]
	fmt.Println("My Slice ", myslice)

	//make( []type , len ,cap)
	var myslice1 = make([]int, 5, 10)

	//Iterate slice
	for i := 0; i < len(myslice1); i++ {
		fmt.Scanf("%d", &myslice1[i])
	}
	fmt.Println("MY SLICE USING LOOP AND MAKE METHOD", myslice1)

	myslice2 := []string{"This", "is", "the", "NITIN",
		"CHAUHAN"}

	// Iterate slice
	// using range in for loop
	for index, ele := range myslice2 {
		fmt.Printf("Index = %d and element = %s\n", index, ele)
	}

	for _, ele1 := range myslice2 {
		fmt.Printf("Element : %s \n", ele1)
	}

	// multi- dimensional slice
	s1 := [][]int{{1, 2},
		{3, 4},
		{5, 6}}

	fmt.Println("printing the s1 slice:", s1)

	// this is the operation for the sorting
	slc1 := []string{"s", "m", "i", "l", "e"}
	slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

	fmt.Println("Before sorting:")
	fmt.Println("Slice 1: ", slc1)
	fmt.Println("Slice 2: ", slc2)

	// Performing sort operation on the
	// slice using sort function
	sort.Strings(slc1)
	sort.Ints(slc2)

	fmt.Println("\nAfter sorting:")
	fmt.Println("Slice 1: ", slc1)
	fmt.Println("Slice 2: ", slc2)

	// trim of slice
	slice1 := []byte("***hello trim slice***")
	slice2 := []byte{'*', '*', 'a', 'b', 'c', '@', '@'}

	result1 := bytes.Trim(slice1, "*")
	result2 := bytes.Trim(slice2, "*@")
	result3 := bytes.Trim([]byte("!!!hello nitin chauhan $$$"), "!&")

	fmt.Printf("\nresult1 : %s", result1)
	fmt.Printf("\nresult2 : %s", result2)
	fmt.Printf("\nresult3 : %s", result3)

}
