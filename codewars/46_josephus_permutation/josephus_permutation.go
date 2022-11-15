package josephus_permutation

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func Josephus(items []interface{}, place int) []interface{} {
	result := []interface{}{}

	for remainingItems := len(items); remainingItems > 0; remainingItems -= 1 {
		offsetPlace := place % remainingItems

		if offsetPlace < 1 {
			offsetPlace = remainingItems
		}

		tail := items[:offsetPlace-1]
		target := items[offsetPlace-1]
		head := items[offsetPlace:]
		result = append(result, target)
		items = append(head, tail...)
	}

	return result
}
