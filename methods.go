package main

import "fmt"

//Method with struct
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Method with non Struct
type data int

// Method with Pointer reciever
type team struct {
	name    string
	country string
	salary  int
}

func main() {
	// Method with struct
	res1 := author{
		name:      "Nitin Chauhan",
		branch:    "CSE",
		particles: 100,
		salary:    35000,
	}
	res1.show()

	// Method with non struct
	value1 := data(10)
	value2 := data(13)

	res2 := value1.multiply(value2)
	fmt.Println("multiply of two number : ", res2)

	// Method with Pointer reciever
	res3 := team{
		name:    "Virat Kolhi",
		country: "India",
		salary:  100000,
	}

	fmt.Println("name: ", res3.name)
	fmt.Println("country: ", res3.country)
	fmt.Println("salary : ", res3.salary)

	p := &res3

	p.info(200000)

	fmt.Println("name: ", res3.name)
	fmt.Println("country: ", res3.country)
	fmt.Println("salary : ", res3.salary)

}

// Method with struct
func (a author) show() {
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Method with non struct
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

// Method with Pointer reciever
func (t *team) info(newsalary int) {
	t.salary = newsalary
}
