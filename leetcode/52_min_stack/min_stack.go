package min_stack

type MinStack struct {
	values    []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func (minStack *MinStack) Push(value int) {
	minStack.values = append(minStack.values, value)

	if len(minStack.minValues) == 0 || minStack.GetMin() >= value {
		minStack.minValues = append(minStack.minValues, value)
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func (minStack *MinStack) Pop() {
	value := minStack.values[len(minStack.values)-1]
	minStack.values = minStack.values[:len(minStack.values)-1]

	if minStack.GetMin() == value {
		minStack.minValues = minStack.minValues[:len(minStack.minValues)-1]
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func (minStack *MinStack) Top() int {
	return minStack.values[len(minStack.values)-1]
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func (minStack *MinStack) GetMin() int {
	return minStack.minValues[len(minStack.minValues)-1]
}
