package array

import "math/rand"

func GenerateIntSlice(len int, bound int) []int {
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = rand.Intn(bound)
	}

	return a
}
