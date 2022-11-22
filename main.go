package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {

	s := []int{345, 78, 135, 10, 76, 2, 57, 5}
	sort.Ints(s)
	fmt.Println("sorted slices :", s)
	fmt.Println("Index Values :", strings.Index("Hello this is Nitin Chauhan", "au"))
	fmt.Println("Time", time.Now().Local())
}
