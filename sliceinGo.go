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

	for i := 0; i < len(myslice1); i++ {
		fmt.Scanf("%d", &myslice1[i])
	}
	fmt.Println("MY SLICE USING LOOP AND MAKE METHOD", myslice1)
}
