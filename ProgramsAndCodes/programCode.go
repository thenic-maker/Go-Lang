package main

import "fmt"

func sort(arr []int) {
	size := len(arr)
	isSwapped := false
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				isSwapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// vartemp := arr[j]
				// arr[j] = arr[j+1]
				// arr[j+1] = vartemp
			}
		}
		if !isSwapped {
			break
		}

	}
}

func main() {
	// complex and imaginary
	a := complex(2, 5)

	fmt.Printf(" %v \n %v \n %v \n ", a, real(a), imag(a))

	var arr = [...]int{1, 2, 3, 3, 4, 5, 4654, 54}
	arr[2] = 10

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[2:3])

	var myarray [10]int = [10]int{1, 2, 3, 3, 4, 5, 4654, 54}
	fmt.Println(myarray)
	//multi dimension

	matrix := [3][3]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println(matrix)
	fmt.Println(len(matrix))

	//map

	// empsal := make(map[string]int)

	// empsal["Ram"] = 1000000
	// empsal["Shiv"] = 123422

	// fmt.Println(empsal)
	empsal := map[string]int{
		"Ram":  123321,
		"Shiv": 89457,
	}
	fmt.Println(empsal)

	Newarr := [10]int{55, 99, 2, 11, 33, 77, 88, 2, 7, 1}
	sort(Newarr[:])
	for i := 0; i < len(Newarr); i++ {
		fmt.Println(Newarr[i])
	}

	//slicing
	slice := []int{1, 2, 3, 4}

	slice1 := slice

	fmt.Println(slice)
	fmt.Println(slice1)
	// here the concept of slice is simple understand i.e. its copy the address of the array
	// Not only value of array
	slice[1] = 100
	fmt.Println(slice)
	fmt.Println(slice1)

	emp := Employee{
		Id:            12,
		Name:          "Nitin chauhan",
		Company:       "EE",
		Mobile_Number: []string{"9623017823", "7346912873"},
	}
	fmt.Println(emp)

}

// STRUCT IN GO
type Employee struct {
	Id            int
	Name          string
	Company       string
	Mobile_Number []string
}
