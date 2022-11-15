package common_denominators

import (
	"strconv"
	"strings"
)

// Average time complexity: O(n log n)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n)
func ConvertFracts(input [][]int) string {
	fractions := make([]Fraction, len(input))
	commonDenominators := make([]int, len(input))

	for index, value := range input {
		fractions[index].Numerator = value[0]
		fractions[index].Denominator = value[1]
		fractions[index].Simplify()
		commonDenominators[index] = value[1]
	}

	previousCommon := 0
	smallestCommon := 0

	for {
		if smallestCommon != 0 && previousCommon == smallestCommon {
			break
		}

		previousCommon = smallestCommon

		for index := range commonDenominators {
			denominator := fractions[index].Denominator

			if commonDenominators[index] < smallestCommon {
				difference := smallestCommon - commonDenominators[index]
				multiple := difference / denominator

				if difference%denominator > 0 {
					multiple += 1
				}

				commonDenominators[index] += denominator * multiple
			}

			if commonDenominators[index] > smallestCommon {
				smallestCommon = commonDenominators[index]
			}
		}
	}

	var builder strings.Builder

	for _, fraction := range fractions {
		fraction.ToCommonDenominator(smallestCommon)
		builder.WriteString("(")
		builder.WriteString(strconv.Itoa(fraction.Numerator))
		builder.WriteString(",")
		builder.WriteString(strconv.Itoa(fraction.Denominator))
		builder.WriteString(")")
	}

	return builder.String()
}

type Fraction struct {
	Numerator, Denominator int
}

func (fraction *Fraction) Simplify() {
	gcf := getGreatestCommonFactor(fraction.Numerator, fraction.Denominator)

	fraction.Numerator /= gcf
	fraction.Denominator /= gcf
}

func (fraction *Fraction) ToCommonDenominator(denominator int) {
	fraction.Numerator *= denominator / fraction.Denominator
	fraction.Denominator = denominator
}

func getGreatestCommonFactor(first, second int) int {
	for second != 0 {
		first, second = second, first%second
	}

	return first
}
