package ugly_numbers

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
// Approach: Dynamic Programming - Bottom Up (Tabulation)
func GetUglyNumber(position int) int {
	uglyNumbers := []int{1}

	if position == 1 {
		return uglyNumbers[0]
	}

	multipleIndexes := map[int]int{
		2: 0,
		3: 0,
		5: 0,
	}
	multiples := map[int]int{
		2: 2,
		3: 3,
		5: 5,
	}

	for currentPosition := 1; currentPosition < position; currentPosition += 1 {
		value := min(min(multiples[2], multiples[3]), multiples[5])
		uglyNumbers = append(uglyNumbers, value)

		if multiples[2] == value {
			multipleIndexes[2] += 1
			multiples[2] = uglyNumbers[multipleIndexes[2]] * 2
		}

		if multiples[3] == value {
			multipleIndexes[3] += 1
			multiples[3] = uglyNumbers[multipleIndexes[3]] * 3
		}

		if multiples[5] == value {
			multipleIndexes[5] += 1
			multiples[5] = uglyNumbers[multipleIndexes[5]] * 5
		}
	}

	return uglyNumbers[len(uglyNumbers)-1]
}

func min(first, second int) int {
	if first < second {
		return first
	}

	return second
}
