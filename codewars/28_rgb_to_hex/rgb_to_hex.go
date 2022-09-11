package rgb_to_hex

import "fmt"

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func RGB(r, g, b int) string {
	return fmt.Sprintf("%02X%02X%02X", normalize(r), normalize(g), normalize(b))
}

func normalize(number int) int {
	if number < 0 {
		return 0
	}

	if number > 255 {
		return 255
	}

	return number
}
