package array

import (
	"fmt"
	"math"
)

func StockMaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min, max := math.MaxInt, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		profit := prices[i] - min
		if max < profit {
			max = profit
		}
	}

	return max
}

func StockMaxProfitII(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	totalProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			totalProfit += prices[i] - prices[i-1]
		}
	}

	return totalProfit
}

func MaxProfit(prices []int, fee int) int {
	if len(prices) <= 0 {
		return 0
	}

	buy := -prices[0]
	sell := 0
	for i := 1; i < len(prices); i++ {
		tmp := buy
		buy = max(buy, sell-prices[i])
		sell = max(sell, tmp+prices[i]-fee)
	}

	return max(buy, sell)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func StockMaxProfitHard(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	hold1 := math.MinInt
	hold2 := math.MinInt
	release1 := 0
	release2 := 0

	for _, i := range prices {
		release2 = max(release2, hold2+i)
		hold2 = max(hold2, release1-i)
		release1 = max(release1, hold1+i)
		hold1 = max(hold1, -i)
	}

	return release2
}

func StockMaxProfitEHard(n int, prices []int) int {
	len := len(prices)
	if n >= len/2 {
		return StockMaxProfitII(prices)
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, len)
	}

	for k := 1; k <= n; k++ {
		localMax := math.MinInt
		for i := 1; i < len; i++ {
			localMax = max(localMax, dp[k-1][i-1]-prices[i-1])
			dp[k][i] = max(dp[k][i-1], prices[i]+localMax)
		}
		fmt.Println(dp)
	}

	return dp[n][len-1]
}

/*
prices = [3, 3, 5, 0, 0, 3, 1, 4]
k = 1
localMax = 0
i = 5
dp = {[0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 2, 2, 2, 3, 3, 4], [0, 0, 2, 2, 2, 5, 5, 6]}
*/
