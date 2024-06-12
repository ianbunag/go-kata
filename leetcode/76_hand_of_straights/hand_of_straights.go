package hand_of_straights

import (
	"sort"
)

// IsNStraightHand
// Worst time complexity:   O(n log n)
// Average time complexity: O(n log n)
// Space complexity:        O(n)
func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)
	counts := make(map[int]int)

	for _, card := range hand {
		counts[card] += 1
	}

	for _, card := range hand {
		if counts[card] == 0 {
			continue
		}

		for index := 0; index < groupSize; index += 1 {
			if counts[card+index] == 0 {
				return false
			}

			counts[card+index] -= 1
		}
	}

	return true
}
