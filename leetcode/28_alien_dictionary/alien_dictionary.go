package alien_dictionary

import (
	"strings"
)

type OrderStatus string

const (
	ORDER_VALID   OrderStatus = "VALID"
	ORDER_INVALID OrderStatus = "INVALID"
)

const DICTIONARY_INVALID = ""

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func AlienOrder(words []string) string {
	wordCharacters := make([][]rune, len(words))
	successors := map[rune][]rune{}

	for word := range words {
		wordCharacters[word] = []rune(words[word])

		for _, character := range wordCharacters[word] {
			successors[character] = []rune{}
		}
	}

	for word := 0; word < len(wordCharacters)-1; word += 1 {
		word1, word2 := wordCharacters[word], wordCharacters[word+1]
		word1Len, word2Len := len(word1), len(word2)
		minLen := min(word1Len, word2Len)

		if word1Len > word2Len && sameWord(word1[:minLen], word2[:minLen]) {
			return DICTIONARY_INVALID
		}

		for character := 0; character < minLen; character += 1 {
			if word1[character] != word2[character] {
				successors[word1[character]] = append(
					successors[word1[character]],
					word2[character],
				)
				break
			}
		}
	}

	tracker := Tracker{}
	reverseAlienOrder := []rune{}
	var order func(character rune) OrderStatus
	order = func(character rune) OrderStatus {
		if tracker.IsVisited(character) {
			if tracker.IsCurrentPath(character) {
				return ORDER_INVALID
			}

			return ORDER_VALID
		}

		tracker.SetIsCurrentPath(character, true)
		for _, successor := range successors[character] {
			if order(successor) == ORDER_INVALID {
				return ORDER_INVALID
			}
		}
		tracker.SetIsCurrentPath(character, false)
		reverseAlienOrder = append(reverseAlienOrder, character)

		return ORDER_VALID
	}

	for character := range successors {
		if order(character) == ORDER_INVALID {
			return DICTIONARY_INVALID
		}
	}

	var alienOrder strings.Builder

	for character := len(reverseAlienOrder) - 1; character >= 0; character-- {
		alienOrder.WriteRune(reverseAlienOrder[character])
	}

	return alienOrder.String()
}

type Tracker map[rune]bool

func (tracker Tracker) IsVisited(value rune) bool {
	_, isVisited := tracker[value]

	return isVisited
}

func (tracker Tracker) IsCurrentPath(value rune) bool {
	return tracker[value]
}

func (tracker Tracker) SetIsCurrentPath(value rune, isCurrentPath bool) {
	tracker[value] = isCurrentPath
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func sameWord(first, second []rune) bool {
	if len(first) != len(second) {
		return false
	}

	for index := range first {
		if first[index] != second[index] {
			return false
		}
	}

	return true
}
