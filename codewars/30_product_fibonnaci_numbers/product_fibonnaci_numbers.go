package product_fibonnaci_numbers

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func ProductFib(prod uint64) [3]uint64 {
	var fibonacci Fibonacci
	var previous uint64
	var current uint64
	var product uint64
	var isExact uint64

	for ; product < prod; product = previous * current {
		previous = current
		current = fibonacci.Next()
	}

	if product == prod {
		isExact = 1
	}

	return [3]uint64{previous, current, isExact}
}

type Fibonacci struct {
	current uint64
	next    uint64
}

func (fibonacci *Fibonacci) Next() uint64 {
	next := fibonacci.next

	if next == 0 {
		fibonacci.next = 1
	} else {
		fibonacci.next = next + fibonacci.current
	}

	fibonacci.current = next

	return next
}
