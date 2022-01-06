package list

func (l *List) Middle() *node {
	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	if fast.next == nil {
		return slow
	}

	if fast.next.next == nil {
		return slow.next
	}

	return nil
}
