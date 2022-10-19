package rot13

import (
	"strings"
)

const (
	CHARACTERS = 'z' - 'a' + 1
	ROTATION   = CHARACTERS / 2
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Rot13(msg string) string {
	var rotated strings.Builder

	for _, value := range msg {
		rotatedValue := value

		if value >= 'a' && value <= 'z' {
			rotatedValue += ROTATION

			if rotatedValue > 'z' {
				rotatedValue -= CHARACTERS
			}
		}

		if value >= 'A' && value <= 'Z' {
			rotatedValue += ROTATION

			if rotatedValue > 'Z' {
				rotatedValue -= CHARACTERS
			}
		}

		rotated.WriteRune(rotatedValue)
	}

	return rotated.String()
}
