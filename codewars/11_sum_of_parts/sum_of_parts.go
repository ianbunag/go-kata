package sum_of_parts

// Average time complexity: O(n)
// Worst time complexity: 	O(n)
// Space complexity: 				O(1)
func PartsSums(list []uint64) []uint64 {
	list = append(list, 0)

	for index := len(list) - 2; index >= 0; index -= 1 {
		list[index] = list[index] + list[index+1]
	}

	return list
}
