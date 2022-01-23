package tree

import (
	"data-structures/queue"
	"fmt"
	"math"
)

type (
	// BinaryTree representation.
	Tree struct {
		root *node
	}
	node struct {
		value int
		left  *node
		right *node
	}
)

func New() *Tree {
	return &Tree{nil}
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &node{val, nil, nil}
		return
	}

	t.root.insert(val)
}

func (n *node) insert(val int) {
	if n == nil {
		return
	}

	q := queue.New()
	q.Enqueue(n)

	for q.Len() != 0 {
		n := q.Dequeue().(*node)
		if n.left != nil {
			q.Enqueue(n.left)
		} else {
			n.left = &node{val, nil, nil}
			return
		}

		if n.right != nil {
			q.Enqueue(n.right)
		} else {
			n.right = &node{val, nil, nil}
			return
		}
	}
}

func (t *Tree) TraverseInOrder() {
	if t.root == nil {
		return
	}

	t.root.traverseInOrder()
	fmt.Println()
}

func (n *node) traverseInOrder() {
	switch {
	case n == nil:
		return
	case n.left != nil:
		n.left.traverseInOrder()
		fallthrough
	default:
		fmt.Printf("%d\t", n.value)
	}

	n.right.traverseInOrder()
}

func (t *Tree) TraversePreOrder() {
	if t.root == nil {
		return
	}

	t.root.traversePreOrder()
	fmt.Println()
}

func (n *node) traversePreOrder() {
	if n == nil {
		return
	}

	fmt.Printf("%d\t", n.value)
	if n.left != nil {
		n.left.traversePreOrder()
	}

	if n.right != nil {
		n.right.traversePreOrder()
	}
}

func (t *Tree) TraversePostOrder() {
	if t.root == nil {
		return
	}

	t.root.traversePostOrder()
	fmt.Println()
}

func (n *node) traversePostOrder() {
	if n == nil {
		return
	}

	if n.left != nil {
		n.left.traversePostOrder()
	}

	if n.right != nil {
		n.right.traversePostOrder()
	}

	fmt.Printf("%d\t", n.value)
}

func (t *Tree) TraverseLevelOrder() {
	if t.root == nil {
		return
	}

	t.root.traverseLevelOrder()
	fmt.Println()
}

func (n *node) traverseLevelOrder() {
	if n == nil {
		return
	}

	q := queue.New()
	q.Enqueue(n)
	for q.Len() != 0 {
		n := q.Dequeue().(*node)
		fmt.Printf("%d\t", n.value)
		if n.left != nil {
			q.Enqueue(n.left)
		}

		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
}

func (t *Tree) TraverseLevelOrderSeparateLine() {
	if t.root == nil {
		return
	}

	t.root.traverseLevelOrderSeparateLine()
}

func (n *node) traverseLevelOrderSeparateLine() {
	if n == nil {
		return
	}

	q := queue.New()
	q.Enqueue(n)
	for q.Len() != 0 {
		level := q.Len()
		for level > 0 {
			n := q.Dequeue().(*node)
			fmt.Printf("%d\t", n.value)
			if n.left != nil {
				q.Enqueue(n.left)
			}

			if n.right != nil {
				q.Enqueue(n.right)
			}
			level--
		}
		fmt.Println("")
	}
}

// Height calculates the height of the given binary tree.
func (t *Tree) Height() int {
	if t.root == nil {
		return 0
	}

	return t.root.height()
}

func (n *node) height() int {
	if n == nil {
		return 0
	}

	lh := n.left.height()
	rh := n.right.height()
	if lh > rh {
		return lh + 1
	}

	return rh + 1
}

// Print prints 2-D view of a binary tree.
func (t *Tree) Print() {
	if t.root == nil {
		return
	}

	t.root.print(0)
}

// print is the helper function for Print.
func (n *node) print(space int) {
	if n == nil {
		return
	}

	// Increase distance between levels
	space += 5

	// Process right child.
	n.right.print(space)

	// Print the element
	fmt.Println("")
	for i := 5; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%d\n", n.value)

	// Process left child.
	n.left.print(space)
}

// PrintMirror print mirror 2-D view of a binary tree.
func (t *Tree) PrintMirror() {
	if t.root == nil {
		return
	}

	t.root.printMirror(0)
}

// printMirror is the helper function for PrintMirror.
func (n *node) printMirror(space int) {
	if n == nil {
		return
	}

	// Increase distance between levels
	space += 5

	// Process left child.
	n.left.printMirror(space)

	// Print the element
	fmt.Println("")
	for i := 5; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%d\n", n.value)

	// Process right child.
	n.right.printMirror(space)
}

// IsBST checks whether a given binary tree can be qualified as BST.
func (t *Tree) IsBST() bool {
	if t.root == nil {
		return true
	}

	return t.root.isBST(math.MinInt, math.MaxInt)
}

func (n *node) isBST(min, max int) bool {
	if n == nil {
		return true
	}

	if n.value < min || n.value > max {
		return false
	}

	return n.left.isBST(min, n.value) && n.right.isBST(n.value, max)
}

// IsBalanced checks whether a given binary tree is balanced binary tree.
func (t *Tree) IsBalanced() bool {
	if t.root == nil {
		return true
	}

	if t.root.isBalanced() > 0 {
		return true
	}

	return false
}

func (n *node) isBalanced() int {
	if n == nil {
		return 0
	}

	lh := n.left.isBalanced()
	if lh == -1 {
		return -1
	}

	rh := n.right.isBalanced()
	if rh == -1 {
		return -1
	}

	var diff int
	if lh > rh {
		diff = lh - rh
		if diff > 1 {
			return -1
		}

		return 1 + lh
	}

	diff = rh - lh
	if diff > 1 {
		return -1
	}

	return 1 + rh
}

// GetNodeWithValue fetches the node in the Binary tree that matches the value given.
// Return nil if no node value matches the given value.
func (t *Tree) GetNodeWithValue(val int) *node {
	if t.root == nil {
		return nil
	}

	return t.root.getNodeWithValue(val)
}

func (n *node) getNodeWithValue(val int) *node {
	if n == nil {
		return nil
	}

	// Check fast and exit on success. The later part of the code becomes less readable due to this.
	if n.value == val {
		return n
	}

	q := queue.New()
	if n.left != nil {
		q.Enqueue(n.left)
	}

	if n.right != nil {
		q.Enqueue(n.right)
	}

	// Alternatively create a queue and enqueue the node itself and make the code more readable.
	/*
		q := queue.New()
		q.Enqueue(n)
	*/
	// Remove line 364-375 and the rest of the code follows. More readable code but always
	// creates a queue even if the root node matches.

	for q.Len() != 0 {
		n := q.Dequeue().(*node)
		if n.value == val {
			return n
		}

		if n.left != nil {
			q.Enqueue(n.left)
		}

		if n.right != nil {
			q.Enqueue(n.right)
		}
	}

	return nil
}

// changes By Prerna 09/01/2022
// 1. Is validate BST
func (t *Tree) ValidBST() bool {
	if t == nil {
		return true
	}
	return t.root.chkNode(math.MaxInt, math.MinInt)
}

func (n *node) chkNode(max int, min int) bool {
	if n == nil {
		return true
	}
	if n.value > max && n.value < min {
		return false
	}
	return n.left.chkNode(n.value, min) && n.right.chkNode(max, n.value)
}

// 2. determine if the given BST is balance

func (t *Tree) IsBal() bool {
	if t == nil {
		// fmt.Println("Balanced")
		return true
	}
	// h := 0
	h1 := t.root.left.getHght(0)
	h2 := t.root.right.getHght(0)
	if h1-h2 > 1 || h2-h1 > 1 {
		// fmt.Println("Un-balanced")
		return false
	}
	return true
	// fmt.Println("Balanced")
}

func (n *node) getHght(h int) int {
	if n == nil {
		return h
	}
	lh := n.left.getHght(h + 1)
	rh := n.right.getHght(h + 1)
	if lh > rh {
		return lh //n.left.GetHght(h + 1)
	}
	return rh //n.right.GetHght(h + 1)
}

//  commited by prerna 18/01/2022
// 1. Finding the cousin nodes
func (t *Tree) IsCousinNode(n1 int, n2 int) bool {
	if t == nil {
		return false
	}
	return t.root.isCousinNode(t.GetNodeWithValue(n1), t.GetNodeWithValue(n2))
}
func (n *node) isCousinNode(n1 *node, n2 *node) bool {
	if n == nil {
		return true
	}

	if n1.height() != n2.height() {
		return false
	}

	if (n.left == n1 && n.right == n2) || (n.left == n2 && n.right == n1) {
		return false
	}
	return n.left.isCousinNode(n1, n2) && n.right.isCousinNode(n1, n2)
}

//2. Printing the top view of a tree

func (t *Tree) PrintTopView() {
	if t == nil {
		return
	}
	t.root.printTopView()
}

func (n *node) printTopView() {
	if n == nil {
		return
	}
	type lvlNode struct {
		n   *node
		lvl int
	}
	nd := lvlNode{n, 0}
	// // creating an empty queue
	q := queue.New()

	// inserting the first node into the q
	mp := make(map[int]bool)
	mp[0] = true
	q.Enqueue(nd)
	for q.Len() != 0 {
		m := q.Dequeue().(lvlNode)

		fmt.Printf("%v\t", m.n.value)
		if m.n.left != nil {
			ln := lvlNode{m.n.left, m.lvl + 1}
			if _, ok := mp[ln.lvl]; ok {
			} else {
				q.Enqueue(ln)
			}
		}

		if m.n.right != nil {
			rn := lvlNode{m.n.right, m.lvl - 1}
			if _, ok := mp[rn.lvl]; ok {
			} else {
				q.Enqueue(rn)
			}
		}
	}
}

//18/01/2022
// 3. printing in level order diff lines as per the level

func (t *Tree) TreeLevelPrint() {
	if t == nil {
		return
	}

	t.root.treeLevelPrint()
}
func (n *node) treeLevelPrint() {
	if n == nil {
		return
	}
	q := queue.New()
	q.Enqueue(n)

	for q.Len() != 0 {
		lvl := q.Len()
		for lvl > 0 {
			n = q.Dequeue().(*node)
			fmt.Printf("%v\t", n.value)

			if n.left != nil {
				q.Enqueue(n.left)
			}
			if n.right != nil {
				q.Enqueue(n.right)
			}
			lvl--
		}
		fmt.Println("")
	}
}
