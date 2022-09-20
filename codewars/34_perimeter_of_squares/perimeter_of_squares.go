package perimeter_of_squares

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Perimeter(n int) (perimeter int) {
	var fibonacci Fibonacci

	for n >= -1 {
		perimeter += fibonacci.Next()
		n -= 1
	}

	return perimeter * 4
}

type Fibonacci struct {
	current int
	next    int
}

func (fibonacci *Fibonacci) Next() int {
	next := fibonacci.next

	if next == 0 {
		fibonacci.next = 1
	} else {
		fibonacci.next = next + fibonacci.current
	}

	fibonacci.current = next

	return next
}
