package main

import "fmt"

type Employee struct {
	name  string
	empId int
}

func pointerptr(a *int) *int {
	number := 342
	a = &number
	return a
}
func main() {

	// taking a pointer
	num := 1213
	var p *int
	p = &num

	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)
	fmt.Println("&p = ", &p)
	fmt.Println("num = ", num)
	fmt.Println("&num = ", &num)
	fmt.Println("p == &num", p == &num)
	fmt.Println("Pointer in functions  : ", *(pointerptr(p)))

	//pointer to structure
	emp1 := Employee{"Nitin Chauhan", 15}
	emp2 := Employee{"Nagesh Kumar", 16}
	emp3 := Employee{"Rudra Pratap Singh", 17}

	p1 := &emp1
	p2 := &emp2
	p3 := &emp3

	fmt.Println("EMPLOYEE 1 ", p1.name)
	fmt.Println("EMPLOYEE 2 ", &p2.empId)
	fmt.Println("EMPLOYEE 3 ", &p2)
	fmt.Println("EMPLOYEE 1 ", (*p1).name)
	fmt.Println("EMPLOYEE 2 ", *&p2.empId)
	fmt.Println("EMPLOYEE 3 ", *p3)

}
