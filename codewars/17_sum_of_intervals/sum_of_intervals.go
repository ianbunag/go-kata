package sum_of_intervals

import (
	"sort"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n log n)
func SumOfIntervals(intervals [][2]int) (sum int) {
	intervalPoints := []IntervalPoint{}
	intervalStack := []IntervalPoint{}
	var lowerInterval IntervalPoint

	for _, currentInterval := range intervals {
		intervalPoints = append(
			intervalPoints,
			IntervalPoint{Value: currentInterval[0], Bound: LowerBound},
			IntervalPoint{Value: currentInterval[1], Bound: UpperBound},
		)
	}

	sort.Slice(intervalPoints, func(x, y int) bool {
		return intervalPoints[x].Value < intervalPoints[y].Value
	})

	for _, interval := range intervalPoints {
		if interval.Bound == LowerBound {
			intervalStack = append(intervalStack, interval)
			continue
		}

		lowerIntervalStackLen := len(intervalStack)
		lowerInterval, intervalStack = intervalStack[lowerIntervalStackLen-1],
			intervalStack[:lowerIntervalStackLen-1]

		if len(intervalStack) == 0 {
			sum += interval.Value - lowerInterval.Value
		}
	}

	return
}

type IntervalBound int

type IntervalPoint struct {
	Value int
	Bound IntervalBound
}

const (
	LowerBound IntervalBound = iota
	UpperBound
)
