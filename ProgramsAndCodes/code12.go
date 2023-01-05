package main

import (
	"fmt"
)

func main() {
	planet := []string{"mercury", "mars", "jupiter", "bb", "ba", "saturn", "uranus", "venus", "earth", "ambr", "a", "anil"}
	fmt.Println(planet)
	fmt.Println(customSorting(planet))
}
func customSorting(strArr []string) []string {
	//taking temp slice to store all elements present in strArr array
	temp := make([]string, len(strArr), len(strArr))
	count := 0 //variable to fill elements one by one in temp array
	// first storing elements if they are of odd length
	for i := 0; i < len(strArr); i++ {
		if (len(strArr[i]) % 2) != 0 {
			temp[count] = strArr[i]
			count++
		}
	}
	mid := count // iterate 0 to mid (here all odd elements reside b/w 0 to mid)
	// sorting odd length elements in ascending order according to their length
	for i := 0; i < mid-1; i++ {
		for j := 0; j < mid-1; j++ {
			var left string = temp[j]
			var right string = temp[j+1]
			if len(right) < len(left) {
				temp[j+1], temp[j] = left, right
			}
			if len(right) == len(left) { //if both elements have equal length then whomever element comes first in alphabetical order will be placed first
				if right < left {
					temp[j+1], temp[j] = left, right
				}
			}
		}
	}
	//Now storing elements if they are of even length and appending it to existing temp slice
	for i := 0; i < len(strArr); i++ {
		if (len(strArr[i]) % 2) == 0 {
			temp[count] = strArr[i]
			count++
		}
	}
	high := count // iterate mid to high (here all even elements reside b/w mid to high)
	//Now sorting even length elements in descending order
	for i := mid; i < high-1; i++ {
		for j := mid; j < high-1; j++ {
			var left string = temp[j]
			var right string = temp[j+1]
			if len(right) > len(left) {
				temp[j+1], temp[j] = left, right
			}
			if len(right) == len(left) { //if both elements have equal length then whomever element comes first in alphabetical order will be placed first
				if right < left {
					temp[j+1], temp[j] = left, right
				}
			}
		}
	}
	return temp
}
