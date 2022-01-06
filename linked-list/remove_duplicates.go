package list

func (l *List) RemoveConsecutiveDuplicates() {
	if l == nil || l.head == nil {
		return
	}

	c := l.head
	for c.next != nil {
		if c.value == c.next.value {
			c.next = c.next.next
		} else {
			c = c.next
		}
	}
}

func (l *List) RemoveDuplicates() {
	if l == nil || l.head == nil {
		return
	}

	exists := make(map[interface{}]bool)
	exists[l.head.value] = true
	slow := l.head
	fast := l.head.next
	for fast != nil {
		if _, ok := exists[fast.value]; ok {
			slow.next = slow.next.next
			fast = fast.next
		} else {
			exists[fast.value] = true
			fast = fast.next
			slow = slow.next
		}
	}
}

// removing non-consecutive duplicates
func (l *List) RemDups() {
	if l.head == nil || l.head.next == nil {
		return
	}
	a := l.head

	for a.next != nil {
		b := a
		c := b.next
		for c.next != nil {
			if a.value == c.value {
				c = c.next
				b.next = c
			} else {
				b = b.next
				c = c.next
			}
			if a.value == c.value {
				b.next = nil
			}
		}
		if a.value == a.next.value {
			a.next = nil
		} else {
			a = a.next
		}

	}
}

func (l *List) RemConsDups() {
	if l.head == nil || l.head.next == nil {
		return
	}
	a := l.head
	b := l.head.next

	for b.next != nil {
		if a.value == b.value {
			b = b.next
			a.next = b
		} else {
			a = a.next
			b = b.next
		}
		if a.value == b.value {
			a.next = nil
		}
	}
}
