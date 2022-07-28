package duplicate_count

import "strings"

func CountDuplicate(value string) (duplicates int) {
	duplicateMap := make(map[rune]bool)

	for _, current := range strings.ToLower(value) {
		counted, found := duplicateMap[current]

		if found && !counted {
			duplicates += 1
			duplicateMap[current] = true
		}

		if !found {
			duplicateMap[current] = false
		}
	}

	return duplicates
}
