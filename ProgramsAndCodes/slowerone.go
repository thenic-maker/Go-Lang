package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	sample:= ""
	for x:=0;x<10;x++{
		sample+= "0"
	}
	fmt.Println(sample)

	var buffer bytes.Buffer
	for x := 0; x < 10; x++ {
		buffer.WriteString("0")
	}
	fmt.Println(buffer.String())

	sample1 := []string{"0","0","0","0","0","0","0","0","0","0"}
	fmt.Printf(strings.Join(sample1, ""))
}
