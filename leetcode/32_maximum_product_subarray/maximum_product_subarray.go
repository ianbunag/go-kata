package maximum_product_subarray

import "math"

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func MaxProduct(values []int) int {
	result, currentMax, currentMin := values[0], values[0], values[0]

	for _, value := range values[1:] {
		nextMax := value * currentMax
		currentMax = max(nextMax, value*currentMin, value)
		currentMin = min(nextMax, value*currentMin, value)
		result = max(result, currentMax)
	}

	return result
}

func max(values ...int) int {
	result := math.MinInt

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}

func min(values ...int) int {
	result := math.MaxInt

	for _, value := range values {
		if value < result {
			result = value
		}
	}

	return result
}
