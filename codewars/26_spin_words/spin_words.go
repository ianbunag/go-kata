package spin_words

import (
	"strings"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func SpinWords(str string) string {
	var wordSpinner WordSpinner

	for _, character := range str {
		if character != ' ' {
			wordSpinner.Queue(character)
			continue
		}

		wordSpinner.Flush().Add(' ')
	}

	return wordSpinner.Flush().String()
}

type WordSpinner struct {
	sentence strings.Builder
	queue    []rune
}

func (spinner *WordSpinner) Add(character rune) {
	spinner.sentence.WriteRune(character)
}

func (spinner *WordSpinner) Queue(character rune) {
	spinner.queue = append(spinner.queue, character)
}

func (spinner *WordSpinner) Flush() *WordSpinner {
	queueLength, queue := len(spinner.queue), spinner.queue
	spinner.queue = []rune{}

	if queueLength > 4 {
		for ctr := queueLength - 1; ctr > -1; ctr -= 1 {
			spinner.Add(queue[ctr])
		}

		return spinner
	}

	for _, flushCharacter := range queue {
		spinner.Add(flushCharacter)
	}

	return spinner
}

func (spinner *WordSpinner) String() string {
	return spinner.sentence.String()
}
