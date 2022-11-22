package main

import "fmt"

//Mehods with struct
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Methods wit non Struct
type data int

func main() {
	res1 := author{
		name:      "Nitin Chauhan",
		branch:    "CSE",
		particles: 100,
		salary:    35000,
	}
	res1.show()

	value1 := data(10)
	value2 := data(13)

	res2 := value1.multiply(value2)
	fmt.Println("multiply of two number : ", res2)
}

func (a author) show() {
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
