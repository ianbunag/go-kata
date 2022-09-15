package weight_for_weight

import (
	"sort"
	"strings"
)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(n)
func OrderWeight(strng string) string {
	weights := strings.Split(strng, " ")
	sortedWeights := sortWeightsAscending(weights)

	return strings.Join(sortedWeights, " ")
}

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n)
func sortWeightsAscending(weights []string) []string {
	hashedWeights := make([]HashedWeight, len(weights))

	for index, weight := range weights {
		hashedWeights[index].setWeight(weight)
	}

	sort.Slice(hashedWeights, func(first, second int) bool {
		if hashedWeights[first].getHash() == hashedWeights[second].getHash() {
			return hashedWeights[first].getWeight() < hashedWeights[second].getWeight()
		}

		return hashedWeights[first].getHash() < hashedWeights[second].getHash()
	})

	sortedWeights := make([]string, len(hashedWeights))

	for index, hashedWeight := range hashedWeights {
		sortedWeights[index] = hashedWeight.getWeight()
	}

	return sortedWeights
}

type HashedWeight [2]interface{}

func (hashedWeight *HashedWeight) setWeight(weight string) {
	var hash rune

	for _, part := range weight {
		hash += part - '0'
	}

	hashedWeight[0] = weight
	hashedWeight[1] = hash
}

func (hashedWeight *HashedWeight) getWeight() string {
	return hashedWeight[0].(string)
}

func (hashedWeight *HashedWeight) getHash() rune {
	return hashedWeight[1].(rune)
}
