package merge_intervals

import (
	"sort"
)

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n)
func Merge(intervals [][]int) (mergedIntervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for current := 1; current < len(intervals); current += 1 {
		// Previous interval does not intersect with current interval
		if intervals[current-1][1] < intervals[current][0] {
			mergedIntervals = append(mergedIntervals, intervals[current-1])
			continue
		}

		// Merge previous interval into current interval
		intervals[current][0] = intervals[current-1][0]
		intervals[current][1] = max(intervals[current-1][1], intervals[current][1])
	}

	return append(mergedIntervals, intervals[len(intervals)-1])
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
