package first_non_repeating_character

import "unicode"

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func FirstNonRepeating(str string) string {
	repeatTracker := make(RepeatTracker)

	for _, part := range str {
		repeatTracker.track(part)
	}

	for _, part := range str {
		if !repeatTracker.isRepeated(part) {
			return string(part)
		}
	}

	return ""
}

const LETTER_RUNE_OFFSET = 'a' - 'A'

type RepeatTracker map[rune]bool

func (repeatTracker RepeatTracker) track(value rune) {
	_, tracked := repeatTracker[repeatTracker.hash(value)]

	if !tracked {
		repeatTracker[repeatTracker.hash(value)] = false

		return
	}

	repeatTracker[repeatTracker.hash(value)] = true
}

func (repeatTracker RepeatTracker) isRepeated(value rune) bool {
	return repeatTracker[repeatTracker.hash(value)]
}

func (repeatTracker RepeatTracker) hash(value rune) rune {
	return unicode.ToLower(value)
}
