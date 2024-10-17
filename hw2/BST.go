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

func (bst *BinSearchTree) is_exist(value int) bool {
	return bst.is_exist_recursive(bst.root, value)
}

func (bst *BinSearchTree) is_exist_recursive(Node *node, value int) bool {
	if Node == nil {
		return false
	}
	if value == Node.value {
		return true
	}
	if value < Node.value {
		return bst.is_exist_recursive(Node.left, value)
	} else {
		return bst.is_exist_recursive(Node.right, value)
	}
}

func (bst *BinSearchTree) delete(value int) {
	bst.root = bst.delete_recursive(bst.root, value)
}

func (bst *BinSearchTree) delete_recursive(Node *node, value int) *node {
	if Node == nil {
		return nil
	}
	if value < Node.value {
		Node.left = bst.delete_recursive(Node.left, value)
	} else if value > Node.value {
		Node.right = bst.delete_recursive(Node.right, value)
	} else {
		if Node.left == nil && Node.right == nil {
			return nil
		}
		if Node.left == nil {
			return Node.right
		}
		if Node.right == nil {
			return Node.left
		}
		smallest_value := bst.find_min_value(Node.right)
		Node.value = smallest_value
		Node.right = bst.delete_recursive(Node.right, smallest_value)
	}
	return Node
}

func (bst *BinSearchTree) find_min_value(Node *node) int {
	current := Node
	for current.left != nil {
		current = current.left
	}
	return current.value
}
