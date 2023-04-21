package decode_ways

import (
	"strconv"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func NumDecodings(value string) (decoded int) {
	decodings := make(Decodings, len(value))
	decoded += calculateDecodings(value, decodings, 0, 1)
	decoded += calculateDecodings(value, decodings, 0, 2)

	return
}

func calculateDecodings(value string, decodings Decodings, from, to int) int {
	if len(value) < to || !isValid(value[from:to]) {
		return 0
	}

	if len(value) == to {
		return 1
	}

	if !decodings.isDecoded(to) {
		decoded := calculateDecodings(value, decodings, to, to+1)
		decoded += calculateDecodings(value, decodings, to, to+2)
		decodings.setDecoded(to, decoded)
	}

	return decodings.getDecoded(to)
}

type Decodings []int

func (decodings Decodings) isDecoded(index int) bool {
	return decodings[index] != 0
}

func (decodings Decodings) getDecoded(index int) int {
	return decodings[index] - 1
}

func (decodings Decodings) setDecoded(index, value int) {
	decodings[index] = value + 1
}

func isValid(value string) bool {
	if value[0] == '0' {
		return false
	}

	intValue, err := strconv.Atoi(value)

	return err == nil && intValue >= 1 && intValue <= 26
}
