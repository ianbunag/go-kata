package permutations

// Worst time complexity:   O(n!)
// Average time complexity: O(n!)
// Space complexity:        O(n!)
func Permute(nums []int) (permutations [][]int) {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	for index, num := range nums {
		rest := append([]int{}, nums[:index]...)
		rest = append(rest, nums[index+1:]...)

		for _, permutation := range Permute(rest) {
			permutations = append(permutations, append(permutation, num))
		}
	}

	return
}
