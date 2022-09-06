package directions_reduction

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func DirReduc(arr []string) []string {
	reduction := []string{}

	for _, direction := range arr {
		reductionLength := len(reduction)

		if reductionLength > 0 && isOpposite(reduction[reductionLength-1], direction) {
			reduction = reduction[:reductionLength-1]
			continue
		}

		reduction = append(reduction, direction)
	}

	return reduction
}

var oppositeMap = map[string]string{
	"NORTH": "SOUTH",
	"SOUTH": "NORTH",
	"EAST":  "WEST",
	"WEST":  "EAST",
}

func isOpposite(first, second string) bool {
	if first == oppositeMap[second] {
		return true
	}

	return false
}
