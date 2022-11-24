package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{12, 6, 76, 98, 5}

	sort.Ints(arr)

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
}
