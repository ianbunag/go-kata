package merge_triplets_to_form_target_triplet

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func MergeTriplets(triplets [][]int, target []int) bool {
	startIndex, lastIndex := 0, len(triplets)-1

	for index := range triplets {
		startIndex = index

		if triplets[index][0] <= target[0] &&
			triplets[index][1] <= target[1] &&
			triplets[index][2] <= target[2] {
			break
		}
	}

	for index := startIndex + 1; index <= lastIndex; index += 1 {
		next0, next1, next2 :=
			max(triplets[index-1][0], triplets[index][0]),
			max(triplets[index-1][1], triplets[index][1]),
			max(triplets[index-1][2], triplets[index][2])

		if next0 <= target[0] && next1 <= target[1] && next2 <= target[2] {
			triplets[index][0], triplets[index][1], triplets[index][2] =
				next0,
				next1,
				next2
		} else {
			triplets[index] = triplets[index-1]
		}
	}

	return triplets[lastIndex][0] == target[0] &&
		triplets[lastIndex][1] == target[1] &&
		triplets[lastIndex][2] == target[2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
