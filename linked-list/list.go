package list

import "fmt"

type (
	// Singly ...
	List struct {
		head *node
	}
	node struct {
		value interface{}
		next  *node
	}
)

func New() *List {
	return &List{nil}
}

func (l *List) Insert(val interface{}) error {
	n := &node{value: val, next: nil}
	if l.head == nil {
		l.head = n
		return nil
	}

	c := l.head
	for c.next != nil {
		c = c.next
	}

	c.next = n
	return nil
}

func (l *List) DeleteHead() {
	f := l.head
	l.head = l.head.next
	f.next = nil
}

func (l *List) Parse() {
	if l == nil || l.head == nil {
		return
	}

	c := l.head
	for c.next != nil {
		fmt.Printf("%v\t", c.value)
		c = c.next
	}

	fmt.Printf("%v\t\n", c.value)
}

func (l *List) DeleteTail() {
	c := l.head
	for c.next != nil && c.next.next != nil {
		c = c.next
	}

	c.next = nil
}

func (l *List) DeleteAtIndex(n int) error {
	if n < 0 {
		return fmt.Errorf("invalid index")
	}

	if l.head == nil {
		return fmt.Errorf("delete operation on empty list")
	}

	if n == 0 {
		l.DeleteHead()
		return nil
	}

	c := l.head
	for i := 1; i < n; i++ {
		if c.next == nil {
			return fmt.Errorf("index out of range")
		}

		c = c.next
	}

	if c.next == nil {
		return fmt.Errorf("index out of range")
	}

	c.next = c.next.next

	return nil
}

func (l *List) Get(n int) (interface{}, error) {
	if n < 0 {
		return nil, fmt.Errorf("invalid index")
	}

	if l.head == nil {
		return nil, fmt.Errorf("get operation on empty list")
	}

	c := l.head
	for i := 1; i <= n; i++ {
		if c.next == nil {
			return nil, fmt.Errorf("index out of range")
		}

		c = c.next
	}

	return c.value, nil
}
