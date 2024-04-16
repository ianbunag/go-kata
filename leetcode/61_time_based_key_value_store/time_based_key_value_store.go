package time_based_key_value_store

type TimeValue struct {
	timestamp int
	value     string
}

type TimeMap map[string][]TimeValue

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(1)
func Constructor() TimeMap {
	return TimeMap{}
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (timeMap *TimeMap) Set(key string, value string, timestamp int) {
	(*timeMap)[key] = append((*timeMap)[key], TimeValue{timestamp, value})
}

// Worst time complexity:   O(log n)
// Average time complexity: O(log n)
// Space complexity:        O(1)
func (timeMap *TimeMap) Get(key string, timestamp int) string {
	timeValues := (*timeMap)[key]
	left, right, closest := 0, len(timeValues)-1, ""

	for left <= right {
		middle := (left + right) / 2

		if timeValues[middle].timestamp == timestamp {
			return timeValues[middle].value
		}

		if timeValues[middle].timestamp < timestamp {
			left, closest = middle+1, timeValues[middle].value
			continue
		}

		right = middle - 1
	}

	return closest
}
