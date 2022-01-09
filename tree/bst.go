package tree

import (
	"fmt"
)

type (
	// BinarySearchTree representation.
	BinarySearchTree struct {
		root *node
	}
)

// NewBST return an empty binary search tree.
func NewBST() *BinarySearchTree {
	return &BinarySearchTree{nil}
}

// Insert operation on BST.
func (t *BinarySearchTree) Insert(val int) {
	if t.root == nil {
		t.root = &node{
			value: val,
			left:  nil,
			right: nil,
		}

		return
	}

	t.root.insertBST(val)
}

func (n *node) insertBST(val int) {
	switch {
	case n == nil:
		return
	case val <= n.value:
		if n.left == nil {
			n.left = &node{value: val, left: nil, right: nil}
			return
		}
		n.left.insertBST(val)
	default:
		if n.right == nil {
			n.right = &node{value: val, left: nil, right: nil}
			return
		}
		n.right.insertBST(val)
	}
}

// TraverseInOrder traverses the BST in In-order.
func (t *BinarySearchTree) TraverseInOrder() {
	if t.root == nil {
		return
	}

	t.root.traverseInOrder()
	fmt.Println()
}

// TraversePreOrder traverses the BST in Pre-order.
func (t *BinarySearchTree) TraversePreOrder() {
	if t.root == nil {
		return
	}

	t.root.traversePreOrder()
	fmt.Println()
}

// TraversePostOrder traverses the BST in Post-order.
func (t *BinarySearchTree) TraversePostOrder() {
	if t.root == nil {
		return
	}

	t.root.traversePostOrder()
	fmt.Println()
}

// TraverseLevelOrder traverses the BST in Level-order.
func (t *BinarySearchTree) TraverseLevelOrder() {
	if t.root == nil {
		return
	}

	root := t.root
	root.traverseLevelOrder()
	fmt.Println()
}

// Height calculates the height of BST.
func (t *BinarySearchTree) Height() int {
	if t.root == nil {
		return 0
	}

	return t.root.height()
}

// Print print the 2-D lateral view of BST.
func (t *BinarySearchTree) Print() {
	if t.root == nil {
		return
	}

	t.root.print(0)
}

// CreateWithMinimalHeight creates a BST from sorted array of minimal height.
func (t *BinarySearchTree) CreateWithMinimalHeight(sortedArr []int) {
	len := len(sortedArr)
	if len == 0 {
		return
	}

	mid := (len - 1) / 2
	t.Insert(sortedArr[mid])
	t.CreateWithMinimalHeight(sortedArr[0:mid])
	t.CreateWithMinimalHeight(sortedArr[mid+1 : len])
}

// IsBalanced checks if the BST is Balanced tree.
func (t *BinarySearchTree) IsBalanced() bool {
	if t.root == nil {
		return true
	}

	if t.root.isBalanced() > 0 {
		return true
	}

	return false
}

// GetNodeWithValue fetches the node in the BST that matches the value given.
// Return nil if no node value matches the given value.
func (t *BinarySearchTree) GetNodeWithValue(val int) *node {
	if t.root == nil {
		return nil
	}

	return t.root.getBSTNodeWithValue(val)
}

func (n *node) getBSTNodeWithValue(val int) *node {
	if n == nil {
		return nil
	}

	if n.value == val {
		return n
	}

	if val < n.value {
		return n.left.getBSTNodeWithValue(val)
	}

	return n.right.getBSTNodeWithValue(val)
}
