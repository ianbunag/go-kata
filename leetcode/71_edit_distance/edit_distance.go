package edit_distance

// Worst time complexity:   O(n*m)
// Average time complexity: O(n*m)
// Space complexity:        O(n*m)
func MinDistance(word1 string, word2 string) int {
	lenWord1, lenWord2 := len(word1), len(word2)
	// Initialize distance matrix.
	minDistances := make([][]int, lenWord1+1)
	for i := range minDistances {
		minDistances[i] = make([]int, lenWord2+1)
	}
	// Initialize base cases.
	for i := 0; i < lenWord1+1; i += 1 {
		minDistances[i][lenWord2] = lenWord1 - i
	}
	for j := 0; j < lenWord2+1; j += 1 {
		minDistances[lenWord1][j] = lenWord2 - j
	}
	// Calculate min distances.
	for i := lenWord1 - 1; i >= 0; i -= 1 {
		for j := lenWord2 - 1; j >= 0; j -= 1 {
			if word1[i] == word2[j] {
				minDistances[i][j] = minDistances[i+1][j+1]
			} else {
				minDistances[i][j] = 1 + min(
					minDistances[i+1][j],
					minDistances[i][j+1],
					minDistances[i+1][j+1],
				)
			}
		}
	}

	return minDistances[0][0]
}

func min(values ...int) (minValue int) {
	if len(values) == 0 {
		return
	}

	minValue = values[0]

	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return
}
