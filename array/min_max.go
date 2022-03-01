package array

import "fmt"

func MinMax(a []int) (int, int, error) {
	if a == nil {
		return 0, 0, fmt.Errorf("slice is empty")
	}

	min := 0
	max := len(a) - 1

	for i := min; i < max; i++ {

	}

	return a[min], a[max], nil
}
