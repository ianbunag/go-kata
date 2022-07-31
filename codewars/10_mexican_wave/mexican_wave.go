package mexican_wave

import (
	"unicode"
)

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
