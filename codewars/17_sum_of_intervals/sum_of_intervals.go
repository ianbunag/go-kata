package sum_of_intervals

import (
	"github.com/yvnbunag/go-kata/lib"
)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)

// @TODO review BST
// @TODO review SOI
func SumOfIntervals(intervals [][2]int) (sum int) {
	intervalsLength := len(intervals)
	isUpperMap := make(map[int]bool, intervalsLength)
	treeNode := lib.TreeNode{Value: intervals[0][0]}

	for _, interval := range intervals {
		lowerBound := interval[0]
		upperBound := interval[1]

		// @TODO investigate intersecting upper and lower
		isUpperMap[lowerBound] = false
		isUpperMap[upperBound] = true

		lowerInserted := treeNode.Insert(lowerBound)
		upperInserted := treeNode.Insert(upperBound)

		if !lowerInserted && !upperInserted {
			continue
		}

		boundSlice := treeNode.ToRangeList(lowerBound, upperBound)

		if boundSlice[0] == lowerBound && boundSlice[1] == upperBound {
			sum += upperBound - lowerBound
			continue
		}

		intervalStack := 0

		for nextIndex, next := range boundSlice {
			if nextIndex == 0 {
				continue
			}

			current := boundSlice[nextIndex-1]
			nextIsUpper, _ := isUpperMap[next]

			if intervalStack == 0 && next == upperBound {
				sum += next - current
				// @TODO Do something with interval stack?
				break
			}

			if intervalStack == 0 && !nextIsUpper {
				sum += next - current
				intervalStack += 1
				continue
			}

			if intervalStack > 0 {
				if nextIsUpper {
					intervalStack -= 1
				}

				if !nextIsUpper {
					intervalStack += 1
				}
			}
		}
	}

	return
}

type OrderedSlice []int
