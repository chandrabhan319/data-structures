package array

import "fmt"

func SumCombinationOf134(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 1
	dp[3] = 2
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-3] + dp[i-4]
	}

	fmt.Printf("Possible number of sequence array for %d to be represented as sum of 1, 3 and 4: %v", n, dp)
	return dp[n]
}
