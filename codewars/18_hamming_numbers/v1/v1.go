package hamming_numbers_v1

import (
	"sort"
	"sync"

	lib_uint "github.com/yvnbunag/go-kata/lib/uint"
)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func Hammer(n int) uint {
	var exponent uint = 1
	var from, to uint = 0, lib_uint.Power(2, exponent)
	position := uint(n) - 1

	for {
		hammingNumbers := FindHammingNumbers(from, to)
		hammingNumbersLen := uint(len(hammingNumbers))

		if (hammingNumbersLen - 1) >= position {
			return hammingNumbers[position]
		}

		position -= hammingNumbersLen
		exponent = exponent + 1
		from = to + 1
		to = to + lib_uint.Power(2, exponent)
	}
}

func FindHammingNumbers(from, to uint) (hammingNumbers []uint) {
	defer func() {
		sort.Slice(hammingNumbers, func(x, y int) bool {
			return hammingNumbers[x] < hammingNumbers[y]
		})
	}()

	var waitGroup sync.WaitGroup
	hammingNumbersChannel := make(chan uint, int(to-(from-1)))

	for ; from <= to; from += 1 {
		waitGroup.Add(1)
		go func(value uint) {
			defer waitGroup.Done()

			if IsHammingNumber(value) {
				hammingNumbersChannel <- value
			}
		}(from)
	}
	go func() {
		waitGroup.Wait()
		close(hammingNumbersChannel)
	}()

	for value := range hammingNumbersChannel {
		hammingNumbers = append(hammingNumbers, value)
	}

	return
}

func IsHammingNumber(value uint) bool {
	if value == 0 {
		return false
	}

	if value == 1 {
		return true
	}

	if value%2 == 0 {
		return IsHammingNumber(value / 2)
	}

	if value%3 == 0 {
		return IsHammingNumber(value / 3)
	}

	if value%5 == 0 {
		return IsHammingNumber(value / 5)
	}

	return false
}

func CalculateHammingNumber(i, j, k uint) uint {
	var result uint = 1
	resultsChannel := make(chan uint, 3)
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	go func() {
		defer waitGroup.Done()
		resultsChannel <- lib_uint.Power(2, i)
	}()
	go func() {
		defer waitGroup.Done()
		resultsChannel <- lib_uint.Power(3, j)
	}()
	go func() {
		defer waitGroup.Done()
		resultsChannel <- lib_uint.Power(5, k)
	}()
	go func() {
		waitGroup.Wait()
		close(resultsChannel)
	}()

	for value := range resultsChannel {
		result *= value
	}

	return result
}
