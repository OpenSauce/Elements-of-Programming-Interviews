package chapter_5

import "math"

func BuySellStockOnce(prices []int) int {
	min := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		todaysProfit := price - min
		maxProfit = int(math.Max(float64(todaysProfit), float64(maxProfit)))
		min = int(math.Min(float64(min), float64(price)))
	}
	return maxProfit
}
