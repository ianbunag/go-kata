package is_valid_ip

import (
	"strconv"
	"strings"
)

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func Is_valid_ip(ip string) bool {
	parts := strings.Split(ip, ".")

	if len(parts) < 4 {
		return false
	}

	for _, part := range parts {
		if !isIpPart(part) {
			return false
		}
	}

	return true
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(n)
func isIpPart(part string) bool {
	_, err := strconv.ParseUint(part, 10, 8)

	if err != nil {
		return false
	}

	runePart := []rune(part)

	if len(runePart) > 1 && runePart[0] == '0' {
		return false
	}

	return true
}
