package evaluate_reverse_polish_notation

import "strconv"

var operations = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func EvalRPN(tokens []string) int {
	values := []int{}

	for _, token := range tokens {
		if operate, ok := operations[token]; ok {
			a, b := values[len(values)-2], values[len(values)-1]
			values = append(values[:len(values)-2], operate(a, b))
			continue
		}

		value, _ := strconv.Atoi(token)
		values = append(values, value)
	}

	return values[0]
}
