package array

import (
	"math"
)

//[1,0,3,7,9,0,11,13]=>[1,3,7,9,11,13,0,0]
/*
1. [1,0,3,7,9,0,11,13]
2. [1,3,0,7,9,0,11,13]
3. [1,3,7,0,9,0,11,13]
4. [1,3,7,9,0,0,11,13]
5. [1,3,7,9,11,13,0,0]
*/

func MoveZeros(a []int) []int {
	l := len(a)
	b := make([]int, l)
	j := 0
	for k, v := range a {
		if v != 0 {
			b[j] = a[k]
			j++
		}
	}

	return b
}

//2.  max sum array- kadane's algorithm
// [-2,1,-2,4,-2,1,3,-7,1,3,2,1,1] = 6

func MaxSum(a []int) int {
	max := math.MinInt
	sum := 0
	for _, v := range a {
		sum = sum + v
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

//3. find the subarray with the max sum

func MaxSubarray(a []int) int {
	b := make([]int, len(a))
	max := a[0]
	b[0] = a[0]
	for i := 1; i < len(a); i++ {
		if a[i]+b[i-1] > a[i] {
			b[i] = a[i] + b[i-1]
		} else {
			b[i] = a[i]
		}

		if b[i] > max {
			max = b[i]
		}
	}
	return max
}

// [][]int {{1,2,3},{}}

// 29 jan 2022

// 4. given a list of prices of stock over a number of days. u hav to find the max profit u can get by buying and selling a stock
// u can only buy stock once and hold it for n number of days.
// u can decide not to buy stock if there is no profit.
// u cannot sell a stock before buying
// if there is no profit u hav to return 0.
// where n > 0 , prices [i] >= 0
// test case : [7,1,3,5,2,4,2]  output(max profit) = 4

func StockProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	buy := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		if price < buy {
			buy = price
		}
		profit := price - buy
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func StockProfitII(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s0[0] = 0
	s1[0] = -prices[0]

	for i := 1; i < len(prices); i++ {

		s1[i] = max(s1[i-1], -prices[i])
		s0[i] = max(s0[i-1], s1[i-1]+prices[i])
	}
	return s0[len(prices)-1]
}

//2. u can only hold one stock at a time.
// intraday transactions are allowed.
//u cannot buy multiple stock on the same day
// [6,4,3,5,0]
func StcoProfitIntrdDay(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

// wallet capacity : n
// mutiple stock purchase allowed in addition with all above conditions
// calc max profit
//[7,1,3,5,2,4,2]
func StockMultPurchase(prices []int, wallet int) int {
	if len(prices) <= 1 {
		return 0
	}

	currWallet := wallet
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			noOfStocks := currWallet / prices[i-1]
			currWallet = currWallet + (prices[i]-prices[i-1])*noOfStocks + (currWallet % prices[i-1])
		}
	}

	return currWallet - wallet
}

// transaction fee = 2

//[7,3,2,5,2,4,2]
// wallet = 5

func StockProfitTransactionFee(prices []int, fee int) int {
	if len(prices) < 1 {
		return 0
	}
	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s0[0] = 0
	s1[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s1[i-1]+prices[i]-fee)
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
	}
	return s0[len(prices)-1]
}
func StockProfitTransactionFeeII(prices []int, fee int) int {
	if len(prices) < 1 {
		return 0
	}

	sell := 0
	buy := -prices[0]
	for i := 1; i < len(prices); i++ {
		prevBuy := buy
		buy = max(buy, sell-prices[i])
		sell = max(sell, prevBuy+prices[i]-fee)
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// [1,3,2,8,4,9], fee = 2 , o/p = 8
//prevBuy = 1
// buy = 1
// sell = 8

// s0 = {-1,1,1,6,6,9}
// s1 = {0,0,0,0,2,2}

func StcoProfitIntrdDayII(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s0[0] = 0
	s1[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s1[i-1]+prices[i])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
	}
	return s0[len(prices)-1]
}

// cooldown one day

func ProfitRestDay(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s2 := make([]int, len(prices))

	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = math.MinInt

	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	return max(s0[len(prices)-1], s2[len(prices)-1])
}

func ProfitRestDayII(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	sell := 0
	buy := -prices[0]
	rest := math.MinInt

	for i := 1; i < len(prices); i++ {
		prevBuy := buy
		prevRest := rest
		buy = max(buy, sell-prices[i])
		rest = prevBuy + prices[i]
		sell = max(sell, prevRest)
	}

	return max(sell, rest)
}
func ProfitRestDayIII(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	sell := 0
	buy := -prices[0]
	rest := math.MinInt

	for i := 1; i < len(prices); i++ {
		prevRest := rest
		rest = buy + prices[i]
		buy = max(buy, sell-prices[i])
		sell = max(sell, prevRest)
	}

	return max(sell, rest)
}
