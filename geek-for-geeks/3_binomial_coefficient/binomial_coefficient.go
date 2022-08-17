package binomial_coefficient

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func GetBinomialCoefficient(n, k int) int {
	if k > n {
		return 0
	}

	factorialTable := make([]int, n+1)
	factorialTable[0] = 1

	for position := 1; position <= n; position += 1 {
		factorialTable[position] = factorialTable[position-1] * position
	}

	return factorialTable[n] / (factorialTable[k] * factorialTable[n-k])
}
