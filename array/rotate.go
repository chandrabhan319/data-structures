package array

func RotateRight(arr []int, pos int) []int {
	pos = pos % len(arr)
	if pos > 0 {
		return append(arr[len(arr)-pos:], arr[:len(arr)-pos]...)
	}

	return arr
}

func RotateLeft(arr []int, pos int) []int {
	pos = pos % len(arr)
	if pos > 0 {
		return append(arr[pos:], arr[:pos]...)
	}
	return arr
}
