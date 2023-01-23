package main

import (
	//"bufio"
	"fmt"
	//"os"
	"strings"
)

func process(input string) {
	// fmt.Println(input)
	wordSlice := strings.Split(input, ", ")
	// fmt.Println(wordSlice)
	availkeys := []byte{'e', 'r', 't', 'a', 'b', 'd', 'f', 'y', 'h'}
	// fmt.Println(availkeys)
	result := ""

	for i := 0; i < len(wordSlice); i++ {
		str := wordSlice[i]
		// fmt.Println(str)
		validinput := ""
		for j := 0; j < len(str); j++ {
			for k := 0; k < len(availkeys); k++ {
				if str[j] == availkeys[k] {
					validinput += string(str[j])
				}
			}
		}
		if validinput == str {
			if len(validinput) >= len(result) {
				result = validinput
			}
		}
	}
	fmt.Println(result)
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	input := "aterdfyh, tad, tabd, hfdbaereftre"
	process(input)
}
