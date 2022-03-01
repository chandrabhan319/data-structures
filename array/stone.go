package array

func StoneGame(piles []int) bool {
	aliceSum := 0
	bobSum := 0
	len := len(piles)
	for len > 0 {
		if len%2 == 0 {
			if piles[0] >= piles[len-1] {
				aliceSum += piles[0]
				piles = piles[1:]
			} else {
				aliceSum += piles[len-1]
				piles = piles[:len-1]
			}
		} else {
			if piles[0] > piles[len-1] {
				bobSum += piles[0]
				piles = piles[1:]
			} else {
				bobSum += piles[len-1]
				piles = piles[:len-1]
			}
		}
		len--
	}

	return aliceSum > bobSum
}

func StoneGameII(piles []int) bool {
	len := len(piles)
	dp := make([][]int, len)
	for i := range dp {
		dp[i] = make([]int, len)
	}

	for i := range dp {
		dp[i][i] = piles[i]
	}

	for j := 1; j < len; j++ {
		for i := 0; i < len-j; i++ {
			dp[i][i+j] = max(piles[i]-dp[i+1][i+j], piles[i+j]-dp[i][i+j-1])
		}
	}

	return dp[0][len-1] > 0
}

func StoneGameIII(piles []int) bool {
	len := len(piles)
	dp := piles
	for j := 1; j < len; j++ {
		for i := 0; i < len-j; i++ {
			dp[i] = max(piles[i]-dp[i+1], piles[i+j]-dp[i])
		}
	}

	return dp[0] > 0
}
