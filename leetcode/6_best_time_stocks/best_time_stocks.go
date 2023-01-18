package best_time_stocks

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func MaxProfit(prices []int) int {
	buyValue, maximumProfit := prices[0], 0

	for _, price := range prices[1:] {
		if price < buyValue {
			buyValue = price
			continue
		}

		if profit := price - buyValue; maximumProfit < profit {
			maximumProfit = profit
		}
	}

	return maximumProfit
}
