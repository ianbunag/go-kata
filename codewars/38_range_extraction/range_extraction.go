package range_extraction

import (
	"strconv"
	"strings"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n log n)
func Solution(list []int) string {
	var rangeBuilder RangeBuilder
	listLength := len(list)
	lowerBoundary := 0
	upperBoundary := 0

	for index := 1; index <= listLength; index += 1 {
		if index < listLength && list[index-1]+1 == list[index] {
			upperBoundary = index
			continue
		}

		if upperBoundary-lowerBoundary > 1 {
			rangeBuilder.AddRange(
				list[lowerBoundary],
				list[upperBoundary],
			)
		} else if upperBoundary != lowerBoundary {
			rangeBuilder.Add(list[lowerBoundary])
			rangeBuilder.Add(list[upperBoundary])
		} else {
			rangeBuilder.Add(list[lowerBoundary])
		}

		lowerBoundary = index
		upperBoundary = index
	}

	return rangeBuilder.String()
}

type RangeBuilder struct {
	builder strings.Builder
}

func (rangeBuilder *RangeBuilder) prefix() {
	if rangeBuilder.builder.Len() > 0 {
		rangeBuilder.builder.WriteString(",")
	}
}

func (rangeBuilder *RangeBuilder) AddRange(lower, upper int) {
	rangeBuilder.prefix()
	rangeBuilder.builder.WriteString(strconv.Itoa(lower))
	rangeBuilder.builder.WriteString("-")
	rangeBuilder.builder.WriteString(strconv.Itoa(upper))
}

func (rangeBuilder *RangeBuilder) Add(value int) {
	rangeBuilder.prefix()
	rangeBuilder.builder.WriteString(strconv.Itoa(value))
}

func (rangeBuilder *RangeBuilder) String() string {
	return rangeBuilder.builder.String()
}
