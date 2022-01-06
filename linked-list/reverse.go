package list

func (l *List) Reverse() {
	switch {
	case l.head == nil:
		return
	case l.head.next == nil:
		return
	}

	a := l.head
	tmpHead := l.head.next
	c := l.head.next.next // i.e tmpHead.next
	a.next = nil

	for c != nil {
		tmpHead.next = a
		a = tmpHead
		tmpHead = c
		c = c.next
	}

	tmpHead.next = a
	l.head = tmpHead
}
