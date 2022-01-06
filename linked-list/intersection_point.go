package list

func IntersectionPoint(l1 *List, l2 *List) *node {
	if l1 == nil || l1.head == nil || l2 == nil || l2.head == nil {
		return nil
	}

	p1 := l1.head
	p2 := l2.head

	for p1 != p2 {
		switch {
		case p1 == nil:
			p1 = l2.head
		case p2 == nil:
			p2 = l1.head
		default:
			p1 = p1.next
			p2 = p2.next
		}
	}

	return p1
}
