package find_outlier

func isEven(integer int) bool {
	return integer%2 == 0
}

func FindOutlier(integers []int) int {
	first := integers[0]
	second := integers[1]
	firstIsEven := isEven(first)
	secondIsEven := isEven(second)

	if firstIsEven == secondIsEven {
		for _, value := range integers[2:] {
			if isEven(value) != secondIsEven {
				return value
			}
		}
	}

	third := integers[2]
	thirdIsEven := isEven(third)

	if secondIsEven == thirdIsEven {
		return first
	}

	return second
}
