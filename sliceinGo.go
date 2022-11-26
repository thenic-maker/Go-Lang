package main

import "fmt"

func main() {

	myarr := [5]string{"Hello ,", "This", "is", " Nitin", "Chauhan"}

	fmt.Println("MyArr", myarr)

	myslice := myarr[1:5]
	fmt.Println("My Slice ", myslice)
}
