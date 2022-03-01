package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"data-structures/array"
	"data-structures/array/sorting"
	"data-structures/bitwise"
	_ "data-structures/graph"
	"data-structures/queue"
	"data-structures/stack"
	"data-structures/tree"

	list "data-structures/linked-list"
)

type ABC struct {
	mu sync.Mutex
}

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

	n1 := bst.GetNodeWithValue(7)
	n2 := bst.GetNodeWithValue(12)

	if bst.AreCousins(n1, n2) {
		fmt.Printf("Node: %+v and  Node: %+v are cousins\n", n1, n2)
	} else {
		fmt.Printf("Node: %+v and  Node: %+v are not cousins\n", n1, n2)
	}

	bst.TopView()
	printSection()
	t1 := tree.New()
	t1.Insert(1)
	t1.Insert(2)
	t1.Insert(3)
	t1.Insert(4)
	t1.Insert(5)
	t1.Insert(6)
	t1.Insert(7)
	t1.Insert(8)
	t1.Insert(9)
	t1.Insert(10)
	t1.Insert(11)
	t1.Insert(12)
	t1.Insert(13)
	t1.Insert(14)
	t1.Insert(15)

	t1.Print()
	printSection()
	fmt.Printf("Top view of the tree: \t")
	t1.TopView()
	printSection()

	t1.DiagonalParse()
	printSection()

	t1.PrintRootToLeafPaths()
	printSection()

	t1.SumTree()
	t1.Print()
	printSection()

	// a := []int{1, 2, 3}
	// b := make([]int, len(a))
	// for i, v := range a {
	// 	b[i] = v
	// }

	// fmt.Printf("%+v", b)
	a := array.GenerateIntSlice(15, 1000)
	a = sorting.MergeSort(a)
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println("")
	printSection()

	x := 10
	fmt.Println(x)
	if x&1 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	//fmt.Println(bitwise.Divide(-1067865265873265732, 10))
	fmt.Println(bitwise.DivideInt8(10, 3))
	// ab := 7
	// fmt.Println(int8(56) << ab)
	// fmt.Printf("%v\n", fizzBuzzNew(15))
	// fmt.Println(fibonacci(5))
	fmt.Println(fibonacciTab(5))
	// fmt.Println(fibonacciMem(100))
	// fmt.Println(stringNew.ReverseN("Hello⌘"))
	// fmt.Println(stringNew.Reverse("Hello⌘"))
	// abc := []rune{'⌘', 120}
	// adc := 'c'
	// fmt.Printf("%T\t %T\n", abc, adc)
	fmt.Printf("%+v\n", generate(5))
	fmt.Println(array.RotateRight([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(array.RotateLeft([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(array.StockMaxProfit([]int{7, 1, 4, 3, 5, 3}))
	fmt.Println(array.StockMaxProfitII([]int{7, 1, 4, 3, 5, 3}))
	fmt.Println(array.MaxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	printSection()
	fmt.Println(array.StoneGame([]int{5, 3, 4, 5}))
	printSection()
	array.StoneGameII([]int{1, 2, 3})
	array.StoneGameIII([]int{2, 2, 5, 4, 3})
	if err := dir("tree"); err != nil {
		panic(err)
	}

	fmt.Println(filepath.Dir("dbh"))
	// arr := []int{1, 2}
	// arr = arr[1 : len(arr)-1]
	// fmt.Println(arr)
	// if strings.Contains(strings.ToLower("find: ‘/tmp/uiTestExecution/cypress/results’: No such file or directory"), "no such file or directory") {
	// 	fmt.Println("true")
	// }
	fmt.Println(array.SumCombinationOf134(5))
	fmt.Println(array.StockMaxProfitHard([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(array.StockMaxProfitEHard(8, []int{3, 3, 5, 0, 0, 3, 1, 4}))
	// fileInfo, err := os.Lstat("/Users/chsingh/go/src/github.com/fake-go-log-generator/temp/Dockerfile")
	// if err != nil {
	// 	panic(err)
	// }

	// switch fileInfo.Mode() & os.ModeType {
	// case os.ModeDir:
	// 	fmt.Println("Directory")
	// case os.ModeSymlink:
	// 	fmt.Println("Symlink")
	// default:
	// 	fmt.Println("File")
	// }
	if err := createTarGz("/Users/chsingh/go/src/github.com/chandrabhan319/data-structures/tree", "", []string{"tree/.del"}); err != nil {
		panic(err)
	}
}

func printSection() {
	fmt.Println("###################################################################################################################")
}

func fizzBuzz(n int) []string {
	a := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			a[i-1] = "FizzBuzz"
		case i%3 == 0:
			a[i-1] = "Fizz"
		case i%5 == 0:
			a[i-1] = "Buzz"
		default:
			a[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return a
}

func fizzBuzzNew(n int) []string {
	a := []string{}
	dict := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	for i := 1; i <= n; i++ {
		numStr := ""

		for k, v := range dict {
			if i%k == 0 {
				numStr += v
			}
		}

		if numStr == "" {
			numStr = fmt.Sprintf("%d", i)
		}

		a = append(a, numStr)
	}

	return a
}

func fibonacci(n int) int {
	fmt.Printf("Calling fib(%d)\n", n)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// Tabulation method.
func fibonacciTab(n int) int {
	a := []int{0, 1}
	for i := 2; i <= n; i++ {
		a = append(a, a[i-1]+a[i-2])
	}

	return a[n]
}

var a = make([]int, 1000)

// Memoization method.
func fibonacciMem(n int) int {
	fmt.Printf("Calling fibMem(%d)\n", n)
	if n == 0 {
		a[n] = 0
		return 0
	}

	if n == 1 {
		a[n] = 1
		return 1
	}

	if a[n] != 0 {
		return a[n]
	}

	f := a[n]
	if f == 0 {
		f = fibonacciMem(n-1) + fibonacciMem(n-2)
		a[n] = f
	}

	return f
}

func generate(numRows int) [][]int {
	var result [][]int
	for i := 1; i <= numRows; i++ {
		a := make([]int, i)
		a[0], a[len(a)-1] = 1, 1
		for j := 1; j < len(a)-1; j++ {
			a[j] = result[i-2][j] + result[i-2][j-1]
		}
		result = append(result, a)
	}
	return result
}

func dir(name string) error {
	entries, err := ioutil.ReadDir(name)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(name, entry.Name())
		fmt.Println(sourcePath)
	}

	return nil
}

func createTarGz(source, target string, ignoreFiles []string) error {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar.gz", filename))
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}

	defer tarfile.Close()
	tarball := tar.NewWriter(tarfile)
	gz := gzip.NewWriter(tarfile)
	defer gz.Close()
	tarball = tar.NewWriter(gz)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			for i := range ignoreFiles {
				var ignoreFile string
				if !filepath.IsAbs(ignoreFiles[i]) {
					ignoreFile, err = filepath.Abs(ignoreFiles[i])
					if err != nil {
						return err
					}
				} else {
					ignoreFile = ignoreFiles[i]
				}

				fmt.Printf("Path: %s, IgnorePath: %s\n", absPath, ignoreFile)
				switch {
				case info.IsDir() && absPath == ignoreFile:
					ignoreFiles = remove(ignoreFiles, i)
					return filepath.SkipDir
				case absPath == ignoreFile:
					ignoreFiles = remove(ignoreFiles, i)
					return nil
				}
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			// Skip copying data from special files like directories.
			if !info.Mode().IsRegular() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//get the filepath for the symbolic link
// func handleSymlink(path string) (string, int64, error) {
// 	//read the link
// 	link, err := os.Readlink(path)
// 	if err != nil {
// 		return "", -1, err
// 	}

// 	//if the filepath isn't an absolute path we need to
// 	//reconstruct it so that it is
// 	if !filepath.IsAbs(link) {
// 		link = filepath.Join(filepath.Dir(path), link)
// 		if err != nil {
// 			return "", -1, err
// 		}
// 	}
// 	fi, err := os.Stat(link)
// 	if err != nil {
// 		return "", -1, err
// 	}
// 	size := fi.Size()
// 	return link, size, nil
// }
