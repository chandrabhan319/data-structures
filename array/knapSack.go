package array

//[]int{90, 100, 120}, []int{2, 3, 3}, 5
func KnapSack01(wt, val []int, cap int) int {
	// return knapSack(wt, val, cap, len(wt))
	// return knapSackDP01(wt, val, cap, len(wt))
	return knapSackDP01II(wt, val, cap, len(wt))

}

func knapSack(wt, val []int, cap, n int) int {
	if n == 0 || cap == 0 {
		return 0
	}

	if cap < wt[n-1] {
		return knapSack(wt, val, cap, n-1)
	}

	return max(val[n-1]+knapSack(wt, val, cap-wt[n-1], n-1), knapSack(wt, val, cap, n-1))
}

func knapSackDP01(wt, val []int, cap, n int) int {
	if cap == 0 || n == 0 {
		return 0
	}

	k := make([][]int, n+1)
	for i := range k {
		k[i] = make([]int, cap+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < cap+1; j++ {
			if wt[i-1] > j {
				k[i][j] = k[i-1][j]
			} else {
				k[i][j] = max(val[i-1]+k[i-1][j-wt[i-1]], k[i-1][j])
			}
		}
	}

	return k[n][cap]
}

func knapSackDP01II(wt, val []int, cap, n int) int {
	if cap == 0 || n == 0 {
		return 0
	}

	k := make([]int, cap+1)
	for i := 1; i < n+1; i++ {
		for j := cap; j > 0; j-- {
			if wt[i-1] <= j {
				k[j] = max(k[j], val[i-1]+k[j-wt[i-1]])
			}
		}
	}

	return k[cap]
}

// given a slice of odd and even numbers, seggregate the slice such that all the even comes
// before the odd numbers
