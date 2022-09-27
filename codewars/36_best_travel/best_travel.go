package best_travel

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n)
func ChooseBestSum(max, towns int, distances []int) int {
	best := -1

	if towns > len(distances) {
		return best
	}

	for distanceIndex, distance := range distances {
		currentBest := distance

		if towns > 1 {
			currentBest += ChooseBestSum(
				max-distance,
				towns-1,
				distances[distanceIndex+1:],
			)
		}

		if currentBest >= distance && currentBest > best && currentBest <= max {
			best = currentBest
		}
	}

	return best
}
