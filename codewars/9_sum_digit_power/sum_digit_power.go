package sum_digit_power

import (
	"fmt"
)

func SumDigPow(min, max uint64) (result []uint64) {
	for ; min <= max; min += 1 {
		if min%10 == 0 {
			continue
		}

		if min == sumConsecutivePow(min) {
			result = append(result, min)
		}
	}

	return
}

func sumConsecutivePow(num uint64) (sum uint64) {
	exponent := uint64(getNumberOfDigits(num))

	for ; exponent > 0; exponent -= 1 {
		currentDigit := num % 10
		sum += uIntPow(currentDigit, exponent)
		num = num / 10
	}

	return
}

func getNumberOfDigits[t interface{}](num t) int {
	str := fmt.Sprint(num)
	return len([]rune(str))
}

func uIntPow(value, exponent uint64) (result uint64) {
	result = 1

	for ctr := uint64(0); ctr < exponent; ctr += 1 {
		result = result * value
	}

	return
}
