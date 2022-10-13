package sum_by_factors

import (
	"fmt"
	"math"
	"strings"
)

// Average time complexity: O(n log n)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n log n)
func SumOfDivided(numbers []int) string {
	var max float64

	for _, value := range numbers {
		if absoluteValue := math.Abs(float64(value)); absoluteValue > max {
			max = absoluteValue
		}
	}

	primes := getPrimesUntil(int(max))
	sums := make([]interface{}, len(primes))
	maxIndex := 0

	for _, value := range numbers {
		divisors, lastIndex := getDivisors(value, primes)

		if lastIndex > maxIndex {
			maxIndex = lastIndex
		}

		for index := 0; index <= lastIndex; index += 1 {
			if divisors[index] {
				if sums[index] == nil {
					sums[index] = 0
				}

				sums[index] = sums[index].(int) + value
			}
		}
	}

	var result strings.Builder

	for index := 0; index <= maxIndex; index += 1 {
		if sums[index] != nil {
			result.WriteString(fmt.Sprintf("(%d %d)", primes[index], sums[index]))
		}
	}

	return result.String()
}

func getDivisors(value int, numbers []int) ([]bool, int) {
	absoluteNumber := math.Abs(float64(value))
	lastIndex := 0
	isDivisor := make([]bool, len(numbers))

	for index, number := range numbers {
		if float64(number) > absoluteNumber {
			break
		}

		if value%number == 0 {
			isDivisor[index] = true
			lastIndex = index
		}
	}

	return isDivisor, lastIndex
}

// Sieve of Eratosthenes algorithm.
//
//	See https://stackoverflow.com/a/21923233/13235279
func getPrimesUntil(number int) []int {
	max := number + 1
	isPrime := make([]bool, max)
	primes := []int{}

	for x := 2; x < max; x += 1 {
		if isPrime[x] == true {
			continue
		}

		primes = append(primes, x)

		for y := x * x; y < max; y += x {
			isPrime[y] = true
		}
	}

	return primes
}
