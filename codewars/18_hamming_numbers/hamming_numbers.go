package hamming_numbers

import lib_unsigned "github.com/yvnbunag/go-kata/lib/unsigned"

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func Hammer(n int) uint {
	hammerSequence := make([]uint, n)
	hammerSequence[0] = 1
	var twoMultiplier, threeMultiplier, fiveMultiplier uint = 2, 3, 5
	var twoCtr, threeCtr, fiveCtr uint = 0, 0, 0

	for index := 1; index < n; index += 1 {
		hammerSequence[index] = lib_unsigned.Min(
			twoMultiplier,
			lib_unsigned.Min(threeMultiplier, fiveMultiplier),
		)

		if hammerSequence[index] == twoMultiplier {
			twoCtr += 1
			twoMultiplier = 2 * hammerSequence[twoCtr]
		}
		if hammerSequence[index] == threeMultiplier {
			threeCtr += 1
			threeMultiplier = 3 * hammerSequence[threeCtr]
		}
		if hammerSequence[index] == fiveMultiplier {
			fiveCtr += 1
			fiveMultiplier = 5 * hammerSequence[fiveCtr]
		}
	}

	return hammerSequence[n-1]
}
