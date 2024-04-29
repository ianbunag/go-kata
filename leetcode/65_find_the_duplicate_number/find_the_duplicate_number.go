package find_the_duplicate_number

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func FindDuplicate(nums []int) int {
	slow, fast, slow2 := 0, 0, 0

	for {
		slow, fast = nums[slow], nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	for slow != slow2 {
		slow, slow2 = nums[slow], nums[slow2]
	}

	return slow
}
