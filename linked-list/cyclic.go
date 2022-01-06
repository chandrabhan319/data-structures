package list

import (
	"fmt"
)

func (l *List) FindCycle() bool {
	if l.head == nil && l.head.next == nil {
		return false
	}

	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}

	return false
}

func (l *List) FindCyclicNode() *node {
	if l.head == nil && l.head.next == nil {
		return nil
	}

	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	slow = l.head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow
}

func (l *List) MakeCyclic(n int) error {
	if n < 0 {
		return fmt.Errorf("invalid index")
	}

	if l.head == nil {
		return fmt.Errorf("empty list")
	}

	ci := l.head
	for i := 1; i <= n; i++ {
		if ci.next == nil {
			return fmt.Errorf("index out of range")
		}

		ci = ci.next
	}

	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}

	tail.next = ci
	return nil
}
