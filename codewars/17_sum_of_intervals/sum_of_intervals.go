package sum_of_intervals

import (
	"github.com/yvnbunag/go-kata/lib"
)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func SumOfIntervals(intervals [][2]int) (sum int) {
	intervalsLength := len(intervals)
	isUpperMap := make(map[int]bool, intervalsLength)
	treeNode := lib.TreeNode{Value: intervals[0][0]}

	for _, interval := range intervals {
		lowerBound := interval[0]
		upperBound := interval[1]

		isUpperMap[lowerBound] = false
		isUpperMap[upperBound] = true

		lowerExists := treeNode.Insert(lowerBound)
		upperExists := treeNode.Insert(upperBound)

		if lowerExists != nil && upperExists != nil {
			continue
		}

		walkerNode := treeNode.Get(lowerBound)

		if walkerNode.Next().Value == upperBound {
			sum += upperBound - lowerBound
			continue
		}
	}

	return
}
