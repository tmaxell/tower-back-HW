package main

type node struct {
	value int
	left  *node
	right *node
}

type BinSearchTree struct {
	root *node
}

func (bst *BinSearchTree) add_recursive(Node *node, value int) *node {
	if Node == nil {
		return &node{value: value}
	}
	if value < Node.value {
		Node.left = bst.add_recursive(Node.left, value)
	} else if value > Node.value {
		Node.right = bst.add_recursive(Node.right, value)
	}

	return Node
}
