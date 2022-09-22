package square_sum_of_squared_divisors

import (
	"math"
	"sync"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n log n)
func ListSquared(from, to int) [][]int {
	var waitGroup sync.WaitGroup
	limiter := make(chan bool, 10)
	squareTotals := make(chan []int)

	for ; from <= to; from += 1 {
		waitGroup.Add(1)
		limiter <- true

		go func(number int) {
			defer func() {
				waitGroup.Done()
				<-limiter
			}()

			squareTotal := sumSquares(getDivisors(number))

			if isSquareRoot(squareTotal) {
				squareTotals <- []int{number, squareTotal}
			}
		}(from)
	}

	go func() {
		waitGroup.Wait()
		close(squareTotals)
	}()

	result := [][]int{}

	for value := range squareTotals {
		result = append(result, value)
	}

	return result
}

func getDivisors(number int) []int {
	divisors := []int{}

	for current := 1; current <= number; current += 1 {
		if number%current == 0 {
			divisors = append(divisors, current)
		}
	}

	return divisors
}

func sumSquares(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number * number
	}

	return total
}

func isSquareRoot(number int) bool {
	squareRoot := math.Sqrt(float64(number))

	return squareRoot == math.Floor(squareRoot)
}
