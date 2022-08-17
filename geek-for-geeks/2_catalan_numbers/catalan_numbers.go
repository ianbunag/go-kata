package catalan_numbers

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func GetCatalanNumber(position int) int {
	if position < 2 {
		return 1
	}

	cache := map[int]int{}

	return getCatalanFactorial(position*2, &cache) / (getCatalanFactorial(position+1, &cache) * getCatalanFactorial(position, &cache))
}

func getCatalanFactorial(position int, cache *map[int]int) int {
	if position < 2 {
		return 1
	}

	cacheValue, isCached := (*cache)[position]

	if isCached {
		return cacheValue
	}

	value := getCatalanFactorial(position-1, cache) * position
	(*cache)[position] = value

	return value
}
