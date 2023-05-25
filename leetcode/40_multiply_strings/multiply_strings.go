package multiply_strings

import (
	"fmt"
	"strings"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func Multiply(first string, second string) string {
	return NewOperatable(first).Multiply(NewOperatable(second)).String()
}

type Operatable []string

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func NewOperatable(from string) Operatable {
	length := len(from)
	operatable := make(Operatable, length)

	for index := length - 1; index >= 0; index -= 1 {
		operatable[length-index-1] = from[index : index+1]
	}

	return operatable
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func (to Operatable) Add(from Operatable) Operatable {
	sum := make(Operatable, len(to))

	copy(sum, to)

	if len(from) > len(sum) {
		sum, from = from, sum
	}

	for fromIndex := range from {
		carryIndex, carry := fromIndex, from[fromIndex]

		for {
			if len(sum) <= carryIndex {
				sum = append(sum, "0")
			}
			sum[carryIndex] = AddDigits(sum[carryIndex], carry)

			if len(sum[carryIndex]) == 1 {
				break
			}

			carry = sum[carryIndex][0:1]
			sum[carryIndex] = sum[carryIndex][1:2]
			carryIndex += 1
		}
	}

	return sum
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func (to Operatable) Multiply(from Operatable) Operatable {
	product := NewOperatable("0")

	if len(to) == 1 && to[0] == "0" || len(from) == 1 && from[0] == "0" {
		return product
	}

	for fromIndex, fromDigit := range from {
		fromProduct := NewOperatable("0")

		for toIndex, toDigit := range to {
			fromProduct = fromProduct.Add(
				NewOperatable(MultiplyDigits(fromDigit, toDigit)).ShiftLeft(toIndex),
			)
		}

		product = product.Add(fromProduct.ShiftLeft(fromIndex))
	}

	return product
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func (operatable Operatable) ShiftLeft(shift int) Operatable {
	shifted := make([]string, len(operatable)+shift)
	copy(shifted[shift:], operatable)
	for index := 0; index < shift; index += 1 {
		shifted[index] = "0"
	}
	return shifted
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func (operatable Operatable) String() string {
	var result strings.Builder

	for index := len(operatable) - 1; index >= 0; index -= 1 {
		result.WriteString(operatable[index])
	}

	return result.String()
}

var additionTable map[string]string = make(map[string]string, 100)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func AddDigits(num1 string, num2 string) string {
	return additionTable[fmt.Sprintf("%s:%s", num1, num2)]
}

var multiplicationTable map[string]string = make(map[string]string, 100)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func MultiplyDigits(num1 string, num2 string) string {
	return multiplicationTable[fmt.Sprintf("%s:%s", num1, num2)]
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func init() {
	for row, columns := range [][]string{
		{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
		{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		{"0", "2", "4", "6", "8", "10", "12", "14", "16", "18"},
		{"0", "3", "6", "9", "12", "15", "18", "21", "24", "27"},
		{"0", "4", "8", "12", "16", "20", "24", "28", "32", "36"},
		{"0", "5", "10", "15", "20", "25", "30", "35", "40", "45"},
		{"0", "6", "12", "18", "24", "30", "36", "42", "48", "54"},
		{"0", "7", "14", "21", "28", "35", "42", "49", "56", "63"},
		{"0", "8", "16", "24", "32", "40", "48", "56", "64", "72"},
		{"0", "9", "18", "27", "36", "45", "54", "63", "72", "81"},
	} {
		for column, product := range columns {
			multiplicationTable[fmt.Sprintf("%d:%d", row, column)] = product
		}
	}

	sums := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"10", "11", "12", "13", "14", "15", "16", "17", "18",
	}
	for row := 0; row <= 9; row += 1 {
		for column := 0; column <= 9; column += 1 {
			additionTable[fmt.Sprintf("%d:%d", row, column)] = sums[row+column]
		}
	}
}
