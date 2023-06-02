package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty, cannot dequeue the value.")
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty, cannot dequeue the value.")
	}
	value := q.data[0]
	return value, nil
}

func main() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Is queue empty: ", queue.IsEmpty())

	value, err := queue.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front value: ", value)
	}

	value, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front value: ", value)
	}
}
