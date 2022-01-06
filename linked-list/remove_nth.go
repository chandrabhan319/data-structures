package list

import "fmt"

func (l *List) RemoveNthFromEnd(n int) error {
	if l == nil || l.head == nil {
		return fmt.Errorf("delete operation on empty list")
	}

	c := l.head
	i := 0
	for c != nil {
		c = c.next
		i++
	}

	if n > i {
		return fmt.Errorf("index out of range")
	}

	if err := l.DeleteAtIndex(i - n); err != nil {
		return err
	}

	return nil
}

func (l *List) RemoveNthFromEndInSinglePass(n int) error {
	if n == 0 {
		return nil
	}

	if l == nil || l.head == nil {
		return fmt.Errorf("delete operation on empty list")
	}
	slow := l.head
	fast := l.head

	for i := 0; i < n; i++ {
		if fast.next == nil {
			if (n - i) == 1 {
				l.head = l.head.next
				return nil
			}

			return fmt.Errorf("index out of range")
		}

		fast = fast.next
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
	return nil
}
