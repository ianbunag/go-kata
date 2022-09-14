package pete_the_baker

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func Cakes(recipe, available map[string]int) int {
	maximum := 0

	for ingredient, quantity := range recipe {
		ingredientMaximum := available[ingredient] / quantity

		if ingredientMaximum == 0 {
			return 0
		}

		if ingredientMaximum < maximum || maximum == 0 {
			maximum = ingredientMaximum
		}
	}

	return maximum
}
