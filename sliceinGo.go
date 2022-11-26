package main

import "fmt"

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
}
