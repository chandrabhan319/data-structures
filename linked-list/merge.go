package list

func MergeSorted(l1, l2 *List) *List {
	if l1 == nil || l1.head == nil {
		return l2
	}

	if l2 == nil || l2.head == nil {
		return l1
	}

	l := New()
	if l1.head.value.(int) <= l2.head.value.(int) {
		l.head = l1.head
		l1.head = l1.head.next
	} else {
		l.head = l2.head
		l2.head = l2.head.next
	}

	c := l.head
	for l1.head != nil && l2.head != nil {
		if l1.head.value.(int) <= l2.head.value.(int) {
			c.next = l1.head
			l1.head = l1.head.next
		} else {
			c.next = l2.head
			l2.head = l2.head.next
		}
		c = c.next
	}

	if l1.head != nil {
		c.next = l1.head
	} else {
		c.next = l2.head
	}

	return l
}
