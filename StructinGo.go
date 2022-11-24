package main

import "fmt"

type employee struct {
	name   string
	salary int
}

// Creating structure
type Student struct {
	name   string
	branch string
	year   int
}

// Creating nested structure
type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

func main() {

	a := &employee{"Nitin chauhan", 999999}

	fmt.Println("Name: ", (*a).name)
	fmt.Println("Salary: ", a.salary)

	stu := Student{"Nitin Chauhan", "CSE", 2019}

	teachr := Teacher{"Manoj", "C#", 10, stu}

	fmt.Println("Student: ", stu)
	fmt.Println("Teacher: ", teachr)

}
