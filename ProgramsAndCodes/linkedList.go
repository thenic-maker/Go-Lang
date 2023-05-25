package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if ll.IsEmpty() {
		ll.head = newNode
	} else {
		currNode := ll.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = newNode
	}
}

func (ll *LinkedList) Prepend(data int) {
	newNode := &Node{data: data}

	if ll.IsEmpty() {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

func (ll *LinkedList) InsertAfter(prevNode *Node, data int) {
	if prevNode.next == nil {
		fmt.Println("Previous Node must not be nil")
		return
	}
	newNode := &Node{data: data}
	newNode.next = prevNode.next
	prevNode.next = newNode
}

func (ll *LinkedList) Delete(data int) {
	if ll.IsEmpty() {
		fmt.Println("List is empty, No deletion is performed")
		return
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}

	currNode := ll.head
	var prevNode *Node
	for currNode != nil && currNode.data != data {
		prevNode = currNode
		currNode = currNode.next
	}

	if currNode == nil {
		fmt.Println("Item not found in the list. No deletion performed")
	} else {
		prevNode.next = currNode.next
	}
}

func (ll *LinkedList) Search(data int) bool {
	currNode := ll.head
	for currNode != nil {
		if currNode.data == data {
			return true
		}
		currNode = currNode.next
	}
	return false
}

func (ll *LinkedList) GetSize() int {
	count := 0
	currNode := ll.head
	for currNode != nil {
		count++
		currNode = currNode.next
	}
	return count
}

func (ll *LinkedList) Display() {
	currNode := ll.head
	for currNode != nil {
		fmt.Printf("%d ", currNode.data)
		currNode = currNode.next
	}
	fmt.Println()
}

func main() {
	linkedList := LinkedList{}
	linkedList.Append(10)
	linkedList.Append(20)
	linkedList.Append(30)
	linkedList.Append(40)
	linkedList.Prepend(50)

	linkedList.Display()

	fmt.Println("Side: ", linkedList.GetSize())
	fmt.Println("Search 30 : ", linkedList.Search(30))

	linkedList.Delete(20)
	linkedList.Display()

}
