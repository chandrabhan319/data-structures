package list

func (l *List) MergeSort() {
	if l == nil || l.head == nil {
		return
	}

	l.head = l.head.mergeSort()
}

func (h *node) mergeSort() *node {
	if h == nil || h.next == nil {
		return h
	}

	mid := h.getMiddle()
	midNext := mid.next
	mid.next = nil
	left := h.mergeSort()
	right := midNext.mergeSort()
	ll := New()
	ll.head = left
	rl := New()
	rl.head = right
	sortedList := MergeSorted(ll, rl)
	return sortedList.head
}

func (h *node) getMiddle() *node {
	if h == nil {
		return h
	}

	slow := h
	fast := h

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
