package minimum_window_substring

import "strings"

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func MinWindow(source string, substring string) string {
	sourceRunes := []rune(source)
	window := NewWindow(substring)
	minimumWindow := MinimumWindow{}
	leftPointer := 0

	for rightPointer, character := range sourceRunes {
		// Track character counts
		window.Add(character)

		// If current window contains substring, compare and update minimum window with current window
		if window.Satisfied() {
			minimumWindow.Update(sourceRunes[leftPointer : rightPointer+1])
		}

		// While current window contains substring, increase left pointer
		for window.Satisfied() {
			window.Remove(sourceRunes[leftPointer])
			leftPointer += 1
		}

		// If current window no longer contains substring, compare and update minimum window with current window
		if leftPointer > 0 {
			minimumWindow.Update(sourceRunes[leftPointer-1 : rightPointer+1])
		}
	}

	return minimumWindow.String()
}

type MinimumWindow []rune

func (minimumWindow *MinimumWindow) Update(window []rune) {
	// If minimum window is shorter than window, keep minimum window
	if maxLength := len(*minimumWindow); maxLength != 0 && maxLength < len(window) {
		return
	}

	*minimumWindow = window
}

func (minimumWindow *MinimumWindow) String() string {
	var result strings.Builder

	for _, character := range *minimumWindow {
		result.WriteRune(character)
	}

	return result.String()
}

func NewWindow(substring string) Window {
	need := make(map[rune]int, len(substring))
	have := make(map[rune]int, len(substring))

	for _, character := range substring {
		need[character] += 1
	}

	missing := len(need)

	return Window{need: need, have: have, missing: missing}
}

type Window struct {
	need    map[rune]int
	have    map[rune]int
	missing int
}

func (window *Window) update(character rune, movement int) {
	_, found := window.need[character]

	// If character is not needed, skip
	if !found {
		return
	}

	oldHave := window.have[character]
	window.have[character] += movement
	needed := window.need[character]
	newHave := window.have[character]

	// If character count transitioned from insufficient to satisfied, decrease missing
	if oldHave < needed && newHave == needed {
		window.missing -= 1
	}

	// If character count transitioned from satisfied to insufficient, increase missing
	if oldHave == needed && newHave < needed {
		window.missing += 1
	}
}

func (window *Window) Add(character rune) {
	window.update(character, 1)
}

func (window *Window) Remove(character rune) {
	window.update(character, -1)
}

func (window *Window) Satisfied() bool {
	return window.missing == 0
}
