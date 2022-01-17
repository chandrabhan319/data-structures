package main

import (
	"fmt"

	"data-structures/array"
	"data-structures/array/sorting"
	"data-structures/queue"
	"data-structures/stack"
	"data-structures/tree"

	list "data-structures/linked-list"
)

func main() {
	// stack operations.
	printSection()
	fmt.Println("Stack operations")
	s := stack.New()
	fmt.Printf("Length of stack before push operations: %d\n", s.Len())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	fmt.Printf("Length of stack after push operations: %d\n", s.Len())
	val, err := s.Peek()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Peek stack got: %d\n", val)
	// Parse the full stack without popping.
	s.Parse()

	// Pop all the values.
	for s.Len() > 0 {
		val, err := s.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Value: %d\t", val)
	}

	fmt.Println("")
	fmt.Printf("Length of stack after pop operations: %d\n", s.Len())
	printSection()

	// queue operations.
	q := queue.New()
	fmt.Printf("Length of queue before enqueue operations: %d\n", q.Len())
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	fmt.Printf("Peek: %d\n", q.Peek())
	fmt.Printf("Length of queue after enqueue operations: %d\n", q.Len())
	for q.Len() > 0 {
		fmt.Printf("Value: %d\n", q.Dequeue())
	}
	fmt.Printf("Length of queue after dequeue operations: %d\n", q.Len())
	printSection()

	fmt.Println("List operations")
	// list operations.
	l := list.New()

	if err := l.Insert(4); err != nil {
		panic(err)
	}

	if err := l.Insert(5); err != nil {
		panic(err)
	}

	if err := l.Insert(6); err != nil {
		panic(err)
	}

	if err := l.Insert(7); err != nil {
		panic(err)
	}

	l.Parse()
	if err := l.RotateLeft(3); err != nil {
		panic(err)
	}

	l.Parse()
	if err := l.RotateRight(7); err != nil {
		panic(err)
	}

	l.Parse()
	if err := l.Insert(7); err != nil {
		panic(err)
	}

	l.Parse()
	l.RemoveConsecutiveDuplicates()
	l.Parse()
	if err := l.DeleteAtIndex(3); err != nil {
		panic(err)
	}

	l.Parse()
	if err := l.RemoveNthFromEndInSinglePass(3); err != nil {
		panic(err)
	}

	l.Parse()
	if err := l.Insert(7); err != nil {
		panic(err)
	}

	l.Parse()

	if err := l.Insert(8); err != nil {
		panic(err)
	}

	l.Parse()

	if err := l.Insert(8); err != nil {
		panic(err)
	}

	l.Parse()
	l.Reverse()
	l.Parse()
	if err := l.Insert(4); err != nil {
		panic(err)
	}

	l.Parse()
	l.Reverse()
	l.Parse()
	if err := l.RotateSpecialR(4); err != nil {
		panic(err)
	}
	l.Parse()

	if err := l.Insert(4); err != nil {
		panic(err)
	}
	l.Parse()

	if err := l.Insert(7); err != nil {
		panic(err)
	}

	l.Parse()
	l.RemDups()
	l.Parse()
	if err := l.RotateLeft(3); err != nil {
		panic(err)
	}

	l.Parse()
	l.Partition(7)
	l.Parse()

	printSection()

	treeOperations()
	printSection()
	sliceOperations()
	printSection()
	// fmt.Printf("%d\n", leftShift(10, 2))
	// fmt.Printf("%d\n", rightShift(10, 2))
	// fmt.Printf("%d\n", pow2(10))
	// fmt.Printf("%d\n", leftS(1, 6))
	// a := leftS(1, 6)
	// a = a | 0x0F
	// b := leftS(3, 5)
	// a = a | b
	// fmt.Printf("%d %d\n", a, b)
	// // a ^= a
	// // fmt.Printf("%d\n", a)
	// // ^a = 1 ^ a
	// fmt.Printf("%d\n", ^a)

}

func treeOperations() {
	// BST Operations.
	fmt.Println("Binary Search Tree operations")
	bst := tree.NewBST()
	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(8)
	bst.Insert(11)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(12)
	bst.Print()
	if bst.IsBalanced() {
		fmt.Println("Given BST is balanced.")
	} else {
		fmt.Println("Given BST is not balanced.")
	}
	fmt.Printf("Height of the binary search tree: %d\n", bst.Height())
	fmt.Print("Traverse Inorder:\t")
	bst.TraverseInOrder()
	fmt.Print("Traverse PreOrder:\t")
	bst.TraversePreOrder()
	fmt.Print("Traverse PostOrder:\t")
	bst.TraversePostOrder()
	fmt.Print("Traverse LevelOrder:\t")
	bst.TraverseLevelOrder()

	// Binary Tree Operations
	printSection()
	fmt.Println("Binary Tree operations")
	t := tree.New()
	t.Insert(9)
	t.Insert(5)
	t.Insert(12)
	t.Insert(4)
	t.Insert(6)
	t.Insert(7)
	t.Insert(13)
	fmt.Printf("Height of the binary tree: %d\n", t.Height())
	printSection()
	fmt.Println("2-D view of the binary tree.")
	t.Print()
	printSection()
	fmt.Println("2-D mirror view of the binary tree.")
	t.PrintMirror()
	printSection()
	fmt.Print("Traverse Inorder:\t")
	t.TraverseInOrder()
	fmt.Print("Traverse PreOrder:\t")
	t.TraversePreOrder()
	fmt.Print("Traverse PostOrder:\t")
	t.TraversePostOrder()
	fmt.Print("Traverse LevelOrder:\t")
	t.TraverseLevelOrder()
	printSection()
	fmt.Println("Traverse LevelOrder with each level on separate line:")
	t.TraverseLevelOrderSeparateLine()
	printSection()
	if t.IsBalanced() {
		fmt.Println("Given Binary tree is balanced.")
	} else {
		fmt.Println("Given Binary tree is not balanced.")
	}

	if t.IsBST() {
		fmt.Println("Given Binary tree is BST.")
	} else {
		fmt.Println("Given Binary tree is not BST.")
	}

	printSection()
	fmt.Println("Binary Search Tree creation with minimal height from a sorted array.")
	bst = tree.NewBST()
	bst.CreateWithMinimalHeight([]int{2, 3, 6, 7, 8, 9, 12, 15, 16, 18, 20})
	bst.Print()

	printSection()
	// top view print
	fmt.Println("Print Top View")
	t.PrintTopView()
	fmt.Println("")
	printSection()
	fmt.Println("Print Level order lines")
	t.TreeLevelPrint()
	fmt.Println("")
	printSection()
	fmt.Println("Cousin Node")
	if t.IsCousinNode(7, 0) {
		fmt.Println("Cousin Node true")
	} else {
		fmt.Println("Cousin Node False")

	}

}

func printSection() {
	fmt.Println("###################################################################################################################")
}
func sliceOperations() {
	a := array.GenerateIntSlice(10, 1000)
	fmt.Printf("%+v\n", a)
	a = sorting.MergeSort(a)
	fmt.Printf("%+v\n", a)
}
func OddEven(num int) bool {
	return num%2 == 0

}
func IsEven(num int) bool {
	return num&1 != 1
}
func leftShift(num int, shift int) int {
	return num << shift
}
func rightShift(num int, shift int) int {
	return num >> shift
}
func pow2(num int) int {
	if num == 0 {
		return 0
	}
	return 1 << num
}
func leftS(num int8, shift int) int8 {
	return num << shift
}
