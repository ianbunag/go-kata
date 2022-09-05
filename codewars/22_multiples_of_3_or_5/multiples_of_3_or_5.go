package multiples_of_3_or_5

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Multiple3And5(number int) int {
	if number < 3 {
		return 0
	}

	multipleSum := MultipleSum{}
	threeMultiple := 3
	fiveMultiple := 5

	for threeMultiple < number || fiveMultiple < number {
		if threeMultiple < fiveMultiple {
			multipleSum.Add(threeMultiple)
			threeMultiple += 3
			continue
		}

		if fiveMultiple < threeMultiple {
			multipleSum.Add(fiveMultiple)
			fiveMultiple += 5
			continue
		}

		// threeMultiple == fiveMultiple
		multipleSum.Add(threeMultiple)
		fiveMultiple += 5
	}

	return multipleSum.Value()
}

type MultipleSum struct {
	sum  int
	last int
}

func (multiple *MultipleSum) Add(value int) {
	if value <= multiple.last {
		return
	}

	multiple.last = value
	multiple.sum += value
}

func (multiple *MultipleSum) Value() int {
	return multiple.sum
}
