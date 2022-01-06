package list

func (l *List) Partition(x int) {
	if l == nil || l.head == nil {
		return
	}

	l1 := &node{}
	l2 := &node{}
	c1 := l1
	c2 := l2

	for l.head != nil {
		if l.head.value.(int) < x {
			c1.next = l.head
			c1 = l.head
		} else {
			c2.next = l.head
			c2 = l.head
		}

		l.head = l.head.next
	}

	c2.next = nil
	c1.next = l2.next
	l.head = l1.next
}
