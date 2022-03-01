package array

func RotateRight(a []int, k int) []int {
	rk := (len(a) + k) % len(a)
	if rk > 0 {
		return append(a[len(a)-rk:], a[:len(a)-rk]...)
	}

	return a
}

func RotateLeft(a []int, k int) []int {
	lk := (len(a) + k) % len(a)
	if lk > 0 {
		return append(a[lk:], a[:lk]...)
	}

	return a
}
