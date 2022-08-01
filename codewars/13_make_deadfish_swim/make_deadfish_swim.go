package make_deadfish_swim

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(log n)
func Parse(data string) (output []int) {
	value := 0

	for _, command := range data {
		switch command {
		case 'i':
			value += 1
		case 'd':
			value -= 1
		case 's':
			value = value * value
		case 'o':
			output = append(output, value)
		default:
		}
	}

	return
}
