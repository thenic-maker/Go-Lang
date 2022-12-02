package main

import (
	"fmt"
	"strings"
)

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
	/* if you trying to change
	      the value of the string
	      then the compiler will
	      throw an error, i.e,
	     cannot assign to mystr[1]

	   str[1]= 'N'
	*/
	str1 = "how are you?"
	fmt.Printf("string1 : %c", str1[5])
	//iterate
	for index, s := range str3 {
		fmt.Printf("The index of a character %c is %d\n", s, index)
	}

	for i := 0; i < len(str1); i++ {
		fmt.Printf("\nThe bytes of a character %c is %v \n", str1[i], str1[i])
	}
	//trimming of string
	str5 := "**helloNitin&&Chauhan&&"

	res := strings.Trim(str5, "*&")

	fmt.Println("trimming od string : --- ", res)

	//split of string
	str6 := "Hello, My name is Nitin Chauhan"

	res6 := strings.Split(str6, "a")

	fmt.Println("split of string : ", res6)

	res7 := strings.SplitAfter(str6, "a")

	fmt.Println("splitAfter of string : ", res7)

	res8 := strings.SplitAfterN(str6, "a", 3)

	fmt.Println("splitAfterN of string : ", res8)
}
