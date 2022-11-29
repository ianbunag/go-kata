package josephus_survivor

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func JosephusSurvivor(initialSurvivors, step int) int {
	if initialSurvivors == 1 {
		return 1
	}

	// See https://cs.stackexchange.com/questions/7048/a-recursive-formula-for-generalized-josephus-problem
	return ((JosephusSurvivor(initialSurvivors-1, step) + step - 1) % initialSurvivors) + 1
}
