package sorting

// given a slice of odd and even numbers, seggregate the slice such that all the even comes
// before the odd numbers
//[1,2,3,4,5,6]
// low(even) = 2
// high(odd) = 5
func SeggregateEvenOdd(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	low := 0
	high := len(arr) - 1

	for low < high {
		for (arr[low]%2 == 0) && (low < high) {
			low++
		}
		for (arr[high]%2 != 0) && (low < high) {
			high--
		}
		if low < high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}

	return arr
}
