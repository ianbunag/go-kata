package longest_common_subsequence

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func LongestCommonSubsequence(text1 string, text2 string) int {
	subsequences := make([]int, len(text1)+1)

	for position2 := len(text2) - 1; position2 >= 0; position2 -= 1 {
		nextSubsequences := make([]int, len(subsequences))

		for position1 := len(text1) - 1; position1 >= 0; position1 -= 1 {
			if text1[position1] == text2[position2] {
				// Get diagonal value (longest subsequence before this match) plus one
				nextSubsequences[position1] = subsequences[position1+1] + 1
				continue
			}

			// Get max of bottom and right values (longest subsequence across
			// 	subproblems before this comparison)
			nextSubsequences[position1] = max(
				nextSubsequences[position1+1],
				subsequences[position1],
			)
		}

		subsequences = nextSubsequences
	}

	return subsequences[0]
}

func max(first, second int) int {
	if first > second {
		return first
	}

	return second
}
