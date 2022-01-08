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

// changes By Prerna 08/01/2022

// 1. Is validate BST
func (t *BinarySearchTree) ValidBST() {
	if t.isBST() {
		fmt.Println("Valid BST")

	} else {
		fmt.Println("Invalid BST")
	}
}

func (t *BinarySearchTree) isBST() bool {
	if t == nil {
		return true
	}
	return t.root.chkNode(100, 100)
}
func (n *node) chkNode(max int, min int) bool {
	if n == nil {
		return true
	}
	if n.val > max && n.val < min {
		return false
	}
	if n.left.chkNode(n.val, min) && n.right.chkNode(max, n.val) {
		return true
	}
	return false

}

// 2. determine if the given BST is balance

func (t *BinarySearchTree) IsBal() {
	if t == nil {
		fmt.Println("Balanced")
		return
	}
	h := 0
	h1 := t.root.left.GetHght(h)
	h2 := t.root.right.GetHght(h)
	if h1-h2 > 1 || h2-h1 > 1 {
		fmt.Println("Un-balanced")
		return
	}
	fmt.Println("Balanced")
}

func (n *node) GetHght(h int) int {
	if n == nil {
		return h
	}
	n.left.GetHght(h + 1)
	n.right.GetHght(h + 1)
	if n.left.GetHght(h+1) > n.right.GetHght(h+1) {
		return n.left.GetHght(h + 1)
	} else {
		return n.right.GetHght(h + 1)
	}

}

// 3. given a sorted array, construct a bst with minimal height
// note: following mthods has error

func ConstructBST(a []int) *BinarySearchTree {

	t := NewBST()
	t.root.GetNode(a)
	t.Print()
	return t
}
func (n *node) GetNode(arr []int) {
	var mid = 0

	if len(arr) > 0 {
		mid = len(arr) / 2
		n = &node{arr[mid-1], nil, nil}

	}
	// else {
	// 	//return n
	// }
	arrL := []int{}
	arrR := []int{}
	// var arrL []int
	// var arrL []int
	if len(arr)-mid > 1 {
		for i := 0; i < mid-1; i++ {
			arrL = append(arrL, arr[i])
		}
		//fmt.Println(arrL)
		for i := mid; i < len(arr); i++ {

			//arrR[i] = arr[i]
			arrR = append(arrR, arr[i])

		}
		// n.left = n.left.GetNode(arrL)
		// n.right = n.right.GetNode(arrR)
		n.left.GetNode(arrL)
		n.right.GetNode(arrR)
	} else {
		if arr[mid] > n.val {
			n.right = &node{arr[mid], nil, nil}
		} else {
			n.left = &node{arr[mid], nil, nil}
		}
	}

	//return n
}
