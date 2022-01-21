package main

import (
	"fmt"
)

type Node struct {
	key   int
	right *Node
	left  *Node
}

func main() {
	tree := &Node{key: 100}
	fmt.Println(tree)
	tree.insert(20)
	tree.insert(240)
	tree.insert(280)
	tree.insert(790)
	tree.insert(890)
	tree.insert(70)
	tree.insert(205)
	if tree.search(295) {
		fmt.Println("Value is Found ")

	} else {
		fmt.Println("Value not found")
	}

}

//Insert method to insert value in node

func (n *Node) insert(k int) {
	if n.key < k {
		//Move right
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		//Move left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}
func (n *Node) search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		//Move right
		return n.right.search(k)

	} else if n.key > k {
		//Move left
		return n.left.search(k)
	}
	return true
}
