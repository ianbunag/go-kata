package human_readable_time

import "fmt"

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func HumanReadableTime(seconds int) string {
	minutes, remainingSeconds := seconds/60, seconds%60
	hours, remainingMinutes := minutes/60, minutes%60

	return fmt.Sprintf("%02d:%02d:%02d", hours, remainingMinutes, remainingSeconds)
}
