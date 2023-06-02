package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(data int) {
	s.data = append(s.data, data)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty ,Cannot pop value")
	}
	size := len((s.data)) - 1
	value := s.data[size]
	s.data = s.data[:size]

	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty ,Cannot peek value")
	}
	index := len(s.data) - 1
	return s.data[index], nil
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	fmt.Println("Is stack empty: ", stack.IsEmpty())

	value, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top value: ", value)
	}

	value, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped value: ", value)
	}

	for _, val := range stack.data {
		fmt.Println(val)
	}
}
