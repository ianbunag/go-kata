package partition_labels

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func PartitionLabels(s string) (partitions []int) {
	lastPositions := make(map[rune]int, 26)
	size, end := 0, 0

	for position, char := range s {
		lastPositions[char] = position
	}

	for position, char := range s {
		size += 1

		if lastPositions[char] > end {
			end = lastPositions[char]
		}

		if position == end {
			partitions, size = append(partitions, size), 0
		}
	}

	return
}
