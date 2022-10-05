package permutations

import (
	"strings"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n!)
func Permutations(value string) []string {
	values := strings.Split(value, "")

	combinations := findCombinations(values)
	uniqueCombinations := []string{}
	uniqueMap := make(map[string]interface{})

	for _, combination := range combinations {
		_, found := uniqueMap[combination]

		if found {
			continue
		}

		uniqueCombinations = append(uniqueCombinations, combination)
		uniqueMap[combination] = nil
	}

	return uniqueCombinations
}

// Average time complexity: O(n!)
// Worst time complexity:   O(n!)
// Space complexity:        O(n!)
func findCombinations(values []string) []string {
	if len(values) == 1 {
		return values
	}

	combinations := []string{}

	for index, value := range values {
		restValues := append([]string{}, values[0:index]...)
		restValues = append(restValues, values[index+1:]...)

		restCombinations := findCombinations(restValues)

		for _, restCombination := range restCombinations {
			combinations = append(combinations, value+restCombination)
		}
	}

	return combinations
}
