package jump_game

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func CanJump(nums []int) bool {
	goal := len(nums) - 1

	for index := len(nums) - 2; index >= 0; index -= 1 {
		if index+nums[index] >= goal {
			goal = index
		}
	}

	return goal == 0
}
