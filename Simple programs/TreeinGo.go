package main

import "fmt"

type Node struct {
	data     int
	children []*Node
}

type Tree struct {
	root *Node
}

func (t *Tree) IsEmpty() bool {
	return t.root == nil
}

func (t *Tree) Insert(data int, parent *Node) {
	newNode := &Node{data: data}

	if parent == nil {
		if t.IsEmpty() {
			t.root = newNode
		} else {
			fmt.Println("Tree already has a root.")
		}
	} else {
		parent.children = append(parent.children, newNode)
	}
}

func (t *Tree) Display() {
	t.displayNode(t.root, 0)
}

func (t *Tree) displayNode(node *Node, level int) {
	if node == nil {
		return
	}
	fmt.Printf("%*s%d\n", level*4, "", node.data)
	for _, child := range node.children {
		t.displayNode(child, level+1)
	}
}

func main() {
	tree := Tree{}
	tree.Insert(1, nil)
	tree.Insert(10, nil)
	tree.Insert(2, tree.root)
	tree.Insert(3, tree.root)
	tree.Insert(4, tree.root.children[0])
	tree.Insert(20, tree.root.children[0].children[0])
	tree.Insert(6, tree.root)
	tree.Insert(7, tree.root.children[1])
	tree.Insert(7, tree.root.children[1])
	tree.Insert(7, tree.root.children[1])
	tree.Insert(9, tree.root.children[2])

	tree.Display()
}
