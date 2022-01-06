package list

import "fmt"

func (l *List) RotateRight(n int) error {
	switch {
	case n < 0:
		return fmt.Errorf("invalid argument")
	case l == nil || l.head == nil:
		return fmt.Errorf("empty list")
	case n == 0:
		return nil
	}

	slow := l.head
	fast := l.head
	i := 1
	for l.head.next != nil {
		l.head = l.head.next
		i++
	}

	if n > i {
		n = n % i
	}

	if i == n || n == 0 {
		l.head = slow
		return nil
	}

	for j := 1; j < (i - n); j++ {
		fast = fast.next
	}

	l.head.next = slow
	l.head = fast.next
	fast.next = nil
	return nil
}

func (l *List) RotateRightNew(n int) error {
	switch {
	case n < 0:
		return fmt.Errorf("invalid argument")
	case l == nil || l.head == nil:
		return fmt.Errorf("empty list")
	case n == 0:
		return nil
	}

	slow := l.head
	fast := l.head
	tmpHead := l.head
	i := 1
	for l.head.next != nil {
		l.head = l.head.next
		i++
	}

	if n > i {
		n = n % i
	}

	if i == n || n == 0 {
		l.head = tmpHead
		return nil
	}

	for i := 0; i < n; i++ {
		if fast.next == nil {
			l.head = tmpHead
			return fmt.Errorf("index out of range")
		}

		fast = fast.next
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	l.head = slow.next
	slow.next = nil
	c := l.head
	for c.next != nil {
		c = c.next
	}

	c.next = tmpHead
	return nil
}

func (l *List) RotateLeft(n int) error {
	switch {
	case n < 0:
		return fmt.Errorf("invalid argument")
	case l == nil || l.head == nil:
		return fmt.Errorf("empty list")
	case n == 0:
		return nil
	}

	slow := l.head
	fast := l.head
	tmpHead := l.head
	i := 1
	for fast.next != nil {
		fast = fast.next
		i++
	}

	if n > i {
		n = n % i
	}

	if i == n || n == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		l.head = l.head.next
	}

	for slow.next != l.head {
		slow = slow.next
	}

	slow.next = fast.next
	for tmpHead.next != nil {
		fast.next = tmpHead
		fast = fast.next
		tmpHead = tmpHead.next
	}

	fast.next = tmpHead
	return nil
}

func (l *List) RotateSpecialR(n int) error {
	if l.head == nil {
		return fmt.Errorf("empty list")
	}

	a := l.head
	b := l.head

	for i := 0; i < n; i++ {
		switch {
		case b.next != nil:
			b = b.next
		case i == (n - 1):
			return nil
		default:
			return fmt.Errorf("n is greater than number of nodes")
		}
	}

	for b.next != nil {
		b = b.next
		a = a.next
	}

	c := a.next
	for c != b {
		d := c.next
		c.next = l.head
		l.head = c
		c = d
	}

	c.next = l.head
	l.head = c
	a.next = nil
	return nil
}

// left rotate
func (l *List) LeftRotate(n int) {
	switch {
	case l.head == nil:
		return
	case l.head.next == nil:
		return
	}
	a := l.head
	for a.next != nil {
		a = a.next
	}
	for n > 0 {
		c := l.head
		l.head = l.head.next
		a.next = c
		c.next = nil
		a = a.next
		n--
	}
}
