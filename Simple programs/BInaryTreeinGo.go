package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) IsEmpty() bool {
	return bt.root == nil
}

func (bt *BinaryTree) Insert(data int) {
	newNode := &Node{data: data, left: nil, right: nil}
	if bt.IsEmpty() {
		bt.root = newNode
	} else {
		bt.insertNode(bt.root, newNode)
	}
}

func (bt *BinaryTree) insertNode(root *Node, newNode *Node) {
	if newNode.data < root.data {
		if root.left == nil {
			root.left = newNode
		} else {
			bt.insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			bt.insertNode(root.right, newNode)
		}
	}
}

func (bt *BinaryTree) Search(data int) bool {
	return bt.searchNode(bt.root, data)
}

func (bt *BinaryTree) searchNode(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if data == root.data {
		return true
	} else if data < root.data {
		return bt.searchNode(root.left, data)
	} else {
		return bt.searchNode(root.right, data)
	}
}

func (bt *BinaryTree) Delete(data int) {
	bt.root = bt.deleteNode(bt.root, data)
}

func (bt *BinaryTree) deleteNode(root *Node, data int) *Node {
	if root == nil {
		return root
	}

	if data < root.data {
		root.left = bt.deleteNode(root.left, data)
	} else if data > root.data {
		root.right = bt.deleteNode(root.right, data)
	} else {
		//  Zero or 1 child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		//2 children

		minNode := bt.findMinNode(root.right)
		root.data = minNode.data
		root.right = bt.deleteNode(root.right, minNode.data)

	}
	return root
}

func (bt *BinaryTree) findMinNode(root *Node) *Node {
	if root.left != nil {
		root.left = bt.findMinNode(root.left)
	}
	return root
}

func (bt *BinaryTree) PreOrderTraversal() {
	bt.preOrder(bt.root)
	fmt.Println()
}

func (bt *BinaryTree) preOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		bt.preOrder(root.left)
		bt.preOrder(root.right)
	}
}

func (bt *BinaryTree) InOrderTraversal() {
	bt.inOrder(bt.root)
	fmt.Println()
}

func (bt *BinaryTree) inOrder(root *Node) {
	if root != nil {
		bt.inOrder(root.left)
		fmt.Printf("%d ", root.data)
		bt.inOrder(root.right)
	}
}

func (bt *BinaryTree) PostOrderTraversal() {
	bt.postOrder(bt.root)
	fmt.Println()
}

func (bt *BinaryTree) postOrder(root *Node) {
	if root != nil {
		bt.postOrder(root.left)
		bt.postOrder(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func main() {
	binaryTree := &BinaryTree{}
	binaryTree.Insert(50)
	binaryTree.Insert(30)
	binaryTree.Insert(70)
	binaryTree.Insert(20)
	binaryTree.Insert(40)
	binaryTree.Insert(60)
	binaryTree.Insert(80)

	fmt.Println("Pre-Order Traversal")
	binaryTree.PreOrderTraversal()

	fmt.Println("In-Order Traversal")
	binaryTree.InOrderTraversal()

	fmt.Println("Post-Order Traversal")
	binaryTree.PostOrderTraversal()

	fmt.Println("Search 40:", binaryTree.Search(40))
	fmt.Println("Search 90:", binaryTree.Search(90))

	binaryTree.Delete(40)

	fmt.Println("In-order Traversal after deleting 40:")
	binaryTree.InOrderTraversal()
}
