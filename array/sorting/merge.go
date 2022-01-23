package sorting

// func MergeSort(a []int) []int {
// 	return mergeSort(a)
// }

// func mergeSort(a []int) []int {
// 	if len(a) < 2 {
// 		return a
// 	}

// 	la := mergeSort(a[:len(a)/2])
// 	ra := mergeSort(a[len(a)/2:])

// 	return merge(la, ra)
// }

// func merge(a, b []int) []int {
// 	final := []int{}
// 	i := 0
// 	j := 0
// 	for i < len(a) && j < len(b) {
// 		if a[i] < b[j] {
// 			final = append(final, a[i])
// 			i++
// 		} else {
// 			final = append(final, b[j])
// 			j++
// 		}
// 	}

// 	for ; i < len(a); i++ {
// 		final = append(final, a[i])
// 	}
// 	for ; j < len(b); j++ {
// 		final = append(final, b[j])
// 	}

// 	return final
// }

func MergeSort(arr []int) []int {
	return mergeSort(arr)
}
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	ls := mergeSort(arr[:len(arr)/2])
	rs := mergeSort(arr[len(arr)/2:])

	return merge(ls, rs)
}

func merge(ls, rs []int) []int {
	i := 0
	j := 0
	final := []int{}
	for i < len(ls) && j < len(rs) {
		if ls[i] < rs[j] {
			final = append(final, ls[i])
			i++
		} else {
			final = append(final, rs[j])
			j++
		}

	}
	for ; i < len(ls); i++ {
		final = append(final, ls[i])
	}
	for ; j < len(rs); j++ {
		final = append(final, rs[j])
	}
	return final
}
