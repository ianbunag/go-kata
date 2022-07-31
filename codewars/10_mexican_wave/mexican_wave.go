package mexican_wave

import (
	"unicode"
)

// Average time complexity: O(log n)
// Worst time complexity: 	O(n)
// Space complexity: 				O(n log n)
func Wave(words string) []string {
	wave := []string{}

	runeLetters := []rune(words)

	for index, runeLetter := range runeLetters {
		if runeLetter == ' ' {
			continue
		}

		currentRuneLetters := append([]rune{}, runeLetters...)
		currentRuneLetters[index] = unicode.ToUpper(runeLetter)
		wave = append(wave, string(currentRuneLetters))
	}

	return wave
}
