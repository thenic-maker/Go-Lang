package main

import "fmt"

type employee struct {
	name   string
	salary int
}

func main() {

	a := &employee{"Nitin chauhan", 999999}

	fmt.Println("Name: ", (*a).name)
	fmt.Println("Salary: ", a.salary)
}
