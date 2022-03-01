package sorting

func MergeSort(a []int) []int {
	return mergeSort(a)
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	ls := mergeSort(a[:len(a)/2])
	rs := mergeSort(a[len(a)/2:])

	return merge(ls, rs)
}

func merge(a, b []int) []int {
	sortedArray := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sortedArray = append(sortedArray, a[i])
			i++
		} else {
			sortedArray = append(sortedArray, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		sortedArray = append(sortedArray, a[i])
	}

	for ; j < len(b); j++ {
		sortedArray = append(sortedArray, b[j])
	}

	return sortedArray
}
