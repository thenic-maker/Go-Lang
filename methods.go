package main

import "fmt"

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func main() {
	res := author{
		name:      "Nitin Chauhan",
		branch:    "CSE",
		particles: 100,
		salary:    35000,
	}
	res.show()
}

func (a author) show() {
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}
