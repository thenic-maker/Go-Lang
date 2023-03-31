package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numWords = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func numToWords(num int) string {
	if num < 0 || num > 9999999 {
		return "invalid number"
	}
	if num == 0 {
		return "zero"
	}
	words := ""
	if num/100000 > 0 {
		words += numToWords(num/100000) + " lakh "
		num %= 100000
	}
	if num/1000 > 0 {
		words += numToWords(num/1000) + " thousand "
		num %= 1000
	}
	if num/100 > 0 {
		words += numWords[num/100] + " hundred "
		num %= 100
	}
	if num > 0 {
		if words != "" {
			words += "and "
		}
		if num < 20 {
			words += numWords[num]
		} else {
			words += numWords[num/10*10]
			if num%10 > 0 {
				words += "-" + numWords[num%10]
			}
		}
	}
	return strings.TrimSpace(words)
}

func main() {
	num := 9999909
	str := strconv.Itoa(num)
	words := numToWords(num)
	fmt.Printf("%s in words is %s\n", str, words)
}
