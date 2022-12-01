package main

import "fmt"

func main() {

	var str1 = "Hello, Nitin Chauhan"
	// RAW LITERALS
	var str2 = `Hello, Nitin Chauhan`

	var str3 = "Hello,\nNitin Chauhan"
	// RAW LITERALS
	var str4 = `Hello,\nNitin Chauhan`

	fmt.Println("string1 : ", str1)
	fmt.Println("string2 : ", str2)
	fmt.Println("string3 : ", str3)
	fmt.Println("string4 : ", str4)

	str1 = "how are you?"
	fmt.Printf("string1 : %c", str1[5])
}
