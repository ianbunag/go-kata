package count_ip_addresses

import (
	"fmt"
)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func IpsBetween(start, end string) int {
	startCount := 0
	endCount := 0

	for index, value := range ipToArray(start) {
		startCount += value << ((3 - index) * 8)
	}
	for index, value := range ipToArray(end) {
		endCount += value << ((3 - index) * 8)
	}

	return endCount - startCount
}

func ipToArray(ip string) [4]int {
	values := [4]int{0, 0, 0, 0}

	fmt.Sscanf(ip, "%d.%d.%d.%d", &values[0], &values[1], &values[2], &values[3])

	return values
}
